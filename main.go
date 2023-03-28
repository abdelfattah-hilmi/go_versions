package main

import (
	"context"
	"fmt"
	"log"

	m "example/go_versions/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI = "mongodb://localhost:27017"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	// call BindJSON to bind the new json to newAlbum
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
		fmt.Println(err)

	}
	defer client.Disconnect(ctx)

	demodb := client.Database("test")

	// err = demodb.CreateCollection(ctx, "packages")
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println(err)
	// }

	pkgCollection := demodb.Collection("packages")
	insertResult, err := pkgCollection.InsertOne(ctx, m.CollectedPackageInstance)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)

	var objects []m.CollectedPackage
	cursor, err := pkgCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &objects); err != nil {
		panic(err)
	}
	fmt.Println(objects)

	// ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
	// 	Connection: "local",
	// }

	// ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
	// 	Inventory: "./inventory,",
	// }

	// playbook := &playbook.AnsiblePlaybookCmd{
	// 	Playbooks:         []string{"site.yml"},
	// 	ConnectionOptions: ansiblePlaybookConnectionOptions,
	// 	Options:           ansiblePlaybookOptions,
	// }

	// err := playbook.Run(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }

	// router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)

	// router.Run("localhost:8080")
}
