package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	if err := p.db.Create(&product).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {
	products := []model.Product{}

	if err := p.db.Table("products").Where("deleted_at IS NULL").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	products := model.Product{}

	if err := p.db.Where("id = ?", id).Delete(&products).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	if err := p.db.Table("products").Where("id = ?", id).Update("id", id).Error; err != nil {
		return err
	}

	return nil
}
