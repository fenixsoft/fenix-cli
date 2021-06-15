module github.com/fenixsoft/fenix-cli

require (
	docker.io/go-docker v1.0.0
	github.com/AlecAivazis/survey/v2 v2.2.12 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/c-bata/go-prompt v0.2.6
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-retryablehttp v0.6.4
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sys v0.0.0-20200918174421-af09f7315aff // indirect
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/klog v1.0.0 // indirect
)

// hack for feature & bugfix
replace github.com/c-bata/go-prompt v0.2.6 => ./lib/go-prompt

go 1.15
