# Github changelog generator

This command line utility generates changelog for a Github repository.
The changelog generation is fully customizable via golang [templates](https://golang.org/pkg/text/template/).

## Example
Github requires OAUTH token to access its API. [Generate a token here.](https://github.com/settings/tokens/new?description=GitHub%20Changelog%20Generator%20token)

```bash
gch -h
gch --oauth-token <github-oauth-token> --repo jaeger-operator
gch --oauth-token <github-oauth-token> --repo jaeger-operator --template ./templates/chrono-list.md
```

The binary contains predefined templates directly in the executable, however `--template` flag
can be used to supply any template. The `main.go` contains a definition of objects which is passed to 
the template. It is basically a list of commits with attached labels and pull requests and tags.

## Develop

```bash
make build
```
