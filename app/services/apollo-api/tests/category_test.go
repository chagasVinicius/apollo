package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chagasVinicius/apollo/app/services/apollo-api/handlers"
	"github.com/chagasVinicius/apollo/internal/core/category"
	"github.com/chagasVinicius/apollo/internal/data/dbtest"
	"github.com/chagasVinicius/apollo/internal/sys/validate"
	v1Web "github.com/chagasVinicius/apollo/internal/web/v1"
	"github.com/chagasVinicius/apollo/kit/docker"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var c *docker.Container

func TestMain(m *testing.M) {
	var err error
	c, err = dbtest.StartDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbtest.StopDB(c)

	m.Run()
}

// CategoryTests holds methods for each category subtest. This type allows
// passing dependencies for tests while still providing a convenient syntax
// when subtests are registered.
type CategoryTests struct {
	app http.Handler
}

func TestCategories(t *testing.T) {
	t.Parallel()

	test := dbtest.NewIntegration(t, c, "inttestprods")
	t.Cleanup(test.Teardown)

	shutdown := make(chan os.Signal, 1)
	tests := CategoryTests{
		app: handlers.APIMux(handlers.APIMuxConfig{
			Shutdown: shutdown,
			Log:      test.Log,
			DB:       test.DB,
		}),
	}

	t.Run("postCategories400EmptyJSON", tests.postCategories400EmptyJSON)
	t.Run("postCategories400WrongJson", tests.postCategories400WrongJSON)
	t.Run("getCategories400", tests.getCategories400)
	t.Run("getCategories404", tests.getCategories404)
	t.Run("category200Flow", tests.category200Flow)
}

// postCategories400 validates a category can't be created with the endpoint
// unless a valid category payload is submitted.
func (ct *CategoryTests) postCategories400EmptyJSON(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/v1/categories", strings.NewReader(`{}`))
	w := httptest.NewRecorder()

	ct.app.ServeHTTP(w, r)

	t.Log("Given the need to validate a new category can't be created with an invalid payload.")
	{
		t.Log("\t When using an incomplete category value.")
		{
			if w.Code != http.StatusBadRequest {
				t.Fatalf("\t [ERROR] Should receive a status code of 400 for the response : %v", w.Code)
			}
			t.Log("\t [SUCCESS] Should receive a status code of 400 for the response.")

			// Inspect the response.
			var got v1Web.ErrorResponse
			if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
				t.Fatalf("\t [ERROR] Should be able to unmarshal the response to an error type : %v", err)
			}
			t.Log("\t [SUCCESS] Should be able to unmarshal the response to an error type.")

			fields := validate.FieldErrors{
				{Field: "name", Err: "name is a required field"},
				{Field: "short_desc", Err: "short_desc is a required field"},
			}
			exp := v1Web.ErrorResponse{
				Error:  "data validation error",
				Fields: fields.Fields(),
			}

			// We can't rely on the order of the field errors so they have to be
			// sorted. Tell the cmp package how to sort them.
			sorter := cmpopts.SortSlices(func(a, b validate.FieldError) bool {
				return a.Field < b.Field
			})

			if diff := cmp.Diff(got, exp, sorter); diff != "" {
				t.Fatalf("\t [ERROR] Should get the expected result. Diff:\n%s", diff)
			}
			t.Log("\t [SUCCESS] Should get the expected result.")
		}
	}
}

func (ct *CategoryTests) postCategories400WrongJSON(t *testing.T) {
	WrongBody := map[string]interface{}{
		"foo": "bar",
	}
	body, _ := json.Marshal(WrongBody)
	r := httptest.NewRequest(http.MethodPost, "/v1/categories", bytes.NewReader(body))
	w := httptest.NewRecorder()

	ct.app.ServeHTTP(w, r)

	t.Log("Given the need to validate a new category can't be created with an invalid payload.")
	{
		t.Log("\t When using an incomplete category value.")
		{
			if w.Code != http.StatusBadRequest {
				t.Fatalf("\t [ERROR] Should receive a status code of 400 for the response : %v", w.Code)
			}
			t.Log("\t [SUCCESS] Should receive a status code of 400 for the response.")

			// Inspect the response.
			var got v1Web.ErrorResponse
			if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
				t.Fatalf("\t [ERROR] Should be able to unmarshal the response to an error type : %v", err)
			}
			t.Log("\t [SUCCESS] Should be able to unmarshal the response to an error type.")

			exp := v1Web.ErrorResponse{
				Error:  "invalid json data",
				Fields: nil,
			}

			// // we can't rely on the order of the field errors so they have to be
			// // sorted. Tell the cmp package how to sort them.
			sorter := cmpopts.SortSlices(func(a, b validate.FieldError) bool {
				return a.Field < b.Field
			})

			if diff := cmp.Diff(got, exp, sorter); diff != "" {
				t.Fatalf("\t [ERROR] Should get the expected result. Diff:\n%s", diff)
			}
			t.Log("\t [SUCCESS] Should get the expected result.")
		}
	}
}

// getCategories400 validates a category request for a malformed id.
func (ct *CategoryTests) getCategories400(t *testing.T) {
	id := "12345"

	r := httptest.NewRequest(http.MethodGet, "/v1/categories/"+id, nil)
	w := httptest.NewRecorder()

	ct.app.ServeHTTP(w, r)

	t.Log("Given the need to validate getting a product with a malformed id.")
	{
		t.Logf("\t When using the new category %s.", id)
		{
			if w.Code != http.StatusBadRequest {
				t.Fatalf("\t [ERROR] Should receive a status code of 400 for the response : %v", w.Code)
			}
			t.Log("\t [SUCCESS] Should receive a status code of 400 for the response.")

			got := w.Body.String()
			exp := `{"error":"ID is not in its proper form"}`
			if got != exp {
				t.Fatalf("\t [ERROR] Should get the expected result.\n\t\t Got: %s.\n\t\t Exp: %s", got, exp)
			}
			t.Log("\t [SUCCESS] Should get the expected result.")
		}
	}
}

func (ct *CategoryTests) getCategories404(t *testing.T) {
	id := "112262f1-1a77-4374-9f22-39e575aa6348"

	r := httptest.NewRequest(http.MethodGet, "/v1/categories/"+id, nil)
	w := httptest.NewRecorder()

	ct.app.ServeHTTP(w, r)

	t.Log("Given the need to validate deleting a product that does not exist.")
	{
		t.Log("\t Given the need to validate getting a category with an unknown id.")
		{
			if w.Code != http.StatusNotFound {
				t.Fatalf("\t [ERROR] Should receive a status code of 404 for the response : %v", w.Code)
			}
			t.Log("\t [SUCCESS] Should receive a status code of 404 for the response.")
		}
	}
}

func (ct *CategoryTests) category200Flow(t *testing.T) {
	var got string
	var err any
	t.Log("Given correct input should be able to create a category")
	{
		Body := map[string]interface{}{
			"name":       "category_name",
			"short_desc": "short desc",
		}
		body, _ := json.Marshal(Body)
		r := httptest.NewRequest(http.MethodPost, "/v1/categories", bytes.NewReader(body))
		w := httptest.NewRecorder()

		ct.app.ServeHTTP(w, r)

		if w.Code != http.StatusCreated {
			t.Fatalf("\t [ERROR] Should receive a status code of 201 for the response: %v", w.Code)
		}
		t.Log("\t [SUCCESS] Should receive a status code of 201 for the response.")

		got = w.Body.String()
		created := category.Category{}
		err = json.Unmarshal([]byte(got), &created)
		if err != nil {
			t.Fatalf("\t [ERROR] Should be created a Category entity: %v", err)
		}
		t.Log("\t [SUCCESS] Should be created a Category entity.")

		id := created.ID
		r2 := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/categories/%v", id), nil)
		w2 := httptest.NewRecorder()

		ct.app.ServeHTTP(w2, r2)

		if w2.Code != http.StatusOK {
			t.Fatalf("\t [ERROR] Should receive a status code of 200 for the response: %v", w2.Code)
		}
		t.Log("\t [SUCCESS] Should receive a status code of 200 for the response.")

		got = w2.Body.String()
		saved := category.Category{}
		err = json.Unmarshal([]byte(got), &saved)

		if err != nil {
			t.Fatalf("\t [ERROR] Should retrieved category: %v", err)
		}
		t.Log("\t [SUCCESS] Should retrieved category.")

		r3 := httptest.NewRequest(http.MethodGet, "/v1/categories", nil)
		w3 := httptest.NewRecorder()

		ct.app.ServeHTTP(w3, r3)

		if w3.Code != http.StatusOK {
			t.Fatalf("\t [ERROR] Should receive a status code of 200 for the response: %v", w3.Code)
		}
		t.Log("\t [SUCCESS] Should receive a status code of 200 for the response.")

		got = w3.Body.String()
		listed := []category.Category{}
		err = json.Unmarshal([]byte(got), &listed)
		clisted := listed[0]

		if err == nil {
			t.Fatalf("\t [ERROR] Should retrieved categories: %v", err)
		}
		t.Log("\t [SUCCESS] Should retrieved categories.")


		created.CreatedAt = time.Time{}
		saved.CreatedAt = time.Time{}
		clisted.CreatedAt = time.Time{}
		diffcs := cmp.Diff(created, saved)
		diffcl := cmp.Diff(created, clisted)

		if (diffcs != "") && (diffcl != "") {
				t.Fatalf("\t [ERROR] Should get back the same category : %s", diffcs)
				t.Fatalf("\t [ERROR] Should get back the same category : %s", diffcl)
			}
		t.Logf("\t [SUCCESS] Should get back the same category.")
	}
}
