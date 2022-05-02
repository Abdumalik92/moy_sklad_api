package employee

import (
	"github.com/Abdumalik92/moy_sklad_api/internal/models"
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/service/employee"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEmployeeList(c *gin.Context) {
	var (
		request  models.Request
		response models.Response
		err      error
	)
	request.Limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Limit param is not digit"})
	}
	request.Offset, err = strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "Limit param is not digit"})
	}
	if err := employee.GetEmployeeList(request, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, response)
}
