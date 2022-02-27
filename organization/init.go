package organization

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Response []struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
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
	} `json:"owner"`
	HTMLURL          string      `json:"html_url"`
	Description      interface{} `json:"description"`
	Fork             bool        `json:"fork"`
	URL              string      `json:"url"`
	ForksURL         string      `json:"forks_url"`
	KeysURL          string      `json:"keys_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	TeamsURL         string      `json:"teams_url"`
	HooksURL         string      `json:"hooks_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	EventsURL        string      `json:"events_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BranchesURL      string      `json:"branches_url"`
	TagsURL          string      `json:"tags_url"`
	BlobsURL         string      `json:"blobs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	TreesURL         string      `json:"trees_url"`
	StatusesURL      string      `json:"statuses_url"`
	LanguagesURL     string      `json:"languages_url"`
	StargazersURL    string      `json:"stargazers_url"`
	ContributorsURL  string      `json:"contributors_url"`
	SubscribersURL   string      `json:"subscribers_url"`
	SubscriptionURL  string      `json:"subscription_url"`
	CommitsURL       string      `json:"commits_url"`
	GitCommitsURL    string      `json:"git_commits_url"`
	CommentsURL      string      `json:"comments_url"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	ContentsURL      string      `json:"contents_url"`
	CompareURL       string      `json:"compare_url"`
	MergesURL        string      `json:"merges_url"`
	ArchiveURL       string      `json:"archive_url"`
	DownloadsURL     string      `json:"downloads_url"`
	IssuesURL        string      `json:"issues_url"`
	PullsURL         string      `json:"pulls_url"`
	MilestonesURL    string      `json:"milestones_url"`
	NotificationsURL string      `json:"notifications_url"`
	LabelsURL        string      `json:"labels_url"`
	ReleasesURL      string      `json:"releases_url"`
	DeploymentsURL   string      `json:"deployments_url"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	PushedAt         time.Time   `json:"pushed_at"`
	GitURL           string      `json:"git_url"`
	SSHURL           string      `json:"ssh_url"`
	CloneURL         string      `json:"clone_url"`
	SvnURL           string      `json:"svn_url"`
	Homepage         interface{} `json:"homepage"`
	Size             int         `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorURL        interface{} `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		SpdxID string `json:"spdx_id"`
		URL    string `json:"url"`
		NodeID string `json:"node_id"`
	} `json:"license"`
	AllowForking  bool          `json:"allow_forking"`
	IsTemplate    bool          `json:"is_template"`
	Topics        []interface{} `json:"topics"`
	Visibility    string        `json:"visibility"`
	Forks         int           `json:"forks"`
	OpenIssues    int           `json:"open_issues"`
	Watchers      int           `json:"watchers"`
	DefaultBranch string        `json:"default_branch"`
	Permissions   struct {
		Admin    bool `json:"admin"`
		Maintain bool `json:"maintain"`
		Push     bool `json:"push"`
		Triage   bool `json:"triage"`
		Pull     bool `json:"pull"`
	} `json:"permissions"`
	Commits int
}
type Activity []struct {
	Total int   `json:"total"`
	Week  int   `json:"week"`
	Days  []int `json:"days"`
}
type Author struct {
	Login      string `json:"login"`
	Avatar_Url string `json:"avatar_url"`
}

type OrgCard struct {
	Title        string   `json:"title"`
	Organization string   `json:"score"`
	Styles       Styles   `json:"styles"`
	Body         []string `json:"body"`
}

type Styles struct {
	Title      string
	Border     string
	Background string
	Text       string
	Textfont   string
	Box        string
}

func MostactivityCard(title string, org string, style Styles, github_token string) OrgCard {

	userurl := "https://api.github.com/orgs/" + org + "/repos"
	// userurl := "https://api.github.com/orgs/devco-morkjebla/repos"

	// Create a new request using http
	req, err := http.NewRequest("GET", userurl, nil)

	// add authorization header to the req
	req.Header.Set("Authorization", "Token "+github_token)
	if err != nil {
		panic(err.Error())
	}
	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	height := 700
	width := 600
	boxwidth := 270
	boxheight := 100
	titleboxheight := 50
	padding := 10
	strokewidth := 3

	body := []string{
		`<style>`,
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.text { font: 20px sans-serif; fill: #` + style.Text + `; font-family: ` + style.Textfont + `; text-decoration: underline;}`,
		`.large {
			font: 25px sans-serif; 
			fill: black
		}`,
		`.title { font: 25px sans-serif; fill: #` + style.Title + `}`,
		`.repobox { 
			fill: #` + style.Box + `;
			border: ` + strconv.Itoa(strokewidth) + `px solid #` + style.Border + `;
		}`,
		`.repobox:hover { fill: rgba(255,0,0,0.8);}`,
		`.repobox:hover rect {filter: blur(30px);}`,
		`.box {
			fill: #` + style.Background + `;
			border: 3px solid #` + style.Border + `;
			stroke: #` + style.Border + `;
			stroke-width: ` + strconv.Itoa(strokewidth) + `px;
		}`,
		`</style>`,
		fmt.Sprintf(`<text x="20" y="35" class="title">%s</text>`, ToTitleCase(org)),
	}
	txt := func(content string, posX int, posY int, class ...string) string {
		return fmt.Sprintf(`<text x="%v" y="%v" class="%v">%v</text>`, posX, posY, strings.Join(class, " "), content)
	}
	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}

	// Calculate where repositoryboxes should begin
	posY := titleboxheight + padding
	originalpos := posY
	posX := 0
	row := func(content []string) {
		bodyAdd(`<g class="repobox" transform="translate(` + strconv.Itoa(posX+padding) + `,` + strconv.Itoa(posY) + `) rotate(0)">`)
		for _, v := range content {
			bodyAdd(v)
		}
		bodyAdd(`</g>`)

		// check if next box will fit into card
		if posY+boxheight+(boxheight+padding) >= height {
			posX += boxwidth + padding
			posY = originalpos - (boxheight + padding)
		}
	}

	// Make sure it is not longer than 10 repos
	if len(responseObject) > 10 {
		responseObject = responseObject[:10]
	}

	bodyAdd(`<g>`)
	for i, r := range responseObject {
		response2, err := http.Get("https://api.github.com/repos/" + org + "/" + r.Name + "/stats/commit_activity")
		if err != nil {
			panic(err.Error())
		}

		responseData2, err := ioutil.ReadAll(response2.Body)
		if err != nil {
			panic(err)
		}
		var responseContent Activity
		json.Unmarshal(responseData2, &responseContent)

		totalCommits := 0
		// Slice to last 4. Get commit activity from last 4 weeks
		if len(responseContent) > 4 {
			responseContent = responseContent[len(responseContent)-4:]
			for _, g := range responseContent {
				totalCommits += g.Total
			}
		}
		fmt.Printf("%v - %v with %v total commits", i+1, r.Name, totalCommits)
		r.Commits = totalCommits
		row([]string{
			fmt.Sprintf(`<rect x="0" y="0" rx="5" class="repobox" width="%v" height="%v" />`, boxwidth, boxheight),
			`<a href="` + r.HTMLURL + `"><text x="5" y="30" class="title">` + r.Name + `</text></a>`,
			`<text x="5" y="50" class="text">Language - ` + r.Language + `</text>`,
			txt(`Language - `+r.Language, 5, 50, `text`),
			`<text x="5" y="70" class="text">Issues - ` + strconv.Itoa(r.OpenIssuesCount) + `</text>`,
			`<text x="5" y="90" class="text">Commits - ` + strconv.Itoa(totalCommits) + `</text>`,
		})
		if posX+boxwidth >= width {
			width = posX + padding
			break
		}
		posY += boxheight + padding
	}
	bodyAdd(`</g>`)
	body = append([]string{fmt.Sprintf(`<rect x="0" y="%v" width="%v" height="%v" fill="#%v"/>`, titleboxheight, width, strokewidth, style.Border)}, body...)
	body = append([]string{fmt.Sprintf(`<rect x="0" y="0" class="box" width="%v" height="%v" rx="15"  />`, width, height)}, body...)
	svgTag := fmt.Sprintf(`<svg width="%v" height="%v" fill="none" viewBox="0 0 %v %v" xmlns="http://www.w3.org/2000/svg">`, width+strokewidth, height+strokewidth, width+strokewidth, height+strokewidth)
	body = append([]string{svgTag}, body...)
	bodyAdd(`</svg>`)
	newcard := OrgCard{title, org, style, body}
	return newcard

}
func ToTitleCase(str string) string {
	return strings.Title(str)
}
