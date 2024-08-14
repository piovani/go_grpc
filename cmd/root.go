package cmd

import (
	"github.com/piovani/go_grpc/ui/grpc"
	"github.com/piovani/go_grpc/ui/http"
	"github.com/spf13/cobra"
)

func Execute() error {
	cmd := &cobra.Command{
		Use:   "go_gprc",
		Short: "this aplication this test to gprc vs rest",
	}

	cmd.AddCommand(
		Grpc,
		Http,
	)

	return cmd.Execute()
}

var Grpc = &cobra.Command{
	Use: "grpc",
	Run: func(cmd *cobra.Command, args []string) {
		grpc.Execute()
	},
}

var Http = &cobra.Command{
	Use: "http",
	Run: func(cmd *cobra.Command, args []string) {
		http.Execute()
	},
}
