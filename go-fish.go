package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func fish(c *gin.Context) {
	var resp struct {
		Message string `json:"message"`
		Time    string `json:"time"`
		User    struct {
			Email string `json:"email"`
			Phone string `json:"phone"`
		} `json:"user"`
	}
	resp.Message = "Hello " + c.Param("name") + "!"
	resp.Time = time.Now().Format(time.RFC3339)
	resp.User.Email = "test@example.com"
	resp.User.Phone = "555 555-5555"
	c.JSON(http.StatusOK, resp)
}

func index(c *gin.Context) {
	var resp struct {
		Message string `json:"message"`
		Time    string `json:"time"`
	}
	resp.Message = "Today's fish is trout al a creme. Enjoy your meal!"
	resp.Time = time.Now().Format(time.RFC3339)

	c.JSON(http.StatusOK, resp)
}

func main() {
	r := gin.Default()
	r.GET("/fish/:name", fish)
	r.GET("/", index)
	r.Run()
}
