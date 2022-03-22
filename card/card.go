package card

import (
	"fmt"
	"githubembedapi/card/style"
	"githubembedapi/card/themes"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Title string       `json:"title"`
	Style themes.Theme `json:"colors"`
	Body  []string     `json:"body"`
}

func (card Card) GetStyles(customStyles ...string) string {
	var style = []string{
		`<style>`,
		`.title { font: 25px sans-serif; fill: ` + card.Style.Title + `}`,
		`.text { font: 16px sans-serif; fill: ` + card.Style.Text + `; font-family: ` + card.Style.Font + `;}`,
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
	gridLayout := []string{fmt.Sprintf(`<g data-testid="flexbox" transform="translate(%v,%v)">`, posX, posY)}
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

		pathData = append(pathData, fmt.Sprintf(`%v,%v`, float64(centerX)+x, float64(centerY)+y))
		radarChart = append(radarChart, fmt.Sprintf(`<circle cx="%v" cy="%v" fill="%v" r="3" />`, float64(centerX)+x, float64(centerY)+y, radar.Color))
	}
	radarChart = append(radarChart, fmt.Sprintf(`<polygon points="%v" fill="%v" opacity="0.5" stroke="%v" stroke-width="3"/>`, strings.Join(pathData, " "), radar.Color, radar.Color))
	radarChart = append(radarChart, `</g></svg>`)
	return strings.Join(radarChart, " ")
}

type Bar struct {
	Labels   []string `json:"labels"`
	Values   []int    `json:"values"`
	Colors   []string `json:"colors"`
	Grid     bool
	Vertical bool
	Width    int
	Height   int
}

func BarChartVertical(bar Bar) string {
	_, max := FindMinAndMax(bar.Values)

	barGap := 10
	columnWidth := 30
	totalWidth := (columnWidth + barGap) * len(bar.Values)

	barChart := []string{
		fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %v %v" width="%v" height="%v" version="1"><g>`,
			totalWidth, max+120, totalWidth, max+120),
	}
	BarChartAdd := func(content string) {
		barChart = append(barChart, content)
	}

	// Grid
	BarChartAdd(fmt.Sprintf(`<g data-testid="grid"><path d="M %v %v V %v H %v" fill="none" stroke="#333333" stroke-width="3" stroke-linecap="round"/>`,
		barGap*2, barGap, (max + 30), totalWidth))
	if bar.Grid {
		path := []string{}
		for i := 0; i < ((max / 10) + 3); i++ {
			path = append(path, fmt.Sprintf(`M %v %v L %v %v`,
				totalWidth, (barGap*i)+10, barGap*2, (barGap*i)+10))

			// Generate numbers along axis
			if (i*10)%4 == 0 {
				BarChartAdd(fmt.Sprintf(`<text style="font-size: 12px" text-anchor="middle" class="text" fill="#333333" x="%v" y="%v">%v</text>`,
					10, max-(barGap*i)+30, (i * 10)))
			}
		}
		BarChartAdd(fmt.Sprintf(`<path d="%v" stroke="#333333" opacity="0.3" stroke-width="1" />`, strings.Join(path, " ")))
	}
	BarChartAdd(`</g>`)

	// Add columns
	content := []string{}
	for i, v := range bar.Values {
		if len(bar.Colors) > i {
			if len(bar.Labels) > i {
				content = append(content, fmt.Sprintf(`<g transform="translate(%v,%v)" class="bar"><rect width="%v" height="%v" fill="#%v"/><text x="%v" y="%v" fill="#000000" text-anchor="middle">%v</text></g>`,
					(20+barGap)*i, (max+30)-v, v, v, bar.Colors[i], (v+20)+len(bar.Labels[i]), -5, bar.Labels[i]))
			} else {
				content = append(content, fmt.Sprintf(`<g transform="translate(%v,%v)" class="bar"><rect width="%v" height="%v" fill="#%v"/></g>`,
					(20+barGap)*i, (max+30)-v, v, v, bar.Colors[i]))
			}
		} else {
			if len(bar.Labels) > i {
				content = append(content, fmt.Sprintf(`<g transform="translate(%v,%v)" class="bar"><rect width="%v" height="%v" fill="%v"/><text transform="rotate(90)" x="%v" y="%v" fill="#000000" text-anchor="middle">%v</text></g>`,
					(20+barGap)*i, (max+30)-v, 20, v, "#ff0000", (v+20)+len(bar.Labels[i]), -5, bar.Labels[i]))
			} else {
				content = append(content, fmt.Sprintf(`<g transform="translate(%v,%v)" class="bar"><rect width="%v" height="%v" fill="%v"/></g>`,
					(20+barGap)*i, (max+30)-v, 20, v, "#ff0000"))
			}

		}
	}

	// Generate Row for columns
	BarChartAdd(FlexBox(bar.Height, barGap*3, 0, barGap, content, true))

	BarChartAdd(`</g></svg>`)
	return strings.Join(barChart, "\n")
}

// Bar Chart Horizontal
func BarChartHorizontal(bar Bar) string {

	_, max := FindMinAndMax(bar.Values)

	barGap := 10
	columnHeight := 30

	// Calculate total height of all columns including padding between columns
	totalHeight := (columnHeight + barGap) * len(bar.Values)

	barChart := []string{
		fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %v %v" width="%v" height="%v" version="1"><g>`,
			((max/10)+3)*10, totalHeight+60, ((max/10)+3)*10, totalHeight+60),
	}
	BarChartAdd := func(content string) {
		barChart = append(barChart, content)
	}

	// ----------------------------
	// Grid
	// ----------------------------
	BarChartAdd(fmt.Sprintf(`<g data-testid="grid"><path d="M %v %v V %v H %v" fill="none" stroke="#333333" stroke-width="3" stroke-linecap="round"/>`,
		barGap, barGap, totalHeight+20, ((max/10)+3)*10))
	if bar.Grid {
		path := []string{}
		for i := 0; i < (max/10)+3; i++ {
			path = append(path, fmt.Sprintf(`M %v %v L %v %v`,
				(barGap*i)+10, totalHeight+20, (barGap*i)+10, barGap))

			// Lines at bottom
			path = append(path, fmt.Sprintf(`M %v %v L %v %v`,
				(barGap*i)+10, totalHeight+20, (barGap*i)+10, totalHeight+30))

			// Generate numbers along axis
			if (i*10)%4 == 0 {
				BarChartAdd(fmt.Sprintf(`<text style="font-size: 12px" text-anchor="middle" class="text" fill="#333333" x="%v" y="%v">%v</text>`,
					(barGap*i)+10, totalHeight+40, (i * 10)))
			}
		}
		BarChartAdd(fmt.Sprintf(`<path d="%v" stroke="#333333" opacity="0.3" stroke-width="1" />`, strings.Join(path, " ")))
	}
	BarChartAdd(`</g>`)

	// ----------------------------
	// Add Columns
	// ----------------------------
	content := []string{}
	for i, v := range bar.Values {
		if len(bar.Colors) > i {
			content = append(content, fmt.Sprintf(`<g transform="translate(0,0)" class="bar"><rect width="%v" height="%v" fill="#%v"/><text x="%v" y="20" fill="#000000" text-anchor="middle">%v</text></g>`, v, columnHeight, bar.Colors[i], v+10, v))
		} else {
			content = append(content, fmt.Sprintf(`<g transform="translate(0,0)" class="bar"><rect width="%v" height="%v" fill="%v"/><text x="%v" y="20" fill="#000000" text-anchor="middle">%v</text></g>`, v, columnHeight, "#ff0000", v+10, v))
		}
	}
	// Generate Row for columns
	BarChartAdd(FlexBox(bar.Height, 5, barGap, barGap, content, false))

	BarChartAdd(`</g></svg>`)
	return strings.Join(barChart, "\n")
}

type Line struct {
	Values         []int `json:"values"`
	GridVertical   bool
	GridHorizontal bool
	Color          string
	Width          int
	Height         int
}

func LineChart(line Line) string {
	_, max := FindMinAndMax(line.Values)
	width := len(line.Values) * 20
	padding := 10
	lineChart := []string{
		fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %v %v" width="%v" height="%v" version="1"><g>`,
			width+padding, line.Height, width+padding, line.Height),
	}
	LineChartAdd := func(content string) {
		lineChart = append(lineChart, content)
	}

	// Grid
	LineChartAdd(fmt.Sprintf(`<g data-testid="grid"><path d="M %v %v V %v H %v" fill="none" stroke="#333333" stroke-width="3" stroke-linecap="round"/>`,
		padding*2, padding, max+(padding*4), width+padding))
	if line.GridHorizontal || line.GridVertical {
		grid := []string{}
		for i := 0; i < (max/10)+3; i++ {
			// Generate numbers along axis
			if (i*10)%4 == 0 {
				LineChartAdd(fmt.Sprintf(`<text style="font-size: 12px" text-anchor="middle" class="text" fill="#333333" x="%v" y="%v">%v</text>`,
					padding, max+(padding*4)-(padding*i), (i * 10)))
			}

			// Generate Lines
			if line.GridHorizontal {
				grid = append(grid, fmt.Sprintf(`M %v %v L %v %v`, width+padding, (padding*i)+10, padding*2, (padding*i)+10))
			}
		}
		for i := 0; i < (width / 10); i++ {
			if line.GridVertical {
				grid = append(grid, fmt.Sprintf(`M %v %v L %v %v`, (padding*i)+padding*2, max+(padding*4), (padding*i)+padding*2, padding))
			}
		}
		LineChartAdd(fmt.Sprintf(`<path d="%v" stroke="#333333" opacity="0.3" stroke-width="1" stroke-linecap="round"/>`,
			strings.Join(grid, " ")))
	}
	LineChartAdd(`</g>`)

	// Add points
	polylinePoints := []string{}
	for i, val := range line.Values {
		polylinePoints = append(polylinePoints, fmt.Sprintf("%v,%v", (i*20)+(padding*2), max+(padding*3)-(val-10)))
	}
	LineChartAdd(fmt.Sprintf(`<polyline
		fill="none"
		stroke="%v"
		stroke-width="3"
		points="%v" stroke-dasharray="1000"
		stroke-dashoffset="1000"><animate attributeName="stroke-dashoffset" repeatCount="once" fill="freeze" dur="10s" values="1000;0"/></polyline>`, line.Color, strings.Join(polylinePoints, " ")))

	LineChartAdd(`</g></svg>`)
	return strings.Join(lineChart, "\n")
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
func GenerateCard(cardstyle themes.Theme, defs []string, body []string, width, height int, customStyles ...string) []string {
	var card Card
	card.Style = cardstyle

	// Calculate rotation
	/*
		// Rotation can be 0 to 360
		var anglePI = (angle) * (Math.PI / 180);
		var angleCoords = {
		    'x1': Math.round(50 + Math.sin(anglePI) * 50) + '%',
		    'y1': Math.round(50 + Math.cos(anglePI) * 50) + '%',
		    'x2': Math.round(50 + Math.sin(anglePI + Math.PI) * 50) + '%',
		    'y2': Math.round(50 + Math.cos(anglePI + Math.PI) * 50) + '%',
		}

	*/
	if cardstyle.Name == "retro" {
		defs = append(defs, style.LinearGradient("retro", 0, []string{"#fc00ff", "#00dbde"}))
	}
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
