module github.com/mazesoul87/go-library/utils/gophp

go 1.23

toolchain go1.23.2

replace github.com/mazesoul87/go-library/utils/gostring => ../../utils/gostring

replace github.com/mazesoul87/go-library/utils/gotime => ../../utils/gotime

replace github.com/mazesoul87/go-library/utils/gorandom => ../../utils/gorandom

require github.com/mazesoul87/go-library/utils/gostring v1.0.20

require (
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/mazesoul87/go-library/utils/gojson v1.0.7 // indirect
	github.com/mazesoul87/go-library/utils/gorandom v1.0.4 // indirect
	github.com/mazesoul87/go-library/utils/gotime v1.0.12 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)
