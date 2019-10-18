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
	Use:   "build",
	Short: "Builds your container image",
	Long: `Build command builds your container images based on the configuration in the kubeoid.yaml file.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		BuildImage(".", "test:v1")
	},
}

var cmdDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys your container image to minikube",
	Long: `Deploy command deploys your container image to the local minikube cluster.`,
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
		fmt.Println("Hello from kubeoid to get started type 'kubeoid --help'")
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
