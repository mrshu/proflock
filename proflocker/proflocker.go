package proflocker

import (
        "../iwscanner"
        "os"
        "encoding/json"
        "io/ioutil"
        "strings"
        "os/exec"
        "math"
)

type APscore struct {
        Essid string
        Address string
        Score float64
        Score_total float64
}

type Location struct {
        Name string
        Path string
        Aps []APscore
        Aps_score float64
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
                                Essid: ap.Essid,
                                Address: ap.Address,
                                Score: scores[ap.Essid].Score + float64(ap.Quality),
                                Score_total: scores[ap.Essid].Score_total + 70,
                        }

                        scores[ap.Address] = score
                }
        }

        for _, value := range scores {
                // making sure the resulting score is the weighted average
                value.Score = value.Score/(value.Score_total/70)
                value.Score_total = value.Score_total/(value.Score_total/70)

                location.Aps = append(location.Aps, value)
                location.Aps_score += value.Score
        }

        return location, nil
}


func ParseLocationsDir(dir string) (Locations, error) {
        locations := Locations{}

        contents, err := ioutil.ReadDir(dir)
        if err != nil {
                return nil, err
        }

        total_sum := 0.0
        for _, f := range contents {
                if f.IsDir() {
                        loc, err := ParseLocation(dir + "/" + f.Name() + "/data", f.Name())
                        if err == nil {
                                total_sum += loc.Aps_score
                                locations = append(locations, loc)
                        }
                }
        }

        for _, loc := range locations {
                loc.Aps_score = loc.Aps_score/total_sum
        }

        return locations, nil
}

func BuildFrequecyScores(location Location) (map[string]APscore) {
        scores := make(map[string]APscore)
        for _, ap := range location.Aps {
                score := APscore{
                        Essid: ap.Essid,
                        Address: ap.Address,
                        Score: scores[ap.Essid].Score + ap.Score,
                        Score_total: scores[ap.Essid].Score_total + ap.Score_total,
                }
                scores[ap.Address] = score
        }

        return scores
}

func ApproximateScore(ap iwscanner.AP, frequencies map[string]APscore) (float64) {
        prob := frequencies[ap.Address].Score
        if prob == 0.0 {
                return 0.0
        }

        prob = 70-math.Abs(prob-float64(ap.Quality))
        return prob
}

func RunHook(hook string, profile string, profiles_dir string) (error) {
        _, err := exec.Command(profiles_dir + "/" + profile + "/hooks.d/" + hook + ".sh").Output()
        return err
}
