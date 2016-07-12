// Copyright © 2016 Robert Ross <robert@creativequeries>

package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/bobbytables/gangway/server"
	"github.com/bobbytables/gangway/store/etcd"
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

var (
	// listenAddr is the addres the server will be started on
	listenAddr string

	// etcdAddr is the etcd endpoint to use for storage
	etcdAddr string
)

func startServer() {
	estore, _ := etcdstore.NewStore([]string{etcdAddr})
	s := server.NewServer(server.Config{}, estore)

	logrus.Infof("starting server on %s", listenAddr)
	s.Listen(listenAddr)
}

func init() {
	RootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&listenAddr, "addr", ":8080", "the address to start the server on")
	startCmd.Flags().StringVar(&etcdAddr, "etcd-addr", "0.0.0.0:4001", "the address to start the server on")
}
