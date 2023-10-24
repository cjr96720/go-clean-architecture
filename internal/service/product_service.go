package service

import (
	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/repository"
	"go-clean-architecture/internal/request"
	"go-clean-architecture/internal/response"
)

type ProductServiceInterface interface {
	AddProduct(product request.AddProductRequest)
	GetAllProduct() (products []response.ProductResponse)
	GetProductById(productId int) (response.ProductResponse, error)
	UpdateProduct(product request.UpdateProductRequest)
	DeleteProduct(productId int)
}

type ProductService struct {
	productRepository repository.ProductRepositoryInterface
}

func NewProductService(pr repository.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{productRepository: pr}
}

func (ps *ProductService) AddProduct(product request.AddProductRequest) {
	productModel := model.Product{
		ProductName: product.ProductName,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}

	ps.productRepository.Add(productModel)
}

func (ps *ProductService) GetAllProduct() []response.ProductResponse {
	result := ps.productRepository.GetAll()

	var products []response.ProductResponse
	for _, value := range result {
		product := response.ProductResponse{
			ID:          value.ID,
			ProductName: value.ProductName,
			Price:       value.Price,
			Quantity:    value.Quantity,
		}
		products = append(products, product)
	}

	return products
}

func (ps *ProductService) GetProductById(productId int) (response.ProductResponse, error) {
	productData, err := ps.productRepository.GetById(productId)
	productResponse := response.ProductResponse{
		ID:          productData.ID,
		ProductName: productData.ProductName,
		Price:       productData.Price,
		Quantity:    productData.Quantity,
	}

	if err != nil {
		return productResponse, err
	} else {
		return productResponse, nil
	}
}

func (ps *ProductService) UpdateProduct(product request.UpdateProductRequest) {
	productModel := model.Product{
		ID:          product.ID,
		ProductName: product.ProductName,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}
	ps.productRepository.Update(productModel)
}

func (ps *ProductService) DeleteProduct(productId int) {
	ps.productRepository.Delete(productId)
}
