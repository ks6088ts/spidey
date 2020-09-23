package repository

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ks6088ts/spidey/api/domains"
)

func NewMockTodoRepository() TodoRepository {
	return &mockTodoRepository{
		db: map[domains.TodoID]*domains.Todo{},
	}
}

var _ TodoRepository = (*mockTodoRepository)(nil)

type mockTodoRepository struct {
	sync.RWMutex
	db map[domains.TodoID]*domains.Todo
}

func (repo *mockTodoRepository) Get(ctx context.Context, id domains.TodoID) (*domains.Todo, error) {
	repo.RLock()
	defer repo.RUnlock()

	todo, ok := repo.db[id]
	if !ok {
		return nil, ErrNoSuchEntity
	}

	return todo, nil
}

func (repo *mockTodoRepository) GetAll(ctx context.Context) ([]*domains.Todo, error) {
	repo.RLock()
	defer repo.RUnlock()
	list := make([]*domains.Todo, 0, len(repo.db))

	for _, todo := range repo.db {
		list = append(list, todo)
	}

	sort.Slice(list, func(i, j int) bool {
		a := list[i]
		b := list[j]
		return a.CreatedAt.UnixNano() > b.CreatedAt.UnixNano()
	})
	return list, nil
}
func (repo *mockTodoRepository) Create(ctx context.Context, todo *domains.Todo) (*domains.Todo, error) {
	if todo.ID != "" {
		return nil, ErrBadRequest
	}

	repo.Lock()
	defer repo.Unlock()

	todo.ID = domains.TodoID(uuid.New().String())
	todo.CreatedAt = time.Now()

	repo.db[todo.ID] = todo
	return todo, nil
}

func (repo *mockTodoRepository) Update(ctx context.Context, todo *domains.Todo) (*domains.Todo, error) {
	if todo.ID == "" {
		return nil, ErrBadRequest
	}

	repo.Lock()
	defer repo.Unlock()

	_, ok := repo.db[todo.ID]
	if !ok {
		return nil, ErrNoSuchEntity
	}

	repo.db[todo.ID] = todo
	copyTodo := *todo
	return &copyTodo, nil
}
