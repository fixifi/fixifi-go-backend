

package response 

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type Response struct {
	Status string `json:"status"`
	Message  string `json:"error message"`
}




func GetErrorResponse(err error)Response{
	return Response{
        Status:  StatusError,
        Message: err.Error(),
    }
}





