package some_package

import (
	"fmt"
	utils "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"testing"
)

func init() {
	router := gin.Default() // This needs to be written to init to start the gin framework
	router.POST("/login", LoginHandler)
	utils.SetRouter(router) //Pass the started engine object into the test framework
}
// parse the error type returned
type OrdinaryResponse struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}
// The real test unit
func TestLoginHandler(t *testing.T) {
	// Define the content of the POST request
	user := map[string]interface{}{
		"username": "Valiben",
		"password": "123456",
		"age":      13,
	}
	// parse the returned response into resp
	resp := OrdinaryResponse{}
	// Call the function to initiate an http request
	err := utils.TestHandlerUnMarshalResp("POST", "/login", "json", user, &resp)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}
	// Get the returned data structure, so far, a post request test is perfectly completed,
	// If you need to benchmark output performance report is also possible
	fmt.Println("result:", resp)
}