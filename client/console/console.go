package console

import (
	"GeeRpc/client/assets"
	"GeeRpc/client/protobufs/clientpb"
	"github.com/reeflective/console"
	"log/slog"
	"sync"
)

const (
	defaultTimeout = 60
)

const (
	// ANSI Colors.
	Normal    = "\033[0m"
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Orange    = "\033[33m"
	Blue      = "\033[34m"
	Purple    = "\033[35m"
	Cyan      = "\033[36m"
	Gray      = "\033[37m"
	Bold      = "\033[1m"
	Clearln   = "\r\x1b[2K"
	UpN       = "\033[%dA"
	DownN     = "\033[%dB"
	Underline = "\033[4m"

	// Info - Display colorful information.
	Info = Bold + Cyan + "[*] " + Normal
	// Warn - Warn a user.
	Warn = Bold + Red + "[!] " + Normal
	// Debug - Display debug information.
	Debug = Bold + Purple + "[-] " + Normal
	// Woot - Display success.
	Woot = Bold + Green + "[$] " + Normal
	// Success - Diplay success.
	Success = Bold + Green + "[+] " + Normal
)

// Observer - A function to call when the sessions changes.
type (
	Observer           func(*clientpb.Session, *clientpb.Beacon)
	BeaconTaskCallback func(*clientpb.BeaconTask)
)

type SliverClient struct {
	App                      *console.Console
	Rpc                      rpcpb.SliverRPCClient
	ActiveTarget             *ActiveTarget
	EventListeners           *sync.Map
	BeaconTaskCallbacks      map[string]BeaconTaskCallback
	BeaconTaskCallbacksMutex *sync.Mutex
	Settings                 *assets.ClientSettings
	IsServer                 bool
	IsCLI                    bool

	jsonHandler slog.Handler
	printf      func(format string, args ...any) (int, error)
}

type ActiveTarget struct {
	session    *clientpb.Session
	beacon     *clientpb.Beacon
	observers  map[int]Observer
	observerID int
	con        *SliverClient
}
