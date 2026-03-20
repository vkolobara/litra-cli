package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"litra/light"
)

var brightnessCmd = &cobra.Command{
	Use:   "brightness <value>",
	Short: "Set, increase, or decrease the brightness (0–100)",
	Long: `Set the brightness to a specific percentage (0–100), or adjust it
relatively using --increase or --decrease.

Examples:
  litra brightness 80
  litra brightness 10 --increase
  litra brightness 10 --decrease`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		value, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid value %q: must be an integer", args[0])
		}

		increase := viper.GetBool("brightness.increase")
		decrease := viper.GetBool("brightness.decrease")
		if increase && decrease {
			return fmt.Errorf("--increase and --decrease are mutually exclusive")
		}

		l, err := light.NewLitraGlow()
		if err != nil {
			return err
		}
		defer l.Close()

		switch {
		case increase:
			if err := l.BrightnessIncrease(value); err != nil {
				return err
			}
			fmt.Printf("Brightness increased by %d%%.\n", value)
		case decrease:
			if err := l.BrightnessDecrease(value); err != nil {
				return err
			}
			fmt.Printf("Brightness decreased by %d%%.\n", value)
		default:
			if err := l.BrightnessSet(value); err != nil {
				return err
			}
			fmt.Printf("Brightness set to %d%%.\n", value)
		}
		return nil
	},
}

func init() {
	brightnessCmd.Flags().Bool("increase", false, "Increase brightness by the given amount")
	brightnessCmd.Flags().Bool("decrease", false, "Decrease brightness by the given amount")
	viper.BindPFlag("brightness.increase", brightnessCmd.Flags().Lookup("increase"))
	viper.BindPFlag("brightness.decrease", brightnessCmd.Flags().Lookup("decrease"))
	rootCmd.AddCommand(brightnessCmd)
}
