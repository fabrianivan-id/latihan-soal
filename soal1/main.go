package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type apiResponse struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       []match `json:"data"`
}
type match struct {
	Team1      string `json:"team1"`
	Team2      string `json:"team2"`
	Team1Goals string `json:"team1goals"`
	Team2Goals string `json:"team2goals"`
}

func totalGoals(team string, year int) (int, error) {
	sum := 0
	base := "https://jsonmock.hackerrank.com/api/football_matches"

	// helper to walk pages and accumulate from a particular team slot and goals field
	fetch := func(paramTeam string, goalsField string) error {
		page := 1
		for {
			q := url.Values{}
			q.Set("year", strconv.Itoa(year))
			q.Set(paramTeam, team)
			q.Set("page", strconv.Itoa(page))

			resp, err := http.Get(base + "?" + q.Encode())
			if err != nil {
				return err
			}
			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				return err
			}
			var ar apiResponse
			if err := json.Unmarshal(body, &ar); err != nil {
				return err
			}
			for _, m := range ar.Data {
				var gs string
				if goalsField == "team1goals" {
					gs = m.Team1Goals
				} else {
					gs = m.Team2Goals
				}
				v, _ := strconv.Atoi(gs)
				sum += v
			}
			if page >= ar.TotalPages {
				break
			}
			page++
		}
		return nil
	}

	if err := fetch("team1", "team1goals"); err != nil {
		return 0, err
	}
	if err := fetch("team2", "team2goals"); err != nil {
		return 0, err
	}
	return sum, nil
}

func main() {
	total, err := totalGoals("Barcelona", 2011)
	if err != nil {
		panic(err)
	}
	fmt.Println(total)
}
