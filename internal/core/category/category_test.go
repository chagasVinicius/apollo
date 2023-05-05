package category_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/chagasVinicius/apollo/internal/core/category"
	"github.com/chagasVinicius/apollo/internal/core/category/stores/categorydb"
	"github.com/chagasVinicius/apollo/internal/data/dbtest"
	"github.com/chagasVinicius/apollo/kit/docker"
	"github.com/google/go-cmp/cmp"
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

func TestCategory(t *testing.T) {
	log, db, teardown := dbtest.NewUnit(t, c, "testcategory")
	t.Cleanup(teardown)

	core := category.NewCore(categorydb.NewStore(log, db))

	t.Log("Given the need to work with Category records.")
	{
		t.Logf("\tWhen handling a single Category.")
		{
			ctx := context.Background()

			nc := category.NewCategory{
				Name:      "Test Category",
				ShortDesc: "Test Short Description",
			}

			category, err := core.Create(ctx, nc)
			if err != nil {
				t.Fatalf("\t [ERROR] Should be able to add a category : %s", err)
			}
			t.Logf("\t [SUCCESS] Should be able to add a category.")

			saved, err := core.QueryByID(ctx, category.ID)
			if err != nil {
				t.Fatalf("\t [ERROR] Should be able to query a category by id : %s", err)
			}
			t.Logf("\t [SUCCESS] Should be able to query a category by id.")

			category.CreatedAt = time.Time{}
			saved.CreatedAt = time.Time{}

			if diff := cmp.Diff(category, saved); diff != "" {
				t.Fatalf("\t [ERROR] Should get back the same category : %s", diff)
			}
			t.Logf("\t [SUCCESS] Should get back the same category.")

			categories, err := core.Query(ctx, 1, 10)
			if err != nil {
				t.Fatalf("\t [ERROR] Should be able to query categories : %s", err)
			}
			t.Logf("\t [SUCCESS] Should be able to query categories.")
			fmt.Println(categories)

			if len(categories) == 0 {
				t.Fatalf("\t [ERROR] Should get back at least one category.")
			}
			t.Logf("\t [SUCCESS] Should get back at least one category.")
		}
	}
}
