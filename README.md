# Github changelog generator

This is a command line utility which generates changelog for a Github repository.
The changelog generation is customizable via golang [templates](https://golang.org/pkg/text/template/).

## Parameters
The executable supports these environmental properties
* `OWNER` - Github user name or organization
* `REPO` - repository
* `OAUTH_TOKEN` - Github oauth token
* `TEMPLATE` - template name, [Generate token here.](https://github.com/settings/tokens/new?description=GitHub%20Changelog%20Generator%20token)

