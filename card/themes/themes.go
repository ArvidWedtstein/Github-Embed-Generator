package themes

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Theme struct {
	Name       string `json:"name"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	Border     string `json:"border"`
	Background string `json:"background"`
	Box        string `json:"box"`
	Font       string `json:"font"`
}

func LoadTheme(themeName string) Theme {
	jsonFile, err := os.Open("card/themes/themes.json")

	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var themes []Theme
	json.Unmarshal(byteValue, &themes)

	var selectedTheme Theme
loop:
	for _, theme := range themes {
		if theme.Name == themeName {
			selectedTheme = theme
			break loop
		}
	}
	return selectedTheme
}
