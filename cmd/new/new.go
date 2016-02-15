package new

import (
    "log"
    "fmt"
    "os"
    "github.com/spf13/cobra"
    "github.com/fatih/color"
)

func New() *cobra.Command {
    var cmd = &cobra.Command{
        Use: "new [AppName]",
        Short: "Create new Gospel app in you current GOPATH/src/[AppName]",
        Long: `
        Create new Gospel app in your currente GOPATH
        i.e gospel new HelloGospel \n 
            cd GOPATH/src/HelloGospel
        `,
        Run: exec,
    }
    // Helper
    cmd.SetUsageTemplate("new [AppName]")    
    return cmd
}

func exec(cmd *cobra.Command, args []string){
    if len(args) == 1 {
        color.Green("Creando nueva aplicacion " + args[0])
        gopath := os.Getenv("GOPATHs")
        if gopath == "" {
            color.Set(color.FgRed)
            defer color.Unset()
            log.Fatalln("GOPATH no found :(")
            
            os.Exit(2)
        }
        src := fmt.Sprintf("%s\\src", gopath)
        appName := args[0]
        appDir := fmt.Sprintf("%s\\%s", src, appName)
        
        createAppFolder(appDir, []string{})
        fmt.Printf("appDir: %s\n",appDir)
        createAppFolder(fmt.Sprintf("%s\\%s", appDir, "public"), []string{"assets"})
        createAppFolder(fmt.Sprintf("%s\\%s", appDir, "app"), []string{"controllers", "models"})
        
        createSubFolder(fmt.Sprintf("%s\\%s\\%s", appDir, "public", "assets"), []string{"js", "scss", "img", "fonts"})
        // creamos la estructura basica
        
    }
}

func createAppFolder(basePath string, subdirs []string) {
    fmt.Printf("basePath: %s\n",basePath)
    err := os.Mkdir(basePath, 0777)
    if err != nil {
        log.Fatalln("%s",err)
    }
    if len(subdirs) != 0 {
        for _, subdir := range subdirs {
            temp := fmt.Sprintf("%s\\%s", basePath, subdir)
            err := os.Mkdir(temp, 0777)
            if err != nil {
                log.Fatalln("%s",err)
            }
        }
    }
}

func createSubFolder(basePath string, subdirs []string) {
    fmt.Printf("basePath: %s\n",basePath)
    if len(subdirs) != 0 {
        for _, subdir := range subdirs {
            temp := fmt.Sprintf("%s\\%s", basePath, subdir)
            err := os.Mkdir(temp, 0777)
            if err != nil {
                log.Fatalln("%s",err)
            }
        }
    }
}

func copyBaseFiles(){
    // copia el main.go
    // crea un package.json en public
    // crea un index.jade
}