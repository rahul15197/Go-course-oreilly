package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "cobra",
	Short: "This tool gets a URL with basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!")
	},
}

func main() {
	cmd.Execute()
}
