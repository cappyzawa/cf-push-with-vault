package plug

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/cli/plugin"
	"github.com/cappyzawa/cf-push-with-vault/vault"
	"github.com/cloudfoundry/bosh-cli/director/template"
)

// Command is struct
type Command struct {
	CliConnection plugin.CliConnection
	Variables     vault.Variables
}

// PushArgs is information for "cf push"
type PushArgs struct {
	AppName      string
	ManifestFile string
	Hostname     string
	Domain       string
}

func (p *PushArgs) toStringArray() []string {
	args := []string{"push"}
	if p.AppName != "" {
		args = append(args, p.AppName)
	}
	if p.Hostname != "" {
		args = append(args, "-n", p.Hostname)
	}
	if p.Domain != "" {
		args = append(args, "-d", p.Domain)
	}
	return append(args, "-f", p.ManifestFile)
}

// Push pushes cf app based on manifest
func (c *Command) Push(args PushArgs) error {
	// read file
	absFile, err := filepath.Abs(args.ManifestFile)
	if err != nil {
		return err
	}
	bytes, err := c.readManifest(absFile)
	if err != nil {
		return err
	}

	// evaluate variables
	var obj interface{}
	result, err := c.evaluate(bytes, &obj)
	if err != nil {
		return err
	}

	// write result to tmp file
	tmpFile, err := c.writeTmpFile(absFile, result)
	if err != nil {
		return err
	}
	args.ManifestFile = tmpFile.Name()

	// cf push
	if _, err := c.CliConnection.CliCommand(args.toStringArray()...); err != nil {
		return err
	}

	if err := c.removeTmpFile(absFile); err != nil {
		return err
	}
	return nil
}

func (c *Command) readManifest(manifest string) ([]byte, error) {
	f, err := os.Open(manifest)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (c *Command) evaluate(bytes []byte, out interface{}) ([]byte, error) {
	tpl := template.NewTemplate(bytes)
	bytes, err := tpl.Evaluate(c.Variables, nil, template.EvaluateOpts{
		ExpectAllKeys: true,
	})
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (c *Command) writeTmpFile(manifest string, result []byte) (*os.File, error) {
	tmpFileName := fmt.Sprintf("%s.tmp", manifest)
	tmpFile, err := os.Create(tmpFileName)
	if err != nil {
		return nil, err
	}
	tmpFile.Write(result)
	defer tmpFile.Close()
	return tmpFile, nil
}

func (c *Command) removeTmpFile(manifest string) error {
	tmpFileName := fmt.Sprintf("%s.tmp", manifest)
	return os.Remove(tmpFileName)
}
