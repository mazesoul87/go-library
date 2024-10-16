module github.com/mazesoul87/go-library/utils/gorequest

go 1.23

replace github.com/mazesoul87/go-library/utils/gojson => ../../utils/gojson

replace github.com/mazesoul87/go-library/utils/gotime => ../../utils/gotime

replace github.com/mazesoul87/go-library/utils/gostring => ../../utils/gostring

replace github.com/mazesoul87/go-library/utils/gorandom => ../../utils/gorandom

require (
	github.com/MercuryEngineering/CookieMonster v0.0.0-20180304172713-1584578b3403
	github.com/mazesoul87/go-library/utils/gojson v0.0.0-00010101000000-000000000000
	github.com/mazesoul87/go-library/utils/gostring v0.0.0-00010101000000-000000000000
	github.com/mazesoul87/go-library/utils/gotime v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.55.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/otel/trace v1.30.0
)

require (
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/mazesoul87/go-library/utils/gorandom v1.0.4 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)
