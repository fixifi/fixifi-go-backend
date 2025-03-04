package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)


func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgList []string
	for _, err := range errs {
		switch err.ActualTag() { 
		case "required":
			errMsgList = append(errMsgList, fmt.Sprintf("fild %s is required", err.Field()))
		default:
			errMsgList = append(errMsgList, fmt.Sprintf("Invalid field %s", fmt.Sprint(err.Field())))
		}
	}
	//now we join all the eror list in string that we was adding in errMsg
	errMsg := strings.Join(errMsgList, ",")

	return Response{
		Status: StatusError,
		Message:  errMsg,
	}

}