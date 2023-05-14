package main

import (
	"encoding/json"
	"fmt"
)

type code struct {
	Number int16  `json:"Code"`
	Desc   string `json:"Descrip"`
}

type codes []code

func main() {
	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`
	var data codes
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		fmt.Println("error: ", err)
	}

	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(data[i].Number, data[i].Desc)
	// }

	for _, code := range data {
		fmt.Println(code.Number, code.Desc)
	}
}
