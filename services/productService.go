package services

import "gs-1/repositories"

type ProductService interface {
	Get(id int) (*repositories.Product, error)
}

type productService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (p *productService) Get(id int) (*repositories.Product, error) {
	return p.repo.Get(id) // just pass this up the stack, but in "real-life" would likely do something here (maybe log/telemetry/metrics)
}
