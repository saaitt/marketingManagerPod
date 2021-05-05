package main

import (
	"fmt"
	"os"

	customMiddleware "github.com/saaitt/marketingManagerPod/middleware"

	"github.com/saaitt/marketingManagerPod/model"
	"github.com/saaitt/marketingManagerPod/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saaitt/marketingManagerPod/handler"
	"github.com/saaitt/marketingManagerPod/service"

	"github.com/labstack/echo"
)

func main() {
	pass := os.Getenv("DB_PASSWORD")
	url := fmt.Sprintf("host=localhost port=5432 user=mmp dbname=marketingmanagerpod password=%s sslmode=disable", pass)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	product := handler.ProductHandler{
		Service: service.ProductService{
			Repo: &sql.ProductRepo{
				DB: db,
			},
		},
	}
	marketing := handler.MarketingHandler{
		Service: service.MarketingService{
			MarketingRepo: &sql.MarketingRepo{
				DB: db,
			},
			ProductRepo: &sql.ProductRepo{
				DB: db,
			},
		},
	}
	userService := service.UserService{Repo: &sql.UserRepo{DB: db}}
	userHandler := handler.UserHandler{Service: userService}
	adminGroup := e.Group("/admin")
	adminGroup.GET("", product.FindByUser)
	adminGroup.POST("", product.Create)
	adminGroup.Use(customMiddleware.BasicAuthMiddlewareForUserType(userService, model.UserTypeAdmin))
	marketingGroup := e.Group("/marketing")
	marketingGroup.GET("/:id", marketing.ListAllMarketingProducts)
	marketingGroup.POST("/:user_id", marketing.CreateProduct)
	adminGroup.Use(customMiddleware.BasicAuthMiddlewareForUserType(userService, model.UserTypeMarketer))
	e.GET("/:marketing_product", marketing.Redirect)
	e.POST("/users/", userHandler.Create)
	e.Logger.Fatal(e.Start(":9876"))
}
