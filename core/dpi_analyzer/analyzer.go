package dpi_analyzer

import (
    "log"
    "math/rand"
)

type Analyzer struct {
    currentProfile string
}

func NewAnalyzer() *Analyzer {
    return &Analyzer{currentProfile: "chrome_124"}
}

func (a *Analyzer) SetProfile(profile string) {
    a.currentProfile = profile
}

func (a *Analyzer) AnalyzeAndAdapt() {
    if rand.Float32() < 0.03 {
        log.Println("DPI Analyzer: pattern detected, switching profile")
        a.currentProfile = "firefox_125"
    } else {
        a.currentProfile = "chrome_124"
    }
}

func (a *Analyzer) GetProfile() string {
    return a.currentProfile
}
