package new

import (
    "github.com/spf13/cobra"
)

func New() cobra.Command{
    var cmd cobra.Command{}
    cmd.Use = ""
    return cmd
}