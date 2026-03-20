package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"litra/light"
)

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn the Litra light off",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := light.NewLitraGlow()
		if err != nil {
			return err
		}
		defer l.Close()

		if err := l.TurnOff(); err != nil {
			return err
		}
		fmt.Println("Light turned off.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
