module github.com/mazesoul87/go-library/service/pinduoduo

go 1.23.1

replace github.com/mazesoul87/go-library/utils/gojson => ../../utils/gojson

replace github.com/mazesoul87/go-library/utils/gotime => ../../utils/gotime

replace github.com/mazesoul87/go-library/utils/gostring => ../../utils/gostring

replace github.com/mazesoul87/go-library/utils/gorandom => ../../utils/gorandom

replace github.com/mazesoul87/go-library/utils/gorequest => ../../utils/gorequest

replace github.com/mazesoul87/go-library/utils/godecimal => ../../utils/godecimal

require (
	github.com/mazesoul87/go-library/utils/godecimal v1.0.11
	github.com/mazesoul87/go-library/utils/gojson v1.0.7
	github.com/mazesoul87/go-library/utils/gorequest v1.0.83
	github.com/mazesoul87/go-library/utils/gostring v1.0.21
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
)

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403 // indirect
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/mazesoul87/go-library/utils/gorandom v1.0.4 // indirect
	github.com/mazesoul87/go-library/utils/gotime v1.0.12 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.55.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)
