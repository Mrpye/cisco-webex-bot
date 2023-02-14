package cmd

import (
	_ "github.com/Mrpye/cisco-webex-bot/docs"
	"github.com/Mrpye/cisco-webex-bot/modules/api"
	"github.com/spf13/cobra"
)

func Cmd_WebServer() *cobra.Command {
	// webserverCmd represents the webserver command
	var web_port string
	var web_ip string

	var cmd = &cobra.Command{
		Use:   "web",
		Short: "Start a API Web-Server",
		Long:  `Start a API Web-Server`,
		Run: func(cmd *cobra.Command, args []string) {
			api.StartWebServer(web_ip, web_port)
		},
	}
	cmd.Flags().StringVarP(&web_port, "port", "p", "8080", "The Port to listen on")
	cmd.Flags().StringVarP(&web_ip, "ip", "i", "localhost", "IP to listen on")

	return cmd
}
func init() {
	rootCmd.AddCommand(Cmd_WebServer())
}
