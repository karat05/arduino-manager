package main

import "testing"

func TestGuessBoardNames(t *testing.T) {
	vendor, arch := guessBoardNames("esp8266/Arduino")
	if vendor != "esp8266" || arch != "esp8266" {
		t.Errorf("expected esp8266/esp8266, got %s/%s", vendor, arch)
	}
}

func TestWordRank(t *testing.T) {
	if wordRank("Adafruit_SSD1306", "ssd1306") != 0 {
		t.Error("expected rank 0 for prefix match")
	}
}
