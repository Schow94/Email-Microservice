package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type EmailResponseInput struct {
	SenderEmail  string `json:"senderEmail" binding:"required"`
	Password     string `json:"password" binding:"required"`
	SenderName   string `json:"senderName" binding:"required"`
	Subject      string `json:"subject" "binding:"required"`
	EmailBody    string `json:"body" "binding:"required"`
	HtmlTemplate string `json:"htmlTemplate" binding:"required"`
	Recipients   []struct {
		Name  string `json:"name" "binding:"required"`
		Email string `json:"email" binding:"required"`
	} `json:"recipients" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/email", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "/GET request to /email route made. Please make a /POST request to send email(s)",
		})
	})

	r.POST("/email", func(c *gin.Context) {
		var data EmailResponseInput

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// User can provide:
		// - email/pw
		// - email provider (I can choose the correct prnewPaovider in backend (Gmail, Yahoo, Hotmail))
		// - Template user wants to use
		// - Msg/Data to be injected into template
		// - Potentially allow user to send an html template as a string
		// - Need to validate template/data provided in backend & send error back to user if
		// 	 template/data do not work together

		msg, err := Email(data)

		// Error sending email
		if err != nil {
			fmt.Println("ERROR: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Successfully sent email
		c.JSON(200, gin.H{
			"message": msg,
		})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(fmt.Sprintf(":%s", port))
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
