package handler

type CreateVersionReq struct {
	Name                      string   `json:"name"`
	Channel                   string   `json:"channel"`
	Changelog                 string   `json:"changelog"`
	SupportedPlatformVersions []string `json:"supported_platform_versions"`
}
