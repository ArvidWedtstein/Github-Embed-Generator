package commit_activity

import (
	"encoding/json"
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"githubembedapi/card/themes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RepoActivity []struct {
	Total int   `json:"total"`
	Week  int   `json:"week"`
	Days  []int `json:"days"`
}

func recoverFromError() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func RepositoryCommitActivity(title, user, repo string, hide_week string, cardstyle themes.Theme) string {
	apiurl := "https://api.github.com/repos/" + user + "/" + repo + "/stats/commit_activity"

	reqAPI, err := http.NewRequest("GET", apiurl, nil)
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

	var resObjectAPI RepoActivity
	json.Unmarshal(responseDataAPI, &resObjectAPI)

	customstyles := []string{
		`.circle {
		transform: rotate(-90deg);
		}`,
		`.rank-circle-rim {
			stroke: #333333;
			fill: none;
			opacity: 0.4;	
		}`,
	}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", 0, []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
	}

	var gitcolors = map[string]string{
		"default": "#161b22",
		"commit1": "#0e4429",
		"commit2": "#006d32",
		"commit3": "#26a641",
		"commit4": "#39d353",
	}
	paddingX := 30
	paddingY := 30
	body := []string{
		fmt.Sprintf(`<g id="Box"><rect x="0" y="0" rx="15" fill="%v" width="%v" height="%v" /></g>`, cardstyle.Colors.Background, 700, 200),
		`<g data-testid="card-text">`,
		fmt.Sprintf(`<text x="%v" y="%v" id="Stats" class="title">%v Stats</text>`, paddingX, paddingY, card.ToTitleCase(repo)),
		fmt.Sprintf(`<line id="gradLine" x1="%v" y1="40" x2="400" y2="40" stroke="url(#paint0_angular_0_1)"/>`, paddingX),
		`</g>`,
	}
	totalCommits := 0
	gridX := 30
	gridY := 100
	gridYstartPos := 100

	gridPadding := 2
	gridBoxSize := 10

	isHideWeek, _ := strconv.ParseBool(hide_week)
	if isHideWeek {
		gridYstartPos = 80
		gridY = gridYstartPos
	}
	grid := []string{`<g data-testid="card-grid">`}

	for _, data := range resObjectAPI {
		tm := time.Unix(int64(data.Week), 0)
		_, week := tm.ISOWeek()

		totalCommits += data.Total

		grid = append(grid, fmt.Sprintf(`<g id="week%v">`, week))
		if hide_week == "false" || len(hide_week) <= 0 {
			grid = append(grid, fmt.Sprintf(`<text style="font-size: 9px; font-family: Helvetica;" class="text" x="%v" y="%v">%v</text>`, gridX, gridY-5, week))
		}

		for days, commits := range data.Days {
			color := gitcolors["default"]
			if commits <= 5 && commits >= 1 {
				color = gitcolors["default"]
			} else if commits <= 10 && commits >= 5 {
				color = gitcolors["commit1"]
			} else if commits <= 15 && commits >= 10 {
				color = gitcolors["commit2"]
			} else if commits <= 20 && commits >= 15 {
				color = gitcolors["commit3"]
			} else if commits >= 20 {
				color = gitcolors["commit4"]
			}
			grid = append(grid, fmt.Sprintf(`<rect id="day%v" fill="%v" width="%v" height="%v" rx="2" x="%v" y="%v"></rect>`, days, color, gridBoxSize, gridBoxSize, gridX, gridY))
			gridY += gridBoxSize + gridPadding
		}
		grid = append(grid, `</g>`)
		gridX += gridBoxSize + gridPadding
		gridY = gridYstartPos
	}
	grid = append(grid, `</g>`)
	body = append(body, fmt.Sprintf(`<text style="font-family: Helvetica" x="30" y="70" class="text">Total commits past year: %v</text>`, totalCommits))
	body = append(body, strings.Join(grid, "\n"))
	return strings.Join(card.GenerateCard(cardstyle, defs, body, 800, 200, customstyles...), "\n")
}
