package cmd

import (
	"fmt"

	"github.com/sottey/redomvc/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "domaincheck [domain]",
	Short: "Check domain availability using the Name.com API",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Load()

		if file := viper.GetString("file"); file != "" {
			pkg.CheckFromFile(file)
		} else if len(args) == 1 {
			pkg.CheckSingleDomain(args[0])
		} else {
			fmt.Println("You must provide a domain or use the -f flag.")
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("username", "u", "", "API username")
	rootCmd.Flags().StringP("token", "t", "", "API token")
	rootCmd.Flags().StringP("file", "f", "", "File containing list of domains")
	rootCmd.Flags().IntP("workers", "w", 5, "Number of concurrent workers")
	rootCmd.Flags().IntP("delay", "d", 250, "Delay between requests (ms)")
	rootCmd.Flags().StringP("api", "a", "https://api.name.com/v4/domains:checkAvailability", "Name.com API URL")
	rootCmd.Flags().BoolP("verbose", "v", false, "More verbose output")

	viper.BindPFlags(rootCmd.Flags())
}
