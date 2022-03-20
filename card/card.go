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
func FlexBox(heightWidth, posX, posY, padding int, content []string, horizontal bool) string {
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
		if strings.Contains(item, "<g") {
			item = strings.ReplaceAll(item, `translate(0,0)`, fmt.Sprintf(`translate(%v,%v)`, x, y))
		} else {
			item = strings.ReplaceAll(item, `x=""`, fmt.Sprintf(`x="%v"`, x))
			item = strings.ReplaceAll(item, `y=""`, fmt.Sprintf(`y="%v"`, y))
		}
		gridLayout = append(gridLayout, item)

		/* Switch between Horizontal and vertical flexbox */
		if horizontal {
			if x+((itemwidth*2)+padding) >= heightWidth {
				y += itemheight + padding
				x = posX - (itemwidth + padding)
			}
			x += itemwidth + padding
		} else {
			if y+((itemheight*2)+padding) >= heightWidth {
				x += itemwidth + padding
				y = posY - (itemheight + padding)
			}
			y += itemheight + padding
		}

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
// Charts Section
// ----------------------------

type Radar struct {
	Labels []string `json:"labels"`
	Values []int    `json:"values"`
	Color  string   `json:"color"`
	Grid   bool
	Width  int
	Height int
}

func (radar Radar) generateGrid() string {
	centerX := (radar.Width / 2)
	centerY := (radar.Height / 2)
	var rings int = 7

	gridgap := 20
	path := []string{}
	labels := []string{}

	sectionDegree := 360 / (len(radar.Values))
	for ring := 1; ring <= rings; ring++ {
		// Create a new ring
		for point := 1; point <= len(radar.Values); point++ {
			// Calculate degree to radian. Then do some magic and get coordinates for points
			x := (float64(ring*gridgap) * math.Cos(DegreeToRadians(float64(sectionDegree*point))))
			y := (float64(ring*gridgap) * math.Sin(DegreeToRadians(float64(sectionDegree*point))))
			x2 := (float64(ring*gridgap) * math.Cos(DegreeToRadians(float64(sectionDegree*(point+1)))))
			y2 := (float64(ring*gridgap) * math.Sin(DegreeToRadians(float64(sectionDegree*(point+1)))))
			x3 := (float64((ring-1)*gridgap) * math.Cos(DegreeToRadians(float64(sectionDegree*(point+1)))))
			y3 := (float64((ring-1)*gridgap) * math.Sin(DegreeToRadians(float64(sectionDegree*(point+1)))))

			if len(radar.Labels) == len(radar.Values) && ring == rings {
				labels = append(labels, fmt.Sprintf(`<text x="%v" y="%v" text-anchor="middle">%v</text>`, float64(centerX)+x, float64(centerY)+y, radar.Labels[point-1]))
			}
			// Create path values
			path = append(path, fmt.Sprintf(`M %v %v L %v %v M %v %v L %v %v`,
				float64(centerX)+x, float64(centerY)+y,
				float64(centerX)+x2, float64(centerY)+y2,
				float64(centerX)+x2, float64(centerY)+y2,
				float64(centerX)+x3, float64(centerY)+y3,
			))
		}

	}
	return fmt.Sprintf(`<path d="%v" stroke="#333333" stroke-width="1"/>%v`, strings.Join(path, " "), strings.Join(labels, "\n"))
}
func RadarChart(radar Radar) string {

	centerX := (radar.Width / 2)
	centerY := (radar.Height / 2)
	sectionDegree := 360 / (len(radar.Values))
	radarChart := []string{
		fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %v %v" width="%v" height="%v" version="1"><g width="%v" height="%v">`, radar.Width, radar.Height, radar.Width, radar.Height, radar.Width, radar.Height),
	}

	if radar.Grid {
		radarChart = append(radarChart, radar.generateGrid())
	}

	// Generate Data Chart
	pathData := []string{}
	for i := 0; i < len(radar.Values); i++ {

		// Calculate X and Y pos
		x := ToFixed(float64(radar.Values[i])*math.Cos(DegreeToRadians(float64(sectionDegree*(i+1)))), 2)
		y := ToFixed(float64(radar.Values[i])*math.Sin(DegreeToRadians(float64(sectionDegree*(i+1)))), 2)

		// Move to first position
		// if i == 0 {
		// 	pathData = append(pathData, fmt.Sprintf(`M %v %v`, float64(centerX)+x, float64(centerY)+y))
		// }
		pathData = append(pathData, fmt.Sprintf(`%v,%v`, float64(centerX)+x, float64(centerY)+y))
		radarChart = append(radarChart, fmt.Sprintf(`<circle cx="%v" cy="%v" fill="%v" r="3" />`, float64(centerX)+x, float64(centerY)+y, radar.Color))
	}
	radarChart = append(radarChart, fmt.Sprintf(`<polygon points="%v" fill="%v" opacity="0.5" stroke="%v" stroke-width="3"/>`, strings.Join(pathData, " "), radar.Color, radar.Color))
	radarChart = append(radarChart, `</g></svg>`)
	return strings.Join(radarChart, " ")
}

type Bar struct {
	Labels []string `json:"labels"`
	Values []int    `json:"values"`
	Colors []string `json:"colors"`
	Grid   bool
	Width  int
	Height int
}

func BarChart(bar Bar) string {
	barChart := []string{
		fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %v %v" width="%v" height="%v" version="1"><g>`,
			bar.Width, bar.Height, bar.Width, bar.Height),
	}
	// _, max := FindMinAndMax(bar.Values)
	barGap := 10
	barChart = append(barChart, fmt.Sprintf(`<path d="M 10 10 L 10 %v" stroke="#333333" stroke-width="3" />`, bar.Height))

	content := []string{}
	for _, v := range bar.Values {
		fmt.Println(v)
		content = append(content, `<rect x="" y="" width="30" height="50" fill="#ff0000"/>`)
	}
	barChart = append(barChart, FlexBox(bar.Width, barGap, bar.Height-200, barGap, content, true))
	barChart = append(barChart, `</g></svg>`)
	return strings.Join(barChart, "\n")
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
		var startX, startY = getCoordinatesForPercent(cumulativePercent/float64(100), radius)
		cumulativePercent += slice.Percent

		var endX, endY = getCoordinatesForPercent(cumulativePercent/float64(100), radius)

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
func getCoordinatesForPercent(percent float64, radius int) (float64, float64) {
	// gets coordinates along the edge of a circle
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
func DegreeToRadians(degree float64) float64 {
	return degree * math.Pi / float64(180)
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
