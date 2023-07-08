package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucassarcanjo/pagination-study/pkg/model"
)

type OffsetParams struct {
	Offset int `form:"offset" binding:"required,min=0"`
	Limit  int `form:"limit" binding:"required,min=1"`
}

func (s *Server) HandleGetUsersOffset(ctx *gin.Context) {
	var params OffsetParams

	if err := ctx.Bind(&params); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user model.User

	data := s.DB.Offset(params.Offset).Limit(params.Limit).Find(&user)

	if data.RowsAffected == 0 {
		fmt.Println("0 records found")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if data.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, data.Error)
		return
	}

	ctx.JSON(http.StatusOK, &data)
}
