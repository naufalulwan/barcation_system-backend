package main

import (
	"barcation_be/config"
	"barcation_be/controllers"
	"barcation_be/controllers/cart"
	"barcation_be/controllers/category"
	"barcation_be/controllers/information"
	"barcation_be/controllers/inquiry"
	"barcation_be/controllers/payment"
	"barcation_be/controllers/product"
	"barcation_be/controllers/public"
	"barcation_be/controllers/user"
	"barcation_be/helper"
	"barcation_be/middlewares"
	"barcation_be/models"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Sdk struct {
	Port   string
	Domain string
}

var r *gin.Engine

func main() {
	helper.EnviromentHelper()
	helper.LoggerHelper()

	config.LoadConfig()
	config.SetupDatabase()

	err := config.DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Product{}, &models.Cart{}, &models.Inquiry{}, &models.Payment{}, &models.Information{})
	if err != nil {
		return
	}

	gin.SetMode(gin.ReleaseMode)
	r = gin.New()
	r.Use(gin.Logger())

	r.ForwardedByClientIP = true
	err = r.SetTrustedProxies([]string{"0.0.0.0"})
	if err != nil {
		return
	}

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	setRoute()

	sdk := Sdk{
		Port:   os.Getenv("PORT"),
		Domain: os.Getenv("DOMAIN"),
	}

	dns := fmt.Sprintf("%s:%s", sdk.Domain, sdk.Port)

	helper.Logger.Infof("\x1b[34mListening and serving HTTP on %s\x1b[0m", dns)

	err = r.Run(dns)
	if err != nil {
		return
	}
}

func setRoute() {
	r.GET("temp/product_image/:filename", helper.ServeImageHelper)

	// API FOR PUBLIC ["/api/public"]
	publicParam := r.Group("/api/public")

	publicParam.PUT("/get_info_login", public.GetInfoLoginController)
	publicParam.GET("/get_token_device", public.GetTokenDeviceController)

	// API FOR AUTH ["/api/auth"]
	auth := r.Group("/api/auth")

	auth.POST("/register", controllers.RegisterController)
	auth.POST("/login", controllers.LoginController)

	// API FOR USER ["/api/user"]
	protectedUser := r.Group("/api/user")
	protectedUser.Use(middlewares.AuthTokenMiddleware())

	protectedUser.GET("/get_user", user.GetUserController)
	protectedUser.GET("/get_user_by_id", user.GetUserByIdController)
	protectedUser.PUT("/update_user", user.UpdateUserController)
	protectedUser.PUT("/update_password_user", user.UpdatePasswordUserController)
	protectedUser.PUT("/update_level_user", user.UpdateLevelUserController)
	protectedUser.DELETE("/delete_user", user.DeleteUserController)
	protectedUser.PUT("/recovery_user", user.RecoveryUserController)
	protectedUser.PUT("/forgot_user", user.ForgotPasswordUserController)
	protectedUser.GET("/logout", controllers.LogoutController)

	// API FOR CATEGORY ["/api/category"]
	protectedCategory := r.Group("/api/category")
	protectedCategory.Use(middlewares.AuthTokenMiddleware())

	protectedCategory.GET("/get_category", category.GetCategoryController)
	protectedCategory.POST("/create_category", category.CreateCategoryController)
	protectedCategory.PUT("/update_category", category.UpdateCategoryController)
	protectedCategory.DELETE("/delete_category", category.DeleteCategoryController)

	// API FOR PRODUCT ["/api/product"]
	protectedProduct := r.Group("/api/product")
	protectedProduct.Use(middlewares.AuthTokenMiddleware())

	protectedProduct.GET("/get_product", product.GetProductController)
	protectedProduct.POST("/create_product", product.CreateProductController)
	protectedProduct.PUT("/update_product", product.UpdateProductController)
	protectedProduct.DELETE("/delete_product", product.DeleteProductController)
	protectedProduct.GET("/get_product_by_id", product.GetProductByIdController)

	// API FOR CART ["/api/cart"]
	protectedCart := r.Group("/api/cart")
	protectedCart.Use(middlewares.AuthTokenMiddleware())

	protectedCart.POST("/create_cart", cart.CreateCartController)
	protectedCart.GET("/get_cart", cart.GetCartController)
	protectedCart.DELETE("/delete_cart", cart.DeleteCartController)

	// API FOR INQUIRY ["/api/inquiry"]
	protectedInquiry := r.Group("/api/inquiry")
	protectedInquiry.Use(middlewares.AuthTokenMiddleware())

	protectedInquiry.GET("/get_inquiry", inquiry.GetInquiryController)
	protectedInquiry.GET("/get_inquiry_by_id", inquiry.GetInquiryByIdController)
	protectedInquiry.POST("/create_inquiry", inquiry.CreateInquiryController)
	protectedInquiry.DELETE("/delete_inquiry", inquiry.DeleteInquiryController)
	protectedInquiry.PUT("/update_inquiry", inquiry.UpdateInquiryController)

	// API FOR PAYMENT ["/api/payment"]
	protectedPayment := r.Group("/api/payment")
	protectedPayment.Use(middlewares.AuthTokenMiddleware())

	protectedPayment.GET("/get_payment", payment.GetPaymentController)
	protectedPayment.GET("/get_payment_by_id", payment.GetPaymentByIdController)
	protectedPayment.POST("/create_payment", payment.CreatePaymentController)
	protectedPayment.DELETE("/delete_payment", payment.DeletePaymentController)
	protectedPayment.PUT("/update_payment", payment.UpdatePaymentController)

	protectedInformation := r.Group("/api/information")
	protectedInformation.Use(middlewares.AuthTokenMiddleware())
	protectedInformation.POST("/create_information", information.CreateInformationController)
	r.GET("/api/information/get_information", information.GetInformationController)

}
