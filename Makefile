# Makefile variables

test:
	@go test ./... -v

gotest:
	@echo "==> goal gotest..."
	@gotest ./...

look_update_pkgs:
	# take a look at the newer versions of dependency modules
	@go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all 2> /dev/null