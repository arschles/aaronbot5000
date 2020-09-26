package cmd

import (
	"context"
	"log"

	"github.com/arschles/aaronbot5000/pkg/db"
	"github.com/arschles/aaronbot5000/pkg/helix"
	"github.com/arschles/aaronbot5000/pkg/http"
	"github.com/spf13/cobra"
)

var forceStreamingOn bool

func init() {
	runCmd.Flags().BoolVarP(
		&forceStreamingOn,
		"streaming-on",
		"s",
		false,
		"Whether to force the bot to consider the stream on. Only valid if you don't have the 'OBS' module running",
	)

}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run chatbot server",
	Long:  `Use this command to start up the chatbot server.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		httpSrv := http.NewServer(":8080", "./web")
		go httpSrv.Start()
		helixCl, err := helix.GetHelixClient()
		if err != nil {
			log.Fatalf("Error getting helix client (%s)", err)
		}
		redisCl, err := db.New(ctx)
		if err != nil {
			log.Fatalf("Error getting redis client (%s)", err)
		}
		db.UpdateFollowers(ctx, "arschles", redisCl, helixCl)

		// err := db.InitDatabase(bot.DatabasePath(), 0600)
		// if err != nil {
		// 	if err.Error() == "timeout" {
		// 		log.Fatal("Timeout opening database. Check to ensure another process does not have the database file open")
		// 	}
		// 	log.Fatal("Failed to initialize database: ", err)
		// }

		// sig := make(chan os.Signal, 1)
		// signal.Notify(sig, os.Interrupt)
		// go func() {
		// 	<-sig

		// 	bot.ExecuteTrigger("bot::Shutdown", bot.Params{
		// 		Command: "shutdown",
		// 	})

		// 	if bot.IsModuleEnabled("OBS") {
		// 		obs.Disconnect()
		// 	}
		// 	os.Exit(0)
		// }()

		// // TODO: Handle scenario where startup trigger contains a twitch action
		// bot.ExecuteTrigger("bot::Startup", bot.Params{
		// 	Command: "startup",
		// })

		// if forceStreamingOn {
		// 	log.Printf(
		// 		"Bot started with '--streaming-on', forcing it into streaming status. This won't apply if you've enabled the OBS module.",
		// 	)
		// 	bot.Status.Streaming = true
		// }

		// if err := twitch.Run(); err != nil {
		// 	panic(err)
		// }
	},
}
