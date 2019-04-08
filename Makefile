dev:
	go build -o push-with-vault
	cf install-plugin -f ./push-with-vault
redev:
	cf uninstall-plugin push-with-vault
	go build -o push-with-vault
	cf install-plugin -f ./push-with-vault
