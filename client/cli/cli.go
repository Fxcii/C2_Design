package cli

import (
	"fmt"
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

const (
	logFileName = "sliver-client.log"
)

func initLogging(appDir string) *os.File {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logFile, err := os.OpenFile(path.Join(appDir, logFileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		panic(fmt.Sprintf("[!] Error opening file: %s", err))
	}
	log.SetOutput(logFile)
	return logFile
}

func init() {
	rootCmd.TraverseChildren = true

	// Create the console client, without any RPC or commands bound to it yet.
	// This created before anything so that multiple commands can make use of
	// the same underlying command/run infrastructure.
	con := console.NewConsole(false)

	// Import
	rootCmd.AddCommand(importCmd())

	// Client console.
	// All commands and RPC connection are generated WITHIN the command RunE():
	// that means there should be no redundant command tree/RPC connections with
	// other command trees below, such as the implant one.
	rootCmd.AddCommand(consoleCmd(con))

	// Implant.
	// The implant command allows users to run commands on slivers from their
	// system shell. It makes use of pre-runners for connecting to the server
	// and binding sliver commands. These same pre-runners are also used for
	// command completion/filtering purposes.
	rootCmd.AddCommand(implantCmd(con))

	// No subcommand invoked means starting the console.
	rootCmd.RunE, rootCmd.PostRunE = consoleRunnerCmd(con, true)

	// Completions
	carapace.Gen(rootCmd)
}

var rootCmd = &cobra.Command{
	Use:   "sliver-client",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
