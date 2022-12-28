package jsonReader

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonMap map[string]interface{}

func (j *jsonMap) Read(input string) error {
	err := json.Unmarshal([]byte(input), &j)
	fmt.Println(j)
	return err
}

// using type pointer to fetch data
func (j *jsonMap) ReadFile(location string) error {
	data, err := os.ReadFile(location)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, j)
	if err != nil {
		return err
	}
	return err
}
func (j jsonMap) ToMap() map[string]interface{} {
	return j
}
func New() *jsonMap {
	return &jsonMap{}
}
