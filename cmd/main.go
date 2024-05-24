package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/component/uploadprovider"
	"github.com/orgball2608/helmet-shop-be/config"
	"github.com/orgball2608/helmet-shop-be/middleware"
	localPb "github.com/orgball2608/helmet-shop-be/pubsub/localpub"
	"github.com/orgball2608/helmet-shop-be/route/admin"
	"github.com/orgball2608/helmet-shop-be/route/client"
	"github.com/orgball2608/helmet-shop-be/route/user"
	"github.com/orgball2608/helmet-shop-be/subscriber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal("Load config fail", err)
	}

	db, err := gorm.Open(mysql.Open(cfg.MysqlUri), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect DB failed", err)
	}
	log.Println("Connect DB success", db)

	s3Provider := uploadprovider.NewS3Provider(cfg)
	pubSub := localPb.NewPubSub()
	appContext := appctx.NewAppContext(db, s3Provider, cfg, pubSub)
	db = db.Debug()

	if err := subscriber.NewEngine(appContext).Start(); err != nil {
		log.Fatalln()
	}

	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(middleware.Recover(appContext))

	r.Use(cors.New(cors.Config{
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

	// route
	v1 := r.Group("v1")
	admin.AdminRoute(appContext, v1)
	client.ClientRoute(appContext, v1)
	user.UserRoute(appContext, v1)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: r,
	}

	done := make(chan bool)
	go func() {
		if err := GracefulShutDown(cfg, done, server); err != nil {
			fmt.Printf("Stop server shutdown error: %v\n", err.Error())
			return
		}
		fmt.Println("Stopped serving on Services")
	}()
	fmt.Println("Start HTTP Server Successfully")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Start HTTP Server Failed. Error: %s\n", err.Error())
	}
	<-done
	fmt.Println("Stopped backend application.")
}

func GracefulShutDown(config *config.Config, quit chan bool, server *http.Server) error {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.SystemTimeOutSecond)*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	close(quit)
	return nil
}
