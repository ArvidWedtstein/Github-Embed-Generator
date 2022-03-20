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
		`.title { font: 25px sans-serif; fill: ` + card.Style.Title + `}`,
		`.text { font: 16px sans-serif; fill: ` + card.Style.Text + `; font-family: ` + card.Style.Textfont + `;}`,
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

/*
yAY. I reinvented flexbox
*/
func HorizontalFlexBox(width, posX, posY, padding int, content []string) string {
	x := posX
	y := posY
	// Make new group
	gridLayout := []string{fmt.Sprintf(`<g transform="translate(%v,%v)">`, posX, posY)}
	for _, item := range content {
		items := strings.Split(item, ` `)
		var itemwidth int
		var itemheight int
		for _, v := range items {
			// Get Width
			if strings.Contains(v, `width="`) {
				v = strings.ReplaceAll(v, `width=`, ``)
				v = strings.ReplaceAll(v, `"`, ``)
				iwidth, err := strconv.Atoi(v)
				itemwidth = iwidth
				if err != nil {
					panic(err)
				}
			}
			// Get Height
			if strings.Contains(v, `height="`) {
				v = strings.ReplaceAll(v, `height=`, ``)
				v = strings.ReplaceAll(v, `"`, ``)
				iheight, err := strconv.Atoi(v)
				itemheight = iheight
				if err != nil {
					panic(err)
				}
			}
		}
		item = strings.Replace(item, `x=""`, fmt.Sprintf(`x="%v"`, x), 1)
		item = strings.Replace(item, `y=""`, fmt.Sprintf(`y="%v"`, y), 1)
		gridLayout = append(gridLayout, item)

		if x+((itemwidth*2)+padding) >= width {
			y += itemheight + padding
			x = posX - (itemwidth + padding)
		}
		x += itemwidth + padding

	}
	gridLayout = append(gridLayout, `</g>`)
	return strings.Join(gridLayout, "\n")
}
func VerticalFlexBox(height, posX, posY, padding int, content []string) string {
	x := posX
	y := posY
	gridLayout := []string{fmt.Sprintf(`<g height="300" transform="translate(%v,%v)">`, posX, posY)}
	for _, item := range content {
		items := strings.Split(item, ` `)
		var itemwidth int
		var itemheight int
		for _, v := range items {

			// Get Width
			if strings.Contains(v, `width="`) {
				v = strings.ReplaceAll(v, `width=`, ``)
				v = strings.ReplaceAll(v, `"`, ``)
				iwidth, err := strconv.Atoi(v)
				itemwidth = iwidth
				if err != nil {
					panic(err)
				}
			}

			// Get Height
			if strings.Contains(v, `height="`) {
				v = strings.ReplaceAll(v, `height=`, ``)
				v = strings.ReplaceAll(v, `"`, ``)
				iheight, err := strconv.Atoi(v)
				itemheight = iheight
				if err != nil {
					panic(err)
				}
			}
		}
		if strings.Contains(item, "<g") {
			item = strings.ReplaceAll(item, `translate(0,0)`, fmt.Sprintf(`translate(%v,%v)`, x, y))
		} else {
			item = strings.ReplaceAll(item, `x=""`, fmt.Sprintf(`x="%v"`, x))
			item = strings.ReplaceAll(item, `y=""`, fmt.Sprintf(`y="%v"`, y))
		}

		gridLayout = append(gridLayout, item)

		if y+((itemheight*2)+padding) >= height {
			x += itemwidth + padding
			y = posY - (itemheight + padding)
		}
		y += itemheight + padding

	}
	gridLayout = append(gridLayout, `</g>`)

	return strings.Join(gridLayout, "\n")
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
	progressbar := fmt.Sprintf(`<g data-testid="rank-circle" transform="translate(0,0)"><circle class="rank-circle-rim" cx="0" cy="0" r="%v" fill="transparent" stroke="#000000" stroke-width="%v"/><circle stroke-linecap="butt" filter="filter0_d_0_1" style="animation: CircleProgressbar%v 3s forwards ease-in-out;" class="%v" cx="%v" cy="%v" r="%v" fill="transparent" stroke="%v" stroke-width="%v" stroke-dasharray="%v" stroke-dashoffset="%v"/></g>`,
		radius, strokewidth, radius, strings.Join(class, " "), posX, posY, radius, color, strokewidth, dasharray, dashoffset)
	return progressbar, GetProgressAnimation(progress, radius)
}

// ----------------------------
// Charts
// ----------------------------

type Radar struct {
	Name   string `json:"name"`
	Values []int  `json:"values"`
	Color  string `json:"color"`
	Grid   bool
	Width  int
	Height int
}

func RadarChart(radar Radar) string {
	radarChart := []string{
		fmt.Sprintf(`<svg version="1" xmlns="http://www.w3.org/2000/svg" width="%v" height="%v">`, radar.Width, radar.Height),
		fmt.Sprintf(`<rect width="%v" fill="#ff0000" height="%v" x="0" y="0" />`, radar.Width, radar.Height),
	}

	pathData := []string{fmt.Sprintf(`M %v %v`, (radar.Width / 2), (radar.Height/2)+radar.Values[0])}
	for i := 0; i < len(radar.Values); i++ {
		if i == 0 {
			pathData = append(pathData, fmt.Sprintf(`L %v %v`, (radar.Width/2)+radar.Values[i+1], (radar.Height/2)))
		} else if i == 1 {
			pathData = append(pathData, fmt.Sprintf(`L %v %v`, (radar.Width/2), (radar.Height/2)-radar.Values[i+1]))
		} else if i == 2 {
			pathData = append(pathData, fmt.Sprintf(`L %v %v`, (radar.Width/2)-radar.Values[i+1], (radar.Height/2)))
		} else if i == len(radar.Values)-1 {
			pathData = append(pathData, fmt.Sprintf(`L %v %v`, (radar.Width/2), (radar.Height/2)+radar.Values[0]))
		}
	}
	radarChart = append(radarChart, fmt.Sprintf(`<path d="%v" fill="%v" stroke="%v" stroke-width="3"/>`, strings.Join(pathData, " "), radar.Color, radar.Color))
	radarChart = append(radarChart, `</svg>`)
	return strings.Join(radarChart, " ")
}
func BarChart() string {
	return ``
}
func LineChart() string {
	return ``
}

type PieChartSlice struct {
	Name    string  `json:"name"`
	Percent float64 `json:"percent"`
	Color   string  `json:"color"`
}

func PieChart(slices []PieChartSlice, radius, posX, posY int, color string) string {
	var cumulativePercent float64 = 0
	paths := []string{}
	for _, slice := range slices {
		var startX, startY = GetCoordinatesForPercent(cumulativePercent/float64(100), radius)
		cumulativePercent += slice.Percent

		var endX, endY = GetCoordinatesForPercent(cumulativePercent/float64(100), radius)

		var largeArcFlag = 0
		if slice.Percent > 50 {
			largeArcFlag = 1
		}

		var pathData2 = []string{
			fmt.Sprintf(`M %v %v`, startX, startY),                                        // Move
			fmt.Sprintf(`A %v %v 0 %v 1 %v %v`, radius, radius, largeArcFlag, endX, endY), // Arc
			`L 0 0`, // Line
		}
		pathData := strings.Join(pathData2, " ")

		paths = append(paths, fmt.Sprintf(`<path d="%v" fill="%v"></path>`, pathData, slice.Color))
	}
	piechart := fmt.Sprintf(`
	<g transform="translate(%v,%v)">
		<circle cx="0" cy="0" r="%v" fill="%v" class="circle" stroke="#333333" stroke-width="5"/>
		<g class="circle" transform="translate(0,0)">
		%v
		</g>
	</g>`, posX, posY, radius, color, strings.Join(paths, "\n"))
	return piechart
}
func GetCoordinatesForPercent(percent float64, radius int) (float64, float64) {
	var x = math.Cos(float64(2) * math.Pi * percent)
	var y = math.Sin(float64(2) * math.Pi * percent)
	return ToFixed(x*float64(radius), 8), ToFixed(y*float64(radius), 8)
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

// ----------------------------
// Functional Functions
// ----------------------------

func ToTitleCase(str string) string {
	return strings.Title(str)
}

func CalculatePercent(number, total int) int {
	return int((float64(number) / float64(total)) * float64(100))
}
func CalculatePercentFloat(number, total int) float64 {
	return ToFixed((float64(number)/float64(total))*float64(100), 2)
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func Sum(list []int) int {
	total := 0
	for _, v := range list {
		total += v
	}
	return total
}
func Average(list []int) int {
	return Sum(list) / len(list)
}

func FindMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
