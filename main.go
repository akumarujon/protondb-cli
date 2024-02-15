package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Game struct {
	AppID string `json:"appId"`
	Title string `json:"title"`
}

type Report struct {
	ID            string `json:"id"`
	AppID         string `json:"appId"`
	Timestamp     string `json:"timestamp"`
	Rating        string `json:"rating"`
	Notes         string `json:"notes"`
	OS            string `json:"os"`
	GPUDriver     string `json:"gpuDriver"`
	Specs         string `json:"specs"`
	ProtonVersion string `json:"protonVersion"`
}

const baseURL = "https://protondb.max-p.me/"

func main() {
	game_name := os.Args[1:]

	resp, err := http.Get(baseURL + "games")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var games []Game
	err = json.Unmarshal(body, &games)
	if err != nil {
		panic(err)
	}

	var one_game Game

	for game := range games {
		if games[game].Title == strings.Join(game_name, " ") {
			one_game = games[game]
			break
		}
	}

	resp, err = http.Get(baseURL + "games/" + one_game.AppID + "/reports")

	if err != nil {
		panic(err)
	}

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var reports []Report
	err = json.Unmarshal(body, &reports)
	if err != nil {
		panic(err)
	}

	report := reports[0]

	fmt.Println("ID:", report.ID)
	fmt.Println("AppID:", report.AppID)
	fmt.Println("Timestamp:", report.Timestamp)
	fmt.Println("Rating:", report.Rating)
	fmt.Println("Notes:", report.Notes)
	fmt.Println("OS:", report.OS)
	fmt.Println("GPUDriver:", report.GPUDriver)
	fmt.Println("Specs:", report.Specs)
	fmt.Println("ProtonVersion:", report.ProtonVersion)
}
