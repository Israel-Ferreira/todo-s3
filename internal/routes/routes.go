package routes

import (
	"github.com/Israel-Ferreira/todo-s3/pkg/controllers"
	"github.com/Israel-Ferreira/todo-s3/pkg/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateUserRouter(userController controllers.UserController) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json", "multipart/form-data"))

	router.Use(middleware.Heartbeat("/"))
	router.Use(middlewares.JsonMiddleware)

	router.Route("/users", func(r chi.Router) {

		r.Post("/", userController.Create)

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", userController.GetById)
			r.Delete("/", userController.DeleteById)
			r.Put("/", userController.Update)

			r.Put("/upload-photo", userController.UploadProfilePhoto)
		})
	})

	return router
}
