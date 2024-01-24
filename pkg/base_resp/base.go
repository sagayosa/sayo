package baseresp

import (
	sayoerror "sayo_framework/pkg/sayo_error"
)

type BaseResp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *BaseResp) WithData(data interface{}) *BaseResp {
	r.Data = data
	return r
}

// func (r *BaseResp) ToString() string {
// 	bts, err := json.Marshal(r)
// 	sayolog.Msg("BaseResp Marshal failed").Err(err).Error()
// 	return string(bts)
// }

func NewBaseRespByError(err error) *BaseResp {
	code, msg := sayoerror.GetErrMsgByErr(err)

	return &BaseResp{
		Code: code,
		Msg:  msg,
	}
}

func NewSuccessResp(data interface{}) *BaseResp {
	return &BaseResp{
		Code: 200,
		Msg:  "successful",
		Data: data,
	}
}
