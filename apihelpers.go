package apihelpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// pretty print json string
func PrettyPrintJson(body []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "\t")
	if err != nil {
		log.Fatal("Pretty Print Json: ", err)
		return "", err
	}

	return string(prettyJSON.Bytes()), nil
}

var client = &http.Client{
	Timeout: time.Second * 10, // always configure timeout
}

func GetJson(url string, reqBody []byte) (map[string]interface{}, []byte, error) {

	var reqBodyToReader = bytes.NewBuffer([]byte(``))

	if len(reqBody) != 0 {
		reqBodyToReader = bytes.NewBuffer(reqBody)
	} else {
		reqBodyToReader = bytes.NewBuffer(nil)
	}

	// Build the request
	req, err := http.NewRequest("GET", url, reqBodyToReader)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, nil, err
	}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, nil, err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	respBody := body
	//fmt.Println(reflect.TypeOf(body))
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, nil, err
	}

	// getting json data without a struct
	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatal("Unmarshal: ", err)
		return nil, nil, err
	}

	return data, respBody, nil
}

