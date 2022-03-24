package project

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

type ProjectActivity []struct {
	Total int `json:"total"`
	Weeks []struct {
		Week      int `json:"w"`
		Additions int `json:"a"`
		Deletions int `json:"d"`
		Commits   int `json:"c"`
	} `json:"weeks"`
	Author struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
}
type CommitHistory struct {
	Data struct {
		Repository struct {
			DefaultBranchRef struct {
				Name   string `json:"name"`
				Target struct {
					ID      string `json:"id"`
					History struct {
						Nodes []struct {
							Oid          string `json:"oid"`
							Message      string `json:"message"`
							Additions    int    `json:"additions"`
							Deletions    int    `json:"deletions"`
							ChangedFiles int    `json:"changedFiles"`
						} `json:"nodes"`
					} `json:"history"`
				} `json:"target"`
			} `json:"defaultBranchRef"`
		} `json:"repository"`
	} `json:"data"`
}

func recoverFromError() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}
func Project(user, project string, cardstyle themes.Theme) string {
	goal := 1000

	apiurl := "https://api.github.com/repos/" + user + "/" + project + "/stats/contributors"

	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{
			repository(owner: "%v", name: "%v") {
				defaultBranchRef {
				  name
				  target {
					... on Commit {
					  id
					  history(first: 100) {
						nodes {
						  oid
						  message
						  additions
						  deletions
						  changedFiles
						}
					  }
					}
				  }
				}
			  }
		`, user, project),
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
	var data CommitHistory
	json.Unmarshal(responseData, &data)

	additionsList := []int{}
	deletionsList := []int{}
	fileschanged := []int{}
	for _, v := range data.Data.Repository.DefaultBranchRef.Target.History.Nodes {
		additionsList = append(additionsList, v.Additions)
		deletionsList = append(deletionsList, v.Deletions)
		fileschanged = append(fileschanged, v.ChangedFiles)
	}

	/* old */
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

	var resObjectAPI ProjectActivity
	json.Unmarshal(responseDataAPI, &resObjectAPI)

	i, _ := strconv.ParseInt(strconv.Itoa(resObjectAPI[0].Weeks[len(resObjectAPI[0].Weeks)-1].Week), 10, 64)
	tm := time.Unix(i, 0)
	_, week := tm.ISOWeek()

	additions := resObjectAPI[0].Weeks[len(resObjectAPI[0].Weeks)-1].Additions
	deletions := resObjectAPI[0].Weeks[len(resObjectAPI[0].Weeks)-1].Deletions
	commits := resObjectAPI[0].Weeks[len(resObjectAPI[0].Weeks)-1].Commits

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
		// style.StarsFilter(),
		style.DropShadowRing1(),
		// style.CubePattern(),
	}
	paddingX := 30
	paddingY := 30

	prog1, style1 := card.CircleProgressbar(card.CalculatePercent(additions, goal), 80, 10, 0, 0, "#39d353", "circle")
	prog2, style2 := card.CircleProgressbar(card.CalculatePercent(deletions, goal), 70, 10, 0, 0, "red", "circle")
	prog3, style3 := card.CircleProgressbar(card.CalculatePercent(commits, goal), 60, 10, 0, 0, "blue", "circle")
	customstyles = append(customstyles, style1)
	customstyles = append(customstyles, style2)
	customstyles = append(customstyles, style3)

	body := []string{
		`<g id="Box">`,
		`    <mask id="path-1-inside-1_36_15" fill="white">`,
		`        <path d="M539.343 287.881C545.213 287.881 549.972 283.123 549.972 277.252V248.844C549.972 245.794 547.498 243.338 544.453 243.171C542.626 243.271 540.786 243.322 538.934 243.322C538.892 243.322 538.85 243.322 538.808 243.322C533.308 243.315 527.799 243.315 522.3 243.322C522.258 243.322 522.216 243.322 522.173 243.322C520.804 243.322 519.441 243.294 518.085 243.239C516.73 243.294 515.367 243.322 513.997 243.322C513.956 243.322 513.914 243.322 513.873 243.322C508.644 243.315 503.408 243.315 498.179 243.322C498.137 243.322 498.096 243.322 498.054 243.322C498.012 243.322 497.97 243.322 497.927 243.322C492.292 243.315 486.647 243.315 481.012 243.322C480.969 243.322 480.927 243.322 480.885 243.322C426.021 243.322 381.546 198.847 381.546 143.983C381.546 89.1204 426.021 44.645 480.885 44.645C480.927 44.645 480.97 44.645 481.012 44.6451C486.647 44.6522 492.292 44.6522 497.927 44.6451C497.969 44.645 498.012 44.645 498.054 44.645C498.096 44.645 498.138 44.645 498.179 44.6451C503.408 44.6515 508.644 44.6515 513.872 44.6451C513.914 44.645 513.956 44.645 513.997 44.645C515.367 44.645 516.73 44.6727 518.085 44.7276C519.441 44.6727 520.804 44.645 522.173 44.645C522.216 44.645 522.258 44.645 522.3 44.6451C527.799 44.6519 533.308 44.6519 538.807 44.6451C538.85 44.645 538.892 44.645 538.934 44.645C540.786 44.645 542.626 44.6956 544.453 44.7957C547.498 44.6289 549.972 42.1729 549.972 39.1234V10.7146C549.972 4.84445 545.213 0.0857539 539.343 0.0857544L11.173 0.0858006C5.30285 0.0858011 0.544158 4.84449 0.544159 10.7146L0.544182 277.252C0.544183 283.123 5.30287 287.881 11.173 287.881L539.343 287.881Z"/>`,
		`    </mask>`,
		fmt.Sprintf(`<path fill="%v" stroke="black" stroke-width="6" mask="url(#path-1-inside-1_36_15)" d="M539.343 287.881C545.213 287.881 549.972 283.123 549.972 277.252V248.844C549.972 245.794 547.498 243.338 544.453 243.171C542.626 243.271 540.786 243.322 538.934 243.322C538.892 243.322 538.85 243.322 538.808 243.322C533.308 243.315 527.799 243.315 522.3 243.322C522.258 243.322 522.216 243.322 522.173 243.322C520.804 243.322 519.441 243.294 518.085 243.239C516.73 243.294 515.367 243.322 513.997 243.322C513.956 243.322 513.914 243.322 513.873 243.322C508.644 243.315 503.408 243.315 498.179 243.322C498.137 243.322 498.096 243.322 498.054 243.322C498.012 243.322 497.97 243.322 497.927 243.322C492.292 243.315 486.647 243.315 481.012 243.322C480.969 243.322 480.927 243.322 480.885 243.322C426.021 243.322 381.546 198.847 381.546 143.983C381.546 89.1204 426.021 44.645 480.885 44.645C480.927 44.645 480.97 44.645 481.012 44.6451C486.647 44.6522 492.292 44.6522 497.927 44.6451C497.969 44.645 498.012 44.645 498.054 44.645C498.096 44.645 498.138 44.645 498.179 44.6451C503.408 44.6515 508.644 44.6515 513.872 44.6451C513.914 44.645 513.956 44.645 513.997 44.645C515.367 44.645 516.73 44.6727 518.085 44.7276C519.441 44.6727 520.804 44.645 522.173 44.645C522.216 44.645 522.258 44.645 522.3 44.6451C527.799 44.6519 533.308 44.6519 538.807 44.6451C538.85 44.645 538.892 44.645 538.934 44.645C540.786 44.645 542.626 44.6956 544.453 44.7957C547.498 44.6289 549.972 42.1729 549.972 39.1234V10.7146C549.972 4.84445 545.213 0.0857539 539.343 0.0857544L11.173 0.0858006C5.30285 0.0858011 0.544158 4.84449 0.544159 10.7146L0.544182 277.252C0.544183 283.123 5.30287 287.881 11.173 287.881L539.343 287.881Z"/>`, cardstyle.Colors.Background),
		`</g>`,
		`<g id="Stat" transform="translate(480,145)">`,
		prog1,
		prog2,
		prog3,

		`</g>`,
		`<g data-testid="card-text">`,
		fmt.Sprintf(`<text x="%v" y="%v" id="Stats" class="title">%v Stats</text>`, paddingX, paddingY, card.ToTitleCase(project)),
		fmt.Sprintf(`<line id="gradLine" x1="%v" y1="40" x2="400" y2="40" stroke="url(#paint0_angular_0_1)"/>`, paddingX),
		fmt.Sprintf(`<text x="%v" y="130" id="Average" class="text">Average: %v</text>`, paddingX, goal),
		// fmt.Sprintf(`<text x="%v" y="150" id="Additions" class="text">Additions: %v%v🟩</text>`, paddingX, card.CalculatePercent(additions, goal), "%"),
		// fmt.Sprintf(`<text x="%v" y="170" id="Deletions" class="text">Deletions: %v%v🟥</text>`, paddingX, card.CalculatePercent(deletions, goal), "%"),
		// fmt.Sprintf(`<text x="%v" y="190" id="Commits" class="text">Commits: %v🟦</text>`, paddingX, commits),
		fmt.Sprintf(`<text x="%v" y="150" id="Additions" class="text">Additions: %v%v🟩</text>`, paddingX, card.Average(additionsList), "%"),
		fmt.Sprintf(`<text x="%v" y="170" id="FilesChanged" class="text">Files Changed: %v%v🟥</text>`, paddingX, card.Average(fileschanged), "%"),
		fmt.Sprintf(`<text x="%v" y="190" id="Commits" class="text">Deletions: %v%v🟥</text>`, paddingX, card.Average(deletionsList), "%"),
		fmt.Sprintf(`<text x="%v" y="230" id="Week" class="text">Week: %v</text>`, paddingX, week),
		fmt.Sprintf(`<text x="440" y="130" id="Additions" class="text">Add: %v%v</text>`, card.CalculatePercent(additions, goal), "%"),
		fmt.Sprintf(`<text x="440" y="150" id="Deletions" class="text">Del: %v%v</text>`, card.CalculatePercent(deletions, goal), "%"),
		fmt.Sprintf(`<text x="440" y="170" id="Deletions" class="text">Com: %v%v</text>`, card.CalculatePercent(commits, goal), "%"),
		`</g>`,
	}
	return strings.Join(card.GenerateCard(cardstyle, defs, body, 600, 300, customstyles...), "\n")
}
