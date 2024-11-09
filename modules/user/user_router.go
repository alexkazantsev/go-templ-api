package user

import "github.com/alexkazantsev/templ-api/server"

func RegisterRouter(server *server.Server, userController UserController) {
	var group = server.V1.Group("/users")

	group.GET("/:id", userController.FindOne)
}
