package routes

import (
	"challenge-go/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte(apiKey))))
	grp1 := r.Group("/user-api")
	{

		grp1.POST("login", controller.Login)

		grp1.GET("/pictures", controller.GetPictures)
		//grp1.POST("user", controller.CreateUser)
		//grp1.GET("user/:id", controller.GetUserByID)
		//grp1.PUT("user/:id", controller.UpdateUser)
		//grp1.DELETE("user/:id", controller.DeleteUser)
		//grp1.POST("user", controller.CreateUser)

	}
	return r
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

//func AuthRequired(c *gin.Context) {
//	session := sessions.Default(c)
//	user := session.Get(apiKey)
//	if user == nil {
//		// Abort the request with the appropriate error code
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
//		return
//	}
//}
