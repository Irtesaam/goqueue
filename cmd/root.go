package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var todoFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "goq",
	Short:   "A professional CLI todo application",
	Version: "v1.0.0",
	Long: `A beautiful and professional command-line todo application built with Go.
This application allows you to manage your tasks efficiently with a modern CLI interface.

Complete documentation is available at https://github.com/yourusername/goqueue`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/goqueue/config.yaml)")
	rootCmd.PersistentFlags().StringVar(&todoFile, "file", "", "todo file path (default is $HOME/.config/goqueue/todos.json)")

	// Bind the persistent flags to viper
	viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Create config directory if it doesn't exist
		configDir := filepath.Join(home, ".config", "goqueue")
		os.MkdirAll(configDir, 0755)

		// Search config in config directory with name "config" (without extension).
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// Set default values
	home, _ := os.UserHomeDir()
	defaultTodoFile := filepath.Join(home, ".config", "goqueue", "todos.json")
	viper.SetDefault("file", defaultTodoFile)

	// Set up environment variable handling
	viper.SetEnvPrefix("GOQUEUE") // Environment variables will be GOQUEUE_FILE, etc.
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// GetTodoFile returns the configured todo file path
func GetTodoFile() string {
	if todoFile != "" {
		return todoFile
	}
	return viper.GetString("file")
}
