module github.com/TekClinic/Patients-MicroService/server

go 1.22.0

toolchain go1.22.2

require (
	github.com/TekClinic/MicroService-Lib v0.1.1
	github.com/TekClinic/Patients-MicroService/patients_protobuf v0.100.0-integrated
	github.com/go-playground/validator/v10 v10.22.0
	github.com/sa-/slicefunk v0.1.4
	github.com/uptrace/bun v1.2.1
	github.com/uptrace/bun/dialect/pgdialect v1.2.1
	github.com/uptrace/bun/driver/pgdriver v1.2.1
	github.com/uptrace/bun/extra/bundebug v1.2.1
	google.golang.org/grpc v1.65.0
)

require (
	github.com/coreos/go-oidc/v3 v3.10.0 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.4 // indirect
	github.com/go-jose/go-jose/v4 v4.0.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	k8s.io/apimachinery v0.30.2 // indirect
	mellium.im/sasl v0.3.1 // indirect
)

replace github.com/TekClinic/Patients-MicroService/patients_protobuf => ./../patients_protobuf
