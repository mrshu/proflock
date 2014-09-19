package proflocker

import (
        "../iwscanner"
        "os"
        "encoding/json"
        "io/ioutil"
        "strings"
)

type APscore struct {
        essid string
        address string
        score float64
        score_total float64
}

type Location struct {
        Name string
        Path string
        Aps []APscore
}

type Locations []Location

func RecordLocation(location string, profile_dir string, device string) (error) {
        os.MkdirAll(profile_dir + "/" + location, 755)

        aps, err := iwscanner.GetAPs(device)
        if err != nil {
                return err
        }
        out, err := json.Marshal(aps)
        if err != nil {
                return err
        }

        path := profile_dir + "/" + location + "/data"

        f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
        if err != nil {
                f, err = os.Create(path)
        }

        defer f.Close()

        if _, err = f.WriteString(string(out) + "\n"); err != nil {
                return err
        }

        return nil
}

func ParseLocation(path string, name string) (Location, error) {
        location := Location{Path: path, Name: name}

        data, err := ioutil.ReadFile(path)
        if err != nil {
                return location, err
        }

        scores := make(map[string]APscore)
        lines := strings.Split(string(data), "\n")
        for _, line := range lines {
                tmp := &iwscanner.APs{}
                json.Unmarshal([]byte(line), &tmp)

                for _, ap := range *tmp {
                        score := APscore{
                                essid: ap.Essid,
                                address: ap.Address,
                                score: scores[ap.Essid].score + float64(ap.Quality),
                                score_total: scores[ap.Essid].score_total + 70,
                        }

                        scores[ap.Essid] = score
                }
        }

        for _, value := range scores {
                // making sure the resulting score is the weighted average
                value.score = value.score/(value.score_total/70)
                value.score_total = value.score_total/(value.score_total/70)

                location.Aps = append(location.Aps, value)
        }

        return location, nil
}


func ParseLocationsDir(dir string) (Locations, error) {
        locations := Locations{}

        contents, err := ioutil.ReadDir(dir)
        if err != nil {
                return nil, err
        }

        for _, f := range contents {
                if f.IsDir() {
                        loc, err := ParseLocation(dir + "/" + f.Name() + "/data", f.Name())
                        if err == nil {
                                locations = append(locations, loc)
                        }
                }
        }
        return locations, nil
}
