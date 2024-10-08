/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"database/sql"
	"fmt"
	"github.com/felipeamendes/arq-hexagonal/application"
	"os"
	"github.com/spf13/cobra"
	dbInfra "github.com/felipeamendes/arq-hexagonal/adapters/db"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var db, _ = sql.Open("sqlite3","db.sqlite")
var productDb = dbInfra.NewProductDb(db)
var productService = application.ProductService{Persistence: productDb}

var rootCmd = &cobra.Command{
	Use:   "go-hexagonal",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-hexagonal.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-hexagonal")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
