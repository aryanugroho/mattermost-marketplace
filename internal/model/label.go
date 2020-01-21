package model

// Label represents a label shown in the Plugin Marketplace UI.
type Label struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Color       string `json:"color"`
}

var AllLabels = []Label{OfficialLabel, BetaLabel}

var OfficialLabel Label = Label{
	Name:        "Official",
	Description: "This plugin is maintained by Mattermost",
	URL:         "https://mattermost.com/pl/default-community-plugins",
	Color:       "#166de0",
}

var BetaLabel Label = Label{
	Name:        "Beta",
	Description: "This plugin is currently in Beta and it's not recommended in production use.",
	URL:         "https://mattermost.com/pl/default-beta-plugins",
	Color:       "#166de0",
}
