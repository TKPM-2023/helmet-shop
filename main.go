package main

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/component/uploadprovider"
	"TKPM-Go/middleware"
	localPb "TKPM-Go/pubsub/localpub"
	"TKPM-Go/route/admin"
	"TKPM-Go/route/client"
	"TKPM-Go/route/user"
	"TKPM-Go/subscriber"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_URI")
	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")
	secretKey := os.Getenv("SYSTEM_SECRET")


	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect DB failed", err)
	}
	log.Println("Connect DB success", db)

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	pubsub := localPb.NewPubSub()
	appContext := appctx.NewAppContext(db, s3Provider, secretKey, pubsub)
	db = db.Debug()

	if err := subscriber.NewEngine(appContext).Start(); err != nil {
		log.Fatalln()
	}

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-XSRF-TOKEN",
			"screenId",
			"apiOrder",
		},
		ExposeHeaders: []string{
			"Content-Disposition",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(cors.New(config))

	// route
	v1 := r.Group("v1")
	admin.AdminRoute(appContext, v1)
	client.ClientRoute(appContext, v1)
	user.UserRoute(appContext, v1)

	r.Run()

}
