package repository

import (
	"day2/model"
	"testing"
)

func TestProductRepositoryCRUD(t *testing.T) {

	t.Run(
		"Should not nil after create new instance",
		func(t *testing.T) {

			pr := NewProductRepository()

			if pr == nil {
				t.Errorf("Expected pr not nil, but got %v", pr)
			}
		},
	)

	t.Run(
		"Sould return err nil first save product to repository",
		func(t *testing.T) {

			pr := NewProductRepository()

			product := model.Product{
				ID:    1,
				Name:  "Indomie",
				Stock: 5,
			}

			err := pr.Save(product)

			if err != nil {
				t.Errorf("Expected err nil, but got %v", err)
			}
		})

	t.Run(
		"Sould return err not nil save product same id to repository",
		func(t *testing.T) {

			pr := NewProductRepository()

			product := model.Product{
				ID:    1,
				Name:  "Indomie",
				Stock: 5,
			}

			product2 := model.Product{
				ID:    1,
				Name:  "Indomie",
				Stock: 5,
			}

			err := pr.Save(product)

			err2 := pr.Save(product2)

			if err2 != nil {
				t.Errorf("Expected err not nil, but got %v", err)
			}
		})
}
