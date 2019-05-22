# Changelog

{{range .Commits}}
{{- if .Tag }}
## {{.Tag.Name }} ({{.Commit.Commit.Author.Date.Format "2006-01-02"}})
{{ end }}
{{- if .Pull -}}
 * {{.Pull.Title }} ([#{{.Pull.GetNumber}}]({{.Pull.GetHTMLURL}}), [@{{.Commit.Author.GetLogin}}]({{.Commit.Author.GetHTMLURL}}))
{{ end -}}
{{- end}}
