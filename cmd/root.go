package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdBuild, cmdDeploy)
}

var cmdBuild = &cobra.Command{
	Use:   "build [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		BuildImage(".", "test:v1")
	},
}

var cmdDeploy = &cobra.Command{
	Use:   "deploy [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		CreateDeploy()
	},
}

var rootCmd = &cobra.Command{
	Use:   "kubeoid",
	Short: "kubeoid is a simple build and deployment tool for local development on minikube",
	Long: `A simple tool that lets you build your application container images and deploy them on to your local kubernetes cluster.
			Complete documentation is available at http://kubeoid.io`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Hello this is a test entry")
	},
}

//Execute handles root command execution
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

 // HandleError logs the error and the message to the console
func HandleError(e error, msg string) {
	if e != nil {
		log.Fatal(e, msg)
	}
}
