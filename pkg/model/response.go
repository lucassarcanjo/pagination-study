package model

import "github.com/lucassarcanjo/pagination-study/pkg/helpers"

type Response struct {
	Data interface{}            `json:"data"`
	Info helpers.PaginationInfo `json:"info"`
}
