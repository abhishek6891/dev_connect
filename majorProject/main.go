package main

import (
	"github.com/gin-gonic/gin"
	routes "majorProject/route"
	"majorProject/src/user/forgot"
	"majorProject/src/user/login"
	"majorProject/src/user/signup"
)

func main() {
	route := gin.Default()

	route.GET("/login", login.LoginRequestWithGet)
	route.GET("/signup", signup.SignUpRequestWithGet)
	route.GET("/forgotpassword", forgot.ForgotPasswordWithGet)
	route.GET("/resetpassword", forgot.ResetPasswordWithGet)

	routes.RegisterProjectRoutes(route)
	routes.RegisterClientRoutes(route)
	routes.RegisterDeveloperRoutes(route)
	route.Run() // listen and serve on 0.0.0.0:8080
}
