package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "root-cmd",
		Short: "root cmd",
		Long:  `root cmd`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("rootcmd")
		},
	}
}
