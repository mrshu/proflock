package profile

import (
        "github.com/spf13/cobra"
        "fmt"
        "../iwscanner"
)

var wifi_device string
var ProfilesDir string

var CmdProfile = &cobra.Command{
                Use:   "profile [name of the profile]",
                Short: "Manage profile",
                Long:  `Create, update, delete or in some other way manage the given profile.`,
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("Managing")
                },
        }

var cmdShow = &cobra.Command{
                Use:   "show",
                Short: "Show profiles",
                Long:  `show profiles.`,
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Printf("Showing profiles in %s\n", ProfilesDir)
                },
        }

var cmdCreate = &cobra.Command{
                Use:   "create",
                Short: "Create profiles",
                Long:  `create profiles.`,
                Run: func(cmd *cobra.Command, args []string) {
                        if len(args) < 1 {
                                fmt.Println("Please specify a profile name")
                                return
                        }

                        fmt.Printf("Creating a profile %v\n", args[0])
                        aps, err := iwscanner.GetAPs(wifi_device)
                        if err != nil {
                                panic(err);
                        } else {
                                fmt.Printf("I have these APIs %v\n", aps)
                        }

                },
        }


func init () {
        CmdProfile.AddCommand(cmdShow)
        CmdProfile.AddCommand(cmdCreate)
        CmdProfile.PersistentFlags().StringVarP(&wifi_device, "device", "", "wlp2s0",
                                                "Use this wifi-enabled device.")
}
