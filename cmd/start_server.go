package cmd

import (
	"os"

	"clippr/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	kubeConfigPath string
)

var rootCmd = &cobra.Command{
	Use:          "clippr",
	Short:        "Clippr Is a video management framework used to enable content creators to leverage the same content on many platforms",
	SilenceUsage: true,
}

var log *zap.Logger
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the remediator server",
	RunE: func(cmd *cobra.Command, args []string) error {
		runServer(log)
		return nil
	},
}

func Execute(logger *zap.Logger) {
	log = logger
	rootCmd.AddCommand(CreateServerCommand())
	if err := rootCmd.Execute(); err != nil {
		//nolint:gomnd
		os.Exit(1)
	}
}

func CreateServerCommand() *cobra.Command {
	serverCmd.Flags().StringVar(
		&kubeConfigPath,
		"kube-config",
		"/etc/kubernetes/kubeconfig",
		"Path to the cluster's kubeconfig. Defaults to: /etc/kubernetes/kubeconfig",
	)
	return serverCmd
}

func runServer(logger *zap.Logger) {
	server := server.NewServer(logger)

	if err := server.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}
