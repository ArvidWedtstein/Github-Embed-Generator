package rank

import (
	"encoding/json"
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/themes"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type User struct {
	Avatar string
	Score  int
	Name   string
}
type Kv struct {
	Key   string
	Value User
}

type Response struct {
	Total_Count int     `json:"total_count"`
	Items       []Items `json:"items"`
}
type Items struct {
	Url          string `json:"url"`
	Comments_url string `json:"comments_url"`
	Author       Author `json:"author"`
}
type Author struct {
	Login      string `json:"login"`
	Avatar_Url string `json:"avatar_url"`
}

func Rankcard(title string, users []string, cardstyle themes.Theme) string {
	if title == "" || len(title) < 1 {
		title = "Rank"
	}
	if len(users) > 5 {
		users = users[:5]
	}

	customstyles := []string{
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.text { font: 20px sans-serif; fill: ` + cardstyle.Colors.Text + `; font-family: ` + cardstyle.Font + `; text-decoration: underline;}`,
		`.large { font: 25px sans-serif; fill: black}`,
		`.title { font: 25px sans-serif; fill: ` + cardstyle.Colors.Title + `}`,
		`.box { fill: ` + cardstyle.Colors.Background + `}`,
		`.profileimage { border-radius: 50%}`,
	}
	defs := []string{}
	ss := make(map[string]User)
	for key, i := range users {
		userurl := "https://api.github.com/search/commits?q=author:" + fmt.Sprintf("%v", i) + "&sort=author-date&order=desc&page=1"
		response, err := http.Get(userurl)
		if err != nil {
			panic(err)
		}
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		var responseObject Response
		decodeerr := json.Unmarshal(responseData, &responseObject)

		if decodeerr != nil {
			panic(decodeerr)
		}

		ss[fmt.Sprintf("%v", users[key])] = User{Avatar: responseObject.Items[0].Author.Avatar_Url, Score: responseObject.Total_Count, Name: responseObject.Items[0].Author.Login}
	}

	// Sort Scores
	var score []Kv
	for k, v := range ss {
		score = append(score, Kv{k, v})
	}

	// Sort score
	sort.Slice(score, func(i, j int) bool {
		return score[i].Value.Score > score[j].Value.Score
	})

	totalHeight := 40
	width := 400
	strokewidth := 3

	body := []string{
		fmt.Sprintf(`<rect x="0" y="0" class="box" width="%v" height="200" rx="15" style="stroke-width:3;stroke:%v"/>`, width, cardstyle.Colors.Border),
		fmt.Sprintf(`<rect x="0" y="30" width="%v" height="3" fill="%v"/>`, width, cardstyle.Colors.Border),
		fmt.Sprintf(`<text x="20" y="25" class="title">%s</text>`, card.ToTitleCase(title)),
	}
	// Generate body for the users
	pos := 1
	for _, s := range score {
		var rowx int = 20
		img := fmt.Sprintf(`<image x="%v" y="%v" href="%v" class="profileimage" height="30" width="30"/>`, rowx, totalHeight, s.Value.Avatar)
		text := fmt.Sprintf(`<text x="%v" y="%v" class="text">%v. %v - %v commits</text>`, rowx+40, totalHeight+20, pos, card.ToTitleCase(s.Value.Name), s.Value.Score)
		totalHeight += 30
		pos += 1
		body = append(body, text)
		body = append(body, img)
	}
	return strings.Join(card.GenerateCard(cardstyle, defs, body, width+strokewidth, totalHeight+180, customstyles...), "\n")
}
