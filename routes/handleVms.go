package routes

import (
	"context"
	"example/go_versions/helpers"
	"example/go_versions/models"

	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var vmCollection *mongo.Collection = OpenCollection(Client, "vms")

// Insert VM into collectio
func AddVm(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var vm models.Vm

	if err := c.BindJSON(&vm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		defer cancel()
		return
	}

	validationErr := validate.Struct(vm)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		defer cancel()
		return
	}

	vm.ID = primitive.NewObjectID()

	// collection
	for i := 0; i < len(vm.InstalledPackages); i++ {

		pkg_name := vm.InstalledPackages[i].PackageName
		vm.InstalledPackages[i].LatestReleaseNotes = helpers.GetReleaseNotes(pkg_name)
		repo := helpers.GetRepologyData(pkg_name)
		vm.InstalledPackages[i].LatestVersion = repo.LatestVersion
		vm.InstalledPackages[i].Cves = repo.Cves

	}

	// ---------------
	result, insertErr := vmCollection.InsertOne(ctx, vm)

	if insertErr != nil {
		msg := "order item was not created"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		defer cancel()
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, result)

}

// get all vms
func GetVms(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var vms []bson.M

	cursor, err := vmCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		defer cancel()
		return
	}

	if err = cursor.All(ctx, &vms); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		defer cancel()
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, vms)
}

func GetVmByID(c *gin.Context) {
	vmID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(vmID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var vm bson.M

	if err := vmCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&vm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		defer cancel()
		return
	}

	defer cancel()
	fmt.Println(vm)

	c.JSON(http.StatusOK, vm)
}

func GetVmByIP(c *gin.Context) {
	vmIP := c.Params.ByName("ip")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var vms []bson.M

	cursor, err := vmCollection.Find(ctx, bson.M{"ip": vmIP})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		defer cancel()
		return
	}
	if err = cursor.All(ctx, &vms); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		defer cancel()
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, vms)
}

func GetPackages(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	// var vms []bson.M
	var vms []interface{}

	cursor, err := vmCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		defer cancel()
		return
	}

	if err = cursor.All(ctx, &vms); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		defer cancel()
		return
	}

	defer cancel()
	// get packages from vms
	// var pkg map[string]interface{}
}

func DeleteVm(c *gin.Context) {
	vmID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(vmID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := vmCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	defer cancel()

	c.JSON(http.StatusOK, result.DeletedCount)

}
