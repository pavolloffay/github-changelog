package github

import (
	"context"

	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func GetAllTags(client *github.Client, owner, repo string) ([]*github.RepositoryTag, error) {
	var allTags []*github.RepositoryTag
	opts := github.ListOptions{PerPage: 100}
	for {
		tags, response, err := client.Repositories.ListTags(context.Background(), owner, repo, &opts)
		if err != nil {
			return nil, err
		}
		log.WithFields(log.Fields{"nextPage": response.NextPage}).
			Debug("Getting tags")
		allTags = append(allTags, tags...)
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
	return allTags, nil
}

func GetAllCommits(client *github.Client, owner, repo, branch string) ([]*github.RepositoryCommit, error) {
	var allCommits []*github.RepositoryCommit
	opts := github.ListOptions{PerPage: 100}
	for {
		repositoryCommits, response, err := client.Repositories.ListCommits(context.Background(), owner, repo,
			&github.CommitsListOptions{SHA: branch, ListOptions: opts})
		if err != nil {
			return nil, err
		}
		log.WithFields(log.Fields{"nextPage": response.NextPage}).
			Debug("Getting commits")
		allCommits = append(allCommits, repositoryCommits...)
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
	return allCommits, nil
}

func GetAllPullRequests(client *github.Client, owner, repo, branch string) ([]*github.PullRequest, error) {
	var allPullRequests []*github.PullRequest
	opts := github.ListOptions{Page: 0, PerPage: 100}
	for {
		pulls, response, err := client.PullRequests.List(context.Background(), owner, repo,
			&github.PullRequestListOptions{Base: branch, State: "closed", ListOptions: opts})
		if err != nil {
			return nil, err
		}
		log.WithFields(log.Fields{"nextPage": response.NextPage}).
			Debug("Getting pull requests")
		allPullRequests = append(allPullRequests, pulls...)
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
	return allPullRequests, nil
}

func CreateClient(oauthToken string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: oauthToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}
