package style

import (
	"fmt"
	"githubembedapi/card/themes"
	"math"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type Styles struct {
	Title,
	Border,
	Background,
	Text,
	Textfont,
	Box string
}

func CheckTheme(c *gin.Context) themes.Theme {
	var SelectedTheme themes.Theme
	if len(c.Request.FormValue("theme")) < 1 {
		SelectedTheme = themes.LoadTheme("light")
	} else {
		SelectedTheme = themes.LoadTheme(c.Request.FormValue("theme"))
	}

	colors := map[string]string{
		"Title":      c.Request.FormValue("titlecolor"),
		"Text":       c.Request.FormValue("textcolor"),
		"Background": c.Request.FormValue("backgroundcolor"),
		"Border":     c.Request.FormValue("bordercolor"),
		"Box":        c.Request.FormValue("boxcolor"),
	}

	for key, value := range colors {
		if len(value) > 0 {
			r, _ := regexp.Compile("^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")
			if r.MatchString(value) {
				colors[key] = "#" + value
			}
		} else {
			delete(colors, key)
		}
	}

	// Find a dynamic solution for this crap. Can't iterate through structs
	for k, v := range colors {
		if k == "Title" {
			SelectedTheme.Colors.Title = v
		}
		if k == "Text" {
			SelectedTheme.Colors.Text = v
		}
		if k == "Border" {
			SelectedTheme.Colors.Border = v
		}
		if k == "Background" {
			SelectedTheme.Colors.Background = v
		}
		if k == "Box" {
			SelectedTheme.Colors.Box = v
		}
	}
	if len(c.Request.FormValue("font")) < 1 {
		SelectedTheme.Font = "Helvetica"
	}
	return SelectedTheme
}

func RadialGradient(id string, colors []string) string {
	gradient := []string{
		fmt.Sprintf(`<radialGradient id="%v" gradientUnits="userSpaceOnUse">`, id),
	}

	var offset float64 = 1.0 / float64(cap(colors)-1)
	for i, v := range colors {
		gradient = append(gradient, fmt.Sprintf(`<stop offset="%v" stop-color="%v"/>`, offset*float64(i), v))
	}

	gradient = append(gradient, `</radialGradient>`)
	return strings.Join(gradient, "\n")
}
func LinearGradient(id string, degree int, colors []string) string {

	// Calculate rotation
	anglePI := float64(degree) * (math.Pi / 180)
	angleCoords := map[string]string{
		"x1": fmt.Sprintf("%v%%", math.Round(50+math.Sin(anglePI)*50)),
		"y1": fmt.Sprintf("%v%%", math.Round(50+math.Cos(anglePI)*50)),
		"x2": fmt.Sprintf("%v%%", math.Round(50+math.Sin(anglePI+math.Pi)*50)),
		"y2": fmt.Sprintf("%v%%", math.Round(50+math.Cos(anglePI+math.Pi)*50)),
	}
	gradient := []string{
		fmt.Sprintf(`<linearGradient x1="%v" y1="%v" x2="%v" y2="%v" id="%v" gradientUnits="userSpaceOnUse">`,
			angleCoords["x1"], angleCoords["y1"], angleCoords["x2"], angleCoords["y2"], id),
	}
	if cap(colors) < 2 {
		panic(`Gradient must have 2 colors`)
	}
	var offset float64 = 1.0 / float64(cap(colors)-1)
	for i, v := range colors {
		gradient = append(gradient, fmt.Sprintf(`<stop offset="%v" stop-color="%v"/>`, offset*float64(i), v))
	}

	gradient = append(gradient, `</linearGradient>`)
	return strings.Join(gradient, "\n")
}
func HexagonPattern() string {

	return `<pattern id="pattern-hex" x="0" y="0" width="112" height="190" patternUnits="userSpaceOnUse" viewBox="56 -254 112 190">
	<g id="hexagon">
	<path d="M168-127.1c0.5,0,1,0.1,1.3,0.3l53.4,30.5c0.7,0.4,1.3,1.4,1.3,2.2v61c0,0.8-0.6,1.8-1.3,2.2L169.3-0.3 c-0.7,0.4-1.9,0.4-2.6,0l-53.4-30.5c-0.7-0.4-1.3-1.4-1.3-2.2v-61c0-0.8,0.6-1.8,1.3-2.2l53.4-30.5C167-127,167.5-127.1,168-127.1 L168-127.1z"></path>
	<path d="M112-222.5c0.5,0,1,0.1,1.3,0.3l53.4,30.5c0.7,0.4,1.3,1.4,1.3,2.2v61c0,0.8-0.6,1.8-1.3,2.2l-53.4,30.5 c-0.7,0.4-1.9,0.4-2.6,0l-53.4-30.5c-0.7-0.4-1.3-1.4-1.3-2.2v-61c0-0.8,0.6-1.8,1.3-2.2l53.4-30.5 C111-222.4,111.5-222.5,112-222.5L112-222.5z"></path>
	<path d="M168-317.8c0.5,0,1,0.1,1.3,0.3l53.4,30.5c0.7,0.4,1.3,1.4,1.3,2.2v61c0,0.8-0.6,1.8-1.3,2.2L169.3-191 c-0.7,0.4-1.9,0.4-2.6,0l-53.4-30.5c-0.7-0.4-1.3-1.4-1.3-2.2v-61c0-0.8,0.6-1.8,1.3-2.2l53.4-30.5 C167-317.7,167.5-317.8,168-317.8L168-317.8z"></path>
	</g>

	</pattern>`
}
func CubePattern() string {
	return `<pattern id="pattern-cubes" x="0" y="63" patternUnits="userSpaceOnUse" width="31" height="50" viewBox="0 0 10 16"> 
     
		<g id="cube">
			<path fill="darkblue" class="left-shade" d="M0 0l5 3v5l-5 -3z"></path>
			<path fill="blue" class="right-shade" d="M10 0l-5 3v5l5 -3"></path>
		</g>
   
		<use fill="darkblue" x="5" y="8" href="#cube"></use>
		<use fill="blue" x="-5" y="8" href="#cube"></use>

	</pattern>`
}
func StarPattern() string {
	return `<pattern id="star" viewBox="0,0,10,10" width="10%" height="10%">
	<polygon points="0,0 2,5 0,10 5,8 10,10 8,5 10,0 5,2"/>
	  </pattern>`
}
func WavyFilter() string {
	return `
	<filter id="wavy">
		<feTurbulence type="fractalNoise" baseFrequency="0.009"
        numOctaves="5" result="turbulence">
		<animate attributeName="baseFrequency" dur="30s" values="0.01;0.005;0.02;0.009" repeatCount="indefinite" />
		</feTurbulence>
		<feDisplacementMap in="SourceGraphic" scale="20" />
	</filter>"
	`
}
func SunGradient() string {
	return `<linearGradient id="sunGradient" gradientTransform="rotate(90)">
	<stop offset="0%" stop-color="#ffd319" />
	<stop offset="100%" stop-color="#ff2975" />
  </linearGradient>
  <filter id="shadow">
	<feDropShadow dx="0.2" dy="4" stdDeviation="5" flood-color="#ff2975" flood-opacity="0.3" />
	<feDropShadow dx="0.2" dy="-4" stdDeviation="5" flood-color="#ffd319" flood-opacity="0.3" />
  </filter>`
}

func NeonFilter(color string) string {
	filter := fmt.Sprintf(`<filter id="glow" height="300%%" width="300%%" x="-75%%" y="-75%%">
	<feMorphology operator="dilate" radius="2" in="SourceAlpha" result="thicken" />
	<feGaussianBlur in="thicken" stdDeviation="10" result="blurred" />
	<feFlood flood-color="%v" result="glowColor" />
	<feComposite in="glowColor" in2="blurred" operator="in" result="softGlow_colored" />
	<feMerge><feMergeNode in="softGlow_colored"/>
	<feMergeNode in="SourceGraphic"/></feMerge>
	</filter>`, color)
	return filter
}
func Table() string {
	return `<filter id="table" x="0" y="0" width="100%" height="100%">
	<feComponentTransfer>
	  <feFuncR type="table" tableValues="0 0 1 1"></feFuncR>
	  <feFuncG type="table" tableValues="1 1 0 0"></feFuncG>
	  <feFuncB type="table" tableValues="0 1 1 0"></feFuncB>
	</feComponentTransfer>
  </filter>`
}
func RoughPaper() string {
	return `<filter id='roughpaper' x='0%' y='0%' width='100%' height="100%">
	<feTurbulence type="fractalNoise" baseFrequency='0.04' result='noise' numOctaves="5" />
	<feDiffuseLighting in='noise' lighting-color='white' surfaceScale='2'>
		<feDistantLight azimuth='45' elevation='60' />
</feDiffuseLighting>

</filter>`
}
func StarsFilter() string {
	// feColorMatrix
	//------------------
	//	   R G B A M
	//--------------
	// R | 1 0 0 0 0
	// G | 0 1 0 0 0
	// B | 0 0 1 0 0
	// A | 0 0 0 1 0
	return `<filter id="stars">
	<feTurbulence baseFrequency="0.2"/>
	
	<feColorMatrix values="0 0 0 9 -4
						   0 0 0 9 -4
						   0 0 0 9 -4
						   0 0 0 0 1"/>
		</filter>`
}
func DropShadow() string {
	return `<filter id="filter2_d_0_1" x="406" y="71" width="155" height="154" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
	<feFlood flood-opacity="0" result="BackgroundImageFix"/>
	<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
	<feOffset dy="4"/>
	<feGaussianBlur stdDeviation="2"/>
	<feComposite in2="hardAlpha" operator="out"/>
	<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
	<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_0_1"/>
	<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_0_1" result="shape"/>
	</filter>`
}
func Filter2() string {
	return `<filter id="filter1_d_0_1" x="0" y="0" width="100%" height="100%" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
	<feFlood flood-opacity="0" result="BackgroundImageFix"/>
	<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
	<feOffset dy="4"/>
	<feGaussianBlur stdDeviation="2"/>
	<feComposite in2="hardAlpha" operator="out"/>
	<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
	<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_0_1"/>
	<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_0_1" result="shape"/>
	</filter>`
}
func DropShadowRing1() string {
	return `<filter id="filter0_d_0_1" x="391" y="55" width="185" height="186" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
	<feFlood flood-opacity="0" result="BackgroundImageFix"/>
	<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
	<feOffset dy="4"/>
	<feGaussianBlur stdDeviation="2"/>
	<feComposite in2="hardAlpha" operator="out"/>
	<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
	<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_0_1"/>
	<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_0_1" result="shape"/>
	</filter>`
}
func Blur(amount int) string {
	return fmt.Sprintf(`<filter id="blur%v" x="0" y="0">
	<feGaussianBlur in="SourceGraphic" stdDeviation="%v" />
  </filter>`, amount, amount)
}
func DropShadowColor() string {
	return `<filter id="dropshadowcolor" x="0" y="0" width="200%" height="200%">
	<feOffset result="offOut" in="SourceGraphic" dx="20" dy="20" />
	<feGaussianBlur result="blurOut" in="offOut" stdDeviation="10" />
	<feBlend in="SourceGraphic" in2="blurOut" mode="normal" />
  </filter>`
}
