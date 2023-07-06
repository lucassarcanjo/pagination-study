package utils

type OffsetParams struct {
	Offset uint32 `form:"offset" binding:"required,min=0"`
	Limit  uint32 `form:"limit" binding:"required,min=1"`
}

type KeysetParams struct {
	Limit  uint32 `form:"limit" binding:"required,min=1"`
	Cursor string `form:"cursor"`
}
