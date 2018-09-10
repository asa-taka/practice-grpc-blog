.PHONY: go

PROTO_DIRS := $(wildcard proto/*/.)
PROTO_SRCS := $(shell find proto -type f)

PROTO_PATHS := \
	-I proto

all: go ts

# Go

GO_OUT := $(GOPATH)/src
GO_TRGS := $(PROTO_SRCS:proto/%.proto=$(GO_OUT)/%.pb.go)

go: $(GO_TRGS)

$(GO_OUT)/%.pb.go: proto/%.proto
	protoc $(PROTO_PATHS) \
		--go_out=plugins=grpc:$(GO_OUT) \
		$<

# TypeScript

TS_OUT := out/ts
TS_TRGS := $(PROTO_SRCS:proto/%.proto=$(TS_OUT)/%.pb.ts)

ts: $(TS_TRGS)

$(TS_OUT)/%.pb.ts: proto/%.proto
	protoc $(PROTO_PATHS) \
		--js_out="import_style=commonjs,binary:$(TS_OUT)" \
		--ts_out="$(TS_OUT)" \
		$<

# Misc

clean:
	rm -rf $(GO_OUT)/*
	rm -rf $(TS_OUT)/*
