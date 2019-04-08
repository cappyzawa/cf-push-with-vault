# cf-push-with-vault
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
