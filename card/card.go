package card

import (
	"fmt"
	"githubembedapi/card/style"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Title string       `json:"title"`
	Style style.Styles `json:"colors"`
	Body  []string     `json:"body"`
}

func (card Card) GetStyles(customStyles ...string) string {
	var style = []string{
		`<style>`,
		`.title { font: 25px sans-serif; fill: #` + card.Style.Title + `}`,
		`.text { font: 20px sans-serif; fill: #` + card.Style.Text + `; font-family: ` + card.Style.Textfont + `;}`,
	}
	if cap(customStyles) > 0 {
		style = append(style, customStyles...)
	}

	style = append(style, `</style>`)
	return strings.Join(style, "\n")
}
func (card Card) GetScript(customScripts ...string) string {
	var style = []string{
		`<script>`,
	}
	if cap(customScripts) > 0 {
		style = append(style, customScripts...)
	}

	style = append(style, `</script>`)
	return strings.Join(style, "\n")
}
func (card Card) GetDefs(customDefinitions []string) string {
	var defs = []string{
		`<defs>`,
	}
	if cap(customDefinitions) > 0 {
		defs = append(defs, customDefinitions...)
	}

	defs = append(defs, `</defs>`)
	return strings.Join(defs, "\n")
}
func CircleProgressbar(progress, radius, strokewidth, posX, posY int, color string, class ...string) (string, string) {
	dasharray := (2 * math.Pi * float64(radius))

	if progress < 0 {
		progress = 0
	}
	if progress > 100 {
		progress = 100
	}

	dashoffset := ((100 - float64(progress)) / 100) * dasharray
	progressbar := fmt.Sprintf(`<circle stroke-linecap="butt" filter="filter0_d_0_1" style="animation: CircleProgressbar%v 3s forwards ease-in-out; z-index: 2;" class="%v" cx="%v" cy="%v" r="%v" fill="transparent" stroke="%v" stroke-width="%v" stroke-dasharray="%v" stroke-dashoffset="%v"/>`,
		radius, strings.Join(class, " "), posX, posY, radius, color, strokewidth, dasharray, dashoffset)
	return progressbar, GetProgressAnimation(progress, radius)
}
func GetProgressAnimation(progress, radius int) string {
	dasharray := (2 * math.Pi * float64(radius))

	if progress < 0 {
		progress = 0
	}
	if progress > 100 {
		progress = 100
	}

	dashoffset := ((100 - float64(progress)) / 100) * dasharray
	return `@keyframes CircleProgressbar` + strconv.Itoa(radius) + ` { 
	from { 
		stroke-dashoffset: ` + strconv.Itoa(int(dasharray)) + `
	}
	to { 
		stroke-dashoffset: ` + strconv.Itoa(int(dashoffset)) + `
	}
}`
}

func GenerateCard(style style.Styles, defs []string, body []string, width, height int, customStyles ...string) []string {
	var card Card
	card.Style = style

	card.Body = []string{
		fmt.Sprintf(`<svg width="%v" height="%v" viewBox="0 0 %v %v" xmlns="http://www.w3.org/2000/svg">`, width, height, width, height),
		card.GetStyles(customStyles...),
		card.GetDefs(defs),
		strings.Join(body, "\n"),
		`</svg>`,
	}
	return card.Body
}

func ToTitleCase(str string) string {
	return strings.Title(str)
}
