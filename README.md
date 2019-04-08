# cf-push-with-vault
[![Go Report Card](https://goreportcard.com/badge/github.com/cappyzawa/cf-push-with-vault)](https://goreportcard.com/report/github.com/cappyzawa/cf-push-with-vault)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://concourse.ik.am:14161/api/v1/teams/cappyzawa/pipelines/cf-push-with-vault/jobs/test-master/badge)](https://concourse.ik.am:14161/teams/cappyzawa/pipelines/cf-push-with-vault)

cf plugin to push cf app with vault

# How to install
```
$ cf install-plugin -f <binary>
```

# How to use
```
$ cf push-with-vault -h
NAME:
   push-with-vault - This enable to use (( )) place holders in manifest files. (( )) are evaluated by vault

USAGE:
   $ cf push-with-vault

OPTIONS:
   --file              Path to manifest (default: ./manifest.yml)
   --path-prefix       Path under which to namespace credential lookup
   --vault-addr        Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/. (default: "VAULT_ADDR" env)
   --vault-token       Vault authentication token. (default: "VAULT_TOKEN" env)
```
