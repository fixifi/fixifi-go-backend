
package response

type ResponseWithData struct {
	Status string `json:"status"`
	Message  string `json:"error message"`
	Data interface{} `json:"data"`
}