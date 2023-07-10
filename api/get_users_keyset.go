package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucassarcanjo/pagination-study/pkg/helpers"
	"github.com/lucassarcanjo/pagination-study/pkg/model"
)

type KeysetParams struct {
	Cursor string `form:"cursor"`
	Limit  int    `form:"limit" binding:"required,min=1"`
}

func (s *Server) HandleGetUsersKeyset(ctx *gin.Context) {
	var params KeysetParams

	if err := ctx.Bind(&params); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var users []model.User

	query := s.DB

	fmt.Println("cursor ")
	fmt.Println(params.Cursor)

	if params.Cursor != "" {
		cursor, err := helpers.DecodeCursor(params.Cursor)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		fmt.Println("createdAt ")
		fmt.Println(cursor.CreatedAt)

		query = query.Where("(created_at, id) > ?", []any{cursor.CreatedAt, cursor.ID})
	}

	result := query.Limit(params.Limit + 1).Find(&users)

	if result.RowsAffected == 0 {
		fmt.Println("0 records found")
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if result.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	var response model.Response
	response.Data = users[:len(users)-1]
	response.Info = helpers.GeneratePager(users[len(users)-1].ID, *users[len(users)-1].CreatedAt, int64(len(users)))

	ctx.JSON(http.StatusOK, &response)
}
