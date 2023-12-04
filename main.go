package main

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("Unable to load config file!")
		return v, err
	}

	if err := json.Unmarshal(bytes, &v); err != nil {
		fmt.Println("JSON decode error!")
		return v, err
	}

	return v, nil
}

func main() {
	if v, err := loadJson("videos.json"); err == nil {
		pp.Println(v.Videos[0])
	}

}
