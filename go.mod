module github.com/naturalselectionlabs/pregod

go 1.18

// https://github.com/samber/lo/pull/43
// https://github.com/samber/lo/pull/81
replace github.com/samber/lo => github.com/kallydev/lo v1.26.0-alpha.1

replace github.com/labstack/echo/v4 => github.com/labstack/echo/v4 v4.7.2

require (
	github.com/Khan/genqlient v0.5.0
	github.com/alecthomas/repr v0.1.0
	github.com/avast/retry-go/v4 v4.3.3
	github.com/avvydomains/golang-client v0.2.0
	github.com/deckarep/golang-set v1.8.0
	github.com/ethereum/go-ethereum v1.11.3
	github.com/go-playground/validator/v10 v10.11.2
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/go-querystring v1.1.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/hasura/go-graphql-client v0.8.1
	github.com/k0kubun/pp/v3 v3.2.0
	github.com/labstack/echo-contrib v0.14.1
	github.com/labstack/echo/v4 v4.10.2
	github.com/lib/pq v1.10.7
	github.com/naturalselectionlabs/kurora v0.41.0
	github.com/naturalselectionlabs/kurora/client v0.0.0-20230310200322-2f12151ff58a
	github.com/rabbitmq/amqp091-go v1.6.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/samber/lo v1.37.0
	github.com/shopspring/decimal v1.3.1
	github.com/shurcooL/graphql v0.0.0-20220606043923-3cf50f8a0a29
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.2
	github.com/unstoppabledomains/resolution-go/v2 v2.3.2
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/wealdtech/go-ens/v3 v3.5.5
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/exporters/jaeger v1.14.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.opentelemetry.io/otel/trace v1.14.0
	go.uber.org/ratelimit v0.2.0
	go.uber.org/zap v1.24.0
	golang.org/x/exp v0.0.0-20230307190834-24139beb5833
	golang.org/x/net v0.8.0
	golang.org/x/sync v0.1.0
	gorm.io/driver/postgres v1.4.8
	gorm.io/gorm v1.24.6
)

require (
	github.com/Zilliqa/gozilliqa-sdk v1.2.0 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/form v3.1.4+incompatible // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/holiman/big v0.0.0-20221017200358-a027dc42d04e // indirect
	github.com/holiman/uint256 v1.2.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/ipfs/go-cid v0.3.2 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jgimeno/go-namehash v0.0.0-20180808144722-2773b36cc1f8 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.16.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multibase v0.1.1 // indirect
	github.com/multiformats/go-multihash v0.2.1 // indirect
	github.com/multiformats/go-varint v0.0.7 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.9.5 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vektah/gqlparser/v2 v2.5.1 // indirect
	github.com/wealdtech/go-multicodec v1.4.0 // indirect
	github.com/ybbus/jsonrpc v2.1.2+incompatible // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.29.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/blake3 v1.1.7 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)
