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

	router.POST("/api/vms/installed-packages", r.AddVm)
	router.GET("/api/vms/installed-packages", r.GetVms)
	router.GET("/api/vms/:id/installed-packages", r.GetVmByID)
	router.GET("/api/vm/:ip", r.GetVmByIP)
	router.GET("/api/packages", r.GetPackages)
	router.DELETE("/api/vms/:id/installed-packages", r.DeleteVm)

	router.Run("0.0.0.0:8000")
}
