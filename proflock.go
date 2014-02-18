package main

import(
    "github.com/spf13/cobra"
    "fmt"
)

func main() {

    var cmdScan = &cobra.Command{
        Use:   "scan",
        Short: "Scan current location for APs",
        Long:  `Turns on WiFi and scans for AP around.`,
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Scanning")
        },
    }

    var cmdProfile = &cobra.Command{
        Use:   "profile [name of the profile]",
        Short: "Manage profile",
        Long:  `Create, update, delete or differently manage the given profile.`,
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Managing")
        },
    }
    var rootCmd = &cobra.Command{Use: "proflock"}
    rootCmd.AddCommand(cmdScan, cmdProfile)
    rootCmd.Execute()
}

