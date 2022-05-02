package employee

import (
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/service/employee"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")
	if err := employee.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reason": "Employee deleted"})
}
