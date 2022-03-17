package skills

import (
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"githubembedapi/icons"
	"math"
	"strconv"
	"strings"
)

func Skills(title string, languages []string, cardstyle style.Styles) string {
	if title == "" || len(title) <= 0 {
		title = "Skills"
	}
	fmt.Println(cardstyle.Title)
	height := 700
	width := 600
	titleboxheight := 50
	padding := 10
	strokewidth := 3
	boxwidth := 60
	boxheight := 60

	customstyles := []string{
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.repobox { 
			fill: ` + cardstyle.Box + `;
			border: ` + strconv.Itoa(strokewidth) + `px solid #` + cardstyle.Border + `;
		}`,
		`.box {
			fill: ` + cardstyle.Background + `;
			border: 3px solid #` + cardstyle.Border + `;
			stroke: ` + cardstyle.Border + `;
			stroke-width: ` + strconv.Itoa(strokewidth) + `px;
		}`,
	}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
	}

	body := []string{
		fmt.Sprintf(`<text x="20" y="35" class="title">%s</text>`, card.ToTitleCase(title)),
	}

	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}

	// Algoritm for checking if color is too dark
	colorToDark := func(color string) bool {
		var c = strings.Replace(color, "#", "", -1) // strip #
		rgb, err := strconv.ParseInt(c, 16, 32)     // convert rrggbb to decimal
		if err != nil {
			panic(err.Error())
		}
		r := (rgb >> 16) & 0xff // extract red
		g := (rgb >> 8) & 0xff  // extract green
		b := (rgb >> 0) & 0xff  // extract blue

		rFloat := 0.2126
		gFloat := 0.7152
		bFloat := 0.0722
		r2Float := float64(r)
		g2Float := float64(g)
		b2Float := float64(b)
		luma := math.Sqrt(rFloat*(r2Float*r2Float) +
			gFloat*(g2Float*g2Float) +
			bFloat*(b2Float*b2Float))

		return luma < 80
	}
	fmt.Println(colorToDark("#000000"))
	// Calculate where repositoryboxes should begin
	posY := titleboxheight + padding

	posX := 0

	imgsize := boxwidth - (padding * 2)
	originalpos := posX
	newwidth := width
	newheight := height

	row := func(content []string, lang string) {
		bodyAdd(fmt.Sprintf(`<g class="repobox" title="%v" transform="translate(%v,%v) rotate(0)">`, lang, posX+padding, posY))

		for _, v := range content {
			bodyAdd(v)
		}
		bodyAdd(`</g>`)

		newheight = posY + boxheight + padding
		// check if next box will fit into card
		if posX+boxwidth+(boxwidth+padding) >= width {
			posY += boxheight + padding
			newwidth = posX + boxwidth + (padding * 2)
			posX = originalpos - (boxwidth + padding)
		}
	}

	for _, v := range languages {

		icon := icons.Icons(v)
		img := fmt.Sprintf(`<g x="%v" y="%v" height="%v" width="%v">%v</g>`, boxwidth-imgsize, boxheight-imgsize, imgsize, imgsize, icon)

		row([]string{
			fmt.Sprintf(`<rect x="0" y="0" rx="5" class="" width="%v" height="%v" />`, boxwidth, boxheight),
			img,
		}, v)

		posX += boxwidth + padding
	}

	// adjust the svg size to the content
	if newwidth != width {
		width = newwidth
	}
	if newheight != height {
		height = newheight
	}

	// Line on top
	body = append([]string{fmt.Sprintf(`<rect x="0" y="%v" width="%v" height="%v" fill="%v"/>`, titleboxheight, width, strokewidth, cardstyle.Border)}, body...)
	body = append([]string{fmt.Sprintf(`<rect x="%v" y="%v" class="box" width="%v" height="%v" rx="15"  />`, strokewidth/2, strokewidth/2, width, height)}, body...)

	return strings.Join(card.GenerateCard(cardstyle, defs, body, width+strokewidth, height+strokewidth, customstyles...), "\n")

}
