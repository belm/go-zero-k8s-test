// Code generated by goctl. DO NOT EDIT.
package types

type GetInfoRequest struct {
	Uid int64 `path:"uid"`
}

type GetInfoResponse struct {
	Uid      int64  `form:"uid"`
	Username string `form:"username"`
	Age      int64  `form:"age"`
}
