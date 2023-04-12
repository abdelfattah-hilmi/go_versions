package main

import (
	r "example/go_versions/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// defining endpoints
	router.Use(cors.Default())

	router.POST("/vms/installed-packages", r.AddVm)
	router.GET("/vms/installed-packages", r.GetVms)
	router.GET("/vms/:id/installed-packages", r.GetVmByID)
	router.GET("/vm/:ip", r.GetVmByIP)
	router.GET("/packages", r.GetPackages)
	router.DELETE("/vms/:id/installed-packages", r.DeleteVm)

	router.Run("0.0.0.0:8000")
}
