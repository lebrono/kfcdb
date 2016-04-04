package handlers

import (
	"net/http"
	"time"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "test/sample/api/models"
)

type TransactionHandler struct {
	db *gorm.DB
}

func NewTransactionHandler(db *gorm.DB) *TransactionHandler {
	return &TransactionHandler{db}
}

// Fetch all transactions
func (handler TransactionHandler) Index(c *gin.Context) {
	transactions := []m.Transaction{}	
	handler.db.Table("tbl_transaction").Find(&transactions)
	c.JSON(http.StatusOK, &transactions)
}

// Create new transaction
func (handler TransactionHandler) Create(c *gin.Context) {
	now := time.Now().UTC()
	tx_no := c.PostForm("transaction_no")
	order := c.PostForm("order")
	total_amount := c.PostForm("total_amount")
	grand_total := c.PostForm("grand_total")
	is_points := c.PostForm("is_points")
	points_used := c.PostForm("points_used")
	fun_id := c.PostForm("fun_card_id")
	is_ok_kitchen := c.PostForm("is_ok_kitchen")
	is_ok_pos := c.PostForm("is_ok_pos")
	order_type := c.PostForm("order_type")
	status := c.PostForm("status")
	pick_up_date := "2016-04-05 15:01:45"
	user_id := c.PostForm("user_id")
	user_name := c.PostForm("user_name")
	dc_type := c.PostForm("discount_type")
	dc_total := c.PostForm("discount_total")
	dc_percent := c.PostForm("discount_percent")
	branch_code := c.PostForm("branch_code")
	
	orders := []m.Order{}
	err := json.Unmarshal([]byte(order),&orders);
	if err == nil {
		fmt.Println("conversion successful")
	} else {
		fmt.Println("failed to convert", err)
	}

	for _,item := range orders {
		fmt.Println("item ---> ", item.Menu_Name)
		menu_name := item.Menu_Name
		quantity := string(item.Qty)
		amount := string(item.Menu_Amount)
		handler.db.Exec("INSERT INTO tbl_transaction VALUES (null,?,?,?,?,?,?,?,?,'',?,?,?,?,?,?,?,?,?,?,?,?,?,?)",tx_no,menu_name,quantity,amount,total_amount,grand_total,is_points,points_used,fun_id,is_ok_kitchen,is_ok_pos,order_type,status,pick_up_date,user_id,user_name,dc_type,dc_total,dc_percent,branch_code,now,now)
	}
	respond(http.StatusCreated,tx_no,c,false)	
}

// get transaction by transaction no
func (handler TransactionHandler) Show(c *gin.Context) {
	trans_no := c.Param("transaction_no")
 	transaction := m.Transaction{}
	handler.db.Table("tbl_transaction").Where("transaction_no = ?",trans_no).First(&transaction)
	if transaction.Transaction_No != "" {
		c.JSON(http.StatusOK, &transaction)
	} else {
		respond(http.StatusBadRequest,"Transaction not found",c,true)	
	}
}




