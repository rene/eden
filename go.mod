module github.com/lf-edge/eden

go 1.22

require (
	github.com/Insei/rolgo v0.0.2
	github.com/amitbet/vncproxy v0.0.0-20200118084310-ea8f9b510913
	github.com/containerd/containerd v1.7.13
	github.com/docker/distribution v2.8.3+incompatible
	github.com/docker/docker v25.0.6+incompatible
	github.com/docker/go-connections v0.5.0
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.15.0
	github.com/fsnotify/fsnotify v1.6.0
	github.com/go-redis/redis/v9 v9.0.0-beta.1
	github.com/google/go-containerregistry v0.19.1
	github.com/google/uuid v1.6.0
	github.com/lf-edge/eden/eserver v0.0.0-20220711180217-6e2bfa9c3f67
	github.com/lf-edge/eden/sdn/vm v0.0.0-20250314212201-6f1e0716172f
	github.com/lf-edge/edge-containers v0.0.0-20240207093504-5dfda0619b80
	github.com/lf-edge/eve-api/go v0.0.0-20240829123634-7c8ebda876ff
	github.com/lf-edge/eve/pkg/pillar v0.0.0-20240923082146-6d403aaa5513
	github.com/mcuadros/go-lookup v0.0.0-20200831155250-80f87a4fa5ee
	github.com/moby/term v0.5.0
	github.com/nerd2/gexto v0.0.0-20190529073929-39468ec063f6
	github.com/onsi/gomega v1.29.0
	github.com/packethost/packngo v0.25.0
	github.com/rogpeppe/go-internal v1.11.0
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.8.4
	github.com/thediveo/enumflag v0.10.1
	github.com/tmc/scp v0.0.0-20170824174625-f7b48647feef
	golang.org/x/crypto v0.21.0
	golang.org/x/exp v0.0.0-20220827204233-334a2380cb91
	golang.org/x/net v0.23.0
	golang.org/x/oauth2 v0.16.0
	golang.org/x/term v0.18.0
	golang.org/x/text v0.14.0
	google.golang.org/api v0.160.0
	google.golang.org/protobuf v1.33.0
	gopkg.in/errgo.v2 v2.1.0
	gopkg.in/yaml.v2 v2.4.0
	oras.land/oras-go v1.2.5
)

require (
	cloud.google.com/go/compute v1.23.4 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20230811130428-ced1acdcaa24 // indirect
	github.com/AdamKorcz/go-118-fuzz-build v0.0.0-20231105174938-2b5cbb29f3e2 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/Microsoft/hcsshim v0.11.4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bugsnag/bugsnag-go v1.5.3 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/containerd/cgroups v1.1.0 // indirect
	github.com/containerd/continuity v0.4.3 // indirect
	github.com/containerd/fifo v1.1.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.14.3 // indirect
	github.com/containerd/ttrpc v1.2.3 // indirect
	github.com/containerd/typeurl/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/cli v25.0.3+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.1 // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/klauspost/compress v1.17.6 // indirect
	github.com/lf-edge/eve/libs/depgraph v0.0.0-20220711144346-0659e3b03496 // indirect
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/sys/mountinfo v0.7.1 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/sys/signal v0.7.0 // indirect
	github.com/moby/sys/user v0.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc6 // indirect
	github.com/opencontainers/runtime-spec v1.1.0 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.46.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.4.0 // indirect
	github.com/vbatts/tar-split v0.11.3 // indirect
	github.com/yvasiyarov/go-metrics v0.0.0-20150112132944-c25f46c4b940 // indirect
	github.com/yvasiyarov/gorelic v0.0.7 // indirect
	github.com/yvasiyarov/newrelic_platform_go v0.0.0-20160601141957-9c099fbc30e9 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.48.0 // indirect
	go.opentelemetry.io/otel v1.23.1 // indirect
	go.opentelemetry.io/otel/metric v1.23.1 // indirect
	go.opentelemetry.io/otel/trace v1.23.1 // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/tools v0.18.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20240205150955-31a09d347014 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240205150955-31a09d347014 // indirect
	google.golang.org/grpc v1.61.0 // indirect
	gopkg.in/ini.v1 v1.66.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/lf-edge/eden/sdn/vm => ./sdn/vm

replace github.com/lf-edge/eve/libs/depgraph => github.com/lf-edge/eve/libs/depgraph v0.0.0-20220711144346-0659e3b03496
