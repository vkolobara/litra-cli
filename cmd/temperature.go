package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"litra/light"
)

var temperatureCmd = &cobra.Command{
	Use:   "temperature <value>",
	Short: fmt.Sprintf("Set, increase, or decrease the color temperature (%d–%d K)", light.MinTemperature, light.MaxTemperature),
	Long: fmt.Sprintf(`Set the color temperature in Kelvin (%d–%d), or adjust it
relatively using --increase or --decrease.

Examples:
  litra temperature 4000
  litra temperature 200 --increase
  litra temperature 200 --decrease`, light.MinTemperature, light.MaxTemperature),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		value, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid value %q: must be an integer", args[0])
		}

		increase := viper.GetBool("temperature.increase")
		decrease := viper.GetBool("temperature.decrease")
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
			if err := l.TemperatureIncrease(value); err != nil {
				return err
			}
			fmt.Printf("Temperature increased by %d K.\n", value)
		case decrease:
			if err := l.TemperatureDecrease(value); err != nil {
				return err
			}
			fmt.Printf("Temperature decreased by %d K.\n", value)
		default:
			if err := l.TemperatureSet(value); err != nil {
				return err
			}
			fmt.Printf("Temperature set to %d K.\n", value)
		}
		return nil
	},
}

func init() {
	temperatureCmd.Flags().Bool("increase", false, "Increase temperature by the given amount")
	temperatureCmd.Flags().Bool("decrease", false, "Decrease temperature by the given amount")
	viper.BindPFlag("temperature.increase", temperatureCmd.Flags().Lookup("increase"))
	viper.BindPFlag("temperature.decrease", temperatureCmd.Flags().Lookup("decrease"))
	rootCmd.AddCommand(temperatureCmd)
}
