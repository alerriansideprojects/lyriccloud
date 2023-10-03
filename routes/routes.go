package routes

import (
	handlers "go_word_cloud/handlers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Application
	router.GET("/", handlers.AppIndex)
	router.GET("/artists", handlers.ArtistsIndex)

	// API Index
	router.GET("/api/v0", handlers.APIIndex)

	// Artists Related Routes

	// User related routes
	// router.POST("/api/v0/users", handlers.CreateUser)
	// router.GET("/api/v0/users/:id", handlers.ShowUser)

	// // Playlist related routes
	// router.POST("/api/v0/playlists", handlers.CreatePlaylist)
	// router.GET("/api/v0/user/:user_id/playlists", handlers.UserPlaylists)
	// router.GET("/api/v0/playlists/:id", handlers.ShowPlaylist)
	// router.PATCH("/api/v0/user/:user_id/playlists/:id", handlers.UpdatePlaylist)
	// router.GET("/api/v0/playlist_status", handlers.PlaylistsByStatus)
	// router.DELETE("/api/v0/playlists/:id", handlers.DeletePlaylist)

	// // PlaylistItem routes
	// router.POST("/api/v0/items", handlers.CreatePlaylistItem)
	// router.GET("/api/v0/items/:id", handlers.ShowPlaylistItem)
	// router.GET("/api/v0/playlist/:playlist_id/items", handlers.PlaylistItems)
	// router.PATCH("/api/v0/playlists/:playlist_id/items/:id", handlers.UpdateItem)

	return router
}
