package ginupload

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/module/upload/uploadbusiness"
	"LearnGo/module/upload/uploadstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Remove(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		imageStore := uploadstorage.NewSQLStore(db)
		biz := uploadbusiness.NewDeleteImageBusiness(imageStore)
		if err := biz.RemoveImage(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
