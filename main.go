package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/cli/cf/flags"
	"code.cloudfoundry.org/cli/plugin"

	"github.com/cappyzawa/cf-push-with-vault/plug"
	v1 "github.com/cappyzawa/cf-push-with-vault/vault/v1"
	vaultapi "github.com/hashicorp/vault/api"
)

const (
	// COMMAND is plugin command name
	COMMAND = "push-with-vault"
	// DefaultManifest is default cf manifest path
	DefaultManifest = "./manifest.yml"
)

// CfPushWithVault implements cf plugin
type CfPushWithVault struct {
	VaultAddr  string
	VaultToken string
}

func main() {
	plugin.Start(&CfPushWithVault{
		VaultAddr:  os.Getenv("VAULT_ADDR"),
		VaultToken: os.Getenv("VAULT_TOKEN"),
	})
}

// Run pushes cf app which vault
func (c *CfPushWithVault) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] != COMMAND {
		c.GetMetadata()
		os.Exit(0)
	}
	fc, err := c.parseArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

	vaultConfig := vaultapi.DefaultConfig()
	vaultConfig.ConfigureTLS(&vaultapi.TLSConfig{Insecure: true})
	vaultClient, err := vaultapi.NewClient(vaultConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create vault client: %v\n", err)
	}
	vaultClient.SetAddress(fc.String("vault-addr"))
	vaultClient.SetToken(fc.String("vault-token"))

	variablesFactory := v1.NewFactory(vaultClient.Logical(), fc.String("path-prefix"))
	variables := variablesFactory.NewVariables()
	command := &plug.Command{
		CliConnection: cliConnection,
		Variables:     variables,
	}

	appName := ""
	if len(fc.Args()) == 2 {
		appName = fc.Args()[1]
	}
	pushArgs := plug.PushArgs{
		AppName:      appName,
		ManifestFile: fc.String("file"),
		Hostname:     fc.String("hostname"),
		Domain:       fc.String("domain"),
	}

	if err := command.Push(pushArgs); err != nil {
		fmt.Fprintf(os.Stdout, "failed to push with vault: %v", err)
		os.Exit(1)
	}
}

// GetMetadata informs the CLI of the name of a plugin
func (c *CfPushWithVault) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "push-with-vault",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 3,
		},
		Commands: []plugin.Command{
			{
				Name:     "push-with-vault",
				HelpText: "This enable to use (( )) place holders in manifest files. (( )) are evaluated by vault",
				UsageDetails: plugin.Usage{
					Usage: "$ cf push-with-vault [APP_NAME]",
					Options: map[string]string{
						"-file, -f":         "Path to manifest (default: ./manifest.yml)",
						"-vault-addr, -va":  "Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/. (default: \"VAULT_ADDR\" env)",
						"-vault-token, -vt": "Vault authentication token. (default: \"VAULT_TOKEN\" env)",
						"-path-prefix, -pp": "Path under which to namespace credential lookup",
						"-hostname, -n":     "Hostname (e.g. my-subdomain)",
						"-domain, -d":       "Specify a custom domain (e.g. private-domain.example.com, apps.internal.com) to use instead of the default domain",
					},
				},
			},
		},
	}
}

func (c *CfPushWithVault) parseArgs(args []string) (flags.FlagContext, error) {
	fc := flags.New()
	fc.NewStringFlagWithDefault("file", "f", "Path to manifest", DefaultManifest)
	fc.NewStringFlagWithDefault("vault-addr", "va", "Address of the Vault server expressed as a URL and port", c.VaultAddr)
	fc.NewStringFlagWithDefault("vault-token", "vt", "Vault authentication token", c.VaultToken)
	fc.NewStringFlagWithDefault("path-prefix", "pp", "Path under which to namespace credential lookup", "")
	fc.NewStringFlagWithDefault("hostname", "n", "Hostname", "")
	fc.NewStringFlagWithDefault("domain", "d", "Specify a custom domain", "")
	err := fc.Parse(args...)
	return fc, err
}
