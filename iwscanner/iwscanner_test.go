package iwscanner

import (
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestTurnWifi (t *testing.T) {
        e := TurnWifi("wlan0", "on")
        assert.Equal(t, e, nil)

        ret, err := IsWifiOn("wlan0")
        assert.Equal(t, ret, true)
        assert.Equal(t, err, nil)
}

func TestParseIwlist(t *testing.T) {
        instr := `wlan0     Scan completed :
          Cell 01 - Address: 94:44:52:CC:5A:F0
                    Channel:1
                    Frequency:2.412 GHz (Channel 1)
                    Quality=42/70  Signal level=-68 dBm
                    Encryption key:on
                    ESSID:"ivana"
                    Bit Rates:1 Mb/s; 2 Mb/s; 5.5 Mb/s; 11 Mb/s; 9 Mb/s
                              18 Mb/s; 36 Mb/s; 54 Mb/s
                    Bit Rates:6 Mb/s; 12 Mb/s; 24 Mb/s; 48 Mb/s
                    Mode:Master
                    Extra:tsf=0000001649ecc8df
                    Extra: Last beacon: 73ms ago
          Cell 02 - Address: 34:08:04:BF:BF:7A
                    Channel:1
                    Frequency:2.412 GHz (Channel 1)
                    Quality=33/70  Signal level=-77 dBm
                    Encryption key:on
                    ESSID:"Sanyo"
                    Bit Rates:1 Mb/s; 2 Mb/s; 5.5 Mb/s; 11 Mb/s; 9 Mb/s
                              18 Mb/s; 36 Mb/s; 54 Mb/s
                    Bit Rates:6 Mb/s; 12 Mb/s; 24 Mb/s; 48 Mb/s
                    Mode:Master
                    Extra:tsf=000001078f3a6a43
                    Extra: Last beacon: 73ms ago
                    IE: Unknown: 000553616E796F
                    IE: Unknown: 010882848B961224486C
                    IE: Unknown: 030101`
        o := parseIwlistOutput(instr)
        var aps APs
        aps = append(aps, AP{address: "94:44:52:CC:5A:F0", quality:42, essid:"ivana"})
        aps = append(aps, AP{address: "34:08:04:BF:BF:7A", quality:33, essid:"Sanyo"})

        assert.Equal(t, o, aps)
}
