package languageCard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Data struct {
	Data struct {
		User struct {
			Repositories struct {
				Nodes []struct {
					Name      string `json:"name"`
					Languages struct {
						Edges []struct {
							Size int `json:"size"`
							Node struct {
								Color string `json:"color"`
								Name  string `json:"name"`
							}
						} `json:"edges"`
					} `json:"languages"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
}

type Languages struct {
	Size  int    `json:"size"`
	Color string `json:"color"`
}

func LanguageCard(title, user string, cardstyle style.Styles) string {

	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{
		user(login: "%v") {
			repositories(ownerAffiliations: OWNER, isFork: false, privacy: PUBLIC, first: 100) {
			  nodes {
				name
				languages(first: 10, orderBy: {field: SIZE, direction: DESC}) {
				  edges {
					size
					node {
					  color
					  name
					}
				  }
				}
			  }
			}
		  }
		}
		`, user),
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
	responseData, _ := ioutil.ReadAll(response.Body)

	var data Data
	json.Unmarshal(responseData, &data)

	sum := func(values map[string]Languages) int {
		var sum int
		for _, l := range values {
			sum += l.Size
		}
		return sum
	}
	customstyles := []string{}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
	}

	padding := 30
	body := []string{
		fmt.Sprintf(`<g id="Box"><rect x="0" y="0" rx="15" fill="%v" width="%v" height="%v" /></g>`, cardstyle.Background, 700, 200),
		`<g data-testid="card-text">`,
		fmt.Sprintf(`<text x="%v" y="%v" id="Stats" class="title">%v</text>`, padding, padding, card.ToTitleCase(title)),
		fmt.Sprintf(`<line id="gradLine" x1="%v" y1="40" x2="400" y2="40" stroke="url(#paint0_angular_0_1)"/>`, padding),
		`</g>`,
	}

	// gridX := 30
	// gridY := 100
	// gridYstartPos := 100

	languages := map[string]Languages{}
	for _, v := range data.Data.User.Repositories.Nodes {
		if len(v.Languages.Edges) > 0 {
			for _, langs := range v.Languages.Edges {
				if _, ok := languages[langs.Node.Name]; ok {
					languages[langs.Node.Name] = Languages{Size: languages[langs.Node.Name].Size + langs.Size, Color: langs.Node.Color}
				} else {
					languages[langs.Node.Name] = Languages{Size: langs.Size, Color: langs.Node.Color}
				}
			}
		}
	}
	content := []string{}
	for name, l := range languages {
		content = append(content, fmt.Sprintf(`
		<g height="80" x="" y="">
			<circle cx="5" cy="6" r="5" fill="" />
			<text data-testid="lang-name" x="15" y="10" class='lang-name'>%v - %v</text>
    	</g>`, name, card.CalculatePercentFloat(l.Size, sum(languages))),
		)
		fmt.Printf("%v | Percent %v%v\n", name, card.CalculatePercentFloat(l.Size, sum(languages)), "%")
	}
	body = append(body, card.VerticalFlexBox(300, 20, 20, padding, content))

	return strings.Join(card.GenerateCard(cardstyle, defs, body, 800, 300, customstyles...), "\n")
}
