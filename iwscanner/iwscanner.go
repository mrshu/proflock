package iwscanner

import (
        "os/exec"
        "strings"
        "fmt"
        "bytes"
        "strconv"
        "regexp"
)

func IsWifiOn(device string) (ret bool, err error) {
        var out bytes.Buffer
        testcmd := exec.Command("ifconfig", device)
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

func IsWifiOnIp(device string) (ret bool, err error) {
        var out bytes.Buffer
        testcmd := exec.Command("ip", "link", "show", device)
        testcmd.Stdout = &out
        if e := testcmd.Run(); e != nil {
                err = fmt.Errorf("Error with run: %v", e)
                return
        }

        // we'd like to match the UP in <.....,UP......>
        // that is why we want to match ,UP
        contains_up := strings.Contains(out.String(), ",UP")

        if (contains_up) {
                ret = true
        } else {
                ret = false
        }
        return

}

// in can have strings "on" or "off"
func TurnWifi(device string, in string) error {
        var t string

        on, e := IsWifiOn(device)
        if (on && in == "on" && e == nil) {
                return nil
        } else if (!on && in != "on" && e == nil) {
                return nil
        }

        if in == "on" {
                t = "up"
        } else {
                t = "down"
        }

        cmd := exec.Command("ifconfig", device, t)
        if err := cmd.Run(); err != nil {
                return fmt.Errorf("Error with run: %v", err)
        } else {

                wifi_on, e := IsWifiOn(device)
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
        Address string
        Quality int
        Essid string
}

type APs []AP

func GetAPs(device string) (aps APs, err error) {
        var out bytes.Buffer
        cmd := exec.Command("iwlist", device, "scan")
        cmd.Stdout = &out

        if e := cmd.Run(); e != nil {
                err = fmt.Errorf("Error with run: %v", e)
                return
        } else {
                list := parseIwlistOutput(out.String())
                aps = list
                err = nil
                return
        }
}

func GetAPsAsHash(device string) (map[string]AP, error) {
        aps, err := GetAPs(device)
        if err != nil {
                return nil, err
        }

        out := make(map[string]AP)
        for _, ap := range aps {
                out[ap.Address] = ap
        }

        return out, nil
}

func parseIwlistOutput(in string) (aps APs) {
        splits := strings.Split(in, "Cell")
        address_regex, _ := regexp.Compile("Address: ([0-9A-Z:]*)")
        quality_regex, _ := regexp.Compile("Quality=([0-9]+)")
        essid_regex, _ := regexp.Compile("ESSID:\"(.*)\"")

        first := true
        for _, split := range splits {
                if first {
                        first = false
                        continue
                }

                ap := AP{}
                address_match := address_regex.FindStringSubmatch(split)
                ap.Address = address_match[1]

                quality_match := quality_regex.FindStringSubmatch(split)
                i, _ := strconv.Atoi(quality_match[1])
                ap.Quality = i

                ap.Essid = essid_regex.FindStringSubmatch(split)[1]

                aps = append(aps, ap)
        }
        return
}

func parseIwOutput(in string) (aps APs) {
        splits := strings.Split(in, "BSS")
        address_regex, _ := regexp.Compile(" ([0-9a-z:]*)\\(on (.*)\\)")
        quality_regex, _ := regexp.Compile("signal: -([0-9]+)\\.\\d\\d dBm")
        essid_regex, _ := regexp.Compile("SSID: (.*)")

        first := true
        for _, split := range splits {
                if first {
                        first = false
                        continue
                }

                ap := AP{}
                address_match := address_regex.FindStringSubmatch(split)
                ap.Address = address_match[1]

                quality_match := quality_regex.FindStringSubmatch(split)
                i, _ := strconv.Atoi(quality_match[1])
                ap.Quality = 100 - i

                ap.Essid = essid_regex.FindStringSubmatch(split)[1]

                aps = append(aps, ap)
        }
        return
}
