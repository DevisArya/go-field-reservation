package helper

import "github.com/DevisArya/reservasi_lapangan/models/web"

func NewResponse(code int, msg string, data interface{}) web.WebResponse {
	return web.WebResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}
