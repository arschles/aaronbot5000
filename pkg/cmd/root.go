package cmd

import (
	_ "github.com/erikstmartin/erikbotdev/modules/keylight" // TODO: Remove this after we have cobra cmd
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
	initHueCmd()
}

var rootCmd = &cobra.Command{
	Use:   "erikbotdev",
	Short: "Twitch Bot",
	Long:  `Twitch bot for ErikDotDev`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
