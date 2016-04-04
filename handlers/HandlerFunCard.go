package handlers

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "test/sample/api/models"
)

type FunCardHandler struct {
	db *gorm.DB
}

// NewFunCard factory for FunCard Controller
func NewFunCardHandler(db *gorm.DB) *FunCardHandler {
	return &FunCardHandler{db}
}

// fetch list of fun card holder
func (handler FunCardHandler) Index(c *gin.Context) {
	cards := []m.FunCard{}	
	handler.db.Table("tbl_funpoints").Find(&cards)
	c.JSON(http.StatusOK, &cards)
}

// create new fun card holder
func (handler FunCardHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	fun_id_number := c.PostForm("card_id_number")
	card_holder_name := c.PostForm("card_holder_name")
	fun_points := c.PostForm("fun_points")
	
	card := m.FunCard{}
	handler.db.Table("tbl_funpoints").Where("fun_idnumber = ?",fun_id_number).First(&card)

	if card.Fun_Name != "" {
		respond(http.StatusBadRequest,"Fun card number already used",c,true)	
	} else {
		handler.db.Exec("INSERT INTO tbl_funpoints VALUES (null,?,?,?,?,?)",fun_id_number,card_holder_name,fun_points,now,now)
		respond(http.StatusCreated,"New fun card created!",c,false)	
	}
}

// update card
func (handler FunCardHandler) Update(c *gin.Context) {
	card_id := c.Param("card_id")
	card := m.FunCard{}
	handler.db.Table("tbl_funpoints").Where("fun_idnumber = ?",card_id).First(&card)

	if card.Fun_Name != "" {
		now := time.Now().UTC()
		fun_points := c.PostForm("fun_points")
		handler.db.Exec("UPDATE tbl_funpoints SET fun_points = ?, date_updated = ? WHERE fun_idnumber = ?",fun_points,now,card_id)
		respond(http.StatusOK,"Card successfully updated",c,false)	
	} else {
		respond(http.StatusBadRequest,"Card not found!",c,true)	
	}
}


func (handler FunCardHandler) Show(c *gin.Context) {
	card_id := c.Param("card_id")
	card := m.FunCard{}
	handler.db.Table("tbl_funpoints").Where("fun_idnumber = ?",card_id).First(&card)
	if card.Fun_Name != "" {
		c.JSON(http.StatusOK,&card)
	} else {
		respond(http.StatusBadRequest,"Card not found!",c,true)	
	}
}



