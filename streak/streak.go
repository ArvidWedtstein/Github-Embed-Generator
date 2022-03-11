package streak

import (
	"encoding/json"
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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

func Streak(user string, cardstyle style.Styles) string {
	height := 300
	width := 600
	strokewidth := 3
	customstyles := []string{
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.streakcircle {`,
		`fill: none;`,
		fmt.Sprintf(`stroke: %v;`, cardstyle.Border),
		fmt.Sprintf(`stroke-width: %v;`, strokewidth),
		`}`,
		`.box {
			fill: ` + cardstyle.Background + `;
			border: 3px solid #` + cardstyle.Border + `;
			stroke: ` + cardstyle.Border + `;
			stroke-width: ` + strconv.Itoa(strokewidth) + `px;
		}`,
	}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
	}

	body := []string{
		fmt.Sprintf(`<rect x="%v" y="%v" class="box" width="%v" height="%v" rx="15"  />`, strokewidth/2, strokewidth/2, width, height),
		fmt.Sprintf(`<text x="20" y="35" class="title">%s</text>`, card.ToTitleCase("Streak")),
	}

	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}

	currentDate := time.Now()
	year := currentDate.Year()
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
	for i := len(weeks) - 1; i >= 0; i-- { // Loop through weeks
		for b := len(weeks[i].Days) - 1; b >= 0; b-- { // Loop though days
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
	bodyAdd(fmt.Sprintf(`<svg width="60" height="60" xmlns="http://www.w3.org/2000/svg" fill="red" x="%v" y="%v" viewBox="0 0 384 512"><path stroke="#ffffff" stroke-width="3px" d="M216 23.86c0-23.8-30.65-32.77-44.15-13.04C48 191.85 224 200 224 288c0 35.63-29.11 64.46-64.85 63.99-35.17-.45-63.15-29.77-63.15-64.94v-85.51c0-21.7-26.47-32.23-41.43-16.5C27.8 213.16 0 261.33 0 320c0 105.87 86.13 192 192 192s192-86.13 192-192c0-170.29-168-193-168-296.14z"/></svg>`, (width/2)-25, (height/2)-100))
	bodyAdd(fmt.Sprintf(`<circle class="streakcircle" cx="%v" cy="%v" r="80"></circle>`, width/2, height/2))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" class="title">%v</text>`, (width/2)-13, (height/2)+10, streak))
	fmt.Printf("Streak: %v\n", streak)

	return strings.Join(card.GenerateCard(cardstyle, defs, body, width+strokewidth, height+strokewidth, customstyles...), "\n")
}
