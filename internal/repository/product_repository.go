package repository

import (
	"errors"

	"gorm.io/gorm"

	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/request"
	"go-clean-architecture/internal/utils"
)

type ProductRepositoryInterface interface {
	Add(product model.Product)
	GetAll() []model.Product
	GetById(productId int) (product model.Product, err error)
	Update(product model.Product)
	Delete(productId int)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) Add(product model.Product) {
	result := pr.db.Create(&product)
	utils.ErrorPanic(result.Error)
}

func (pr *ProductRepository) GetAll() []model.Product {
	var products []model.Product
	result := pr.db.Find(&products)
	utils.ErrorPanic(result.Error)

	return products
}

func (pr *ProductRepository) GetById(productId int) (model.Product, error) {
	var product model.Product
	result := pr.db.First(&product, productId).Error

	if result != nil {
		return product, errors.New("product not found")
	} else {
		return product, nil
	}
}

func (pr *ProductRepository) Update(product model.Product) {
	var updateProduct = request.UpdateProductRequest{
		ID:          product.ID,
		ProductName: product.ProductName,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}

	result := pr.db.Model(&product).Updates(updateProduct)
	utils.ErrorPanic(result.Error)
}

func (pr *ProductRepository) Delete(productId int) {
	var product model.Product

	result := pr.db.Where("id = ?", productId).Delete(&product)
	utils.ErrorPanic(result.Error)
}
