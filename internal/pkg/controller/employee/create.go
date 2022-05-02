package employee

import (
	"github.com/Abdumalik92/moy_sklad_api/internal/models"
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/service/employee"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateEmployee(c *gin.Context) {
	var (
		request  models.CreateEmployee
		response models.Row
		err      error
	)
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("CreateEmployee bind json err ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Something went wrong"})
		return
	}

	if err = employee.CreateEmployee(request, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, response)
}
