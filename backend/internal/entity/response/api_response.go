package response

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(gin* c.Context,message string, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}
func Error(gin* c.Context,statusCode int ,message string, err interface{}){
	c.JSON(statusCode, ApiResponse{
		Status:  false,
		Message: message,
		Data:    err,
	})
}