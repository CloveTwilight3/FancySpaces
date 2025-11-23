package versions

import "time"

type Version struct {
	SpaceID                   string        `json:"space_id"`
	ID                        string        `json:"id"`
	Name                      string        `json:"name"`
	Channel                   string        `json:"channel"`
	PublishedAt               time.Time     `json:"published_at"`
	Changelog                 string        `json:"changelog"`
	SupportedPlatformVersions []string      `json:"supported_platform_versions"`
	Files                     []VersionFile `json:"files"`
	Downloads                 int64         `json:"downloads"`
}

type VersionFile struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Size int64  `json:"size"`
}

type Channel string

const (
	ChannelRelease Channel = "release"
	ChannelBeta    Channel = "beta"
	ChannelAlpha   Channel = "alpha"
)
