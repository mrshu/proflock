package iwscanner

import (
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestTurnWifi (t *testing.T) {
        e := TurnWifi("wlp2s0", "on")
        assert.Equal(t, e, nil)

        ret, err := IsWifiOn("wlp2s0")
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

func TestParseIw(t *testing.T) {
        instr := `BSS 00:19:07:34:d7:42(on wlp2s0)
        TSF: 1193552217340 usec (13d, 19:32:32)
        freq: 2422
        beacon interval: 100 TUs
        capability: ESS Privacy ShortPreamble ShortSlotTime APSD (0x0c31)
        signal: -60.00 dBm
        last seen: 0 ms ago
        Information elements from Probe Response frame:
        SSID: eduroam
        Supported rates: 1.0* 2.0 5.5 6.0 9.0 11.0 12.0 18.0
        DS Parameter set: channel 3
        ERP: <no flags>
        RSN:	 * Version: 1
        	 * Group cipher: CCMP
        	 * Pairwise ciphers: CCMP
        	 * Authentication suites: IEEE 802.1X
        	 * Capabilities: 4-PTKSA-RC 4-GTKSA-RC (0x0028)
        Extended supported rates: 24.0 36.0 48.0 54.0
        WMM:	 * Parameter version 1
        	 * u-APSD
        	 * BE: CW 15-1023, AIFSN 3
        	 * BK: CW 15-1023, AIFSN 7
        	 * VI: CW 7-15, AIFSN 2, TXOP 3008 usec
        	 * VO: CW 3-7, AIFSN 2, TXOP 1504 usec
BSS 00:19:07:34:d7:40(on wlp2s0)
        TSF: 1193552350593 usec (13d, 19:32:32)
        freq: 2422
        beacon interval: 100 TUs
        capability: ESS Privacy ShortPreamble ShortSlotTime APSD (0x0c31)
        signal: -59.00 dBm
        last seen: 0 ms ago
        Information elements from Probe Response frame:
        SSID: FMFI_UK
        Supported rates: 1.0* 2.0 5.5 6.0 9.0 11.0 12.0 18.0
        DS Parameter set: channel 3
        ERP: <no flags>
        Extended supported rates: 24.0 36.0 48.0 54.0
        WPA:	 * Version: 1
        	 * Group cipher: TKIP
        	 * Pairwise ciphers: TKIP
        	 * Authentication suites: IEEE 802.1X
        	 * Capabilities: 4-PTKSA-RC 4-GTKSA-RC (0x0028)
        WMM:	 * Parameter version 1
        	 * u-APSD
        	 * BE: CW 15-1023, AIFSN 3
        	 * BK: CW 15-1023, AIFSN 7
        	 * VI: CW 7-15, AIFSN 2, TXOP 3008 usec
        	 * VO: CW 3-7, AIFSN 2, TXOP 1504 usec`
        o := parseIwOutput(instr)
        var aps APs
        aps = append(aps, AP{address: "00:19:07:34:d7:42", quality:40, essid:"eduroam"})
        aps = append(aps, AP{address: "00:19:07:34:d7:40", quality:41, essid:"FMFI_UK"})

        assert.Equal(t, o, aps)
}
