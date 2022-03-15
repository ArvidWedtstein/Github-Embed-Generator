package main

import (
	"fmt"
	"githubembedapi/card/style"
	"githubembedapi/commit_activity"
	"githubembedapi/organization"
	"githubembedapi/project"
	"githubembedapi/rank"
	"githubembedapi/streak"
	"net/http"
	"strings"

	"githubembedapi/skills"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/static", http.Dir("./static"))
	router.GET("/ranklist", rankList)
	router.GET("/skills", getSkills)
	router.GET("/mostactivity", getMostactivity)
	router.GET("/project", projectcard)
	router.GET("/commitactivity", repositoryCommitActivity)
	router.GET("/streak", userstreak)
	router.GET("/resonance", resonance)

	// router.Run("localhost:8080")
	router.Run()
}

func getMostactivity(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	var color style.Styles
	styles := map[string]string{
		"Title":      c.Request.FormValue("titlecolor"),
		"Border":     c.Request.FormValue("bordercolor"),
		"Background": c.Request.FormValue("backgroundcolor"),
		"Text":       c.Request.FormValue("textcolor"),
		"Box":        c.Request.FormValue("boxcolor"),
	}
	color = style.CheckHex(styles)
	org := c.Request.FormValue("org")
	title := c.Request.FormValue("title")

	github_token := ""

	c.String(http.StatusOK, organization.MostactivityCard(title, org, color, github_token))
}
func repositoryCommitActivity(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	var color style.Styles
	styles := map[string]string{
		"Title":      c.Request.FormValue("titlecolor"),
		"Border":     c.Request.FormValue("bordercolor"),
		"Background": c.Request.FormValue("backgroundcolor"),
		"Text":       c.Request.FormValue("textcolor"),
		"Box":        c.Request.FormValue("boxcolor"),
	}
	color = style.CheckHex(styles)
	user := c.Request.FormValue("user")
	repo := c.Request.FormValue("repo")
	title := c.Request.FormValue("title")
	var hide_week string = c.Request.FormValue("hide_week")

	c.String(http.StatusOK, commit_activity.RepositoryCommitActivity(title, user, repo, hide_week, color))
}
func projectcard(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var color style.Styles
	styles := map[string]string{
		"Title":      c.Request.FormValue("titlecolor"),
		"Border":     c.Request.FormValue("bordercolor"),
		"Background": c.Request.FormValue("backgroundcolor"),
		"Text":       c.Request.FormValue("textcolor"),
		"Box":        c.Request.FormValue("boxcolor"),
	}
	user := c.Request.FormValue("user")
	repo := c.Request.FormValue("repo")
	color = style.CheckHex(styles)
	c.String(http.StatusOK, project.Project(user, repo, color))
}
func rankList(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var color style.Styles
	styles := map[string]string{
		"Title":      c.Request.FormValue("titlecolor"),
		"Border":     c.Request.FormValue("bordercolor"),
		"Background": c.Request.FormValue("backgroundcolor"),
		"Text":       c.Request.FormValue("textcolor"),
	}
	color = style.CheckHex(styles)
	users := strings.Split(fmt.Sprintf("%v", c.Request.FormValue("users")), ",")
	title := c.Request.FormValue("title")

	c.String(http.StatusOK, rank.Rankcard(title, users, color))
}
func userstreak(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	user := c.Request.FormValue("user")

	var color style.Styles
	styles := map[string]string{
		"Title":      c.Request.FormValue("titlecolor"),
		"Border":     c.Request.FormValue("bordercolor"),
		"Background": c.Request.FormValue("backgroundcolor"),
		"Text":       c.Request.FormValue("textcolor"),
	}
	color = style.CheckHex(styles)
	hide_title := c.Request.FormValue("hide_title")
	c.String(http.StatusOK, streak.Streak(user, hide_title, color))
}
func getSkills(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	// Define styles
	var color style.Styles
	languages := strings.Split(c.Request.URL.Query().Get("languages"), ",")

	styles := map[string]string{
		"Title":      c.Request.FormValue("titlecolor"),
		"Border":     c.Request.FormValue("bordercolor"),
		"Background": c.Request.FormValue("backgroundcolor"),
		"Text":       c.Request.FormValue("textcolor"),
		"Box":        c.Request.FormValue("boxcolor"),
	}
	// Function that checks all HEX codes
	color = style.CheckHex(styles)
	title := c.Request.FormValue("title")

	c.String(http.StatusOK, skills.Skills(title, languages, color))
}

func resonance(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	c.String(http.StatusOK, `<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 500 500" onload="anim()">
	<style>
	
	#lineGroup line {
	  stroke-width: 1px;
	}
	#earth, #sol, #venus {
		fill: none;
	}
	</style>
	<script>
	  // <![CDATA[
	  function anim() {
	  let earth = document.getElementById('earth')
	  let venus = document.getElementById('venus')
	  let lineGroup = document.getElementById('lineGroup')
	  const earthDeg = 5,
	earthOrbits = 8,
	venusOrbits = 13,
	resonance = earthOrbits / venusOrbits,
	centre = 250,
	earthDist = centre - parseInt(earth.getAttribute("cy"), 10),
	venusDist = centre - parseInt(venus.getAttribute("cy"), 10);
	let i = 0,
	orbitals = setInterval(function(){
	  earth.setAttribute("transform", "rotate("+ i + " " + centre + " " + centre + ")");
	   venus.setAttribute("transform", "rotate("+ i / resonance + " " + centre + " " + centre + ")");
	  let earthX = Math.cos((i*Math.PI/180)) * earthDist + centre,
	  earthY = Math.sin((i*Math.PI/180)) * earthDist + centre;
	  venusX = Math.cos((i/(earthOrbits/13))*Math.PI/180) * venusDist + centre,
	  venusY = Math.sin((i/(earthOrbits/13))*Math.PI/180) * venusDist + centre,
	  resLine = document.createElementNS('http://www.w3.org/2000/svg', 'line');
	  resLine.setAttribute('x1', earthX);
	  resLine.setAttribute('y1', earthY);
	  resLine.setAttribute('x2', venusX);
	  resLine.setAttribute('y2', venusY);
	  resLine.setAttribute('stroke', 'hsla(' + i + ', 50%, 50%, 0.5)');
	  lineGroup.appendChild(resLine);
	  i += earthDeg;
	if (i == (360 * earthOrbits) + earthDeg) {
	  clearInterval(orbitals);
	 }
	}, 60);
	}
	anim()
	// ]]>
	  </script>
	  <g id="orbits">
	  <circle id="venusorbit" cx="250" cy="250" r="120" />
	  <circle id="earthorbit" cx="250" cy="250" r="165" />
	</g>
	  <g id="lineGroup" transform="rotate(-90 250 250)"></g>
	  <circle id="earth" cx="250" cy="85" r="8" />
	<circle id="venus" cx="250" cy="130" r="5" />
	  <circle id="sol" cx="250" cy="250" r="16" /></svg>`)
}
