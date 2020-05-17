package farmerly

import (
	"fmt"
	"net/url"
)

func hasEmptyValues(values map[string][]string) bool {
	if len(values) == 0 {
		return true
	}
	return false
}

func isError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func verifyData(details url.Values) {
	if len(details["category"][0]) != 0 && len(details["username"][0]) != 0 {
		fmt.Println("Submitted successfully, your details are")
		fmt.Println(details)
	}

}
