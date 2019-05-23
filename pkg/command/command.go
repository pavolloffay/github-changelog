package command

import (
	"flag"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	repoFlag       = "repo"
	ownerFlag      = "owner"
	oauthTokenFlag = "oauth-token"
	branchFlag     = "branch"
	templateFlag   = "template"
	logLevelFlag   = "log-level"
)

var (
	compiledTemplates = []string{"/chrono-list.md", "/all-labels.md"}
)

type Opts struct {
	Owner    string
	Repo     string
	Branch   string
	Template string
	Token    string
}

func NewCommand(fce func(otp Opts) error, v *viper.Viper) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "gch",
		Short: "gch is github changelog generator",
		Long:  `Golang template based changelog generator`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fce(getOptions(v))
		},
	}
	cmd.Flags().AddGoFlagSet(addFlags())
	v.BindPFlags(cmd.Flags())
	return cmd
}

func addFlags() *flag.FlagSet {
	flagSet := new(flag.FlagSet)
	flagSet.String(
		ownerFlag,
		"jaegertracing",
		"Github user or organization")
	flagSet.String(
		repoFlag,
		"jaeger-operator",
		"Github repository")
	flagSet.String(
		branchFlag,
		"master",
		"Git branch")
	flagSet.String(
		oauthTokenFlag,
		"",
		"Github OAUTH token")
	flagSet.String(
		templateFlag,
		compiledTemplates[0],
		fmt.Sprintf("Template name. Templates %v are compiled inside the binary", compiledTemplates))
	flagSet.String(
		logLevelFlag,
		"info",
		"Logrus log level")
	return flagSet
}

func getOptions(v *viper.Viper) Opts {
	o := Opts{}
	o.Owner = v.GetString(ownerFlag)
	o.Repo = v.GetString(repoFlag)
	o.Branch = v.GetString(branchFlag)
	o.Template = v.GetString(templateFlag)
	o.Token = v.GetString(oauthTokenFlag)
	return o
}

func GetLogLevel(v *viper.Viper) string {
	return v.GetString(logLevelFlag)
}
