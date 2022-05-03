package routes

import (
	"fmt"
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/controller/employee"
	"github.com/Abdumalik92/moy_sklad_api/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func RunAllRoutes() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	f, err := os.OpenFile(utils.AppSettings.AppParams.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("file create error", err.Error())
		return
	}

	logger := &lumberjack.Logger{
		Filename:   f.Name(),
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	}

	log.SetOutput(logger)
	gin.DefaultWriter = io.MultiWriter(logger, os.Stdout)
	r.GET("/employee", employee.GetEmployeeList)
	r.POST("/employee", employee.CreateEmployee)
	r.DELETE("/employee/:id", employee.DeleteEmployee)
	r.PUT("/employee/:id", employee.UpdateEmployee)
	err = r.Run(utils.AppSettings.AppParams.PortRun)
	if err != nil {
		log.Panic("Error on start server ", err.Error())
	}
}
