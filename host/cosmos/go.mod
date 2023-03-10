module pkg.berachain.dev/polaris/cosmos

go 1.20

// Required for supporting stateful precompiled contracts.
replace (
	cosmossdk.io/x/evidence => cosmossdk.io/x/evidence v0.0.0-20230305125317-672b052f323b
	cosmossdk.io/x/feegrant => cosmossdk.io/x/feegrant v0.0.0-20230305125317-672b052f323b
	github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.46.0-beta2.0.20230305125317-672b052f323b
	github.com/docker/docker => github.com/docker/docker v20.10.3-0.20221013203545-33ab36d6b304+incompatible // 22.06 branch
	github.com/ethereum/go-ethereum => github.com/berachain/go-ethereum v0.0.0-20230306205832-364e922bfe8d
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	pkg.berachain.dev/polaris/eth => ../../eth
	pkg.berachain.dev/polaris/lib => ../../lib
)

require (
	cosmossdk.io/api v0.3.1
	cosmossdk.io/client/v2 v2.0.0-20230220152935-67f04e629623
	cosmossdk.io/core v0.6.0
	cosmossdk.io/depinject v1.0.0-alpha.3
	cosmossdk.io/errors v1.0.0-beta.7
	cosmossdk.io/log v0.0.0-20230305202224-89c956f8ed3a
	cosmossdk.io/math v1.0.0-beta.6.0.20230305125317-672b052f323b
	cosmossdk.io/simapp v0.0.0-20230305125317-672b052f323b
	cosmossdk.io/store v0.0.0-20230305125317-672b052f323b
	cosmossdk.io/tools/confix v0.0.0-20230227140543-7c9245dbdf19
	cosmossdk.io/x/evidence v0.1.0
	cosmossdk.io/x/feegrant v0.0.0-20230305125317-672b052f323b
	cosmossdk.io/x/upgrade v0.0.0-20230305125317-672b052f323b
	github.com/btcsuite/btcd v0.23.4
	github.com/btcsuite/btcd/btcutil v1.1.3
	github.com/cometbft/cometbft v0.37.0
	github.com/cosmos/cosmos-db v1.0.0-rc.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.2
	github.com/cosmos/cosmos-sdk v0.48.0
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/gogoproto v1.4.6
	github.com/ethereum/go-ethereum v1.11.3
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/onsi/ginkgo/v2 v2.9.0
	github.com/onsi/gomega v1.27.3
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.15.0
	github.com/tidwall/btree v1.6.0
	google.golang.org/genproto v0.0.0-20230216225411-c8e22ba71e44
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.29.0
	pkg.berachain.dev/polaris/eth v0.0.1-alpha.0.20230310224229-ede3de74fb93
	pkg.berachain.dev/polaris/lib v0.0.1-alpha.0.20230310224229-ede3de74fb93
)

require (
	cloud.google.com/go v0.110.0 // indirect
	cloud.google.com/go/compute v1.18.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v0.11.0 // indirect
	cloud.google.com/go/storage v1.29.0 // indirect
	cosmossdk.io/collections v0.0.0-20230214153846-b6c6e4e99177 // indirect
	cosmossdk.io/x/tx v0.2.0 // indirect
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/ChainSafe/go-schnorrkel v0.0.0-20200405005733-88cbf1b4c40d // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/VictoriaMetrics/fastcache v1.12.1 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/aws/aws-sdk-go v1.44.203 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bgentry/speakeasy v0.1.1-0.20220910012023-760eaf8b6816 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/bufbuild/protocompile v0.5.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/cockroachdb/apd/v2 v2.0.2 // indirect
	github.com/cockroachdb/errors v1.9.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v0.0.0-20230309163202-51422ae2d449 // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/cometbft/cometbft-db v0.7.0 // indirect
	github.com/confio/ics23/go v0.9.0 // indirect
	github.com/cosmos/btcutil v1.0.5 // indirect
	github.com/cosmos/gogogateway v1.2.0 // indirect
	github.com/cosmos/iavl v0.21.0-alpha.1 // indirect
	github.com/cosmos/ledger-cosmos-go v0.13.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/creachadair/atomicfile v0.2.8 // indirect
	github.com/creachadair/taskgroup v0.4.2 // indirect
	github.com/creachadair/tomledit v0.0.24 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/fjl/memsize v0.0.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/getsentry/sentry-go v0.19.0 // indirect
	github.com/gin-gonic/gin v1.9.0 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/orderedcode v0.0.1 // indirect
	github.com/google/pprof v0.0.0-20230309165930-d61513b1440d // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-bexpr v0.1.11 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.7.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hdevalence/ed25519consensus v0.1.0 // indirect
	github.com/holiman/big v0.0.0-20221017200358-a027dc42d04e // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/holiman/uint256 v1.2.1 // indirect
	github.com/huandu/skiplist v1.2.0 // indirect
	github.com/huin/goupnp v1.1.0 // indirect
	github.com/improbable-eng/grpc-web v0.15.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jhump/protoreflect v1.15.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/linxGnu/grocksdb v1.7.15 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/manifoldco/promptui v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20210601165009-122bf33a46e0 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/pointerstructure v1.2.1 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/petermattis/goid v0.0.0-20221215004737-a150e88a970d // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/rs/cors v1.8.3 // indirect
	github.com/rs/zerolog v1.29.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/status-im/keycard-go v0.2.0 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	github.com/tecbot/gorocksdb v0.0.0-20191217155057-f0fad39f321c // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/ulikunitz/xz v0.5.11 // indirect
	github.com/urfave/cli/v2 v2.25.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	github.com/zondax/hid v0.9.1 // indirect
	github.com/zondax/ledger-go v0.14.1 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/exp v0.0.0-20230310171629-522b1b587ee0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/oauth2 v0.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/term v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.110.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.4.0 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
	pgregory.net/rapid v0.5.5 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
