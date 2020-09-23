package domains

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
)

type TodoID string

type Todo struct {
	ID        TodoID
	Name      string
	CreatedAt time.Time
}

type TodoRepository interface {
	Get(ctx context.Context, id TodoID) (*Todo, error)
	GetAll(ctx context.Context) ([]*Todo, error)
	Create(ctx context.Context, todo *Todo) (*Todo, error)
	Update(ctx context.Context, todo *Todo) (*Todo, error)
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{
		db: map[TodoID]*Todo{},
	}
}

var _ TodoRepository = (*todoRepository)(nil)

type todoRepository struct {
	sync.RWMutex
	db map[TodoID]*Todo
}

func (repo *todoRepository) Get(ctx context.Context, id TodoID) (*Todo, error) {
	repo.RLock()
	defer repo.RUnlock()

	todo, ok := repo.db[id]
	if !ok {
		return nil, ErrNoSuchEntity
	}

	return todo, nil
}

func (repo *todoRepository) GetAll(ctx context.Context) ([]*Todo, error) {
	repo.RLock()
	defer repo.RUnlock()
	list := make([]*Todo, 0, len(repo.db))

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
func (repo *todoRepository) Create(ctx context.Context, todo *Todo) (*Todo, error) {
	if todo.ID != "" {
		return nil, ErrBadRequest
	}

	repo.Lock()
	defer repo.Unlock()

	todo.ID = TodoID(uuid.New().String())
	todo.CreatedAt = time.Now()

	repo.db[todo.ID] = todo
	return todo, nil
}

func (repo *todoRepository) Update(ctx context.Context, todo *Todo) (*Todo, error) {
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
