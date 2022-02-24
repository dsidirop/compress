package arena

import "github.com/apache/thrift/lib/go/thrift"

func NewThriftCompactSerializer() *thrift.TSerializer {
	transport := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{}).GetProtocol(transport)

	return &thrift.TSerializer{
		Transport: transport,
		Protocol:  protocol,
	}
}

func NewThriftCompactDeserializer() *thrift.TDeserializer {
	transport := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTCompactProtocolFactoryConf(&thrift.TConfiguration{}).GetProtocol(transport)

	return &thrift.TDeserializer{
		Transport: transport,
		Protocol:  protocol,
	}
}
