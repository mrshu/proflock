package main

import(
        "github.com/spf13/cobra"
        "fmt"
        "./profile"
        "./iwscanner"
        "github.com/rakyll/globalconf"
)

var wifi_device string

func main() {

        conf, err := globalconf.New("proflock")
        if err != nil {
                panic(err)
        }

        var cmdScan = &cobra.Command{
                Use:   "scan",
                Short: "Scan current location for APs",
                Long:  `Turns on WiFi and scans for AP around.`,
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Printf("Turning %s on\n", wifi_device)

                        if e := iwscanner.TurnWifi(wifi_device, "on"); e != nil {
                                panic(e)
                        }

                        fmt.Println("Scanning")
                        aps, err := iwscanner.GetAPs(wifi_device)
                        if err != nil {
                                panic(err);
                        } else {
                                fmt.Println(aps)
                        }
                },
        }

        var cmdTurnWifi = &cobra.Command{
                Use:   "turn-wifi [on|off]",
                Short: "Turns Wifi On or Off",
                Long:  `Turns Wifi On or Off.`,
                Run: func(cmd *cobra.Command, args []string) {
                        if len(args) == 0 {
                                if e := iwscanner.TurnWifi(wifi_device, "on"); e != nil {
                                        panic(e)
                                }
                        } else {
                                if e := iwscanner.TurnWifi(wifi_device, args[0]); e != nil {
                                        panic(e)
                                }
                        }
                },
        }
        var rootCmd = &cobra.Command{Use: "proflock"}
        rootCmd.PersistentFlags().StringVarP(&wifi_device, "device", "", "wlp2s0",
                                                "Use this wifi-enabled device.")
        rootCmd.AddCommand(cmdScan, profile.CmdProfile, cmdTurnWifi)
        rootCmd.Execute()
}

