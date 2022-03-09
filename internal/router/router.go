package router

import (
	"github.com/blyndusk/flamingops/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	usersRoute(r)
	r.GET("/load_fixtures", controllers.LoadData)

}

func usersRoute(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	// r.GET("/users", middlewares.JWTVerify(controllers.GetAllUsers))
	r.GET("/users/:id", middlewares.JWTVerify(controllers.GetUserById))

	r.PUT("/users/:id", middlewares.JWTVerify((controllers.UpdateUser))

	r.DELETE("/users/:id", middlewares.JWTVerify((controllers.DeleteUser))
}
