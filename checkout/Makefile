CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
MOCKGEN=${BINDIR}/mockgen_${MOCKGENVER}_${GOVER}
MOCKGENVER=v1.6.0
PROTOC=${BINDIR}/protoc_gen_go_${PROTOCVER}_${GOVER}
PROTOCVER=v1.28.1
PROTOCGRPC=${BINDIR}/protoc_gen_go_grpc_${PROTOCGRPCVER}_${GOVER}
PROTOCGRPCVER=v1.2
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=gitlab.ozon.dev/nlnaa/homework-1/checkout/cmd/app

all: format build test lint

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

generate-checkoutv1:
	mkdir -p pkg/checkout_v1
	protoc -I api/checkout_v1 -I vendor_proto \
	--go_out=pkg/checkout_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=${PROTOC} \
	--go-grpc_out=pkg/checkout_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=${PROTOCGRPC} \
	api/checkout_v1/checkout_v1_service.proto

generate-lomsv1:
	mkdir -p pkg/loms_v1
	protoc -I api/loms_v1 -I vendor_proto \
	--go_out=pkg/loms_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=${PROTOC} \
	--go-grpc_out=pkg/loms_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=${PROTOCGRPC} \
	api/loms_v1/loms_v1_service.proto
	rm -rf api/loms_v1

generate-product:
	mkdir -p pkg/product
	protoc -I api/product \
	--go_out=pkg/product --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=${PROTOC} \
	--go-grpc_out=pkg/product --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=${PROTOCGRPC} \
	api/product/product_service.proto
	rm -rf api/product

generate-protoc: install-protoc install-protoc-grpc get-protoc-deps vendor-proto \
	generate-checkoutv1 generate-lomsv1 generate-product

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

get-protoc-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

vendor-proto:
	mkdir -p vendor_proto
	@if [ ! -d vendor_proto/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor_proto/googleapis && \
		mkdir -p  vendor_proto/google/ && \
		mv vendor_proto/googleapis/google/api vendor_proto/google && \
		rm -rf vendor_proto/googleapis ; \
	fi
	@if [ ! -d vendor_proto/google/protobuf ]; then \
		git clone https://github.com/protocolbuffers/protobuf vendor_proto/protobuf && \
		mkdir -p  vendor_proto/google/protobuf && \
		mv vendor_proto/protobuf/src/google/protobuf/*.proto vendor_proto/google/protobuf && \
		rm -rf vendor_proto/protobuf ; \
	fi

install-protoc: bindir
	test -f ${PROTOC} || \
		(GOBIN=${BINDIR} go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOCVER} && \
		mv ${BINDIR}/protoc-gen-go ${PROTOC})

install-protoc-grpc: bindir		
	test -f ${PROTOCGRPC} || \
		(GOBIN=${BINDIR} go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@${PROTOCGRPCVER} && \
		mv ${BINDIR}/protoc-gen-go-grpc ${PROTOCGRPC})

install-mockgen: bindir
	test -f ${MOCKGEN} || \
    	(GOBIN=${BINDIR} go install github.com/golang/mock/mockgen@{MOCKGENVER} && \
   	 	mv ${BINDIR}/mockgen ${MOCKGEN})

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})
