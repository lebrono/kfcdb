package handlers

import (
	"net/http"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "test/sample/api/models"
)

type MenuHandler struct {
	db *gorm.DB
}

// NewMenu factory for Menu Controller
func NewMenuHandler(db *gorm.DB) *MenuHandler {
	return &MenuHandler{db}
}

// fetch list of menus
func (handler MenuHandler) Index(c *gin.Context) {
	//create array instance of Menu model
	menus := []m.Menu{}	
	handler.db.Table("tbl_menu").Find(&menus)
	//send response with http status 200 and the result of the query
	c.JSON(http.StatusOK, &menus)
}

// create new Menu
func (handler MenuHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	menu_name := c.PostForm("menu_name")
	menu_description := c.PostForm("menu_description")
	menu_amount := c.PostForm("menu_amount")
	category_name := c.PostForm("category_name")
	handler.db.Exec("INSERT INTO tbl_menu VALUES (null,?,?,?,?,0,0,?,?)",menu_name,menu_description,menu_amount,category_name,now,now)
	respond(http.StatusCreated,"New Menu created!",c,false)	
}

// update menu
func (handler MenuHandler) Update(c *gin.Context) {
	menu_id := c.Param("menu_id")
 	menu := m.Menu{}
	handler.db.Table("tbl_menu").Where("menu_id = ?",menu_id).First(&menu)

	if menu.Menu_Name != "" {
		now := time.Now().UTC()
		menu_name := c.PostForm("menu_name")

		menu := m.Menu{}
		handler.db.Table("tbl_menu").Where("menu_id <> ? AND menu_name = ?",menu_id,menu_name).First(&menu)

		fmt.Println("menu name ---> ", menu.Menu_Name)

		if menu.Menu_Name != "" {
			respond(http.StatusBadRequest,"Menu name already taken",c,true)
		} else {
			menu_description := c.PostForm("menu_description")
			menu_amount := c.PostForm("menu_amount")
			category_name := c.PostForm("category_name")
			is_sync := c.PostForm("is_sync")
			is_active := c.PostForm("is_active")
			handler.db.Exec("UPDATE tbl_menu SET menu_name = ?, menu_description = ? , menu_amount = ?, category_name = ?, isSync = ? , isActive = ?, date_updated = ? WHERE menu_id = ?",menu_name,menu_description,menu_amount,category_name,is_sync,is_active,now,menu_id)
			respond(http.StatusOK,"Menu successfully updated",c,false)	
		}
	} else {
		respond(http.StatusBadRequest,"Branch not found",c,true)	
	}
}



