[![Build Status][ci-img]][ci] 

# Github changelog generator
This command line utility generates changelog for a Github repository.
The changelog generation is fully customizable via golang [templates](https://golang.org/pkg/text/template/).

## Example
Github requires OAUTH token to access its API. [Generate a token here.](https://github.com/settings/tokens/new?description=GitHub%20Changelog%20Generator%20token)

The executable can be downloaded from Github releases page.
```bash
docker run --rm  pavolloffay/gch:latest --oauth-token <github-oauth-token>
# the binary is in /app folder in the docker image
docker run --rm  -v "${PWD}:/app" pavolloffay/gch:latest --oauth-token <github-oauth-token> --template /app/templates/chrono-list.md
```

Examples can be found in [examples](./examples) directory.

### Templates
The binary contains predefined templates directly in the executable, however `--template` flag
can be used to supply any template.

* `chrono-list.md` - Chronologically order list of all pull requests split into tags.
* `all-labels.md` - Pull requests are split into labels, unlabeled commits are in a separate category.
* `default-labels.md` - Pull requests are split into bugs and enhancements.

The templates splitting pull requests into categories require proper labeling.

## Writing a custom template
Template uses golang [template](https://golang.org/pkg/text/template/) language to render the data.
The default examples can be found in [templates](./templates) directory and these are also compiled
into the binary. The object passed to the template is `TemplateData` defined in [pkg/templates/model.go](./pkg/templates/model.go).
The struct uses objects from [go-github](https://github.com/google/go-github) which are exposed in the templates.

```go
go run ./cmd/main.go --template <your-template>
```

## Develop

```bash
make build
./buld/gch
```

## License

[Apache 2.0 License](./LICENSE).


[ci-img]: https://travis-ci.org/pavolloffay/github-changelog.svg?branch=master
[ci]: https://travis-ci.org/pavolloffay/github-changelog
