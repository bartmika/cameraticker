package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/bartmika/cameraticker/internal/app"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the camera timer",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		runcameraticker()
	},
}

func runcameraticker() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	defer app.StopMainRuntimeLoop()

	// DEVELOPERS CODE:
	// The following code will create an anonymous goroutine which will have a
	// blocking chan `sigs`. This blocking chan will only unblock when the
	// golang app receives a termination command; therfore the anyomous
	// goroutine will run and terminate our running application.
	//
	// Special Thanks:
	// (1) https://gobyexample.com/signals
	// (2) https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	//
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs // Block execution until signal from terminal gets triggered here.
		fmt.Println("Starting graceful shut down now.")
		app.StopMainRuntimeLoop()
	}()

	app.RunMainRuntimeLoop()
}
