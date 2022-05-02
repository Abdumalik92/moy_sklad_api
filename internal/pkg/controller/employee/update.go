package employee

import (
	"github.com/Abdumalik92/moy_sklad_api/internal/models"
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/service/employee"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UpdateEmployee(c *gin.Context) {
	var (
		request  models.CreateEmployee
		response models.Row
		err      error
	)
	err = c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("UpdateEmployee bind json err ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Something went wrong"})
		return
	}
	request.ID = c.Param("id")
	if err = employee.UpdateEmployee(request, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, response)
}
