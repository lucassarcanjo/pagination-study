package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/lucassarcanjo/pagination-study/api"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	srv := &api.Server{}

	dsn := os.Getenv("DB_DSN")

	srv.InitDb(dsn)
	srv.InitGin()

	srv.RegisterRoutes()

	srv.Start(":8080")
}

// func main() {
// 	router := gin.Default()
// 	router.GET("/users-offset", getUsersOffset)
// 	router.GET("/users-keyset", getUsersKeyset)

// 	router.Run("localhost:8080")
// }

// // Routes
// func getUsersOffset(c *gin.Context) {
// 	var params utils.OffsetParams

// 	if err := c.Bind(&params); err != nil {
// 		return
// 	}

// }

// func getUsersKeyset(c *gin.Context) {

// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Contrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	// Call BindJSON to bind the received JSON to newAlbum
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
