module go.dtapp.net/library/utils/gophp

go 1.23

toolchain go1.23.2

replace go.dtapp.net/library/utils/gostring => ../../utils/gostring

replace go.dtapp.net/library/utils/gotime => ../../utils/gotime

replace go.dtapp.net/library/utils/gorandom => ../../utils/gorandom

require go.dtapp.net/library/utils/gostring v1.0.20

require (
	github.com/basgys/goxml2json v1.1.0 // indirect
	go.dtapp.net/library/utils/gojson v1.0.7 // indirect
	go.dtapp.net/library/utils/gorandom v1.0.4 // indirect
	go.dtapp.net/library/utils/gotime v1.0.12 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)
