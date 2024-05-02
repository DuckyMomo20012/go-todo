package cmd

import (
	"os"

	"github.com/DuckyMomo20012/go-todo/cmd/gateway"
	"github.com/DuckyMomo20012/go-todo/cmd/task"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "go-todo",
		Short: "A simple CLI for managing your tasks.",
		Long:  "A simple CLI for managing your tasks.",
	}
}

func Execute() {
	rootCmd := NewRootCmd()
	taskCmds := task.NewTaskCmd()
	gatewayCmd := gateway.NewGatewayCmd()

	rootCmd.AddCommand(taskCmds)
	rootCmd.AddCommand(gatewayCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
