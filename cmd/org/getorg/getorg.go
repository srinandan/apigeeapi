package getorg

import (
	"net/url"
	"../../shared"
	"path"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Show details of an Apigee Org",
	Long:  "Show details of an Apigee Org",
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := url.Parse(shared.BaseURL)
		u.Path = path.Join(u.Path, shared.RootArgs.Org)
		shared.GetHttpClient(u.String(), shared.RootArgs.Token)
	},
}

func init() {

	Cmd.Flags().StringVarP(&shared.RootArgs.Org, "org", "o",
		"", "Apigee organization name")

	Cmd.MarkFlagRequired("org")		
}