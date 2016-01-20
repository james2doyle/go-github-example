package main

import (
	"fmt"
	"github.com/google/go-github/github"
	"gopkg.in/macaron.v1"
)

func fetch_repos(username string) []github.Repository {
	client := github.NewClient(nil)

	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, err := client.Repositories.List(username, opt)
	if err != nil {
		fmt.Printf("error: %v\n\n", err)
	}

	rate, _, err := client.RateLimits()
	if err != nil {
		fmt.Printf("Error fetching rate limit: %#v\n\n", err)
	} else {
		fmt.Printf("API Rate Limit: %#v\n\n", rate)
	}

	return repos
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Message"] = "Hello. Enter Your Github Username To Continue:"
		ctx.HTML(200, "hello") // 200 is the response code.
	})

	m.Get("/repos/:username", func(ctx *macaron.Context) {
		data := fetch_repos(ctx.Params(":username"))
		ctx.JSON(200, &data)
	})

	m.Run()
}
