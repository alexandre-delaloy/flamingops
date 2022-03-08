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

	r.POST("/users/:user_id/display-preferences", controllers.CreateDisplayPreferences)
	
	r.GET("/users/:user_id/display-preferences", controllers.GetDisplayPreferences)
	
	r.PUT("/users/:user_id/display-preferences", controllers.UpdateDisplayPreferences)
	
	r.DELETE("/users/:user_id/display-preferences", controllers.DeleteDisplayPreferences)

	r.POST("/users/:user_id/active-services", controllers.CreateActiveServices)
	
	r.GET("/users/:user_id/active-services", controllers.GetActiveServices)
	
	r.PUT("/users/:user_id/active-services", controllers.UpdateActiveServices)
	
	r.DELETE("/users/:user_id/active-services", controllers.DeleteActiveServices)
}

func secretsManagerRoute(r *gin.Engine) {
	r.POST("/secrets", controllers.CreateSecret)

	r.GET("/secrets/:secretName", controllers.GetSecretByName)

	r.PUT("/secrets/:secretName", controllers.UpdateSecret)

	r.DELETE("/secrets/:secretName", controllers.DeleteSecret)
}
