package profile

import (
        "github.com/spf13/cobra"
        "fmt"
)

var CmdProfile = &cobra.Command{
                Use:   "profile [name of the profile]",
                Short: "Manage profile",
                Long:  `Create, update, delete or differently manage the given profile.`,
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("Managing")
                },
        }
