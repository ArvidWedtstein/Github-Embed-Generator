package card

import (
	"fmt"
	"githubembedapi/card/style"
	"githubembedapi/card/themes"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Card struct {
	Title string       `json:"title"`
	Style themes.Theme `json:"colors"`
	Body  []string     `json:"body"`
}

func (card Card) GetStyles(customStyles ...string) string {
	var style = []string{
		`<style>`,
		`.title { font: 25px sans-serif; fill: ` + card.Style.Colors.Title + `}`,
		`.text { font: 16px sans-serif; fill: ` + card.Style.Colors.Text + `; font-family: ` + card.Style.Font + `;}`,
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
	PosX   int
	PosY   int
}

func (radar Radar) generateGrid() string {
	centerX := (radar.Width / 2)
	centerY := (radar.Height / 2)

	var rings int = 10

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
		fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %v %v" width="%v" height="%v" version="1"><g transform="translate(%v,%v) scale(1)">`, radar.Width, radar.Height, radar.Width, radar.Height, radar.PosX, radar.PosY),
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
			totalWidth, max+120, totalWidth, max+10),
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
				// content = append(content, fmt.Sprintf(`<g transform="translate(%v,%v)" class="bar"><rect width="%v" height="%v" fill="%v"/><text transform="rotate(-90)" x="%v" y="%v" fill="#000000" text-anchor="middle">%v</text></g>`,
				// 	(20+barGap)*i, (max+30)-v, 20, v, "#ff0000", (v+0)+len(bar.Labels[i]), 10, bar.Labels[i]))
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

	// Grid
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

	// Add Columns
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
			width+padding, max+(padding*4), width+padding, max+(padding*4)),
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

// ----------------------------------------
// Retro / Synthwave Mountain Generation
// ----------------------------------------
func mountain(width, height int, fill string, grid bool) string {
	mountain := []string{fmt.Sprintf("M 0 %v", height)}
	mountaingrid := []string{}
	steps := width / 10
	var lastrand int = 0
	var lastrandgrid int = 0
	var maxgrid = 2
	var mingrid = 0
	for i := 0; i < steps; i++ {
		max := lastrand + 41
		min := 40

		randomCoord := rand.Intn(max-min) + min
		mountain = append(mountain, fmt.Sprintf("L %v %v", i*10, height-randomCoord))
		lastrand = randomCoord

		// Wireframe
		if i%2 == 0 {
			maxgrid = (i * 10) + 10
			mingrid = lastrandgrid
		}
		randomCoordGridX := rand.Intn(maxgrid-mingrid) + mingrid
		mountaingrid = append(mountaingrid, fmt.Sprintf("M %v %v L %v %v", i*10, height-randomCoord, randomCoordGridX, height))
		lastrandgrid = randomCoordGridX
	}
	mountain = append(mountain, fmt.Sprintf(`L %v %v Z`, width, height))
	mountaingrid = append(mountaingrid, fmt.Sprintf(`L %v %v Z`, width, height))
	path := []string{}
	if grid {
		// Mountain
		path = append(path, fmt.Sprintf("<path d='%v' stroke='%v' stroke-width='1' fill='%v'/>", strings.Join(mountain, " "), "url(#retro)", fill))

		// Grid
		path = append(path, fmt.Sprintf("<path d='%v' stroke='%v' stroke-width='1' />", strings.Join(mountaingrid, ""), "url(#retro)"))
	} else {
		// Mountain
		path = append(path, fmt.Sprintf("<path d='%v' fill='%v'/>", strings.Join(mountain, " "), fill))
	}
	return strings.Join(path, " ")
}
func generateGrid(width, height float64) string {
	path := []string{}
	coordsLineX := []float64{-7.5}

	startnum := -7.5
	coordsMoveX := []float64{-13.5}
	startMx := -13.5

	for i := 0; startnum <= width; i++ {
		num1 := startnum + 7
		num2 := num1 + 6.5
		num3 := num2 + 7
		num4 := num3 + 6.5
		num5 := num4 + 6.5

		coordsLineX = append(coordsLineX, num1, num2, num3, num4, num5)
		startnum = num5

		num1m := startMx + 41
		num2m := num1m + 41
		num3m := num2m + 41.5
		num4m := num3m + 41
		num5m := num4m + 41
		num6m := num5m + 41.5

		coordsMoveX = append(coordsMoveX, num1m, num2m, num3m, num4m, num5m, num6m)
		startMx = num6m
	}

	for i, v := range coordsLineX {
		path = append(path, fmt.Sprintf(`M %v 212 L %v 1 `, coordsMoveX[i], v))
	}
	startYLine := 5.5
	var step float64 = 2.5

	for i := 0; i < 20; i++ {
		path = append(path, fmt.Sprintf("M1 %v L%v %v", startYLine, width, startYLine))
		if i%3 == 0 {
			step += 2
		}
		startYLine += step
	}
	return strings.Join(path, " ")
}
func wave(width, height int) string {

	path := []string{fmt.Sprintf(`M 1.5 %v`, height)}
	posX := 0

	for i := 0; posX < width; i++ {
		// c dx1, dy1 dx2, dy2 dx, dy
		// Uppercase C uses absolute coordinates
		// Lowercase c uses relative coordinates
		// dy := height - 50

		randomDY := rand.Intn(50) + (height - 50)
		path = append(path, fmt.Sprintf(`C %v,%v %v,%v %v,%v`, posX, randomDY, posX, randomDY, posX, height-100))

		posX += 50
	}
	path = append(path, fmt.Sprintf(`L %v %v Z`, width, height))
	return strings.Join(path, " ")
}

// ----------------------------------------
// Card Generation
// ----------------------------------------
func GenerateCard(cardstyle themes.Theme, defs, body []string, width, height int, hasBox bool, customStyles ...string) []string {
	var card Card
	card.Style = cardstyle

	/**
		TODO
		-------------
		Make some kind of system for handling gradients
	**/
	theme := []string{}

	if cardstyle.Name == "rgb" {
		defs = append(defs, style.LinearGradient("rgb", 180, []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}))
	}
	if cardstyle.Name == "ig9te" {
		defs = append(defs, style.SunGradient())
		defs = append(defs, style.LinearGradient("ig9te", -90, []string{"#8801ee", "#ff6a00"}))
	}
	if cardstyle.Name == "red" {
		defs = append(defs, style.LinearGradient("redgrad", -90, []string{"#c31432", "#240b36"}))
		defs = append(defs, style.LinearGradient("fire", 0, []string{"#f12711", "#f5af19"}))

		wave := wave(width, height)
		theme = append(theme, fmt.Sprintf(`<path transform="translate(0,0)" fill="url(#fire)" d='%v'></path>`, wave))
	}

	if cardstyle.Name == "retro" {
		defs = append(defs, style.SunGradient())

		defs = append(defs, style.LinearGradient("retro", 180, []string{"#fc00ff", "#00dbde"}))
		theme = append(theme, fmt.Sprintf(`<svg width="%v" height="%v" viewbox="0 0 %v %v">`, width, height, width, height))

		grid := generateGrid(float64(width)+float64(width/5)+25, float64(height))
		theme = append(theme, fmt.Sprintf(`<path transform="scale(0.8) translate(0,%v)" fill="none" stroke="url(#retro)" stroke-width="1.0391" stroke-miterlimit="10" d='%v'></path>`, height-30, grid))

		if height >= 300 {
			theme = append(theme, fmt.Sprintf(`<circle cx="%v" cy="%v" filter="url(#shadow)" fill="url(#sunGradient)" r="30"><animate attributeName="cy" dur="5s" values="%v;%v" repeatCount="0" /><animate attributeName="r" dur="5s" values="10;30" repeatCount="0" /></circle>`, (width/2)+(width/4), (height/2), ((height/2)+(height/8))+10, (height/2)))
			theme = append(theme, fmt.Sprintf("<g transform='translate(0,100) scale(0.5)'>%v%v</g>", mountain(width*2, height, "#111111", false), mountain(width*2, height, "#222222", true)))
		}
		theme = append(theme, `</svg>`)
	}

	card.Body = []string{
		fmt.Sprintf(`<svg width="%v" height="%v" viewBox="0 0 %v %v" xmlns="http://www.w3.org/2000/svg">`, width, height, width, height),
		card.GetDefs(defs),
	}
	strokewidth := 3
	if hasBox {
		customStyles = append(customStyles, `
		.box {
			fill: `+cardstyle.Colors.Background+`;
			stroke: `+cardstyle.Colors.Border+`;
			stroke-width: `+strconv.Itoa(strokewidth)+`px;
		}
		`)
		card.Body = append(card.Body, card.GetStyles(customStyles...), fmt.Sprintf(`<rect id="box" x="%v" y="%v" class="box" width="%v" height="%v" rx="15"  />`, strokewidth/2, strokewidth/2, width-strokewidth, height-strokewidth))
	}

	card.Body = append(card.Body, strings.Join(theme, "\n"), strings.Join(body, "\n"), `</svg>`)
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
	if len(list) < 1 {
		return 0
	}
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

func ArrayContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func RemoveFromSlice(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func IndexOf(element string, data []string) int {
	for k, v := range data {
		if strings.Contains(v, ToTitleCase(element)) {
			return k
		}
	}
	return -1 //not found.
}

func UrlSplit(c *gin.Context, index string) []string {
	return strings.Split(c.Request.FormValue(index), ",")
}
