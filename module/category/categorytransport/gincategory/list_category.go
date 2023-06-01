package gincategory

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/category/categorybiz"
	"TKPM-Go/module/category/categorymodel"
	"TKPM-Go/module/category/categoryrepository"
	"TKPM-Go/module/category/categorystorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var filter categorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		pagingData.Fulfill()

		var result []categorymodel.Category
		store := categorystorage.NewSQLStore(db)
		repo := categoryrepository.NewListCategoryRepo(store)
		business := categorybiz.NewListCategoryBusiness(repo)
		result, err := business.ListCategory(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
			products := result[i].Products
			for i := range products {
				products[i].Mask()
				products[i].GenCategoryUID()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
