# Changelog

{{range .Tags}}
{{- if .Tag }}
## {{.Tag.Name }}
{{ end -}}

### New features
{{ with $x := (index .Labeled "enhancement") }}{{range $key, $value := $x }}
{{- if .Pull -}}
* {{.Pull.Title }} ([#{{.Pull.GetNumber}}]({{.Pull.GetHTMLURL}}), [@{{.Commit.Author.GetLogin}}]({{.Commit.Author.GetHTMLURL}}))
{{ end -}}
{{end}}
{{ end -}}
{{ with $x := (index .Labeled "Enhancement") }}{{range $key, $value := $x }}
{{- if .Pull -}}
* {{.Pull.Title }} ([#{{.Pull.GetNumber}}]({{.Pull.GetHTMLURL}}), [@{{.Commit.Author.GetLogin}}]({{.Commit.Author.GetHTMLURL}}))
{{ end -}}
{{end}}
{{ end -}}

### Bug fixes, Minor improvements
{{ with $x := (index .Labeled "bug") }}{{range $key, $value := $x }}
{{- if .Pull -}}
* {{.Pull.Title }} ([#{{.Pull.GetNumber}}]({{.Pull.GetHTMLURL}}), [@{{.Commit.Author.GetLogin}}]({{.Commit.Author.GetHTMLURL}}))
{{ end -}}
{{end}}
{{ end -}}
{{ with $x := (index .Labeled "Bug") }}{{range $key, $value := $x }}
{{- if .Pull -}}
* {{.Pull.Title }} ([#{{.Pull.GetNumber}}]({{.Pull.GetHTMLURL}}), [@{{.Commit.Author.GetLogin}}]({{.Commit.Author.GetHTMLURL}}))
{{ end -}}
{{end}}
{{ end -}}

{{- end}}
