package timer

import (
	"time"

	pkgtimer "github.com/congcongke/timer/pkg/timer"
	"github.com/spf13/cobra"
)

func NewTimerCommand() *cobra.Command {
	conf := pkgtimer.TimerConf{}

	cmd := &cobra.Command {
		Use:   "slottimer",
		Short: "slottimer is a simple executable period",
		Long:  "it is expected to execute the command periodly",
		Run: func(cmd *cobra.Command, args []string) {
			conf.Exec()
		},
	}

	cmd.PersistentFlags().StringVar(&conf.Command, "command", "", "the command need to run")
	cmd.PersistentFlags().DurationVar(&conf.Interval, "interval", time.Second, "inteval to run the command")
	cmd.PersistentFlags().StringSliceVar(&conf.Args, "args", []string{}, "the args of command")
	cmd.PersistentFlags().Int32Var(&conf.TotalTimes, "times", 10, "how many times the command will run")

	return cmd
}