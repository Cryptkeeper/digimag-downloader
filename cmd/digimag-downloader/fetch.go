package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(url, filePath string) error {
	if resp, err := http.Get(url); err != nil {
		return err
	} else if b, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else {
		resp.Body.Close()
		return ioutil.WriteFile(filePath, b, os.ModePerm)
	}
}
