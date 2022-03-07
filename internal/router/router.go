package router

import (
	"github.com/blyndusk/flamingops/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	usersRoute(r)
	secretsManagerRoute(r)
	r.GET("/load_fixtures", controllers.LoadData)

}

func usersRoute(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserById)

	r.PUT("/users/:id", controllers.UpdateUser)

	r.DELETE("/users/:id", controllers.DeleteUser)
}

func secretsManagerRoute(r *gin.Engine) {
	r.POST("/secrets", controllers.CreateSecret)

	r.GET("/secrets/:secretName", controllers.GetSecretByName)
	
	r.PUT("/secrets/:secretName", controllers.UpdateSecret)

	r.DELETE("/secrets/:secretName", controllers.DeleteSecret)
}
