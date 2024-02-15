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

	fmt.Println(one_game.AppID)
	fmt.Println(one_game.Title)
}
