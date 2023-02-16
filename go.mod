module github.com/ipfs/testground

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/containerd/containerd v1.5.18 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/docker v1.4.2-0.20190910181529-415f8ecb65e8
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/hashicorp/go-getter v1.4.0
	github.com/imdario/mergo v0.3.12
	github.com/ipfs/testground/sdk/runtime v0.0.0-00010101000000-000000000000
	github.com/logrusorgru/aurora v0.0.0-20190803045625-94edacc10f9b
	github.com/morikuni/aec v0.0.0-20170113033406-39771216ff4c // indirect
	github.com/otiai10/copy v1.0.1
	github.com/otiai10/curr v0.0.0-20190513014714-f5a3d24e5776 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/urfave/cli v1.22.2
	go.uber.org/zap v1.10.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)

replace (
	github.com/ipfs/testground/sdk/iptb => ./sdk/iptb
	github.com/ipfs/testground/sdk/runtime => ./sdk/runtime
	github.com/ipfs/testground/sdk/sync => ./sdk/sync
	github.com/miekg/dns => github.com/miekg/dns v1.0.14
)
