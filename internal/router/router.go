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
	r.POST("/user", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	// r.GET("/users", controllers.GetAllUsers)

	r.GET("/user", middlewares.JWTVerify(controllers.GetUserById))
	r.PUT("/user", middlewares.JWTVerify(controllers.UpdateUser))
	r.DELETE("/user", middlewares.JWTVerify(controllers.DeleteUser))

	r.GET("/user/requested-regions", middlewares.JWTVerify(controllers.GetRequestedRegions))
	r.PUT("/user/requested-regions", middlewares.JWTVerify(controllers.UpdateRequestedRegions))
	r.DELETE("/user/requested-regions", middlewares.JWTVerify(controllers.DeleteRequestedRegions))

	r.GET("/user/active-services", middlewares.JWTVerify(controllers.GetActiveServices))
	r.PUT("/user/active-services", middlewares.JWTVerify(controllers.UpdateActiveServices))
	r.DELETE("/user/active-services", middlewares.JWTVerify(controllers.DeleteActiveServices))

	r.GET("/user/sw-services-data", middlewares.JWTVerify(controllers.GetSwServicesData))
	r.DELETE("/user/sw-services-data", middlewares.JWTVerify(controllers.DeleteSwServicesData))

	r.GET("/user/aws-services-data", middlewares.JWTVerify(controllers.GetAwsServicesData))
	r.DELETE("/user/aws-services-data", middlewares.JWTVerify(controllers.DeleteAwsServicesData))

	r.POST("/user/send-message", middlewares.JWTVerify(controllers.ManuallySendMessage))
}

func secretsManagerRoute(r *gin.Engine) {
	r.POST("/secrets", middlewares.JWTVerify(controllers.CreateSecret))

	r.PUT("/secrets/:secretName", middlewares.JWTVerify(controllers.UpdateSecret))

	r.DELETE("/secrets/:secretName", middlewares.JWTVerify(controllers.DeleteSecret))
}
