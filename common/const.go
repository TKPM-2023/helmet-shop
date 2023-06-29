package common

import "TKPM-Go/pubsub"

type DbType int

const (
	DbTypeCategory      DbType = 1
	DbTypeUser          DbType = 2
	DbTypeProduct       DbType = 3
	DbTypeOrder         DbType = 4
	DbTypeOrderDetail   DbType = 5
	DbTypeProductRating DbType = 6
	DbTypeContact       DbType = 7
	DbTypeCart          DbType = 8
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetUserEmail() string
	GetUserRole() string
	GetCartId() int
}

const (
	TopicUserAddProduct          pubsub.Topic = "TopicUserAddProduct"
	TopicUserDeleteProduct       pubsub.Topic = "TopicUserDeleteProduct"
	TopicUserRatingProduct       pubsub.Topic = "TopicUserRatingProduct"
	TopicUserDeleteRatingProduct pubsub.Topic = "TopicUserDeleteRatingProduct"
	TopicAddProductsToCart       pubsub.Topic = "TopicAddProductsToCart"
	TopicRemoveProductsFromCart  pubsub.Topic = "TopicRemoveProductsFromCart"
)
