package request

// PageInfoRequest 分页信息
type PageInfoRequest struct {
	Page     int `binding:"required" json:"page"`
	PageSize int `binding:"required" json:"pageSize"`
}
