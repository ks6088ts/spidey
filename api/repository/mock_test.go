package repository

import (
	"context"
	"testing"
	"github.com/ks6088ts/spidey/api/domains"
)

func Test_todoRepository(t *testing.T) {
	ctx := context.Background()
	repo := NewMockTodoRepository()

	todoA, err := repo.Create(ctx, &domains.Todo{
		Name: "test A",
	})

	if err != nil {
		t.Fatal(err)
	}
	if v := todoA.CreatedAt; v.IsZero() {
		t.Errorf("unexpected: %#v", v)
	}

	_, err = repo.Update(ctx, &domains.Todo{
		ID:   todoA.ID,
		Name: "test A!",
	})
	if err != nil {
		t.Fatal(err)
	}

	todoA2, err := repo.Get(ctx, todoA.ID)
	if err != nil {
		t.Fatal(err)
	}
	if v1, v2 := todoA.ID, todoA2.ID; v1 != v2 {
		t.Errorf("unexpected: %#v, %#v", v1, v2)
	}

	todoB, err := repo.Create(ctx, &domains.Todo{
		Name: "test B",
	})
	if err != nil {
		t.Fatal(err)
	}

	list, err := repo.GetAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if v := len(list); v != 2 {
		t.Fatalf("unexpected: %#v", v)
	}
	if v := list[0]; v.ID != todoB.ID {
		t.Fatalf("unexpected: %#v", v)
	}
	if v := list[1]; v.ID != todoA.ID {
		t.Fatalf("unexpected: %#v", v)
	}
}
