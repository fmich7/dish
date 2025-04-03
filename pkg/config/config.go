package config

import (
	"errors"
	"flag"
	"fmt"
)

// TODO: Please change ApiCacheTTLHours from hours to minutes. If we for example add a new socket to be monitored to our remote API endpoint, we do not want to wait up to 24 hours for it to be monitored or have to clear the cache manually.
type Config struct {
	InstanceName         string
	ApiHeaderName        string
	ApiHeaderValue       string
	ApiCacheSockets      bool
	ApiCacheDirectory    string
	ApiCacheTTLHours     uint
	Source               string
	Verbose              bool
	PushgatewayURL       string
	TelegramBotToken     string
	TelegramChatID       string
	TimeoutSeconds       uint
	ApiURL               string
	WebhookURL           string
	TextNotifySuccess    bool
	MachineNotifySuccess bool
}

const (
	defaultInstanceName         = "generic-dish"
	defaultApiHeaderName        = ""
	defaultApiHeaderValue       = ""
	defaultApiCacheSockets      = false
	defaultApiCacheDir          = "cache"
	defaultApiCacheTTLHours     = 24
	defaultVerbose              = false
	defaultPushgatewayURL       = ""
	defaultTelegramBotToken     = ""
	defaultTelegramChatID       = ""
	defaultTimeoutSeconds       = 10
	defaultApiURL               = ""
	defaultWebhookURL           = ""
	defaultTextNotifySuccess    = false
	defaultMachineNotifySuccess = false
)

var ErrNoSourceProvided = errors.New("no source provided")

// defineFlags defines flags on the provided FlagSet. The values of the flags are stored in the provided Config when parsed.
func defineFlags(fs *flag.FlagSet, cfg *Config) {
	// System flags
	fs.StringVar(&cfg.InstanceName, "name", defaultInstanceName, "a string, dish instance name")
	fs.UintVar(&cfg.TimeoutSeconds, "timeout", defaultTimeoutSeconds, "an int, timeout in seconds for http and tcp calls")
	fs.BoolVar(&cfg.Verbose, "verbose", defaultVerbose, "a bool, console stdout logging toggle")

	// Integration channels flags
	//
	// General:
	fs.BoolVar(&cfg.TextNotifySuccess, "textNotifySuccess", defaultTextNotifySuccess, "a bool, specifies whether successful checks with no failures should be reported to text channels")
	fs.BoolVar(&cfg.MachineNotifySuccess, "machineNotifySuccess", defaultMachineNotifySuccess, "a bool, specifies whether successful checks with no failures should be reported to machine channels")

	// API socket source:
	fs.StringVar(&cfg.ApiHeaderName, "hname", defaultApiHeaderName, "a string, custom additional header name")
	fs.StringVar(&cfg.ApiHeaderValue, "hvalue", defaultApiHeaderValue, "a string, custom additional header value")
	fs.BoolVar(&cfg.ApiCacheSockets, "cache", defaultApiCacheSockets, "a bool, specifies whether to cache fetched socket list from API")
	fs.StringVar(&cfg.ApiCacheDirectory, "cacheDir", defaultApiCacheDir, "a string, specifies cache directory for API socket lists")
	fs.UintVar(&cfg.ApiCacheTTLHours, "cacheTTL", defaultApiCacheTTLHours, "an int, time duration (in hours) for which the cached file is valid")

	// Pushgateway:
	fs.StringVar(&cfg.PushgatewayURL, "target", defaultPushgatewayURL, "a string, result update path/URL to pushgateway, plaintext/byte output")

	// Telegram:
	fs.StringVar(&cfg.TelegramBotToken, "telegramBotToken", defaultTelegramBotToken, "a string, Telegram bot private token")
	fs.StringVar(&cfg.TelegramChatID, "telegramChatID", defaultTelegramChatID, "a string, Telegram chat/channel ID")

	// API for pushing results:
	fs.StringVar(&cfg.ApiURL, "updateURL", defaultApiURL, "a string, API endpoint URL for pushing results")

	// Webhooks:
	fs.StringVar(&cfg.WebhookURL, "webhookURL", defaultWebhookURL, "a string, URL of webhook endpoint")
}

// NewConfig returns a new instance of Config.
//
// If a flag is used for a supported config parameter, the config parameter's value is set according to the provided flag. Otherwise, a default value is used for the given parameter.
func NewConfig(fs *flag.FlagSet, args []string) (*Config, error) {
	cfg := &Config{
		InstanceName:      defaultInstanceName,
		ApiHeaderName:     defaultApiHeaderName,
		ApiHeaderValue:    defaultApiHeaderValue,
		ApiCacheSockets:   defaultApiCacheSockets,
		ApiCacheDirectory: defaultApiCacheDir,
		Verbose:           defaultVerbose,
		PushgatewayURL:    defaultPushgatewayURL,
		TelegramBotToken:  defaultTelegramBotToken,
		TelegramChatID:    defaultTelegramChatID,
		TimeoutSeconds:    defaultTimeoutSeconds,
		ApiURL:            defaultApiURL,
		WebhookURL:        defaultWebhookURL,
	}

	defineFlags(fs, cfg)

	// Parse flags
	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf("error parsing flags: %v", err)
	}

	// Parse args
	parsedArgs := flag.CommandLine.Args()

	// If no source is provided, return an error
	if len(parsedArgs) == 0 {
		return nil, ErrNoSourceProvided
	}
	// Otherwise, store the source in the config
	cfg.Source = parsedArgs[0]

	return cfg, nil
}
