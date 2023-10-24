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
//	@Produce		json
//	@Param			product_name	query		string	true	"product name"
//	@Param			price			query		number	false	"price of the product"
//	@Param			quantity		query		int		false	"quantity"
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
//	@Produce		json
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

func (pc *ProductController) Delete(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param(("productId")))
	utils.ErrorPanic(err)

	pc.productService.DeleteProduct(productId)

	response := response.DefaultResponse{
		Message: "SUCCESS",
	}
	c.JSON(http.StatusOK, response)
}
