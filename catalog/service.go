package catalog

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostProduct(ctx context.Context, p Product) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context, skip uint64, take uint64, ids []string) ([]Product, error)
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type catalogService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &catalogService{r}
}

func (s *catalogService) PostProduct(ctx context.Context, p Product) (*Product, error) {
	p.ID = ksuid.New().String()
	if err := s.repository.PutProduct(ctx, p); err != nil {
		return nil, err
	}
	return &Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, nil
}

func (s *catalogService) GetProduct(ctx context.Context, id string) (*Product, error) {
	return s.repository.GetProductByID(ctx, id)
}

func (s *catalogService) GetProducts(ctx context.Context, skip uint64, take uint64, ids []string) ([]Product, error) {
	if take > 100 {
		take = 100
	}
	if len(ids) == 0 {
		return s.repository.ListProducts(ctx, skip, take)
	} else {
		return s.repository.ListProductsWithIDs(ctx, skip, take, ids)
	}
}