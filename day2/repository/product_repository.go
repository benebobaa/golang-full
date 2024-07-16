package repository

import (
	"day2/model"
	"errors"
	"log"
)

type ProductRepository struct {
	Products map[int]model.Product
}

func NewProductRepository() *ProductRepository {
	products := make(map[int]model.Product)
	return &ProductRepository{
		Products: products,
	}
}

func (p *ProductRepository) Save(product model.Product) error {

	_, ok := p.Products[product.ID]

	if ok {
		log.Println("ok -> ", ok)
		return errors.New("ID already exists")
	}

	p.Products[product.ID] = product

	return nil
}

func (p *ProductRepository) GetLastProduct() model.Product {
	var product model.Product

	for _, value := range p.Products {
		product = value
	}

	return product
}

func (p *ProductRepository) FindAll() []model.Product {
	var products []model.Product

	for _, value := range p.Products {
		products = append(products, value)
	}

	return products
}
