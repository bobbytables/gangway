// Copyright © 2016 Robert Ross <robert@creativequeries>

package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/bobbytables/gangway/server"
	"github.com/bobbytables/gangway/store/fake"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a gangway server",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

// listenAddr is the addres the server will be started on
var listenAddr string

func startServer() {
	s := server.NewServer(server.Config{}, &fake.Store{})

	logrus.Infof("starting server on %s", listenAddr)
	s.Listen(listenAddr)
}

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&listenAddr, "addr", ":8080", "the address to start the server on")
}
