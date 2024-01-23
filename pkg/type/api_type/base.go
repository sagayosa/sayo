package apitype

import (
	sayoerror "sayo_framework/pkg/sayo_error"
)

type BaseResp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var successResp BaseResp = BaseResp{
	Code: 200,
	Msg:  "Successful",
}

func NewBaseRespByError(err error) *BaseResp {
	code, msg := sayoerror.GetErrMsgByErr(err)

	return &BaseResp{
		Code: code,
		Msg:  msg,
	}
}

func NewSuccessResp() BaseResp {
	return successResp
}
