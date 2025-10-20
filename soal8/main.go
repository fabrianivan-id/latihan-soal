package main

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	SKU             string  `json:"sku"`
	ProductName     string  `json:"productName"`
	QuantityInStock int     `json:"quantityInStock"`
	Price           float64 `json:"price"`
	Category        string  `json:"category"`
}

var (
	skuRe           = regexp.MustCompile(`^SKU-\d{8}$`)
	allowedCategory = map[string]struct{}{
		"Electronics": {},
		"Books":       {},
		"Apparel":     {},
		"Home Goods":  {},
	}
)

// validateProduct returns messages IN FIELD ORDER.
func validateProduct(p Product) []string {
	msgs := []string{}

	// sku
	sku := strings.TrimSpace(p.SKU)
	if sku == "" {
		msgs = append(msgs, "The sku is a mandatory field")
	} else if !skuRe.MatchString(sku) {
		msgs = append(msgs, "The sku must be in the format SKU-XXXXXXXX")
	}

	// productName
	if strings.TrimSpace(p.ProductName) == "" {
		msgs = append(msgs, "The productName is a mandatory field")
	}

	// quantityInStock
	if p.QuantityInStock < 0 {
		msgs = append(msgs, "The quantityInStock cannot be negative")
	}

	// price
	if p.Price <= 0 {
		msgs = append(msgs, "The price must be greater than zero")
	}

	// category
	cat := strings.TrimSpace(p.Category)
	if cat == "" {
		msgs = append(msgs, "The category is a mandatory field")
	} else {
		if _, ok := allowedCategory[cat]; !ok {
			msgs = append(msgs, "Invalid product category")
		}
	}

	return msgs
}

func main() {
	r := gin.Default()

	r.POST("/products", func(c *gin.Context) {
		var p Product
		if err := c.ShouldBindJSON(&p); err != nil {
			// Invalid JSON
			c.JSON(http.StatusBadRequest, []string{"Invalid JSON payload"})
			return
		}
		if errs := validateProduct(p); len(errs) > 0 {
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		c.Status(http.StatusOK)
	})

	_ = r.Run(":8080")
}
