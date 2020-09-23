package graph

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/ks6088ts/spidey/api/repository"
)

type Resolver struct{
	TodoRepository repository.TodoRepository
}
