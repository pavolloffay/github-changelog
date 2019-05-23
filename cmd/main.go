package main

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/pavolloffay/github-changelog/pkg/command"
	myGithub "github.com/pavolloffay/github-changelog/pkg/github"
	"github.com/pavolloffay/github-changelog/pkg/templates"
)

// TemplateData holds data passed to golang template.
type TemplateData struct {
	Commits []CommitBundle
	Tags    []TagBundle
}

// CommitBundle holds data associated with a commit.
type CommitBundle struct {
	Commit *github.RepositoryCommit
	Pull   *github.PullRequest
	Tag    *github.RepositoryTag
	Labels map[string]bool
}

// TagBundle holds a tag and commits associated with the tag.
// The commits are split per label(of associated PR).
// No labeled commits are added to a separate list.
type TagBundle struct {
	Labeled   map[string][]CommitBundle
	NoLabeled []CommitBundle
	Tag       *github.RepositoryTag
}

func main() {
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	cmd := command.NewCommand(run, v)

	logLevel := command.GetLogLevel(v)
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)

	err = cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func run(opts command.Opts) error {
	client := myGithub.CreateClient(opts.Token)
	tags, err := myGithub.GetAllTags(client, opts.Owner, opts.Repo)
	if err != nil {
		log.Fatal(err)
	}
	commits, err := myGithub.GetAllCommits(client, opts.Owner, opts.Repo, opts.Branch)
	if err != nil {
		log.Fatal(err)
	}
	pullRequests, err := myGithub.GetAllPullRequests(client, opts.Owner, opts.Repo, opts.Branch)
	if err != nil {
		log.Fatal(err)
	}

	shaPullMap := map[string]*github.PullRequest{}
	for _, pull := range pullRequests {
		if pull.MergeCommitSHA != nil {
			shaPullMap[*pull.MergeCommitSHA] = pull
		} else {
			log.WithField("url", pull.GetURL()).
				Debug("Pull request has merge SHA null")
		}
	}
	shaTagMap := map[string]*github.RepositoryTag{}
	for _, tag := range tags {
		if tag.GetCommit() != nil && tag.GetCommit().SHA != nil {
			shaTagMap[*tag.GetCommit().SHA] = tag
		} else {
			log.WithField("name", tag.Name).
				Debug("Tag has commit SHA null")
		}
	}

	var commitBundle []CommitBundle
	for _, commit := range commits {
		pull := shaPullMap[*commit.SHA]
		labels := map[string]bool{}
		if pull != nil {
			for _, label := range pull.Labels {
				labels[*label.Name] = true
			}
		}
		commitBundle = append(commitBundle, CommitBundle{
			Commit: commit,
			Pull:   pull,
			Tag:    shaTagMap[*commit.SHA],
			Labels: labels,
		})
	}

	var tagBundle []TagBundle
	for _, commit := range commitBundle {
		var t *TagBundle
		if commit.Tag != nil || len(tagBundle) == 0 {
			t = &TagBundle{Tag: commit.Tag, Labeled: map[string][]CommitBundle{}}
			tagBundle = append(tagBundle, *t)
		} else {
			t = &tagBundle[len(tagBundle)-1]
		}

		if len(commit.Labels) == 0 {
			t.NoLabeled = append(t.NoLabeled, commit)
		} else {
			for k, _ := range commit.Labels {
				t.Labeled[k] = append(t.Labeled[k], commit)
			}
		}
	}
	return generateOutput(opts.Template, &TemplateData{Commits: commitBundle, Tags: tagBundle})
}

func generateOutput(templateName string, data *TemplateData) error {
	fileContent, err := templates.FSString(false, templateName)
	if err != nil {
		all, err := ioutil.ReadFile(templateName)
		if err != nil {
			return err
		}
		fileContent = string(all)
	}
	t, err := template.New(templateName).Parse(fileContent)
	if err != nil {
		return err
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		return err
	}
	return nil
}
