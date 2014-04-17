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

var cmdShow = &cobra.Command{
                Use:   "show",
                Short: "Show profiles",
                Long:  `show profiles.`,
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("showing profiles")
                },
        }

var cmdCreate = &cobra.Command{
                Use:   "create",
                Short: "Create profiles",
                Long:  `create profiles.`,
                Run: func(cmd *cobra.Command, args []string) {
                        if len(args) < 1 {
                                fmt.Println("Please specify a profile name")
                        }

                        fmt.Printf("Creating a profile %v\n", args[0])
                },
        }


func init () {
        CmdProfile.AddCommand(cmdShow)
        CmdProfile.AddCommand(cmdCreate)
}
