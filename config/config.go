package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var cfgFile string

var cmd = &cobra.Command{
	Use:   "mlp",
	Short: "API for MLP Platform.",
	Long:  `API for Machine Learning Platform Developed By SanDataSystem.`,
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/mlp/mlp.yaml)")
}

func initConfig() {
	confFilePath := "/etc/mlp"
	confFileName := "mlp.yaml"

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(confFilePath)
		viper.SetConfigName(confFileName)
	}
	if Exists(filepath.Join(confFilePath, confFileName)) == false {
		log.Fatal("File: %s not found in %s directory", confFileName, confFilePath)
	}
	viper.AutomaticEnv()
	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal(err)
	}
}
