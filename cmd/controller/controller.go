package controller

import (
    "github.com/spf13/cobra"
)
func Generate() *cobra.Command {
    var cmd = &cobra.Command{
        Use: "generate [command]",
        Short: "Generate a controller, model or view",
        Long: `
        Create new Gospel app in your currente GOPATH
        i.e gospel new HelloGospel \n 
            cd GOPATH/src/HelloGospel
        `,
        Run: exec,
    }
    // Helper
    cmd.SetUsageTemplate("generate [AppName]")    
    return cmd
}

func exec(cmd *cobra.Command, args []string) {
    if (len(args) == 0){
    }
}