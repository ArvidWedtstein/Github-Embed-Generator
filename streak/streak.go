package streak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Commits struct {
	Username      string `json:"username"`
	Year          string `json:"year"`
	Min           int    `json:"min"`
	Max           int    `json:"max"`
	Median        int    `json:"median"`
	P80           int    `json:"p80"`
	P90           int    `json:"p90"`
	P99           int    `json:"p99"`
	Contributions []struct {
		Week int `json:"week"`
		Days []struct {
			Count int `json:"count"`
		} `json:"days"`
	} `json:"contributions"`
}

func Streak(user string) string {
	currentDate := time.Now()
	year := currentDate.Year()
	// url := "https://api.github.com/search/commits?q=author:" + user + "&sort=author-date&order=desc&page=1"
	url := fmt.Sprintf("https://skyline.github.com/%v/%v.json", user, year)
	recoverFromError := func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}
	reqAPI, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}
	clientAPI := &http.Client{}

	responseAPI, err := clientAPI.Do(reqAPI)
	defer recoverFromError()
	if err != nil {
		panic(err.Error())
	}
	defer responseAPI.Body.Close()

	responseDataAPI, err := ioutil.ReadAll(responseAPI.Body)
	if err != nil {
		panic(err)
	}

	var resObjectAPI Commits
	json.Unmarshal(responseDataAPI, &resObjectAPI)

	//
	// days := currentDate.Sub(resObjectAPI.Items[0].Commit.Author.Date).Hours() / 24

	// if days < 1 {
	// 	fmt.Println(days)
	// }

	_, currentweek := currentDate.ISOWeek()
	streak := 0
	var weeks []struct {
		Week int "json:\"week\""
		Days []struct {
			Count int "json:\"count\""
		} "json:\"days\""
	}
	for _, week := range resObjectAPI.Contributions {
		weeks = append(weeks, week)
		// Find correct week.
		if week.Week == currentweek {
			break
		}
	}
out1:
	for i := len(weeks) - 1; i >= 0; i-- {
		for b := len(weeks[i].Days) - 1; b >= 0; b-- {
			if i == len(weeks)-1 && b+1 <= int(currentDate.Weekday()) {
				if weeks[i].Days[b].Count > 0 {
					streak++
				} else if weeks[i].Days[b].Count == 0 {
					break out1
				}
			} else if i != len(weeks)-1 {
				if weeks[i].Days[b].Count > 0 {
					streak++
				} else if weeks[i].Days[b].Count == 0 {
					break out1
				}
			}
		}
	}
	fmt.Printf("STREAk: %v\n", streak)
	// lag algoritme som sjekker streak / om dagene henger sammen

	streaklength := func() {
		currentYear, currentMonth, _ := time.Now().Date()
		currentLocation := time.Now().Location()

		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		fmt.Printf("first date of month %v\n", firstOfMonth.Day())
		fmt.Printf("last date of month: %v\n", lastOfMonth.Day())
	}
	streaklength()
	return ``
}
