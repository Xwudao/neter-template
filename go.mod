module github.com/Xwudao/neter-template

go 1.25.0

// gotossr currently needs the local development checkout for the TanStack
// Router compatibility fixes used by this hybrid SSR integration.
replace github.com/Xwudao/gotossr => /Users/tim/Codes/ai-project/gotossr

require (
	ariga.io/atlas v0.19.1-0.20240203083654-5948b60a8e43
	entgo.io/ent v0.14.0
	github.com/PuerkitoBio/goquery v1.10.3
	github.com/Xwudao/gotossr v1.0.9-0.20260906090231-e6ce1152f227
	github.com/aws/aws-sdk-go v1.34.0
	github.com/gin-contrib/cors v1.7.7
	github.com/gin-gonic/gin v1.12.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.5.0
	github.com/json-iterator/go v1.1.12
	github.com/knadh/koanf/parsers/json v0.1.0
	github.com/knadh/koanf/parsers/yaml v0.1.0
	github.com/knadh/koanf/providers/file v0.1.0
	github.com/knadh/koanf/v2 v2.0.1
	github.com/natefinch/lumberjack/v3 v3.0.0-alpha
	github.com/robfig/cron/v3 v3.0.1
	github.com/snabb/sitemap v1.0.4
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.11.1
	go.uber.org/mock v0.6.0
	go.uber.org/ratelimit v0.3.0
	go.uber.org/zap v1.25.0
	golang.org/x/crypto v0.49.0
	golang.org/x/net v0.52.0
	golang.org/x/time v0.3.0
)

require (
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/andybalholm/cascadia v1.3.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/buke/quickjs-go v0.6.7 // indirect
	github.com/bytedance/gopkg v0.1.4 // indirect
	github.com/bytedance/sonic v1.15.0 // indirect
	github.com/bytedance/sonic/loader v0.5.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/evanw/esbuild v0.27.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.13 // indirect
	github.com/gin-contrib/sse v1.1.1 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.30.2 // indirect
	github.com/goccy/go-json v0.10.6 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/pelletier/go-toml/v2 v2.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/quic-go/qpack v0.6.0 // indirect
	github.com/quic-go/quic-go v0.59.0 // indirect
	github.com/redis/go-redis/v9 v9.17.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rosbit/dukgo v0.8.2 // indirect
	github.com/rosbit/go-embedding-utils v0.4.1 // indirect
	github.com/snabb/diagio v1.0.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tkrajina/go-reflector v0.5.8 // indirect
	github.com/tkrajina/typescriptify-golang-structs v0.2.0 // indirect
	github.com/tommie/v8go v0.34.0 // indirect
	github.com/tommie/v8go/deps/android_amd64 v0.0.0-20250515043113-5dcc98077472 // indirect
	github.com/tommie/v8go/deps/android_arm64 v0.0.0-20250515043113-5dcc98077472 // indirect
	github.com/tommie/v8go/deps/darwin_amd64 v0.0.0-20250515043113-5dcc98077472 // indirect
	github.com/tommie/v8go/deps/darwin_arm64 v0.0.0-20250515043113-5dcc98077472 // indirect
	github.com/tommie/v8go/deps/linux_amd64 v0.0.0-20250515043113-5dcc98077472 // indirect
	github.com/tommie/v8go/deps/linux_arm64 v0.0.0-20250515043113-5dcc98077472 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.1 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.mongodb.org/mongo-driver/v2 v2.5.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/arch v0.25.0 // indirect
	golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/mod v0.33.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	modernc.org/libc v1.66.10 // indirect
	modernc.org/libquickjs v0.12.2 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
	modernc.org/quickjs v0.17.0 // indirect
)
