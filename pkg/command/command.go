package command

import "github.com/spf13/cobra"

var dayRCalCommand = &cobra.Command{
	Use:     "dayRCal [want/have] [goods name] [amount]",
	Aliases: []string{"drc"},
	Short:   "",
}
