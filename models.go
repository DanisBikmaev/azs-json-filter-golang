package main

type Marks struct {
	Marks []Mark `json:"maps"`
}

type Mark struct {
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
	Options    Options    `json:"options"`
	Id         int        `json:"id"`
	Type       string     `json:"type"`
}

type Properties struct {
	HintContent          string `json:"hintContent"`
	BalloonContentFooter string `json:"balloonContentFooter"`
	BalloonContentHeader string `json:"balloonContentHeader"`
	IconCaption          string `json:"iconCaption"`
}

type Geometry struct {
	Coordinates []string `json:"coordinates"`
	Type        string   `json:"type"`
}

type Options struct {
	Preset    string `json:"preset"`
	IconColor string `json:"iconColor"`
}
