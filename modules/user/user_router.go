package user

import "github.com/alexkazantsev/go-templ-api/server"

func RegisterRouter(server *server.Server, userController UserController) {
	var group = server.V1.Group("/users")

	group.POST("/", userController.Create)
	group.PUT("/:id", userController.UpdateOne)
	group.GET("/", userController.FindMany)
	group.GET("/:id", userController.FindOne)
	group.DELETE("/:id", userController.DeleteOne)
}
