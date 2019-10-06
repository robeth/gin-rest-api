package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func HomePage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var products []Product
	db.Find(&products, 1)

	c.JSON(200, gin.H{
		"message": "Hello World",
		"data":    products,
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PostHomePage",
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("Hello World")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Code:  "L1212",
		Price: 1200,
	})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "L1212")

	app := gin.Default()
	app.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	app.GET("/", HomePage)
	app.POST("/", PostHomePage)
	app.GET("/query", QueryStrings)
	app.GET("/path/:name/:age", PathParameters)
	app.Run()
}
