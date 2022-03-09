package router

import (
	"github.com/blyndusk/flamingops/internal/controllers"
	"github.com/blyndusk/flamingops/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	usersRoute(r)
	secretsManagerRoute(r)
	r.GET("/load_fixtures", controllers.LoadData)

}

func usersRoute(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	// r.GET("/users", middlewares.JWTVerify(controllers.GetAllUsers))
	r.GET("/users/:id", middlewares.JWTVerify(controllers.GetUserById))
	r.PUT("/users/:id", middlewares.JWTVerify(controllers.UpdateUser))
	r.DELETE("/users/:id", middlewares.JWTVerify(controllers.DeleteUser))

	r.POST("/users/:id/display-preferences", controllers.CreateDisplayPreferences)
	r.GET("/users/:id/display-preferences", controllers.GetDisplayPreferences)
	r.PUT("/users/:id/display-preferences", controllers.UpdateDisplayPreferences)
	r.DELETE("/users/:id/display-preferences", controllers.DeleteDisplayPreferences)

	r.POST("/users/:id/active-services", controllers.CreateActiveServices)
	r.GET("/users/:id/active-services", controllers.GetActiveServices)
	r.PUT("/users/:id/active-services", controllers.UpdateActiveServices)
	r.DELETE("/users/:id/active-services", controllers.DeleteActiveServices)

	r.POST("/users/:id/sw-services-data", controllers.CreateSwServicesData)
	r.GET("/users/:id/sw-services-data", controllers.GetSwServicesData)
	r.DELETE("/users/:id/sw-services-data", controllers.DeleteSwServicesData)

	r.POST("/users/:id/aws-services-data", controllers.CreateAwsServicesData)
	r.GET("/users/:id/aws-services-data", controllers.GetAwsServicesData)
	r.DELETE("/users/:id/aws-services-data", controllers.DeleteAwsServicesData)
}

func secretsManagerRoute(r *gin.Engine) {
	r.POST("/secrets", controllers.CreateSecret)

	r.GET("/secrets/:secretName", controllers.GetSecretByName)

	r.PUT("/secrets/:secretName", controllers.UpdateSecret)

	r.DELETE("/secrets/:secretName", controllers.DeleteSecret)
}
