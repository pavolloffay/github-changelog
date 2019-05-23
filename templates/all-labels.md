# Changelog

{{range .Tags}}
{{- if .Tag }}
## {{.Tag.Name }}
{{ end -}}

{{ range $key, $value := .Labeled }}
### {{ $key }}
{{ range $value -}}
{{- if .Pull -}}
* {{.Pull.Title }} ([#{{.Pull.GetNumber}}]({{.Pull.GetHTMLURL}}), [@{{.Commit.Author.GetLogin}}]({{.Commit.Author.GetHTMLURL}}))
{{ end -}}
{{ end -}}
{{ end }}

### Without Label
{{range .NoLabeled}}
{{- if .Pull -}}
 * {{.Pull.Title }} ([#{{.Pull.GetNumber}}]({{.Pull.GetHTMLURL}}), [@{{.Commit.Author.GetLogin}}]({{.Commit.Author.GetHTMLURL}}))
{{ end -}}
{{ end -}}
{{- end}}
