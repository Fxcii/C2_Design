package assets

const (
	settingsFileName = "tui-settings.json"
)

// ClientSettings - Client JSON config
type ClientSettings struct {
	TableStyle        string `json:"tables"`
	AutoAdult         bool   `json:"autoadult"`
	BeaconAutoResults bool   `json:"beacon_autoresults"`
	SmallTermWidth    int    `json:"small_term_width"`
	AlwaysOverflow    bool   `json:"always_overflow"`
	VimMode           bool   `json:"vim_mode"`
	UserConnect       bool   `json:"user_connect"`
	ConsoleLogs       bool   `json:"console_logs"`
}
