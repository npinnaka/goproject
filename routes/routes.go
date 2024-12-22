package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.DELETE("/events/:id", deleteEventById)
	server.PUT("/events/:id", updateEventById)
	server.POST("/event", createEvent)

	server.POST("/user", signupUser)
	server.POST("/user/login", login)

	server.GET("/users", getUsers)
	server.DELETE("/user/:id", deleteUserById)
}
