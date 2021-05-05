package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/saaitt/marketingManagerPod/handler"
	"github.com/saaitt/marketingManagerPod/model"
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
			Repo: &model.SQLItemRepo{
				DB: db,
			},
		},
	}
	marketing := handler.MarketingHandler{
		Service: service.MarketingService{
			MarketingRepo: &model.SQLMarketingRepo{
				DB: db,
			},
			ProductRepo: &model.SQLItemRepo{
				DB: db,
			},
		},
	}
	e.GET("/", product.ListAll)
	e.POST("/", product.Create)
	e.GET("/marketing/:id", marketing.ListAllMarketingProducts)
	e.POST("/marketing/:user_id",marketing.CreateProduct)
	e.GET("/:marketing_product", marketing.Redirect)
	e.Logger.Fatal(e.Start(":9876"))
}
