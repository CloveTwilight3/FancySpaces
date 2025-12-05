package handler

import "github.com/fancyinnovations/fancyspaces/src/internal/versions"

type CreateVersionReq struct {
	SpaceID                   string           `json:"space_id"`
	Name                      string           `json:"name"`
	Channel                   versions.Channel `json:"channel"`
	Changelog                 string           `json:"changelog"`
	SupportedPlatformVersions []string         `json:"supported_platform_versions"`
}
