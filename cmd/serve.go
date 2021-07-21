package cmd

import (
	"github.com/b-nova-openhub/fisicus/pkg/rest"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	siteMap string
	filter string

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve filter-sitemap-customizer service",
		Long:  ``,
		Run:   serve,
	}
)

func init() {
	serveCmd.PersistentFlags().StringVar(&siteMap, "sitemap", "", "This is the fully qualified url where the target sitemap resides.")
	serveCmd.PersistentFlags().StringVar(&filter, "filter", "", "These is a list of patterns to filter target sitemap entries. Only entries are shown in the output sitemap if the pattern matches.")
	viper.BindPFlag("sitemap", serveCmd.PersistentFlags().Lookup("sitemap"))
	viper.BindPFlag("filter", serveCmd.PersistentFlags().Lookup("filter"))
}

func serve(ccmd *cobra.Command, args []string) {
	rest.HandleRequests()
}
