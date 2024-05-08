package subscriber

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/component/asyncjob"
	"github.com/orgball2608/helmet-shop-be/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appCtx appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (engine *consumerEngine) Start() error {

	engine.startSubTopic(
		common.TopicUserAddProduct,
		true,
		IncreaseProductTotalAfterAddProduct(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserDeleteProduct,
		true,
		DecreaseProductTotalAfterDeleteProduct(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserRatingProduct,
		true,
		IncreaseTotalRatingAfterUserRating(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicUserDeleteRatingProduct,
		true,
		DecreaseTotalRatingAfterUserRemoveRating(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicAddProductsToCart,
		true,
		IncreaseProductTotalAfterAddProductsToCart(engine.appCtx),
	)

	engine.startSubTopic(
		common.TopicRemoveProductsFromCart,
		true,
		DecreaseProductTotalAfterRemoveProductsFromCart(engine.appCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("setup consumer for: ", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHldArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHldArr[i] = asyncjob.NewJob(jobHdl)
			}

			group := asyncjob.NewGroup(isConcurrent, jobHldArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
