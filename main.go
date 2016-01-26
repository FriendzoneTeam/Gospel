package main

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)
func main() {
    var cmdNew = &cobra.Command{
        Use: "new [AppName]",
        Short: "Create new Gospel app in you current GOPATH/src/[AppName]",
        Long: `Create new Gospel app in your currente GOPATH \n
        i.e gospel new HelloGospel \n 
            cd GOPATH/src/HelloGospel
        `,
        Run: func(cmd *cobra.Command, args []string){
            if len(args) == 1 {
                gopath := os.Getenv("GOPATH")
                if gopath == "" {
                    fmt.Println("GOPATH no found :(")
                    os.Exit(1)
                }
                src := fmt.Sprintf("%s\\src", gopath)
                appName := args[0]
                appDir := fmt.Sprintf("%s\\%s", src, appName)
                err := os.Mkdir(appDir, 0777)
                if err != nil {
                    fmt.Printf("%s",err)
                }
            }
        },
    }
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

