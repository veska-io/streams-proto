streams:
	protoc -I proto proto/binance/klines.proto \
			--go_out=gen/go/streams \
			--go_opt=paths=source_relative \
			--go-grpc_out=gen/go/streams \
			--go-grpc_opt=paths=source_relative \
			--experimental_allow_proto3_optional