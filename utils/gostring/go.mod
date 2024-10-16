module github.com/mazesoul87/go-library/utils/gostring

go 1.23

replace github.com/mazesoul87/go-library/utils/gojson => ../../utils/gojson

replace github.com/mazesoul87/go-library/utils/gotime => ../../utils/gotime

replace github.com/mazesoul87/go-library/utils/gorandom => ../../utils/gorandom

require (
	github.com/mazesoul87/go-library/utils/gojson v0.0.0-00010101000000-000000000000
	github.com/mazesoul87/go-library/utils/gorandom v0.0.0-00010101000000-000000000000
	github.com/mazesoul87/go-library/utils/gotime v0.0.0-00010101000000-000000000000
)

require (
	github.com/basgys/goxml2json v1.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
