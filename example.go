package main

import (
        "./iwscanner"
        "fmt"
)

func main() {
        fmt.Println(iwscanner.GetAPs("wlan0"))
}
