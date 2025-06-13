package test1

import (
	"fmt"
	"time"

	"github.com/segmentfault/pacman/log"
	"github.com/spf13/cobra"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "Example--gotest server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	fmt.Println("api/api.go: 0.init()...start")
	defer fmt.Println("api/api.go: 0.init()...end")

	StartCmd.PersistentFlags().StringVarP(&configYml,
		"config",
		"c",
		"config/settings.yml",
		"Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&apiCheck,
		"api",
		"a",
		false,
		"Start server with check api data")

}

func setup() {
	// 注入配置扩展项

	usageStr := `.....starting api server...`
	log.Info(usageStr)
}

func run() error {

	fmt.Println(string("global.LogoContent"))
	tip()
	fmt.Println(string("Server run at:"))
	fmt.Printf("-  Local:   %s://localhost:%d/ \r\n", "http", 8999)
	fmt.Printf("-  Network: %s://%s:%d/ \r\n", "http", "LocalHost", 8999)
	fmt.Println(string("Swagger run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/swagger/admin/index.html \r\n", 8999)
	fmt.Printf("-  Network: %s://%s:%d/swagger/admin/index.html \r\n", "http", "pkg.GetLocalHost()", 8999)
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", time.Now().String())

	log.Info(".........Shutdown Server ... ")

	log.Info("Server exiting")

	return nil
}

//var Router runtime.Router

func tip() {
	usageStr := `欢迎使用  go-admin v1.88, 可以使用 -h 查看命令`
	fmt.Printf("%s \n\n", usageStr)
}

func init() {
	fmt.Println("api/api.go: 1. init")
}

func init() {
	fmt.Println("api/api.go: 2. init")
}
