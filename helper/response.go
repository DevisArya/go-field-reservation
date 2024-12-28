package helper

import "github.com/DevisArya/reservasi_lapangan/dto"

func NewResponse(code int, msg string, data interface{}) dto.WebResponse {
	return dto.WebResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
