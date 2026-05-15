module go.etcd.io/etcd/v3

go 1.21

require (
	github.com/coreos/go-semver v0.3.1
	github.com/coreos/go-systemd/v22 v22.5.0
	github.com/dustin/go-humanize v1.0.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/protobuf v1.5.3
	github.com/google/btree v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jonboulle/clockwork v0.4.0
	github.com/json-iterator/go v1.1.12
	github.com/prometheus/client_golang v1.17.0
	github.com/sirupsen/logrus v1.9.3
	github.com/soheilhy/cmux v0.1.5
	github.com/spf13/cobra v1.7.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75
	github.com/urfave/cli v1.22.14
	github.com/xiang90/probing v0.0.0-20221125231312-a49e3df8f510
	go.etcd.io/bbolt v1.3.8
	go.etcd.io/raft/v3 v3.0.0-20231026054731-3f4c3b1a7e2c
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.45.0
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.19.0
	go.opentelemetry.io/otel/sdk v1.19.0
	go.uber.org/zap v1.26.0
	golang.org/x/crypto v0.14.0
	golang.org/x/net v0.17.0
	google.golang.org/grpc v1.58.3
	google.golang.org/protobuf v1.31.0
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20231012201019-e917dd12ba7a // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231012201019-e917dd12ba7a // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231012201019-e917dd12ba7a // indirect
)

// Personal fork for learning purposes - studying etcd's raft implementation
// and distributed consensus mechanisms. Not intended for production use.
//
// Study notes:
//   - go.etcd.io/raft/v3: core raft consensus library, focus area for learning
//   - go.etcd.io/bbolt: underlying persistent storage engine (B+ tree)
//   - See /docs/learning/ for personal notes on the raft log replication flow
