package templates

import "github.com/google/go-github/github"

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
