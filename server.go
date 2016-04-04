package main

import (
	"os"
	"fmt"
	"log"

	 "github.com/gin-gonic/gin"
	 _ "github.com/go-sql-driver/mysql"
	 h "test/sample/api/handlers"
	 "test/sample/api/config"
	"github.com/jinzhu/gorm"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *gorm.DB) {
	public := r.Group("/api/v1")

	//manage category
	categoryHandler := h.NewCategoryHandler(db)
	public.GET("/categories", categoryHandler.Index)
	public.POST("/categories", categoryHandler.Create)

	//manage branches
	branchHandler := h.NewBranchHandler(db)
	public.GET("/branch", branchHandler.Index)
	public.GET("/branch/:branch_code", branchHandler.Login)
	public.POST("/branch", branchHandler.Create)

	//manage menus
	menuHandler := h.NewMenuHandler(db)
	public.GET("/menus", menuHandler.Index)
	public.POST("/menus", menuHandler.Create)
	public.PUT("/menus/:menu_id", menuHandler.Update)

	//manage fun cards
	cardHandler := h.NewFunCardHandler(db)
	public.GET("/fun_cards", cardHandler.Index)
	public.POST("/fun_cards", cardHandler.Create)
	public.PUT("/fun_cards/:card_id", cardHandler.Update)
	public.GET("/fun_cards/:card_id", cardHandler.Show)

	//manage transactions
	transactionHandler := h.NewTransactionHandler(db)
	public.GET("/transactions", transactionHandler.Index)
	public.POST("/transactions", transactionHandler.Create)
	public.GET("/transactions/:transaction_no", transactionHandler.Show)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("PORT ---> ",port)
	r.Run(fmt.Sprintf(":%s", port))
}

func InitDB() *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("DB_USER"), config.GetString("DB_PASS"),
		config.GetString("DB_HOST"), config.GetString("DB_PORT"),
		config.GetString("DB_NAME"))
	log.Printf("\nDatabase URL: %s\n", dbURL)

	_db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	_db.DB()
	_db.LogMode(true)
	//_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.Category{})
	_db.Set("gorm:table_options", "ENGINE=InnoDB")
	return _db
}

func GetPort() string {
    var port = os.Getenv("PORT")
    // Set a default port if there is nothing in the environment
    if port == "" {
        port = "8000"
        fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
    }
    fmt.Println("port -----> ", port)
    return ":" + port
}