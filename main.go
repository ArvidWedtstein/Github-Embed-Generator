package main

import (
	"fmt"
	"githubembedapi/card"
	"githubembedapi/card/style"
	"githubembedapi/commit_activity"
	"githubembedapi/languageCard"
	"githubembedapi/organization"
	"githubembedapi/project"
	"githubembedapi/rank"
	"githubembedapi/skills"
	"githubembedapi/streak"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.StaticFS("/static", http.Dir("./static"))
	router.StaticFile("/.env", "/.env")
	router.GET("/ranklist", rankList)
	router.GET("/skills", getSkills)
	router.GET("/mostactivity", getMostactivity)
	router.GET("/project", projectcard)
	router.GET("/commitactivity", repositoryCommitActivity)
	router.GET("/streak", userstreak)
	router.GET("/languageCard", language)
	router.GET("/radar", radar)
	router.GET("/line", line)
	router.GET("/bar", bar)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	// router.Run("localhost:8080")
	router.Run()
}
func radar(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var radar card.Radar
	values := strings.Split(fmt.Sprintf("%v", c.Request.FormValue("values")), ",")
	labels := strings.Split(fmt.Sprintf("%v", c.Request.FormValue("labels")), ",")
	for _, v := range values {
		val, _ := strconv.Atoi(v)
		radar.Values = append(radar.Values, val)
	}
	radar.Labels = labels
	radar.Color = "#0000ff"
	radar.Width = 300
	radar.Height = 300
	radar.Grid = true
	c.String(http.StatusOK, card.RadarChart(radar))
}
func line(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var line card.Line
	values := strings.Split(fmt.Sprintf("%v", c.Request.FormValue("values")), ",")
	for _, v := range values {
		val, _ := strconv.Atoi(v)
		line.Values = append(line.Values, val)
	}
	line.Width = 300
	line.Height = 300
	line.Color = "#0074d9"
	line.GridVertical = true
	line.GridHorizontal = true
	c.String(http.StatusOK, card.LineChart(line))
}
func bar(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")
	var bar card.Bar
	values := strings.Split(fmt.Sprintf("%v", c.Request.FormValue("values")), ",")
	labels := strings.Split(fmt.Sprintf("%v", c.Request.FormValue("labels")), ",")
	for _, v := range values {
		val, _ := strconv.Atoi(v)
		bar.Values = append(bar.Values, val)
	}
	bar.Labels = labels
	bar.Width = 300
	bar.Height = 300
	bar.Grid = true
	bar.Vertical = true
	c.String(http.StatusOK, card.BarChartVertical(bar))
}

func getMostactivity(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	org := c.Request.FormValue("org")
	title := c.Request.FormValue("title")

	var color = style.CheckTheme(c)

	c.String(http.StatusOK, organization.MostactivityCard(title, org, color))
}
func repositoryCommitActivity(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	user := c.Request.FormValue("user")
	repo := c.Request.FormValue("repo")
	title := c.Request.FormValue("title")
	var hide_week string = c.Request.FormValue("hide_week")

	var color = style.CheckTheme(c)

	c.String(http.StatusOK, commit_activity.RepositoryCommitActivity(title, user, repo, hide_week, color))
}
func projectcard(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	user := c.Request.FormValue("user")
	repo := c.Request.FormValue("repo")

	var color = style.CheckTheme(c)

	c.String(http.StatusOK, project.Project(user, repo, color))
}
func language(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	title := c.Request.FormValue("title")
	user := c.Request.FormValue("user")
	org := c.Request.FormValue("organization")
	langs_count := c.Request.FormValue("langs_count")

	if len(org) > 0 && len(user) < 1 {
		user = org
	}
	var color = style.CheckTheme(c)

	c.String(http.StatusOK, languageCard.LanguageCard(title, user, langs_count, color))
}
func rankList(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	users := strings.Split(fmt.Sprintf("%v", c.Request.FormValue("users")), ",")
	title := c.Request.FormValue("title")

	var color = style.CheckTheme(c)

	c.String(http.StatusOK, rank.Rankcard(title, users, color))
}
func userstreak(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	user := c.Request.FormValue("user")
	hide_title := c.Request.FormValue("hide_title")

	var color = style.CheckTheme(c)

	c.String(http.StatusOK, streak.Streak(user, hide_title, color))
}
func getSkills(c *gin.Context) {
	c.Header("Content-Type", "image/svg+xml")

	// Define styles
	languages := strings.Split(c.Request.URL.Query().Get("languages"), ",")
	title := c.Request.FormValue("title")

	var color = style.CheckTheme(c)

	c.String(http.StatusOK, skills.Skills(title, languages, color))
}
