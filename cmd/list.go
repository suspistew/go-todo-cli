package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

import . "go-todo-cli/models"
import . "go-todo-cli/utils"

// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Aliases: []string{"l"}, 
    Short: "List all tasks",
    Long: `List all tasks. Only returns pending ones by default
    `,
    Run: func(cmd *cobra.Command, args []string) {
        MasterInit()
        print("\033[H\033[2J")
        tasks := LoadTasks()
        filteredTasks := FilterTasks(tasks)
        if len(filteredTasks) == 0 {
            fmt.Printf("\n📭 No task registered.\n\n")
            return
        }
        grouped := GroupByCategory(filteredTasks)
        DisplayInline(grouped)
    },
}

func DisplayInline(grouped map[string][]Task) {
    for category, taskList := range grouped {
        fmt.Printf("\n\n")
        c := category
        if c == "" {
            c = "No category"
        }
        fmt.Printf("\033[1m %s :\033[0m\n\n", c)
        for _, task := range taskList {
            fmt.Printf("    %d. %s (%s)\n", task.Id, task.Title, task.Status.String())
        }
    }
    fmt.Printf("\n\n")
}


func GroupByCategory(tasks []Task) map[string][]Task {
    grouped := make(map[string][]Task)
    for _, task := range tasks {
        grouped[task.Category] = append(grouped[task.Category], task)
    }

    return grouped
}

func init() {
    listCmd.Flags().StringVarP(&Filter, "filter", "f", "", "returns tasks matching the given filter")
    listCmd.Flags().BoolVarP(&Detail, "detail", "d", false, "provide detailed view for each task")
    listCmd.Flags().BoolVarP(&All, "all", "a", false, "include all status, event completed tasks")
    listCmd.Flags().StringVarP(&Category, "category", "c", "", "returns tasks matching the given category")
    listCmd.Flags().StringVarP(&Status, "status", "s", "", "returns tesks matching the given status (Pending|InProgress|Completed|Stashed)")
    rootCmd.AddCommand(listCmd)
}
