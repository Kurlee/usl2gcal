package main

import (
	"fmt"
	"usl2gcal/internal/webClient"
)

func main() {
	pageContents := webClient.Get_https("https://www.uslchampionship.com/league-schedule")
	fmt.Printf("%s\n", pageContents)
}
