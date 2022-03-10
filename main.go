package main

import (
	"fmt"
	"githubembedapi/card/style"
	"githubembedapi/commit_activity"
	"githubembedapi/organization"
	"githubembedapi/project"
	"githubembedapi/rank"
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
	router.Run("localhost:8080")
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

	if title == "" || len(title) <= 0 {
		title = "Skills"
	}

	c.String(http.StatusOK, skills.Skills(title, languages, color))
}
