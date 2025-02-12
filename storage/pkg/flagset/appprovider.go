package flagset

import (
	"github.com/micro/cli/v2"
	"github.com/owncloud/ocis/ocis-pkg/flags"
	"github.com/owncloud/ocis/storage/pkg/config"
)

// AppProviderWithConfig applies cfg to the root flagset
func AppProviderWithConfig(cfg *config.Config) []cli.Flag {
	flags := []cli.Flag{

		// debug ports are the odd ports
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.DebugAddr, "0.0.0.0:9165"),
			Usage:       "Address to bind debug server",
			EnvVars:     []string{"APP_PROVIDER_BASIC_DEBUG_ADDR"},
			Destination: &cfg.Reva.AppProvider.DebugAddr,
		},

		// Auth

		// Services

		// AppProvider

		&cli.StringFlag{
			Name:        "network",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.GRPCNetwork, "tcp"),
			Usage:       "Network to use for the storage auth-basic service, can be 'tcp', 'udp' or 'unix'",
			EnvVars:     []string{"APP_PROVIDER_BASIC_GRPC_NETWORK"},
			Destination: &cfg.Reva.AppProvider.GRPCNetwork,
		},
		&cli.StringFlag{
			Name:        "addr",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.GRPCAddr, "0.0.0.0:9164"),
			Usage:       "Address to bind storage service",
			EnvVars:     []string{"APP_PROVIDER_BASIC_GRPC_ADDR"},
			Destination: &cfg.Reva.AppProvider.GRPCAddr,
		},
		&cli.StringFlag{
			Name:        "external-addr",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.ExternalAddr, "127.0.0.1:9164"),
			Usage:       "Address to connect to the storage service for other services",
			EnvVars:     []string{"APP_PROVIDER_BASIC_EXTERNAL_ADDR"},
			Destination: &cfg.Reva.AppProvider.ExternalAddr,
		},
		&cli.StringSliceFlag{
			Name:    "service",
			Value:   cli.NewStringSlice("appprovider"),
			Usage:   "--service appprovider [--service otherservice]",
			EnvVars: []string{"APP_PROVIDER_BASIC_SERVICES"},
		},

		&cli.StringFlag{
			Name:        "driver",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.Driver, ""),
			Usage:       "Driver to use for app provider",
			EnvVars:     []string{"APP_PROVIDER_DRIVER"},
			Destination: &cfg.Reva.AppProvider.Driver,
		},

		// WOPI driver
		&cli.StringFlag{
			Name:        "wopi-driver-app-apikey",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.WopiDriver.AppAPIKey, ""),
			Usage:       "The API key used by the app, if applicable.",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_APP_API_KEY"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.AppAPIKey,
		},
		&cli.BoolFlag{
			Name:        "wopi-driver-app-desktop-only",
			Value:       flags.OverrideDefaultBool(cfg.Reva.AppProvider.WopiDriver.AppDesktopOnly, false),
			Usage:       "Whether the app can be opened only on desktop",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_APP_DESKTOP_ONLY"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.AppDesktopOnly,
		},
		&cli.StringFlag{
			Name:        "wopi-driver-app-icon-uri",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.WopiDriver.AppIconURI, ""),
			Usage:       "IOP Secret (Shared with WOPI server)",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_APP_ICON_URI"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.AppIconURI,
		},
		&cli.StringFlag{
			Name:        "wopi-driver-app-internal-url",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.WopiDriver.AppInternalURL, ""),
			Usage:       "The internal app URL in case of dockerized deployments. Defaults to AppURL",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_APP_INTERNAL_URL"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.AppInternalURL,
		},
		&cli.StringFlag{
			Name:        "wopi-driver-app-name",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.WopiDriver.AppName, ""),
			Usage:       "The App user-friendly name.",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_APP_NAME"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.AppName,
		},
		&cli.StringFlag{
			Name:        "wopi-driver-app-url",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.WopiDriver.AppURL, ""),
			Usage:       "App server URL",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_APP_URL"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.AppURL,
		},
		&cli.BoolFlag{
			Name:        "wopi-driver-insecure",
			Value:       flags.OverrideDefaultBool(cfg.Reva.AppProvider.WopiDriver.Insecure, false),
			Usage:       "Disable SSL certificate verification of WOPI server and WOPI bridge",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_INSECURE"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.Insecure,
		},
		&cli.StringFlag{
			Name:        "wopi-driver-iopsecret",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.WopiDriver.IopSecret, ""),
			Usage:       "IOP Secret (Shared with WOPI server)",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_IOP_SECRET"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.IopSecret,
		},
		&cli.StringFlag{
			Name:        "wopi-driver-wopiurl",
			Value:       flags.OverrideDefaultString(cfg.Reva.AppProvider.WopiDriver.WopiURL, ""),
			Usage:       "WOPI server URL",
			EnvVars:     []string{"APP_PROVIDER_WOPI_DRIVER_WOPI_URL"},
			Destination: &cfg.Reva.AppProvider.WopiDriver.WopiURL,
		},

		// Gateway
		&cli.StringFlag{
			Name:        "gateway-url",
			Value:       flags.OverrideDefaultString(cfg.Reva.Gateway.Endpoint, "localhost:9142"),
			Usage:       "URL to use for the storage gateway service",
			EnvVars:     []string{"STORAGE_GATEWAY_ENDPOINT"},
			Destination: &cfg.Reva.Gateway.Endpoint,
		},
	}

	flags = append(flags, TracingWithConfig(cfg)...)
	flags = append(flags, DebugWithConfig(cfg)...)
	flags = append(flags, SecretWithConfig(cfg)...)

	return flags
}
