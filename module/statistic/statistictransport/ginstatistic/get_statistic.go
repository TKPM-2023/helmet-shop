package ginstatistic

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/category/categorystorage"
	"TKPM-Go/module/order/orderstorage"
	"TKPM-Go/module/product/productstorage"
	"TKPM-Go/module/statistic/statisticmodel"
	"TKPM-Go/module/user/userstorage"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStatistic(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		year := context.Param("year")

		var data statisticmodel.Statistic
		user_store := userstorage.NewSQLStore(db)
		data.UserCount = user_store.CountUser()
		//fmt.Println(user_count)

		order_store := orderstorage.NewSQLStore(db)
		data.OrderCount = order_store.CountOrder()
		//fmt.Println(order_count)

		product_store := productstorage.NewSQLStore(db)
		data.ProductCount = product_store.CountProduct()
		//fmt.Println(product_count)

		category_store := categorystorage.NewSQLStore(db)
		data.CategoryCount = category_store.CountCategory()
		//fmt.Println(category_count)
		year_str, _ := strconv.Atoi(year)
		data.Revenue = order_store.GetRevenue(year_str)

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
