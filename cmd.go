package main

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "train"}

	var logCmd = &cobra.Command{
		Use: "log",
		Short: "Log a new workout",
		Run: func(cmd *cobra.Command, args []string) {
			exercise, _ := cmd.Flags().GetString("exercise")
			weight, _ := cmd.Flags().GetFloat64("weight")
			sets, _ := cmd.Flags().GetInt("sets")
			repsStr, _ := cmd.Flags().GetString("reps")
			date, _ := cmd.Flags().GetString("date")

			reps := []int{}
			for _, rep := range strings.Split(repsStr, ",") {
				repInt, _ := strconv.Atoi(rep)
				reps = append(reps, repInt)
			}

			addLog(exercise, weight, sets, reps, date)
		},
	}

	logCmd.Flags().String("exercise", "", "Exercise name")
	logCmd.Flags().Float64("weight", 0, "Weight in kg")
	logCmd.Flags().Int("sets", 0, "Number of sets")
	logCmd.Flags().String("reps", "", "Comma-separated reps per set (e.g., 8,8,6)")
	logCmd.Flags().String("sets", "", "Date of the workout (YYYY-MM-DD)")

	rootCmd.AddCommand(logCmd)
	rootCmd.Execute()
}
