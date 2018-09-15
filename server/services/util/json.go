package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

// CreateArrayJSONDecoder returns `json.Decoder` to load an array
// of  `grpc.Message` from single JSON file.
// This technique is introduced on [golang/protobuf#675].
// [golang/protobuf#675]: https://github.com/golang/protobuf/issues/675#issuecomment-411131669
func CreateArrayJSONDecoder(filePath string) *json.Decoder {

	jsonBytes, _ := ioutil.ReadFile(filePath)
	jsonString := string(jsonBytes)
	jsonDecoder := json.NewDecoder(strings.NewReader(jsonString))

	// read open bracket
	if _, err := jsonDecoder.Token(); err != nil {
		log.Fatal(err)
	}

	return jsonDecoder
}
