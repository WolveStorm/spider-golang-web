package global

type GameInfoDetail struct {
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Company       string `json:"company"`
	Score         string `json:"score"`
	DownloadTimes string `json:"download_times"`
	Description   string `json:"description"`
	ApkUrl        string `json:"apk_url"`
}

type GameList struct {
	Total int32             `json:"total,omitempty"`
	List  []*GameInfoDetail `json:"list,omitempty"`
}
