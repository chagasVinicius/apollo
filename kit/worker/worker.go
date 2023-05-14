package worker

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type JobFunc func(ctx context.Context)

type Worker struct {
	wg         sync.WaitGroup
	mu         sync.RWMutex
	sem        chan bool
	isShutdown chan struct{}
	running    map[string]context.CancelFunc
}

func New(maxRunningJobs int) (*Worker, error) {
	if maxRunningJobs <= 0 {
		return nil, errors.New("max running jobs must be greater than 0")
	}

	sem := make(chan bool, maxRunningJobs)

	for i := 0; i < maxRunningJobs; i++ {
		sem <- true
	}

	w := Worker{
		sem:        sem,
		isShutdown: make(chan struct{}),
		running:    make(map[string]context.CancelFunc),
	}

	return &w, nil
}

func (w *Worker) Running() int {
	w.mu.RLock()
	defer w.mu.RUnlock()

	return len(w.running)
}

func (w *Worker) Shutdown(ctx context.Context) error {
	close(w.isShutdown)

	w.mu.RLock()
	{
		for _, cancel := range w.running {
			cancel()
		}
	}
	w.mu.RUnlock()

	ch := make(chan struct{})
	go func() {
		w.wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func (w *Worker) Start(ctx context.Context, fn JobFunc) (string, error) {
	select {
	case <-w.isShutdown:
		return "", errors.New("shutting down")
	case <-ctx.Done():
		return "", ctx.Err()
	case <-w.sem:
	}

	workKey := uuid.NewString()

	deadline, ok := ctx.Deadline()
	if !ok {
		deadline = time.Now().Add(time.Second)
	}

	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	w.mu.Lock()
	defer w.mu.Unlock()
	w.running[workKey] = cancel

	w.wg.Add(1)

	go func() {
		defer func() { w.sem <- true }()

		defer func() {
			cancel()
			w.removeWork(workKey)
			w.wg.Done()
		}()

		fn(ctx)
	}()

	return workKey, nil
}

func (w *Worker) Stop(workKey string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	cancel, exists := w.running[workKey]
	if !exists {
		return fmt.Errorf("work[%s] is not running", workKey)
	}

	cancel()

	return nil
}

func (w *Worker) removeWork(workKey string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	delete(w.running, workKey)
}
