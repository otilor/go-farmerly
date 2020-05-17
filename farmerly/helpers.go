// Package farmerly implements a simple farming-algorithm.
// It maps farmers to content based on their selection

package farmerly

import (
	"net/url"
)

func hasEmptyValues(details url.Values) bool {
	if len(details["category"][0]) != 0 && len(details["username"][0]) != 0 {
		return false
	}
	return true

}

func isError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
