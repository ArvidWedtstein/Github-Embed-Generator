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

func Streak(user, hide_title string, cardstyle style.Styles) string {
	height := 200
	width := 400
	strokewidth := 8
	customstyles := []string{
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.streakcircle {`,
		`fill: none;`,
		fmt.Sprintf(`stroke: %v;`, "#e25822"),
		`}`,
		`.box {
			fill: ` + cardstyle.Background + `;
			border: 3px solid #` + cardstyle.Border + `;
			stroke: ` + cardstyle.Border + `;
			stroke-width: ` + strconv.Itoa(strokewidth) + `px;
		}`,
		`.streaktxt {
			font-size: 40px;
			font-family: Helvetica;
			font-weight: 600;
			fill: ` + cardstyle.Text + `;
		}`,
		`.mediantxt {
			font-size: 24px;	
		}`,
		`.titletxt { font-size: 16px;}`,
	}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
		style.WavyFilter(),
	}

	body := []string{
		fmt.Sprintf(`<rect x="%v" y="%v" class="box" width="%v" height="%v" rx="15"  />`, strokewidth/2, strokewidth/2, width, height),
	}

	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}

	hideTitle, _ := strconv.ParseBool(hide_title)

	if !hideTitle {
		bodyAdd(fmt.Sprintf(`<text x="%v" y="35" text-anchor="middle" class="title">%s</text>`, (width / 2), card.ToTitleCase("Streak")))
	}
	currentDate := time.Now()
	year := currentDate.Year()
	url := fmt.Sprintf("https://skyline.github.com/%v/%v.json", user, year)
	/*
		query {
		  viewer {
		    contributionsCollection(
		        from: "2021-03-01T00:00:00Z",
		        to: "2022-03-01T00:00:00Z"
		    ) {
		      contributionCalendar {
		        totalContributions
		        colors
		        weeks {
		          firstDay
		          contributionDays {
		            date
		            weekday
		            contributionCount
		            contributionLevel
		            color
		          }
		        }
		      }
		  }
		  }
		}

	*/
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
			if i == len(weeks)-1 && b <= int(currentDate.Weekday()) {
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
	bodyAdd(fmt.Sprintf(`<circle class="streakcircle" stroke-width="5" cx="%v" cy="%v" r="50"></circle>`, width/2, height/2))
	bodyAdd(fmt.Sprintf(`<circle class="streakcircle" stroke-width="%v" filter="url(#wavy) blur(3px)" cx="%v" cy="%v" r="50"></circle>`, strokewidth, width/2, height/2))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="streaktxt">%v</text>`, (width / 2), (height/2)+15, streak))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="mediantxt text">%v</text>`, (width/2)+130, (height/2)+5, resObjectAPI.Max))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="titletxt text">Highest Commit</text>`, (width/2)+130, (height/2)-30))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="mediantxt text">%v</text>`, (width/2)-130, (height/2)+5, resObjectAPI.Median))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="titletxt text">Median</text>`, (width/2)-130, (height/2)-30))

	return strings.Join(card.GenerateCard(cardstyle, defs, body, width+strokewidth, height+strokewidth, customstyles...), "\n")
}
