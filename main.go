package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/k0kubun/pp"
)

type Videos struct {
	Videos []struct {
		URL    string `json:"url"`
		Nome   string `json:"nome"`
		Cantor string `json:"cantor"`
	} `json:"videos"`
}

func loadJson(arquivo string) (Videos, error) {
	var v Videos
	bytes, err := os.ReadFile(arquivo)

	if err != nil {
		return v, errors.New("Unable to load config file!")
	}

	if err := json.Unmarshal(bytes, &v); err != nil {
		return v, errors.New("JSON decode error!")
	}

	return v, nil
}

func main() {
	if v, err := loadJson("videos.json"); err == nil {
		pp.Println(v.Videos[0])
	}

}
