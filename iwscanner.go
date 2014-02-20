package iwscanner

import (
        "os/exec"
        "strings"
        "fmt"
        "bytes"
)

func IsWifiOn() (ret bool, err error) {
        var out bytes.Buffer
        testcmd := exec.Command("ifconfig", "wlan0")
        testcmd.Stdout = &out
        if e := testcmd.Run(); e != nil {
                err = fmt.Errorf("Error with run: %v", e)
                return
        }

        contains_up := strings.Contains(out.String(), "UP")

        if (contains_up) {
                ret = true
        } else {
                ret = false
        }
        return
}

// in can have strings "on" or "off"
func TurnWifi(in string) error {
        var t string

        if in == "on" {
                t = "up"
        } else {
                t = "down"
        }

        cmd := exec.Command("ifconfig", "wlan0", t)
        if err := cmd.Run(); err != nil {
                return fmt.Errorf("Error with run: %v", err)
        } else {

                wifi_on, e := IsWifiOn()
                if e != nil {
                        return fmt.Errorf("Error with IsWifiOn: %v", e)
                }

                if in == "on" && wifi_on {
                        return nil
                }
                if in == "off" && !wifi_on {
                        return nil
                }
                return fmt.Errorf("Error: something is wrong. This should not happen.")

        }
}

type AP struct {
        address string
        quality int
        essid string
}

func GetAPs () (aps []AP, err error) {
        return
}
