package versions

import (
	"time"
)

type Version struct {
	SpaceID                   string        `json:"space_id" bson:"space_id"`
	ID                        string        `json:"id" bson:"id"`
	Name                      string        `json:"name" bson:"name"`
	Platform                  Platform      `json:"platform" bson:"platform"`
	Channel                   Channel       `json:"channel" bson:"channel"`
	PublishedAt               time.Time     `json:"published_at" bson:"published_at"`
	Changelog                 string        `json:"changelog" bson:"changelog"`
	SupportedPlatformVersions []string      `json:"supported_platform_versions" bson:"supported_platform_versions"`
	Files                     []VersionFile `json:"files" bson:"files"`
}

type VersionFile struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
	Size int64  `json:"size" bson:"size"`
}

type Platform string

const (
	PlatformMinecraftPlugin Platform = "minecraft_plugin"
	PlatformHytalePlugin    Platform = "hytale_plugin"
)

type Channel string

const (
	ChannelRelease Channel = "release"
	ChannelBeta    Channel = "beta"
	ChannelAlpha   Channel = "alpha"
)
