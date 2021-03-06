package model

import (
	"encoding/json"
	"io"
	"time"

	mattermostModel "github.com/mattermost/mattermost-server/model"
)

// Plugin represents a Mattermost plugin in the Plugin Marketplace.
type Plugin struct {
	HomepageURL     string                    `json:"homepage_url"`
	IconData        string                    `json:"icon_data"` // The base64 encoding of an svg image
	DownloadURL     string                    `json:"download_url"`
	ReleaseNotesURL string                    `json:"release_notes_url"`
	Labels          []Label                   `json:"labels"`
	Signature       string                    `json:"signature"` // A signature of a plugin saved in base64 encoding.
	Manifest        *mattermostModel.Manifest `json:"manifest"`
	UpdatedAt       time.Time                 `json:"updated_at"` // The point in time this release of the plugin was added to the Plugin Marketplace
}

// PluginFromReader decodes a json-encoded cluster from the given io.Reader.
func PluginFromReader(reader io.Reader) (*Plugin, error) {
	cluster := Plugin{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&cluster)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return &cluster, nil
}

// PluginsFromReader decodes a json-encoded list of plugins from the given io.Reader.
func PluginsFromReader(reader io.Reader) ([]*Plugin, error) {
	plugins := []*Plugin{}
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&plugins)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return plugins, nil
}

// PluginFilter describes the parameters used to constrain a set of plugins.
type PluginFilter struct {
	Page          int
	PerPage       int
	Filter        string
	ServerVersion string
}
