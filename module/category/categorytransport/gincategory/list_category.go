package gincategory

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/category/categorybiz"
	"github.com/orgball2608/helmet-shop-be/module/category/categorymodel"
	"github.com/orgball2608/helmet-shop-be/module/category/categoryrepository"
	"github.com/orgball2608/helmet-shop-be/module/category/categorystorage"
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
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
