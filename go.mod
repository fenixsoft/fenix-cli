module github.com/fenixsoft/fenix-cli

require (
	cloud.google.com/go v0.54.0 // indirect
	docker.io/go-docker v1.0.0
	github.com/AlecAivazis/survey/v2 v2.2.12
	github.com/Azure/go-autorest/autorest v0.11.12 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/c-bata/go-prompt v0.2.6
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/hashicorp/go-hclog v0.9.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.4
	github.com/imdario/mergo v0.3.5 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/olekukonko/tablewriter v0.0.5
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/schollz/progressbar/v3 v3.8.1 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/client-go v0.21.1
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
	sigs.k8s.io/krew v0.4.1 // indirect
)

// hack for feature & bugfix
replace github.com/c-bata/go-prompt v0.2.6 => ./lib/go-prompt

replace sigs.k8s.io/krew v0.4.1 => ./lib/krew

go 1.16
