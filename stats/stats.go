package stats

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
	"strings"
	"time"
)

type Data struct {
	Data struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions            int `json:"totalCommitContributions"`
				TotalIssueContributions             int `json:"totalIssueContributions"`
				TotalPullRequestContributions       int `json:"totalPullRequestContributions"`
				TotalPullRequestReviewContributions int `json:"totalPullRequestReviewContributions"`
			} `json:"contributionsCollection"`
			Repositories struct {
				Nodes []struct {
					Name           string `json:"name"`
					StargazerCount int    `json:"stargazerCount"`
					DiskUsage      int    `json:"diskUsage"`
					ForkCount      int    `json:"forkCount"`
					Milestones     struct {
						TotalCount int `json:"totalCount"`
					} `json:"milestones"`
					Packages struct {
						TotalCount int `json:"totalCount"`
					} `json:"packages"`
					PullRequests struct {
						TotalCount int `json:"totalCount"`
					} `json:"pullRequests"`
					Releases struct {
						TotalCount int `json:"totalCount"`
					} `json:"releases"`
					Watchers struct {
						TotalCount int `json:"totalCount"`
					} `json:"watchers"`
					Issues struct {
						TotalCount int `json:"totalCount"`
					} `json:"issues"`
				} `json:"nodes"`
			} `json:"repositories"`
			RepositoriesContributedTo struct {
				Nodes []struct {
					NameWithOwner string `json:"nameWithOwner"`
					Owner         struct {
						Login string `json:"login"`
					} `json:"owner"`
					IsInOrganization bool `json:"isInOrganization"`
				} `json:"nodes"`
				TotalCount int `json:"totalCount"`
			} `json:"repositoriesContributedTo"`
		} `json:"user"`
	} `json:"data"`
}

func Stats(title string, user string, cardstyle themes.Theme) string {

	year := time.Now().Year()
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{
			user(login: "%v") {
                contributionsCollection(from: "%v-01-01T00:00:00Z", to: "%v-12-31T23:59:59Z") {
                    totalCommitContributions
					totalIssueContributions
					totalPullRequestContributions
					totalPullRequestReviewContributions
                }
				repositories(last: 100, isFork: false, affiliations: OWNER, privacy: PUBLIC) {
					nodes {
					  name
					  stargazerCount
					  diskUsage
					  forkCount
					  milestones {
						totalCount
					  }
					  packages {
						totalCount
					  }
					  pullRequests {
						totalCount
					  }
					  releases {
						totalCount
					  }
					  watchers {
						totalCount
					  }
					  issues {
						  totalCount
					  }
					}
				  }
				repositoriesContributedTo(
					first: 100
					contributionTypes: [COMMIT, ISSUE, PULL_REQUEST, REPOSITORY]
				  ) {
					nodes {
					  	nameWithOwner
						owner {
							login
						}
						isInOrganization
					}
					totalCount
				}
            }
		}
		`, user, year, year),
	}

	jsonValue, _ := json.Marshal(jsonData)

	// Make request
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
	var data Data
	json.Unmarshal(responseData, &data)

	if title == "" || len(title) <= 0 {
		title = "Stats"
	}

	// TODO
	//
	// Total Contributions - DONE
	// Organizations Contributed To - DONE
	// People Contributed To
	// Repositories Contributed To - DONE
	// Issue Count (Open issues only??)
	// Disk Usage
	// Stargazer Count
	// Milestones Count
	// Packages Count
	// Fork Count
	// Releases Count
	// Watchers Count - DONE
	// Pull Request Count - DONE

	contributionsCollection := data.Data.User.ContributionsCollection
	totalContributions := contributionsCollection.TotalCommitContributions

	totalContributions += contributionsCollection.TotalIssueContributions
	totalContributions += contributionsCollection.TotalPullRequestContributions
	totalContributions += contributionsCollection.TotalPullRequestReviewContributions
	repositoriesContributedTo := data.Data.User.RepositoriesContributedTo.TotalCount

	orgsContributedTo := []string{}
	for _, v := range data.Data.User.RepositoriesContributedTo.Nodes {
		if !card.ArrayContains(orgsContributedTo, v.Owner.Login) && v.IsInOrganization {
			orgsContributedTo = append(orgsContributedTo, v.Owner.Login)
		}
	}
	fmt.Printf("Repos Contributed To: %v\n", repositoriesContributedTo)
	fmt.Printf("Orgs Contributed to: %v\n", len(orgsContributedTo))

	totalMilestones := 0
	totalPackages := 0
	totalForks := 0
	totalReleases := 0
	totalWatchers := 0
	totalStargazers := 0
	totalDiskUsage := 0
	totalPullRequests := 0
	totalIssues := 0

	for _, v := range data.Data.User.Repositories.Nodes {
		totalMilestones += v.Milestones.TotalCount
		totalPackages += v.Packages.TotalCount
		totalForks += v.ForkCount
		totalReleases += v.Releases.TotalCount
		totalWatchers += v.Watchers.TotalCount
		totalStargazers += v.StargazerCount
		totalDiskUsage += v.DiskUsage
		totalPullRequests += v.PullRequests.TotalCount
		totalIssues += v.Issues.TotalCount
	}

	fmt.Printf("TotalContributions: %v\n", totalContributions)
	fmt.Printf("totalMilestones: %v\n", totalMilestones)
	fmt.Printf("totalPackages: %v\n", totalPackages)
	fmt.Printf("totalForks: %v\n", totalForks)
	fmt.Printf("totalReleases: %v\n", totalReleases)
	fmt.Printf("totalWatchers: %v\n", totalWatchers)
	fmt.Printf("totalStargazers: %v\n", totalStargazers)
	fmt.Printf("totalDiskUsage: %v\n", totalDiskUsage)
	fmt.Printf("totalPullRequests: %v\n", totalPullRequests)
	fmt.Printf("totalIssues: %v\n", totalIssues)
	height := 700
	width := 600
	titleboxheight := 50
	// padding := 10
	strokewidth := 3

	customstyles := []string{
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
	}
	defs := []string{
		style.LinearGradient("gradient-fill", 0, []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
	}

	body := []string{
		fmt.Sprintf(`<text x="20" y="35" class="title">%s</text>`, card.ToTitleCase(title)),
	}

	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}
	bodyAdd(`<g></g>`)

	// Line on top
	body = append([]string{fmt.Sprintf(`<rect x="0" y="%v" width="%v" height="%v" fill="%v"/>`, titleboxheight, width, strokewidth, cardstyle.Colors.Border)}, body...)

	return strings.Join(card.GenerateCard(cardstyle, defs, body, width+strokewidth, height+strokewidth, true, customstyles...), "\n")
}
