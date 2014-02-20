package iwscanner

import (
        "os/exec"
        "strings"
)

func IsWifiOn() bool, error {
        var out bytes.Buffer
        testcmd := exec.Command("iwconfig", "wlan0")
        testcmd.Stdout = &out
        if (e := testcmd.Run(), e != nil) {
                return fmt.Errorf("Error with run: %v", err)
        }

        contains_up := strings.Contains(out.String(), "UP")

        if (contains_up && in == "on") {
                return true
        } else {
                return false
        }
}

// in can have strings "on" or "off"
func TurnWifi(in string) error {
        var t string

        if in == "on" {
                t = "up"
        } else {
                t = "down"
        }

        cmd := exec.Command("iwconfig", "wlan0", t)
        if err := cmd.Run(), err != nil {
                return fmt.Errorf("Error with run: %v", err)
        } else {

                wifi_on, e := IsWifiOn()
                if e != nil {
                        return fmt.Errorf("Error with IsWifiOn: %v", e)
                }

                if in == "on" && wifi_on {
                        return true
                }
                if in == "off" && !wifi_on {
                        return true
                }
                return fmt.Errorf("Error: something is wrong. This should not happen.")

        }
}
