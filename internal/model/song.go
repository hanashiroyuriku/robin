package model

type Song struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Artist      string      `json:"artist"`
	Description string      `json:"description"`
	Lyrics      []LyricLine `json:"lyrics"`
}

type LyricLine struct {
	Text      string  `json:"text"`
	CharDelay float64 `json:"char_delay"`
	LinePause float64 `json:"line_pause"`
}
