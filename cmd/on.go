package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"litra/light"
)

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn the Litra light on",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := light.NewLitraGlow()
		if err != nil {
			return err
		}
		defer l.Close()

		if err := l.TurnOn(); err != nil {
			return err
		}
		fmt.Println("Light turned on.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}
