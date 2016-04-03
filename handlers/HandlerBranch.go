package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "test/sample/api/models"
)

type BranchHandler struct {
	db *gorm.DB
}

// NewBranch factory for Branch Controller
func NewBranchHandler(db *gorm.DB) *BranchHandler {
	return &BranchHandler{db}
}

// Index retrieves a list of branch
func (handler BranchHandler) Index(c *gin.Context) {
	//create array instance of Branch model
	branches := []m.Branch{}	
	handler.db.Table("tbl_branch").Find(&branches)
	//send response with http status 200 and the result of the query
	c.JSON(http.StatusOK, &branches)
}

// Create new branch
func (handler BranchHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	branch_code := c.PostForm("branch_code")
	branch_name := c.PostForm("branch_name")
	branch_address := c.PostForm("branch_address")
	branch_email := c.PostForm("branch_email")
	branch_contact_no := c.PostForm("branch_contact_no")
	latitude := c.PostForm("latitude")
	longitude := c.PostForm("longitude")

 	branch := m.Branch{}
	handler.db.Table("tbl_branch").Where("branch_code = ?",branch_code).First(&branch)

	if branch.Branch_Code != "" {
		respond(http.StatusBadRequest,"branch code was already in use",c,true)
	} else {
		handler.db.Exec("INSERT INTO tbl_branch VALUES (null,?,?,?,?,?,?,?,?,?)",branch_code,branch_name,branch_address,branch_email,branch_contact_no,latitude,longitude,now,now)
		respond(http.StatusCreated,"New branch created!",c,false)	
	}
}

func (handler BranchHandler) Login(c *gin.Context) {
	branch_code := c.Param("branch_code")
 	branch := m.Branch{}
	handler.db.Table("tbl_branch").Where("branch_code = ?",branch_code).First(&branch)
	if branch.Branch_Code != "" {
		c.JSON(http.StatusOK, &branch)
	} else {
		respond(http.StatusBadRequest,"Branch not found",c,true)	
	}
}




