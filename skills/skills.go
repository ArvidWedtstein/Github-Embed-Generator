package skills

import (
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"githubembedapi/card/themes"
	"githubembedapi/icons"
	"strings"
)

func Skills(title string, languages []string, cardstyle themes.Theme) string {
	if title == "" || len(title) <= 0 {
		title = "Skills"
	}
	fmt.Println(cardstyle.Colors.Title)
	height := 700
	width := 600
	titleboxheight := 50
	padding := 10
	strokewidth := 3
	boxwidth := 60
	boxheight := 40

	customstyles := []string{
		`@font-face { font-family: Papyrus; src: '../papyrus.TFF'}`,
		`.languagebox { 
			fill: ` + cardstyle.Colors.Box + `;
		}`,
	}
	defs := []string{
		style.RadialGradient("paint0_angular_0_1", []string{"#7400B8", "#6930C3", "#5E60CE", "#5390D9", "#4EA8DE", "#48BFE3", "#56CFE1", "#64DFDF", "#72EFDD"}),
		style.LinearGradient("gradient-fill", 0, []string{"#1f005c", "#5b0060", "#870160", "#ac255e", "#ca485c", "#e16b5c", "#f39060", "#ffb56b"}),
	}

	body := []string{
		fmt.Sprintf(`<text x="20" y="35" class="title">%s</text>`, card.ToTitleCase(title)),
	}

	bodyAdd := func(content string) string {
		body = append(body, content)
		return content
	}

	// Algoritm for checking if color is too dark
	// colorToDark := func(color string) bool {
	// 	var c = strings.Replace(color, "#", "", -1) // strip #
	// 	rgb, err := strconv.ParseInt(c, 16, 32)     // convert rrggbb to decimal
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	r := (rgb >> 16) & 0xff // extract red
	// 	g := (rgb >> 8) & 0xff  // extract green
	// 	b := (rgb >> 0) & 0xff  // extract blue

	// 	rFloat := 0.2126
	// 	gFloat := 0.7152
	// 	bFloat := 0.0722
	// 	r2Float := float64(r)
	// 	g2Float := float64(g)
	// 	b2Float := float64(b)
	// 	luma := math.Sqrt(rFloat*(r2Float*r2Float) +
	// 		gFloat*(g2Float*g2Float) +
	// 		bFloat*(b2Float*b2Float))

	// 	return luma < 80
	// }
	// Calculate where repositoryboxes should begin
	posY := titleboxheight + padding

	posX := 0

	originalpos := posX
	newwidth := width
	newheight := height

	row := func(content []string, lang string) {
		bodyAdd(fmt.Sprintf(`<g class="languagebox" title="%v" transform="translate(%v,%v) rotate(0)">`, lang, posX+padding, posY))

		for _, v := range content {
			bodyAdd(v)
		}
		bodyAdd(`</g>`)

		newheight = posY + boxheight + padding
		// check if next box will fit into card
		if posX+(boxwidth+(len(lang)*6))+((boxwidth+(len(lang)*6))+padding) >= width {
			posY += boxheight + padding
			newwidth = posX + (boxwidth + (len(lang) * 6)) + (padding * 2)
			posX = originalpos - ((boxwidth + (len(lang) * 6)) + padding)
		}
	}

	for _, lang := range languages {

		icon := icons.Icons(lang)

		// Calculate text width somehow.
		img := fmt.Sprintf(`<g data-testid="icon" transform="translate(%v,%v)">%v<text x="%v" y="%v" alignment-baseline="auto" text-anchor="left" class="text">%v</text></g>`,
			0, 0, icon, boxwidth+(len(lang))-(len(lang)+20), (boxheight/2)+5, card.ToTitleCase(lang))

		row([]string{
			fmt.Sprintf(`<rect x="0" y="0" rx="5" class="" width="%v" height="%v" />`, boxwidth+(len(lang)*6), boxheight),
			img,
		}, lang)

		posX += (boxwidth + (len(lang) * 6)) + padding

	}

	// adjust the svg size to the content
	if newwidth != width {
		width = newwidth
	}
	if newheight != height {
		height = newheight
	}

	// Line on top
	body = append([]string{fmt.Sprintf(`<rect x="0" y="%v" width="%v" height="%v" fill="%v"/>`, titleboxheight, width, strokewidth, cardstyle.Colors.Border)}, body...)

	return strings.Join(card.GenerateCard(cardstyle, defs, body, width+strokewidth, height+strokewidth, true, customstyles...), "\n")
}
