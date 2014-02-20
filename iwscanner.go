package iwscanner

import (
        "os/exec"
        "strings"
)

// in can have strings "on" or "off"
func TurnWifi(in string) bool {
        var t string

        if in == "on" {
                t = "up"
        } else {
                t = "down"
        }

        cmd := exec.Command("iwconfig", "wlan0", t)
        if err := cmd.Run(), err != nil {
                return false
        } else {
                var out bytes.Buffer
                testcmd := exec.Command("iwconfig", "wlan0")
                testcmd.Stdout = &out
                if (e := testcmd.Run(), e != nil) {
                        return false
                }

                contains_up := strings.Contains(out.String(), "UP")

                if ((contains_up && in == "on") || (!contains_up && in != "on")) {
                        return true
                } else {
                        return false
                }
        }
}
