# cf-push-with-vault
[![Go Report Card](https://goreportcard.com/badge/github.com/cappyzawa/cf-push-with-vault)](https://goreportcard.com/report/github.com/cappyzawa/cf-push-with-vault)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://concourse.ik.am:14161/api/v1/teams/cappyzawa/pipelines/cf-push-with-vault/jobs/test-master/badge)](https://concourse.ik.am:14161/teams/cappyzawa/pipelines/cf-push-with-vault)

cf plugin to push cf app with vault

## How to install
### From [CF-Community](https://plugins.cloudfoundry.org/#push-with-vault)

```bash
$ cf install-plugin -r CF-Community "push-with-vault"
```

### From [GitHub Release](https://github.com/cappyzawa/cf-push-with-vault/releases)
Download tar.gz file from release page.
```
$ tar -zxvf cf-push-with-vault_*.tar.gz
$ cf install-plugin -f ./cf-push-with-vault 
```

## How to use
```
$ cf push-with-vault -h
NAME:
   push-with-vault - This enable to use (( )) place holders in manifest files. (( )) are evaluated by vault

USAGE:
   $ cf push-with-vault [APP_NAME]

OPTIONS:
   --domain, -d             Specify a custom domain (e.g. private-domain.example.com, apps.internal.com) to use instead of the default domain
   --file, -f               Path to manifest (default: ./manifest.yml)
   --hostname, -n           Hostname (e.g. my-subdomain)
   --path-prefix, -pp       Path under which to namespace credential lookup
   --vault-addr, -va        Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/. (default: "VAULT_ADDR" env)
   --vault-token, -vt       Vault authentication token. (default: "VAULT_TOKEN" env)
```

## Examples 
If you want to push cf app has follow manifest with vault.

```yml
---
applications:
- name: APP-ONE
  path: ./APP-ONE-DIRECTORY
  env:
    bar: ((/foo/bar))
```

You must set `/foo/bar` to vault with `value` field. (inspired by [Credential lookup rules](https://concourse-ci.org/vault-credential-manager.html))

```bash
$ vault write /foo/bar value="cred"
```

_This plugin can only [KV Secrets Engine \- Version 1](https://www.vaultproject.io/docs/secrets/kv/kv-v1.html)_

```bash
$ export VAULT_ADDR=https://your.vault.address
$ export VAULT_TOKEN=xxxxxxxxxxxx
$ cf push-with-vault --path-prefix=/foo -f manifest.yml
```
