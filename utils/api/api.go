package utils

import (
	"mercury_algo_backend/kitex_gen/api"
)

func GenerateBaseResp(err error) *api.BaseResponse {
	if err == nil {
		return &api.BaseResponse{
			Code:    0,
			Message: "Success",
		}
	} else {
		return &api.BaseResponse{
			Code:    -1,
			Message: err.Error(),
		}
	}
}
