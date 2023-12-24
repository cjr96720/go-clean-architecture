package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-clean-architecture/internal/request"
	"go-clean-architecture/internal/response"
	"go-clean-architecture/internal/service"
	"go-clean-architecture/internal/utils"
)

type ProductController struct {
	productService service.ProductServiceInterface
}

func NewProductController(ps service.ProductServiceInterface) *ProductController {
	return &ProductController{productService: ps}
}

// AddProduct    godoc
//
//	@Summary		Add a product
//	@Description	Add a product to database
//	@Accept			json
//	@Produce		application/json
//	@Param			products	body	request.AddProductRequest	true	"product you want to add"
//	@Success		200				{object}	response.DefaultResponse
//	@Router			/api/v1/product [post]
func (pc *ProductController) Add(c *gin.Context) {
	procuct := request.AddProductRequest{}
	err := c.ShouldBindJSON(&procuct)
	utils.ErrorPanic(err)

	pc.productService.AddProduct(procuct)

	response := response.DefaultResponse{
		Message: "SUCCESS",
		Data:    nil,
	}
	c.JSON(http.StatusOK, response)
}

// GetAllProduct    godoc
//
//	@Summary		Get all products
//	@Description	Get all products from database
//	@Accept			json
//	@Produce		application/json
//	@Success		200				{object}	response.DefaultResponse
//	@Router			/api/v1/product [get]
func (pc *ProductController) GetAll(c *gin.Context) {
	productResponse := pc.productService.GetAllProduct()
	response := response.DefaultResponse{
		Message: "SUCCESS",
		Data:    productResponse,
	}
	c.JSON(http.StatusOK, response)
}

// GetProductById    godoc
//
//	@Summary		Get product by its ID
//	@Description	Get product by its ID
//	@Param			productId	path	string	true "product ID"
//	@Produce		application/json
//	@Success		200				{object}	response.DefaultResponse
//	@Router			/api/v1/product/{productId} [get]
func (pc *ProductController) GetById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param(("productId")))
	utils.ErrorPanic(err)

	productResponse, err := pc.productService.GetProductById(productId)

	if err != nil {
		response := response.DefaultResponse{
			Message: "PRODUCT NOT FOUNT",
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := response.DefaultResponse{
			Message: "SUCCESS",
			Data:    productResponse,
		}
		c.JSON(http.StatusOK, response)
	}
}

// UpdateProduct    godoc
//
//	@Summary		Update product by its ID
//	@Description	Update product by its ID
//	@Param			productId	path	string							true	"product ID"
//	@Param			products	body	request.UpdateProductRequest	true	"things you want to update"
//	@Produce		application/json
//	@Success		200				{object}	response.DefaultResponse
//	@Router			/api/v1/product/{productId} [patch]
func (pc *ProductController) Update(c *gin.Context) {
	updateProductRequest := request.UpdateProductRequest{}
	err := c.ShouldBindJSON(&updateProductRequest)
	utils.ErrorPanic(err)

	productId, err := strconv.Atoi(c.Param("productId"))
	utils.ErrorPanic(err)

	updateProductRequest.ID = productId
	pc.productService.UpdateProduct(updateProductRequest)

	response := response.DefaultResponse{
		Message: "SUCCESS",
	}
	c.JSON(http.StatusOK, response)
}

// DeleteProductById    godoc
//
//	@Summary		Delete product by its ID
//	@Description	Delete product by its ID
//	@Param			productId	path	string	true	"product ID"
//	@Produce		application/json
//	@Success		200				{object}	response.DefaultResponse
//	@Router			/api/v1/product/{productId} [delete]
func (pc *ProductController) Delete(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param(("productId")))
	utils.ErrorPanic(err)

	pc.productService.DeleteProduct(productId)

	response := response.DefaultResponse{
		Message: "SUCCESS",
	}
	c.JSON(http.StatusOK, response)
}
