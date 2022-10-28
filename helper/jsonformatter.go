package helper

type ResponseWithData struct {
	Success string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseErrorWithData struct {
	Success string      `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"data"`
}

type ResponseWithOutData struct {
	Success string       `json:"success"`
	Message string       `json:"message"`
	Data    StructKosong `json:"data"`
}

type StructKosong struct {
}

func ApiResponse(status string, message string, data interface{}) interface{} {
	if data == nil {
		var structkosong StructKosong
		jsonResponse := ResponseErrorWithData{
			Success: status,
			Message: message,
			Error:   structkosong,
		}
		return jsonResponse
	} else {
		jsonResponse := ResponseWithData{
			Success: status,
			Message: message,
			Data:    data,
		}
		return jsonResponse
	}

}

func ApiWithOutData(status string, message string) interface{} {
	jsonResponse := ResponseWithOutData{
		Success: status,
		Message: message,
	}
	return jsonResponse
}
