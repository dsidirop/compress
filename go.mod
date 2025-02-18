module github.com/klauspost/compress

go 1.18

// github.com/golang/snappy                   doesnt really offer any gains
// github.com/yasushi-saito/zlibng            doesnt really offer any gains
// github.com/yasushi-saito/cloudflare-zlib   doesnt really offer any gains

require (
	github.com/andybalholm/brotli v1.0.4
	github.com/apache/thrift v0.16.0
	github.com/dsnet/compress v0.0.1
	github.com/fxamacker/cbor/v2 v2.4.0
	github.com/hamba/avro v1.6.6
	github.com/json-iterator/go v1.1.12
	github.com/pierrec/lz4/v4 v4.1.14
	github.com/tinylib/msgp v1.1.6
	github.com/mailru/easyjson v0.7.7
	github.com/vmihailenco/msgpack/v5 v5.3.5
	go.mongodb.org/mongo-driver v1.8.3
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/x448/float16 v0.8.4 // indirect
)
