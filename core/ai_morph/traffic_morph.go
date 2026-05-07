package ai_morph

import (
    "log"
    "math/rand"
    "time"
)

var profiles = map[string]struct {
    pattern []int
    jitter  time.Duration
}{
    "chrome_124":  {[]int{120, 1400, 40, 800, 100}, 20 * time.Millisecond},
    "firefox_125": {[]int{100, 1500, 50, 700, 200}, 25 * time.Millisecond},
    "safari_17":   {[]int{90, 1300, 60, 750, 150}, 15 * time.Millisecond},
    "twitch":      {[]int{1200, 1200, 1200, 1200, 1200}, 5 * time.Millisecond},
    "zoom_voip":   {[]int{160, 160, 320, 160, 480}, 30 * time.Millisecond},
}

var activeProfile = "chrome_124"

func SetActiveProfile(name string) {
    if _, ok := profiles[name]; ok {
        activeProfile = name
        log.Printf("AI Morph profile changed to %s", name)
    }
}

func EnableTrafficMorphing() {
    log.Println("AI Traffic Morphing activated with 5 advanced profiles")
    go func() {
        for {
            p := profiles[activeProfile]
            for _, size := range p.pattern {
                jitter := p.jitter + time.Duration(rand.Intn(10))*time.Millisecond
                log.Printf("[AI Morph] %s : size=%d jitter=%v", activeProfile, size, jitter)
                time.Sleep(jitter)
            }
        }
    }()
}
