package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DiegoPreciadoM/Go-Web/tree/main/02-Go-Web-Afternoon/Practices/Project/internal/domain"
	"github.com/DiegoPreciadoM/Go-Web/tree/main/02-Go-Web-Afternoon/Practices/Project/internal/product"
	"github.com/gin-gonic/gin"
)

type productHandler struct {
	s product.Service
}

func NewProductHandler(s product.Service) *productHandler {
	return &productHandler{
		s: s,
	}
}

func (h *productHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, _ := h.s.GetAll()
		ctx.JSON(http.StatusAccepted, gin.H{
			"data": products,
		})
		return
	}
}

func (h *productHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid id",
			})
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Product not found",
			})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{
			"data": product,
		})
		return
	}
}

func (h *productHandler) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		priceParam := ctx.Query("priceGt")
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid price",
			})
			return
		}
		products, err := h.s.SearchByPriceGt(price)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "No products found",
			})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{
			"data": products,
		})
		return
	}
}

func ValidateExpirationDate(new_product string) bool {
	splited_date := strings.Split(new_product, "/")
	day, _ := strconv.Atoi(splited_date[0])
	month, _ := strconv.Atoi(splited_date[1])
	year, _ := strconv.Atoi(splited_date[2])

	if day < 1 || day > 31 {
		return false
	}

	if month < 1 || month > 12 {
		return false
	}

	if year < 0 || year > 9999 {
		return false
	}

	return true
}

func (h *productHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token := ctx.GetHeader("token")
		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "Invalid token",
		// 	})
		// }

		var product domain.Product

		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Verify body information",
			})
			return
		}

		if !ValidateExpirationDate(product.Expiration) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "Expiration date is incorrect",
			})
			return
		}

		p, err := h.s.Create(product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"data": p,
		})
		return

	}
}

func (h *productHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token := ctx.GetHeader("token")
		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "Invalid token",
		// 	})
		// 	return
		// }

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid id",
			})
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Product Deleted",
		})
		return
	}
}

func (h *productHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// token := ctx.GetHeader("token")
		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "Invalid token",
		// 	})
		// 	return
		// }

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid id",
			})
			return
		}

		var product domain.Product
		if err = ctx.ShouldBindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Verify body information",
			})
			return
		}

		if !ValidateExpirationDate(product.Expiration) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Expiration date is incorrect",
			})
			return
		}
		p, err := h.s.Update(id, product)
		if err != nil {
			ctx.JSON(http.StatusNotAcceptable, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": p,
		})
		return
	}
}

func (h *productHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Quantity    int     `json:"quantity,omitempty"`
		CodeValue   string  `json:"code_value,omitempty"`
		IsPublished bool    `json:"is_published,omitempty"`
		Expiration  string  `json:"expiration,omitempty"`
		Price       float64 `json:"price,omitempty"`
	}
	return func(ctx *gin.Context) {
		// token := ctx.GetHeader("token")
		// if token != os.Getenv("TOKEN") {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "Invalid token",
		// 	})
		// 	return
		// }

		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid id",
			})
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Verify body information",
			})
			return
		}

		update := domain.Product{
			Name:         r.Name,
			Quantity:     r.Quantity,
			Code_value:   r.CodeValue,
			Is_published: r.IsPublished,
			Expiration:   r.Expiration,
			Price:        r.Price,
		}

		if update.Expiration != "" {
			if !ValidateExpirationDate(update.Expiration) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Expiration date is incorrect",
				})
				return
			}
		}

		p, err := h.s.Update(id, update)
		if err != nil {
			ctx.JSON(http.StatusNotAcceptable, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": p,
		})
		return
	}
}
