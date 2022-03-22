package streak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"githubembedapi/card/themes"
	"io/ioutil"
	"net/http"
	"os"
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

type StreakData struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					TotalContributions int `json:"totalContributions"`
					Weeks              []struct {
						FirstDay         string `json:"firstDay"`
						ContributionDays []struct {
							Color             string `json:"color"`
							ContributionCount int    `json:"contributionCount"`
							ContributionLevel string `json:"contributionLevel"`
							Date              string `json:"date"`
							Weekday           int    `json:"weekday"`
						} `json:"contributionDays"`
					} `json:"weeks"`
				} `json:"contributionCalendar"`
			} `json:"contributionsCollection"`
		} `json:"user"`
	} `json:"data"`
}

func Streak(user, hide_title string, cardstyle themes.Theme) string {

	year := time.Now().Year()
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{
			user(login: "%v") {
                contributionsCollection(from: "%v-01-01T00:00:00Z", to: "%v-12-31T23:59:59Z") {
                    contributionCalendar {
                        totalContributions
                        weeks {
                            contributionDays {
                            contributionCount
                            date
                            }
                        }
                    }
                }
            }
		}
		`, user, year, year),
	}

	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonValue))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("GITHUB_TOKEN")))
	if err != nil {
		panic(fmt.Sprintf("Request Failed. Error: %v", err))
	}
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Request Failed. Error: %v", err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var data StreakData
	json.Unmarshal(responseData, &data)

	type Contribution struct {
		Date          time.Time
		Contributions int
	}
	type StreakC struct {
		Start  time.Time
		End    time.Time
		Length int
	}
	type Stats struct {
		TotalContributions int
		FirstContribution  string
		LongestStreak      StreakC
		CurrentStreak      StreakC
	}

	getContributionDates := func() []Contribution {

		var contributions []Contribution
		today := time.Now()
		tomorrow := today.AddDate(0, 0, 1)

		for _, week := range data.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
			for _, day := range week.ContributionDays {
				date, err := time.Parse("2006-01-02", day.Date)
				if err != nil {
					panic(err.Error())
				}
				count := day.ContributionCount

				// count contributions until current date
				// also count next day if user contributed already
				if (date.Before(tomorrow) || date == today) || (date == tomorrow && count > 0) {
					contributions = append(contributions, Contribution{Date: date, Contributions: count})
				}
			}
		}
		return contributions
	}

	getContributionStats := func(contributions []Contribution) Stats {
		if len(contributions) <= 0 {
			panic("No contributions exist")
		}

		today := contributions[len(contributions)-1]
		first := contributions[0].Date
		stats := Stats{
			TotalContributions: 0,
			FirstContribution:  "",
			LongestStreak: StreakC{
				Start:  first,
				End:    first,
				Length: 0,
			},
			CurrentStreak: StreakC{
				Start:  first,
				End:    first,
				Length: 0,
			},
		}

		for _, date := range contributions {
			stats.TotalContributions += date.Contributions

			if date.Contributions > 0 {
				stats.CurrentStreak.Length++
				stats.CurrentStreak.End = date.Date

				if stats.CurrentStreak.Length == 1 {
					stats.CurrentStreak.Start = date.Date
				}

				if stats.FirstContribution == "" {
					stats.FirstContribution = date.Date.Format("2006-01-02")
				}

				if stats.CurrentStreak.Length > stats.LongestStreak.Length {
					stats.LongestStreak.Start = stats.CurrentStreak.Start
					stats.LongestStreak.End = stats.CurrentStreak.End
					stats.LongestStreak.Length = stats.CurrentStreak.Length
				}
			} else if date.Date != today.Date {

				// reset streak
				stats.CurrentStreak.Length = 0
				stats.CurrentStreak.Start = today.Date
				stats.CurrentStreak.End = today.Date
			}
		}
		return stats
	}
	// get stats
	contributions := getContributionDates()
	stats := getContributionStats(contributions)

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
		fmt.Sprintf(`.streaktxt {
			font-size: 40px;
			font-family: Helvetica;
			font-weight: 600;
			fill: %v;
		}`, cardstyle.Text),
		`.mediantxt {
			font-size: 24px;	
		}`,
		`.datetxt {
			font-size: 10px;
		}`,
		`.titletxt { font-size: 16px;}`,
	}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", 0, []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
		style.WavyFilter(),
	}

	body := []string{
		fmt.Sprintf(`<g><rect x="%v" y="%v" class="box" width="%v" height="%v" rx="15"  />`, strokewidth/2, strokewidth/2, width, height),
	}

	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}

	hideTitle, _ := strconv.ParseBool(hide_title)

	if !hideTitle {
		bodyAdd(fmt.Sprintf(`<text x="%v" y="35" text-anchor="middle" class="title">%s</text>`, (width / 2), card.ToTitleCase("Streak")))
	}
	ctstart := stats.CurrentStreak.Start.Format("Jan 2, 2006")
	ctend := stats.CurrentStreak.End.Format("Jan 2, 2006")

	if ctend == time.Now().Format("Jan 2, 2006") {
		ctend = "Today"
	}
	bodyAdd(fmt.Sprintf(`<circle class="streakcircle" stroke-width="5" cx="%v" cy="%v" r="50"></circle>`, width/2, height/2))
	bodyAdd(fmt.Sprintf(`<circle class="streakcircle" stroke-width="%v" filter="url(#wavy) blur(3px)" cx="%v" cy="%v" r="50"></circle>`, strokewidth, width/2, height/2))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="streaktxt">%v</text>`, (width / 2), (height/2)+15, stats.CurrentStreak.Length))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="datetxt text">%v - %v</text>`, (width / 2), (height/2)+70, ctstart, ctend))

	tstart := stats.LongestStreak.Start.Format("Jan 2, 2006")
	tend := stats.LongestStreak.End.Format("Jan 2, 2006")
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="titletxt text">Longest Streak</text>`, (width/2)+130, (height/2)-30))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="mediantxt text">%v</text>`, (width/2)+130, (height/2)+5, stats.LongestStreak.Length))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="datetxt text">%v - %v</text>`, (width/2)+130, (height/2)+20, tstart, tend))

	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="titletxt text">Total Contributions</text>`, (width/2)-120, (height/2)-30))
	bodyAdd(fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle" class="mediantxt text">%v</text>`, (width/2)-130, (height/2)+5, stats.TotalContributions))
	bodyAdd(`</g>`)
	return strings.Join(card.GenerateCard(cardstyle, defs, body, width+strokewidth, height+strokewidth, customstyles...), "\n")
}
