package timer

import (
	"time"

	pkgtimer "github.com/congcongke/timer/pkg"
	"github.com/spf13/cobra"
)

func NewTimerCommand() *cobra.Command {
	conf := pkgtimer.TimerConf{}

	cmd := &cobra.Command {
		Use:   "slottimer",
		Short: "slottimer is a simple executable period",
		Long:  "it is expected to execute the command periodly",
		Run: func(cmd *cobra.Command, args []string) {
			udpTimer, err := pkgtimer.NewUdpTimer(&conf)
			if err != nil {
				panic("create udptimer failed")
			}
			udpTimer.Exec()
		},
	}

	cmd.PersistentFlags().DurationVar(&conf.Interval, "interval", time.Second, "inteval to run the command")
	cmd.PersistentFlags().IntVar(&conf.TotalTimes, "times", 10, "how many times the command will run")
	cmd.PersistentFlags().StringVar(&conf.Filename, "file", "", "file of data")
	cmd.PersistentFlags().IntVar(&conf.Block, "block", 50000, "how many block times to run")
	cmd.PersistentFlags().StringVar(&conf.Destination, "dest", "", "destination in xxx:xx")

	return cmd
}