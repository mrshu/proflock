package main

import(
        "github.com/spf13/cobra"
        "fmt"
        "path"
        "./profile"
        "./iwscanner"
        "github.com/rakyll/globalconf"
        "./proflocker"
)

var wifi_device string
var profiles_dir string
var flag_profiles_dir string
var log_probs bool

func OverrideProfileDir() {
        if flag_profiles_dir != "" {
                profiles_dir = flag_profiles_dir
        }
}

func main() {

        conf, err := globalconf.New("proflock")
        if err != nil {
                panic(err)
        }

        profiles_dir = path.Dir(conf.Filename) + "/profiles"

        log_probs = false

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

         var cmdIsWifi= &cobra.Command{
                Use:   "is-wifi [on|off]",
                Short: "Is Wifi On or Off",
                Long:  `Is Wifi On or Off?`,
                Run: func(cmd *cobra.Command, args []string) {
                        on, e := iwscanner.IsWifiOn(wifi_device)
                        if e != nil {
                                panic(e);
                        }

                        fmt.Println(on)

                },
        }

        var cmdProfiles = &cobra.Command{
                Use:   "profiles",
                Short: "Show profiles",
                Long:  `Show profiles.`,
                Run: func(cmd *cobra.Command, args []string) {
                        OverrideProfileDir()
                        profiles, err := proflocker.ParseLocationsDir(profiles_dir)
                        if err != nil {
                                panic(err);
                        }

                        if len(profiles) == 0 {
                                fmt.Println("fatal: no profile found")
                        } else {
                                fmt.Printf("The following profiles were found in %s:\n", profiles_dir)
                                for _, profile := range profiles {
                                        fmt.Println(profile.Name)
                                }
                        }


                },
        }

        var cmdRecord = &cobra.Command{
                Use:   "record [location]",
                Short: "Record location to use it as a profile afterwards.",
                Long:  `Record location to use it as a profile afterwards.`,
                Run: func(cmd *cobra.Command, args []string) {
                        if len(args) < 1 {
                                fmt.Println("Location name is required.")
                                return
                        }

                        OverrideProfileDir()
                        err := proflocker.RecordLocation(args[0], profiles_dir, wifi_device)
                        if err != nil {
                                panic(err)
                        }
                },
        }

        var cmdDumpProfile = &cobra.Command{
                Use:   "dump-profile [location]",
                Short: "Dump the data we have for a location",
                Long:  "Dump the data we have for a location",
                Run: func(cmd *cobra.Command, args []string) {
                        if len(args) < 1 {
                                fmt.Println("Location name is required.")
                                return
                        }

                        name := args[0]

                        OverrideProfileDir()
                        location, err := proflocker.ParseLocation(profiles_dir + "/" + name + "/data", name)
                        if err != nil {
                                panic(err)
                        }

                        fmt.Println(location.Aps)

                        //TODO: dump the location data in a serious way.
                },
        }

        var cmdShow = &cobra.Command{
                Use:   "show [location]",
                Short: "Shows location's data.",
                Long:  `Shows location's data.`,
                Run: func(cmd *cobra.Command, args []string) {
                        if len(args) < 1 {
                                fmt.Println("Location name is required.")
                                return
                        }

                        name := args[0]

                        OverrideProfileDir()
                        location, err := proflocker.ParseLocation(profiles_dir + "/" + name + "/data", name)
                        if err != nil {
                                panic(err)
                        }

                        fmt.Println(location)
                },
        }

        var cmdLocate = &cobra.Command{
                Use:   "locate",
                Short: "Locate using prerecorded locations.",
                Long:  `Locate using prerecorded locations.`,
                Run: func(cmd *cobra.Command, args []string) {
                        OverrideProfileDir()
                        profiles, err := proflocker.ParseLocationsDir(profiles_dir)
                        if err != nil {
                                panic(err);
                        }


                        var probabs = make(map[string]float64)
                        aps, err := iwscanner.GetAPsAsHash(wifi_device)

                        for _, profile := range profiles {
                                score := profile.Aps_score
                                frequencies := proflocker.BuildFrequecyScores(profile)

                                if err != nil {
                                        panic(err)
                                }

                                for _, ap := range aps {
                                        prob := proflocker.ApproximateScore(ap, frequencies)/100

                                        // fmt.Printf("%s\t%f\n", ap.Essid, prob)

                                        if prob == 0.0 {
                                                prob = 0.01
                                        }

                                        score = score * prob
                                }

                                probabs[profile.Name] = score
                        }

                        max_profile := "no_profile"
                        max_probab := 0.0
                        for profile, probab := range probabs {
                                if log_probs {
                                        fmt.Printf("%s\t\t%f\n", profile, probab)
                                }
                                if probab > max_probab {
                                        max_profile = profile
                                        max_probab = probab
                                }
                        }

                        if max_probab > 1.0 {
                                fmt.Printf("Seems like you are located at '%s'.\n", max_profile)
                                err := proflocker.RunHook("located", max_profile, profiles_dir)
                                if err != nil {
                                        panic(err);
                                }

                        } else {
                                fmt.Println("It is not possible to say where you are located.")
                        }
                },
        }

        cmdLocate.Flags().BoolVarP(&log_probs, "log-probabilities", "v", false,
                                                "log probabilities for profiles while locating")


        var rootCmd = &cobra.Command{Use: "proflock"}
        rootCmd.PersistentFlags().StringVarP(&wifi_device, "device", "", "wlp4s0",
                                                "wifi-enabled device to be used")
        rootCmd.PersistentFlags().StringVarP(&flag_profiles_dir, "profiledir", "P", "",
                                                "the directory with profiles in it")


        rootCmd.AddCommand(cmdScan, profile.CmdProfile, cmdTurnWifi, cmdProfiles, cmdRecord, cmdShow,
                                cmdLocate, cmdIsWifi, cmdDumpProfile)
        rootCmd.Execute()
}

