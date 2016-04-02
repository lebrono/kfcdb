package handlers

import (
	"net/http"
	"strconv"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	m "test/sample/api/models"
	"gopkg.in/mgo.v2"	
	"gopkg.in/mgo.v2/bson"	
)

type CategoryHandler struct {
	sess *mgo.Session
}

// NewCategory factory for CategoryController
func NewCategoryHandler(sess *mgo.Session) *CategoryHandler {
	return &CategoryHandler{sess}
}

//fetch list of categories
func (handler CategoryHandler) Index(c *gin.Context) {
	start := -1
	max := 10

	//check if start exists in url parameters
	if c.Query("start") != ""  {
		i,_ := strconv.Atoi(c.Query("start"))
		start = i;
	} else {
		fmt.Println("cant read start query param")
	}

	if c.Query("max") != ""  {
		i,_ := strconv.Atoi(c.Query("max"))
		max = i;
	} 

	fmt.Printf("offset ---> %d max ---> %d\n", start, max)
	categories := []m.Category{}
	collection := handler.sess.DB("kfcdb").C("category") 
	collection.Find(nil).Sort("-createdat").All(&categories)
	c.JSON(http.StatusOK, categories)
}

// Create new category
func (handler CategoryHandler) Create(c *gin.Context) {
	category := m.Category{}
	c.Bind(&category)
	collection := handler.sess.DB("kfcdb").C("category") 
	result := m.Category{}
	err := collection.Find(bson.M{"categoryname": category.CategoryName}).One(&result)
	//check if barangay name is not existing
	if fmt.Sprintf("%s", err) == "not found" {
		// generate hashed password
		category.Id = bson.NewObjectId()
		category.CreatedAt = time.Now().UTC()
		category.UpdatedAt = time.Now().UTC()
		category.Status = "active"
		collection.Insert(&category)
		c.JSON(http.StatusCreated,category)
	} else {
		respond(http.StatusBadRequest,"Category name was already taken",c,true)
	}
}

// Update category
func (handler CategoryHandler) Update(c *gin.Context) {
	id := c.Param("category_id")
	category := m.Category{}
	c.Bind(&category)
	collection := handler.sess.DB("kfcdb").C("category") 
	result := m.Category{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	//check if category exists
	if fmt.Sprintf("%s", err) == "not found" {
		respond(http.StatusBadRequest,"No matching record found with supplied id",c,true)
	} else {
		//check if category name exists
		otherCategory := m.Category{}
		err := collection.Find(bson.M{"$and": []bson.M{bson.M{"categoryname": category.CategoryName}, 
							bson.M{"_id" : bson.M{"$ne" : bson.ObjectIdHex(id)}}}}).One(&otherCategory)
		if fmt.Sprintf("%s", err) == "not found" {
			change := mgo.Change {
				Update: bson.M{"$set": bson.M{"categoryname": category.CategoryName}},
				ReturnNew: true,
			}
			updatedCategory := m.Category{}
			collection.FindId(bson.ObjectIdHex(id)).Apply(change, &updatedCategory) // Apply
			c.JSON(http.StatusOK,updatedCategory)
		} else {
			respond(http.StatusBadRequest,"Category name was already taken",c,true)
		}
	}
}



