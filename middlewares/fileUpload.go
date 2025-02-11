package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileUploadMiddleware(formKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _, err := c.Request.FormFile(formKey)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
			c.Abort()
			return
		}
		defer file.Close()

		randomFileName := uuid.New().String()

		c.Set(formKey+"_fileName", randomFileName)
		c.Set(formKey+"_file", file)

		c.Next()
	}
}
