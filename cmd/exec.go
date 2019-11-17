/*
Copyright © 2019 Hidenori Sugiyama <madogiwa@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/madogiwa/trigger/watcher"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec [path] [command]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires path argument")
		}
		if len(args) < 2 {
			return errors.New("requires command argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		command := args[1]
		arg := args[2:]

		if err := watcher.Watch(path, command, arg); err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		fmt.Printf("watching for changes in %s ...\n", path)
		watcher.Start()

		exitSignal := make(chan os.Signal)
		signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
		<-exitSignal

		fmt.Println("shutdown...", command, path)
		watcher.Stop()
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
