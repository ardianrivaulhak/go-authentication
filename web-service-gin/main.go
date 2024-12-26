package main

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID string `json"id"`
	Title string `json"title"`
	Artist string `json"artist"`
	Price float64 `json"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album
    // Call BindJSON to bind the received JSON to

    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumID(c *gin.Context)  {
	id := c.Param("id")

	for _, v := range albums {
		// fmt.Println("Album ID:", v)
		if v.ID == id  {
			c.IndentedJSON(http.StatusOK, v)
		} 
	}
	
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumID(c *gin.Context)  {
	id := c.Param("id")
	
    var updatedAlbum album

	// bindJSON untuk mengikat album yang akan di update 
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	for i, v := range albums {
		if v.ID == id {
			// Update album details
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// func deleteAlbumID(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, v := range albums {
// 		if v.ID == id {
// 			// Remove album from the slice
// 			albums = append(albums[:i], albums[i+1:]...)
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }


func main()  {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbumID)

	router.Run("localhost:8000")
}


