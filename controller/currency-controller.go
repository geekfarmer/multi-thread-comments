package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Currency controller
type CurrencyController interface {
	GetCurrency(response http.ResponseWriter, request *http.Request)
}

//Method to create new Currency
func NewCurrencyController() CurrencyController {
	return &controller{}
}

func (*controller) GetCurrency(response http.ResponseWriter, request *http.Request) {
	b, _ := ioutil.ReadFile("./files/currency.json")

	rawIn := json.RawMessage(string(b))
	var objmap map[string]*json.RawMessage
	e := json.Unmarshal(rawIn, &objmap)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(objmap)
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(objmap)

}
