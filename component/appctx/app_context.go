package appctx

import (
	"github.com/orgball2608/helmet-shop-be/component/uploadprovider"
	"github.com/orgball2608/helmet-shop-be/config"
	"github.com/orgball2608/helmet-shop-be/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	GetSecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	config         *config.Config
	ps             pubsub.Pubsub
}

func NewAppContext(
	db *gorm.DB,
	uploadProvider uploadprovider.UploadProvider,
	config *config.Config,
	ps pubsub.Pubsub,
) *appCtx {
	return &appCtx{
		db:             db,
		uploadProvider: uploadProvider,
		config:         config,
		ps:             ps,
	}
}

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.ps
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetSecretKey() string {
	return ctx.config.S3SecretKey
}
