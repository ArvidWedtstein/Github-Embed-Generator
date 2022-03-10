package streak

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Commits struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		URL         string `json:"url"`
		Sha         string `json:"sha"`
		NodeID      string `json:"node_id"`
		HTMLURL     string `json:"html_url"`
		CommentsURL string `json:"comments_url"`
		Commit      struct {
			URL    string `json:"url"`
			Author struct {
				Date  time.Time `json:"date"`
				Name  string    `json:"name"`
				Email string    `json:"email"`
			} `json:"author"`
			Committer struct {
				Date  time.Time `json:"date"`
				Name  string    `json:"name"`
				Email string    `json:"email"`
			} `json:"committer"`
			Message string `json:"message"`
			Tree    struct {
				URL string `json:"url"`
				Sha string `json:"sha"`
			} `json:"tree"`
			CommentCount int `json:"comment_count"`
		} `json:"commit"`
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
		Committer struct {
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
		} `json:"committer"`
		Parents []struct {
			URL     string `json:"url"`
			HTMLURL string `json:"html_url"`
			Sha     string `json:"sha"`
		} `json:"parents"`
		Repository struct {
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
			HTMLURL          string `json:"html_url"`
			Description      string `json:"description"`
			Fork             bool   `json:"fork"`
			URL              string `json:"url"`
			ForksURL         string `json:"forks_url"`
			KeysURL          string `json:"keys_url"`
			CollaboratorsURL string `json:"collaborators_url"`
			TeamsURL         string `json:"teams_url"`
			HooksURL         string `json:"hooks_url"`
			IssueEventsURL   string `json:"issue_events_url"`
			EventsURL        string `json:"events_url"`
			AssigneesURL     string `json:"assignees_url"`
			BranchesURL      string `json:"branches_url"`
			TagsURL          string `json:"tags_url"`
			BlobsURL         string `json:"blobs_url"`
			GitTagsURL       string `json:"git_tags_url"`
			GitRefsURL       string `json:"git_refs_url"`
			TreesURL         string `json:"trees_url"`
			StatusesURL      string `json:"statuses_url"`
			LanguagesURL     string `json:"languages_url"`
			StargazersURL    string `json:"stargazers_url"`
			ContributorsURL  string `json:"contributors_url"`
			SubscribersURL   string `json:"subscribers_url"`
			SubscriptionURL  string `json:"subscription_url"`
			CommitsURL       string `json:"commits_url"`
			GitCommitsURL    string `json:"git_commits_url"`
			CommentsURL      string `json:"comments_url"`
			IssueCommentURL  string `json:"issue_comment_url"`
			ContentsURL      string `json:"contents_url"`
			CompareURL       string `json:"compare_url"`
			MergesURL        string `json:"merges_url"`
			ArchiveURL       string `json:"archive_url"`
			DownloadsURL     string `json:"downloads_url"`
			IssuesURL        string `json:"issues_url"`
			PullsURL         string `json:"pulls_url"`
			MilestonesURL    string `json:"milestones_url"`
			NotificationsURL string `json:"notifications_url"`
			LabelsURL        string `json:"labels_url"`
			ReleasesURL      string `json:"releases_url"`
			DeploymentsURL   string `json:"deployments_url"`
		} `json:"repository"`
		Score float64 `json:"score"`
	} `json:"items"`
}

func Streak(user string) string {
	url := "https://api.github.com/search/commits?q=author:" + user + "&sort=author-date&order=desc&page=1"
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

	var times []time.Time
	for _, v := range resObjectAPI.Items {
		t := v.Commit.Author.Date
		times = append(times, t)
		days := t2.Sub(t1).Hours() / 24
	}
	return ``
}
