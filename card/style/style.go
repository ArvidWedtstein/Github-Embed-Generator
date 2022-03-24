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
		"x1": fmt.Sprintf("%v%v", math.Round(50+math.Sin(anglePI)*50), "%"),
		"y1": fmt.Sprintf("%v%v", math.Round(50+math.Cos(anglePI)*50), "%"),
		"x2": fmt.Sprintf("%v%v", math.Round(50+math.Sin(anglePI+math.Pi)*50), "%"),
		"y2": fmt.Sprintf("%v%v", math.Round(50+math.Cos(anglePI+math.Pi)*50), "%"),
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
	<feDropShadow dx="0.2" dy="6" stdDeviation="5" flood-color="var(--pink)" flood-opacity="0.3" />
	<feDropShadow dx="0.2" dy="-6" stdDeviation="5" flood-color="var(--yellow)" flood-opacity="0.3" />
  </filter>`
}
func RetroGrid() string {
	return `
	<g transform-origin="center" transform="translate(-200,200) scale(0.040000,-0.0400000)" fill="#000000" stroke="none">
<path d="M238 4863 c39 -2 105 -2 145 0 39 1 7 3 -73 3 -80 0 -112 -2 -72 -3z 
M2383 4863 c15 -2 39 -2 55 0 15 2 2 4 -28 4 -30 0 -43 -2 -27 -4z
M6245 4860 c-3 -6 1 -7 9 -4 18 7 21 14 7 14 -6 0 -13 -4 -16 -10z
M6530 4866 c0 -2 7 -7 16 -10 8 -3 12 -2 9 4 -6 10 -25 14 -25 6z
M10363 4863 c15 -2 39 -2 55 0 15 2 2 4 -28 4 -30 0 -43 -2 -27 -4z
M12418 4863 c39 -2 105 -2 145 0 39 1 7 3 -73 3 -80 0 -112 -2 -72 -3z
M6314 4770 l-29 -80 -100 0 -99 0 74 77 c75 77 65 71 -40 -26 l-55
-51 -100 1 c-94 0 -98 1 -72 15 15 8 27 17 27 20 0 3 -16 -4 -36 -15 -40 -24
-6 -22 -474 -20 -102 1 -194 1 -205 1 -11 0 -146 0 -300 0 -154 0 -291 1 -305
0 -14 0 -216 -1 -450 -3 l-425 -3 360 -6 c198 -3 349 -7 335 -8 -14 -2 -20 -5
-15 -8 6 -2 53 2 106 9 114 14 593 12 569 -3 -19 -12 -18 -12 65 3 43 7 132
11 230 9 148 -3 158 -4 135 -18 -17 -10 -18 -13 -5 -9 63 22 95 25 200 23
l117 -3 -130 -72 -129 -73 -184 1 -184 0 143 56 c78 31 141 58 139 60 -2 2
-75 -24 -163 -57 l-159 -60 -278 2 -278 3 231 59 c126 33 228 61 226 63 -2 2
-120 -25 -263 -61 l-260 -66 -424 1 -424 1 354 61 c195 34 349 63 344 64 -6 2
-186 -26 -400 -62 l-390 -65 -634 2 -634 3 195 22 c590 66 863 98 845 101 -11
1 -282 -27 -602 -63 l-581 -65 -964 1 c-530 1 -954 3 -943 5 11 1 313 24 670
51 358 26 638 49 624 51 -15 2 -213 -11 -440 -28 -805 -60 -945 -70 -1037 -76
-243 -18 -12 -24 891 -25 854 -1 968 -3 872 -14 -293 -35 -1825 -205 -1839
-205 -9 0 -16 -4 -16 -10 0 -7 343 -10 1008 -11 554 0 1003 -2 997 -4 -12 -5
-1978 -335 -1994 -335 -6 0 -11 -7 -11 -15 0 -13 121 -15 1007 -17 953 -3
1004 -4 943 -19 -36 -9 -487 -123 -1002 -253 -515 -130 -940 -236 -943 -236
-3 0 -5 -9 -5 -20 0 -20 8 -20 1012 -22 l1013 -3 -70 -27 c-39 -15 -164 -63
-280 -106 -387 -146 -897 -340 -1260 -477 -198 -76 -372 -142 -387 -147 -22
-8 -28 -16 -28 -39 l0 -29 991 -2 991 -3 -987 -555 -986 -555 78 -3 78 -3 970
561 970 560 808 0 c717 0 845 -3 799 -22 -5 -1 -256 -233 -558 -513 -302 -281
-567 -528 -589 -548 l-40 -36 85 0 85 0 575 559 575 560 814 0 c770 0 813 -1
806 -17 -4 -10 -94 -252 -201 -538 -107 -286 -199 -530 -204 -542 -9 -23 -9
-23 74 -23 l85 0 164 502 c91 277 173 529 183 561 l18 57 816 0 816 0 18 -57
c10 -32 92 -284 183 -560 l164 -503 85 0 c47 0 83 3 81 8 -3 4 -95 248 -205
542 -110 294 -203 543 -207 553 -7 16 36 17 806 17 l814 0 575 -560 575 -559
85 0 85 0 -63 57 c-261 238 -1131 1049 -1131 1055 -1 4 362 7 807 7 l807 0
970 -560 970 -561 78 3 78 3 -986 555 -987 555 991 3 991 2 0 29 c0 23 -6 31
-27 39 -16 5 -309 117 -653 247 -344 130 -724 274 -845 320 -121 46 -280 106
-354 134 -74 28 -137 53 -139 56 -3 3 450 5 1006 5 1005 0 1012 0 1012 20 0
11 -2 20 -5 20 -3 0 -428 106 -943 236 -515 130 -966 244 -1002 253 -61 15
-10 16 943 19 886 2 1007 4 1007 17 0 8 -5 15 -11 15 -17 0 -1982 330 -1994
335 -5 2 443 4 998 4 664 1 1007 4 1007 11 0 6 -7 10 -16 10 -14 0 -1546 170
-1839 205 -96 11 18 13 873 14 902 1 1133 7 890 25 -92 6 -232 16 -1037 76
-227 17 -425 30 -440 28 -14 -2 267 -25 624 -51 358 -27 659 -50 670 -51 11
-2 -413 -4 -943 -5 l-964 -1 -581 65 c-320 36 -591 64 -602 63 -18 -3 255 -35
845 -101 l195 -22 -634 -3 -634 -2 -390 65 c-214 36 -394 64 -400 62 -5 -1
149 -30 344 -64 l354 -61 -425 -1 -425 -1 -273 69 c-151 38 -276 72 -279 75
-12 12 482 11 576 -1 53 -7 101 -11 106 -9 6 3 -1 6 -15 8 -14 1 137 5 335 8
l360 6 -425 3 c-234 2 -436 3 -450 3 -14 1 -151 0 -305 0 -154 0 -289 0 -300
0 -11 0 -103 0 -205 -1 -468 -2 -434 -4 -474 20 -20 11 -36 18 -36 15 0 -3 12
-12 28 -20 25 -14 21 -15 -73 -15 l-100 -1 -55 51 c-105 97 -115 103 -40 26
l74 -77 -99 0 -100 0 -29 80 c-16 45 -30 79 -33 77 -2 -3 8 -37 22 -78 14 -40
25 -74 25 -76 0 -2 -45 -3 -100 -3 -55 0 -100 1 -100 3 0 2 11 36 25 76 14 41
24 75 22 78 -3 2 -17 -32 -33 -77z m1639 -171 l248 -64 -278 -3 -278 -2 -159
60 c-88 33 -161 59 -163 57 -2 -2 61 -29 139 -60 l143 -56 -184 0 -184 -1
-111 62 c-61 34 -120 67 -130 74 -16 11 35 13 279 13 164 0 300 2 302 4 7 7
121 -18 376 -84z m-1978 6 l-80 -74 -155 -1 -155 0 130 75 130 74 105 0 105 0
-80 -74z m277 0 l-27 -75 -155 0 -155 1 78 74 79 75 103 0 104 0 -27 -75z
m276 7 c12 -38 22 -71 22 -75 0 -4 -67 -7 -150 -7 -82 0 -150 3 -150 7 0 4 10
37 22 75 l22 68 106 0 106 0 22 -68z m279 -7 l78 -74 -155 -1 -155 0 -27 75
-27 75 104 0 103 0 79 -75z m278 0 l130 -75 -155 0 -155 1 -80 74 -80 74 105
0 105 0 130 -74z m-3705 -110 c-41 -7 -331 -56 -644 -109 l-569 -96 -976 1
-976 1 245 29 c135 15 324 37 420 48 96 12 375 43 620 71 245 27 472 54 505
59 33 5 373 8 755 9 600 0 685 -2 620 -13z m635 -95 l-430 -110 -650 1 -650 1
160 28 c88 16 243 43 345 60 102 17 309 53 460 80 l275 48 460 0 460 1 -430
-109z m1074 106 c-2 -2 -130 -51 -284 -110 l-280 -106 -430 1 c-366 1 -423 3
-385 14 25 7 189 49 365 95 176 45 340 88 365 95 31 9 138 13 349 14 167 0
302 -1 300 -3z m430 0 c-2 -2 -89 -51 -193 -110 l-188 -106 -279 2 -279 3 273
106 272 107 199 1 c110 1 197 -1 195 -3z m269 -74 c-46 -44 -99 -93 -120 -111
l-36 -31 -226 2 -227 3 183 107 183 106 163 1 163 1 -83 -78z m388 -32 l-41
-110 -234 0 -233 0 113 110 114 109 161 1 161 0 -41 -110z m421 5 c18 -55 33
-103 33 -107 0 -5 -103 -8 -230 -8 -126 0 -230 3 -230 6 0 7 61 197 67 207 2
4 77 6 166 5 l161 -3 33 -100z m422 -5 l113 -110 -233 0 -234 0 -41 110 -41
110 161 0 161 -1 114 -109z m415 3 l187 -108 -227 -3 -226 -2 -36 31 c-21 18
-74 67 -120 111 l-83 78 159 0 160 0 186 -107z m508 0 l278 -108 -279 -3 -279
-2 -188 106 c-104 59 -191 108 -193 110 -2 2 83 4 189 4 l194 0 278 -107z
m431 88 c39 -11 207 -54 372 -97 165 -42 320 -82 345 -89 38 -11 -19 -13 -385
-14 l-430 -1 -280 106 c-154 59 -282 108 -284 110 -2 2 129 4 293 4 261 0 305
-2 369 -19z m1345 -61 c257 -44 541 -93 632 -109 l165 -29 -650 -1 -650 -1
-430 110 -430 109 448 0 448 1 467 -80z m1317 42 c273 -33 759 -89 1254 -144
l295 -33 -975 -3 -974 -2 -570 96 c-313 53 -604 102 -645 109 -65 11 24 13
640 11 l715 -2 260 -32z m-7556 -206 c-2 -2 -302 -79 -666 -170 l-661 -166
-984 1 c-540 1 -972 4 -958 6 14 2 450 78 970 168 l945 163 679 1 c374 1 677
-1 675 -3z m941 -13 c-30 -13 -445 -172 -757 -289 l-91 -34 -652 2 -652 3 124
31 c195 50 693 179 938 243 l225 59 450 0 c414 0 447 -1 415 -15z m667 11 c-3
-3 -137 -80 -298 -170 l-292 -164 -423 2 -424 3 428 167 427 166 294 1 c162 1
292 -2 288 -5z m503 -18 c-14 -13 -96 -89 -183 -170 l-158 -146 -351 0 -352 0
294 170 295 170 240 0 240 0 -25 -24z m530 15 c0 -4 -27 -80 -59 -167 l-59
-159 -352 -3 -352 -2 173 169 174 170 238 1 c130 0 237 -4 237 -9z m571 -148
c28 -87 54 -164 57 -170 3 -10 -71 -13 -348 -13 -277 0 -351 3 -348 13 3 6 29
83 57 170 l53 157 238 0 238 0 53 -157z m638 -14 l173 -169 -352 2 -352 3 -59
159 c-32 87 -59 163 -59 167 0 5 107 9 238 9 l237 -1 174 -170z m631 1 l294
-170 -352 0 -351 0 -158 146 c-87 81 -169 157 -183 170 l-25 24 240 0 240 0
295 -170z m757 3 l433 -168 -424 -3 -423 -2 -292 164 c-161 90 -295 167 -298
170 -4 3 124 6 282 6 l289 0 433 -167z m751 106 c130 -34 365 -95 522 -135
157 -41 343 -89 413 -106 l129 -33 -652 -3 -652 -2 -91 34 c-312 117 -727 276
-757 289 -32 14 0 15 408 16 l442 1 238 -61z m2091 -95 c498 -86 940 -162 981
-169 64 -11 -68 -13 -908 -14 l-984 -1 -661 166 c-364 91 -664 168 -666 170
-2 2 296 4 664 4 l668 0 906 -156z m-8094 -223 c-31 -13 -429 -164 -870 -331
-220 -83 -411 -155 -425 -160 -16 -7 -384 -9 -1010 -8 l-984 3 89 22 c50 13
169 43 265 68 96 25 272 71 390 101 118 30 265 68 325 84 61 16 218 57 350 91
132 34 305 78 385 99 l145 38 680 1 c374 0 671 -4 660 -8z m960 0 c-6 -5 -207
-119 -449 -255 l-438 -246 -641 2 -641 3 400 154 c219 85 512 198 649 252
l250 98 440 0 c250 1 435 -3 430 -8z m760 -20 c-17 -16 -140 -131 -274 -255
l-244 -226 -526 0 c-317 0 -520 4 -511 10 8 5 206 119 440 254 l425 246 360 0
360 0 -30 -29z m768 -63 c-19 -51 -62 -166 -95 -255 l-60 -163 -527 0 -526 1
260 254 260 254 361 1 362 0 -35 -92z m892 -163 l81 -250 -263 -3 c-145 -1
-381 -1 -526 0 l-263 3 78 240 c43 132 80 246 83 253 3 10 82 12 366 10 l362
-3 82 -250z m950 0 l260 -254 -526 -1 -527 0 -60 163 c-33 89 -76 204 -95 255
l-35 92 362 0 361 -1 260 -254z m935 9 c234 -135 432 -249 440 -254 9 -6 -194
-10 -511 -10 l-526 0 -244 226 c-134 124 -257 239 -274 255 l-30 29 360 0 360
0 425 -246z m802 132 c161 -62 455 -176 652 -252 l360 -139 -641 -3 -641 -2
-438 246 c-242 136 -443 250 -448 255 -6 5 178 9 426 9 l437 0 293 -114z
m1367 53 c130 -33 296 -76 369 -95 283 -73 405 -105 612 -158 118 -30 294 -76
390 -101 96 -25 215 -55 265 -68 l89 -22 -984 -3 c-626 -1 -994 1 -1010 8 -14
5 -205 77 -425 160 -441 167 -839 318 -870 331 -11 4 283 8 654 8 l674 1 236
-61z m-7454 -498 c-6 -5 -311 -178 -679 -385 l-668 -376 -967 2 c-913 3 -964
4 -921 20 44 16 1670 646 1840 713 l85 34 660 0 c396 1 655 -3 650 -8z m1165
-5 c-8 -8 -195 -181 -415 -385 l-400 -371 -790 0 c-496 0 -784 4 -775 10 8 5
305 178 660 384 l645 374 545 1 c469 1 542 -1 530 -13z m1167 2 c-2 -7 -67
-181 -143 -385 l-139 -373 -795 0 -795 0 30 28 c17 16 195 189 397 385 l367
357 542 0 c429 0 540 -3 536 -12z m1290 -368 c67 -206 123 -378 123 -382 0 -5
-355 -8 -790 -8 -434 0 -790 3 -790 6 0 8 241 747 247 757 2 4 248 6 545 5
l542 -3 123 -375z m1401 23 c202 -196 380 -369 397 -385 l30 -28 -795 0 -795
0 -139 373 c-76 204 -141 378 -143 385 -4 9 107 12 536 12 l542 0 367 -357z
m1423 -9 c349 -201 648 -374 664 -385 30 -18 20 -19 -760 -19 l-790 0 -400
371 c-220 204 -407 377 -415 385 -12 12 59 14 527 14 l540 0 634 -366z m1101
237 c1253 -486 1558 -605 1598 -619 43 -16 -9 -17 -921 -20 l-967 -2 -668 376
c-368 207 -673 380 -678 385 -6 5 251 9 647 9 l656 0 333 -129z"/>
</g>`

}
func NeonFilter(color string) string {
	filter := fmt.Sprintf(`<filter id="glow" height="300%v" width="300%v" x="-75%v" y="-75%v">
	<feMorphology operator="dilate" radius="2" in="SourceAlpha" result="thicken" />
	<feGaussianBlur in="thicken" stdDeviation="10" result="blurred" />
	<feFlood flood-color="%v" result="glowColor" />
	<feComposite in="glowColor" in2="blurred" operator="in" result="softGlow_colored" />
	<feMerge><feMergeNode in="softGlow_colored"/>
	<feMergeNode in="SourceGraphic"/></feMerge>
	</filter>`, "%", "%", "%", "%", color)
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
