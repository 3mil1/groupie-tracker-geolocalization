package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func JsonFromAPI(URL string, str interface{}) {
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	if err := json.Unmarshal(body, str); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON", URL)
	}
}
