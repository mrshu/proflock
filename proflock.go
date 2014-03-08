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

                if e := iwscanner.TurnWifi("wlp2s0", "on"); e != nil {
                        panic(e)
                }

                fmt.Println("Scanning")
                aps, err := iwscanner.GetAPs("wlp2s0")
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

    var cmdTurnWifi = &cobra.Command{
        Use:   "turn-wifi [on|off]",
        Short: "Turns Wifi On or Off",
        Long:  `Turns Wifi On or Off.`,
        Run: func(cmd *cobra.Command, args []string) {
                if len(args) == 0 {
                        iwscanner.TurnWifi("wlp2s0", "on")
                } else {
                        iwscanner.TurnWifi("wlp2s0", args[0])
                }
        },
    }
    var rootCmd = &cobra.Command{Use: "proflock"}
    rootCmd.AddCommand(cmdScan, cmdProfile, cmdTurnWifi)
    rootCmd.Execute()
}

