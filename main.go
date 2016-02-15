package main

import (
    "fmt"
    "github.com/spf13/cobra"
    cmd "Gospel/cmd/new"
)
func main() {
    // Comando new
    var cmdNew = cmd.New()

    var cmdGenerate = &cobra.Command{
        Use: "generate [command]",
        Short: "Generate controller, model or view",
        Long: `generate gospel controller,model or view for
        the application.`,
        Run: func(cmd *cobra.Command, args []string){
            if len(args) == 0 {
                fmt.Println("generate exec help")
                
            }
            if len(args) == 1 {
                switch args[0] {
                case "controller":
                    fmt.Println("gospel generate controller")
                
                case "model":
                    fmt.Println("gospel generate model")
                }
            }
        },
    }
    
    var rootCmd = &cobra.Command{Use: "gospel"}
    rootCmd.AddCommand(cmdNew, cmdGenerate)
    rootCmd.Execute()
}

