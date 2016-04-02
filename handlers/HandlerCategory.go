package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "test/sample/api/models"
)

type CategoryHandler struct {
	db *gorm.DB
}

// NewCategory factory for Category Controller
func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{db}
}

// Index retrieves a list of categories
func (handler CategoryHandler) Index(c *gin.Context) {
	//create array instance of Category model
	categories := []m.Category{}	
	handler.db.Table("tbl_category").Find(&categories)
	//send response with http status 200 and the result of the query
	c.JSON(http.StatusOK, &categories)
}

// Index retrieves a list of categories
func (handler CategoryHandler) Create(c *gin.Context) {
	//create array instance of Category model
	categories := []m.Category{}	
	handler.db.Table("tbl_category").Find(&categories)
	//send response with http status 200 and the result of the query
	c.JSON(http.StatusOK, &categories)
}




