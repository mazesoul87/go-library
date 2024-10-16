module github.com/mazesoul87/go-library/utils/gojobs

go 1.23.0

replace github.com/mazesoul87/go-library/utils/gojson => ../../utils/gojson

replace github.com/mazesoul87/go-library/utils/gotime => ../../utils/gotime

replace github.com/mazesoul87/go-library/utils/gorequest => ../../utils/gorequest

replace github.com/mazesoul87/go-library/utils/gostring => ../../utils/gostring

replace github.com/mazesoul87/go-library/utils/gorandom => ../../utils/gorandom

require (
	entgo.io/ent v0.14.1
	github.com/redis/go-redis/v9 v9.6.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/mazesoul87/go-library/utils/gojson v1.0.7
	github.com/mazesoul87/go-library/utils/gorequest v1.0.84
	github.com/mazesoul87/go-library/utils/gostring v1.0.21
	github.com/mazesoul87/go-library/utils/gotime v1.0.12
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
	gorm.io/gorm v1.25.12
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mazesoul87/go-library/utils/gorandom v1.0.4 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.55.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)
