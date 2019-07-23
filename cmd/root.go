package cmd

import (
	"fmt"
	"os"

	"github.com/google/tcpproxy"

	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tcpproxy",
		Short: `tcp proxy server`,
		Long:  `tcp proxy server`,
	}

	subCommand = &cobra.Command{
		Use:   "start",
		Short: `Run to tcp proxy server`,
		Long:  `Run to tcp proxy server`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if viper.GetString("raddr") == "" || viper.GetString("rport") == "" {
				fmt.Println(Red("[err] raddr or rport args don't exist"))
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var proxy tcpproxy.Proxy
			proxy.AddRoute(fmt.Sprintf(":%s", viper.GetString("lport")),
				tcpproxy.To(fmt.Sprintf("%s:%s", viper.GetString("raddr"), viper.GetString("rport"))))
			if err := proxy.Run(); err != nil {
				fmt.Println(Red(err))
				os.Exit(1)
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(Red(err))
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().String("lport", "10000", "local listen port")
	rootCmd.PersistentFlags().String("rport", "", "[required] remote port")
	rootCmd.PersistentFlags().String("raddr", "", "[required] remote address")

	// mapping viper
	viper.BindPFlag("lport", rootCmd.PersistentFlags().Lookup("lport"))
	viper.BindPFlag("rport", rootCmd.PersistentFlags().Lookup("rport"))
	viper.BindPFlag("raddr", rootCmd.PersistentFlags().Lookup("raddr"))

	rootCmd.AddCommand(subCommand)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.ReadInConfig()
}
