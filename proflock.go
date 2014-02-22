package main

import(
    "github.com/spf13/cobra"
    "fmt"
    "./iwscanner"
)

func main() {

    var cmdScan = &cobra.Command{
        Use:   "scan",
        Short: "Scan current location for APs",
        Long:  `Turns on WiFi and scans for AP around.`,
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Turning wifi on")

                if e := iwscanner.TurnWifi("wlan0", "on"); e != nil {
                        panic(e)
                }

                fmt.Println("Scanning")
                aps, err := iwscanner.GetAPs("wlan0")
                if err != nil {
                        panic(err);
                } else {
                        fmt.Println(aps)
                }
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

