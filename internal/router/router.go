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
	// r.GET("/users", controllers.GetAllUsers)

	r.GET("/users/:id", middlewares.JWTVerify(controllers.GetUserById))
	r.PUT("/users/:id", middlewares.JWTVerify(controllers.UpdateUser))
	r.DELETE("/users/:id", middlewares.JWTVerify(controllers.DeleteUser))

	r.GET("/users/:id/requested-regions", middlewares.JWTVerify(controllers.GetRequestedRegions))
	r.PUT("/users/:id/requested-regions", middlewares.JWTVerify(controllers.UpdateRequestedRegions))
	r.DELETE("/users/:id/requested-regions", middlewares.JWTVerify(controllers.DeleteRequestedRegions))

	r.GET("/users/:id/active-services",  middlewares.JWTVerify(controllers.GetActiveServices))
	r.PUT("/users/:id/active-services",  middlewares.JWTVerify(controllers.UpdateActiveServices))
	r.DELETE("/users/:id/active-services",  middlewares.JWTVerify(controllers.DeleteActiveServices))

	r.GET("/users/:id/sw-services-data",  middlewares.JWTVerify(controllers.GetSwServicesData))
	r.DELETE("/users/:id/sw-services-data",  middlewares.JWTVerify(controllers.DeleteSwServicesData))

	r.GET("/users/:id/aws-services-data", middlewares.JWTVerify(controllers.GetAwsServicesData))
	r.DELETE("/users/:id/aws-services-data",  middlewares.JWTVerify(controllers.DeleteAwsServicesData))

	r.POST("/users/:id/send-message", middlewares.JWTVerify(controllers.ManuallySendMessage))
}

func secretsManagerRoute(r *gin.Engine) {
	r.POST("/secrets/:id", middlewares.JWTVerify(controllers.CreateSecret))

	r.PUT("/secrets/:secretName/:id", middlewares.JWTVerify(controllers.UpdateSecret))

	r.DELETE("/secrets/:secretName/:id", middlewares.JWTVerify(controllers.DeleteSecret))
}
