module github.com/VulpesFerrilata/catan

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/VulpesFerrilata/grpc v0.0.0-20210108084505-964ee54318c0
	github.com/VulpesFerrilata/library v0.0.0-20210108114535-906b7f5cafa7
	github.com/kataras/iris/v12 v12.1.8
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/nats-io/nats-server/v2 v2.1.8 // indirect
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/pkg/errors v0.9.1
	go.uber.org/dig v1.10.0
	gopkg.in/go-playground/validator.v9 v9.31.0
	gorm.io/gorm v1.20.2
)
