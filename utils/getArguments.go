package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// func GetArguments() (*string, *int, *bool) {
// 	problemFile := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
// 	testTime := flag.Int("time", 30, "Specifies the total time the quiz is going to run for.")
// 	shuffle := flag.Bool("shuffle", false, "Whether or not to shuffle questions while asking.")
// 	flag.Parse()
// 	return problemFile, testTime, shuffle
// }

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A quiz application",
	Long:  "A quiz application that reads the questions from a CSV or JSON file, with support for Question andwer type problems or MCQ problems",
	Run: func(cmd *cobra.Command, args []string) {
		//do something ?
	},
}

func init() {
	rootCmd.PersistentFlags().String("csv", "problems.csv", "a csv file in the format of 'question,answer' ")
	rootCmd.PersistentFlags().Int("time", 30, "Specifies the total time the quiz is going to run for")
	rootCmd.PersistentFlags().Bool("shuffle", false, "Whether or not to present the problems in a random order")
	viper.BindPFlag("csv", rootCmd.PersistentFlags().Lookup("csv"))
	viper.BindPFlag("time", rootCmd.PersistentFlags().Lookup("time"))
	viper.BindPFlag("shuffle", rootCmd.PersistentFlags().Lookup("shuffle"))

}

func GetArguments() (*string, *int, *bool) {
	csv := viper.GetString("csv")
	time := viper.GetInt("time")
	shuffle := viper.GetBool("shuffle")
	return &csv, &time, &shuffle

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		ExitWithMessage("Something went wrong", 1)
	}
}
