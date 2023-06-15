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
	dsn := os.Getenv("MYSQL_STR")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	dsn="root:Nguyendepzai2@tcp(127.0.0.1:3306)/TKPM?charset=utf8&parseTime=True&loc=Local"
	s3APIKey="AKIA42E3U7VD3UISMCM4"
	s3BucketName="g04images"
	s3Domain="https://d3pfouzi5at9lt.cloudfront.net"
	s3Region="ap-southeast-1"
	s3SecretKey="asbBuwc34W0n5mXhbM9x2WU019r8/LUbJ5Q/o5Qj"
	secretKey="go-appwaangyeuusus"

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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	// route
	v1 := r.Group("v1")
	admin.AdminRoute(appContext, v1)
	client.ClientRoute(appContext, v1)
	user.UserRoute(appContext, v1)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
