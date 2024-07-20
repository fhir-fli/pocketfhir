package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type SearchParameter struct {
	ResourceType string   `json:"resourceType"`
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Code         string   `json:"code"`
	Base         []string `json:"base"`
	Type         string   `json:"type"`
	Expression   string   `json:"expression"`
}

type Entry struct {
	Resource SearchParameter `json:"resource"`
}

type Bundle struct {
	Entry []Entry `json:"entry"`
}

func loadSearchParameters(filename string) (*Bundle, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var bundle Bundle
	err = json.Unmarshal(bytes, &bundle)
	if err != nil {
		return nil, err
	}
	return &bundle, nil
}
