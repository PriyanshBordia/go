package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)


func ParseBody(router *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(router.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}