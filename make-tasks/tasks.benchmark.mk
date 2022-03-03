# https://github.com/glycerine/golang-thrift-minimal-example/blob/master/serialize.go

cpucount ?= 1

.PHONY:\
benchmark
benchmark:									               \
	benchmark-serialization-performance		               \
	benchmark-serialization-with-compression-performance   \
	benchmark-deserialization-performance                  \
	benchmark-serialization-deserialization-performance    \
	benchmark-serialization-message-size-footprint         \
	merge-output-images-of-plots  #		benchmark-serialization-deserialization-elapsed-time

.PHONY:\
benchmark-serialization-performance
benchmark-serialization-performance:    compile-idl
	@$(call benchmark-performance,a-serialization-performance,$(cpucount))

.PHONY:\
benchmark-serialization-with-compression-performance
benchmark-serialization-with-compression-performance:    compile-idl
	@$(call benchmark-performance,a-serialization-with-compression-performance,$(cpucount))

.PHONY:\
benchmark-deserialization-performance
benchmark-deserialization-performance:    compile-idl
	@$(call benchmark-performance,b-deserialization-performance,$(cpucount))

.PHONY:\
benchmark-serialization-deserialization-performance
benchmark-serialization-deserialization-performance:    compile-idl
	@$(call benchmark-performance,c-serialization-deserialization-performance,$(cpucount))

# .PHONY:\
# benchmark-serialization-deserialization-elapsed-time
# benchmark-serialization-deserialization-elapsed-time:    compile-idl
# 	@$(call benchmark-single-metric,d-serialization-deserialization-elapsed-time,Average Elapsed Time in nsecs - Lower is better,ns,$(cpucount))

.PHONY:\
benchmark-serialization-message-size-footprint
benchmark-serialization-message-size-footprint:    compile-idl
	@$(call benchmark-single-metric,e-serialization-eventual-message-size-footprint,Eventual Size in Bytes - Lower is better,bytes,$(cpucount))

.PHONY:\
merge-output-images-of-plots
merge-output-images-of-plots: # merge all images into one
	@convert     -append       './arena-results/*-cpu$(cpucount)----category-overall-results.png'      './arena-results/x-cpu$(cpucount)-all-results.png'

.PHONY:\
compile-idl
compile-idl:           \
	compile-msgp       \
	compile-avro       \
	compile-thrift     \
	compile-protobuf

compile-msgp: ./arena/fooitem.go
	@cd arena  &&  msgp  --file  fooitem.go
	@touch  $@

compile-avro: ./arena/*.avdl
	@cd arena  &&  java   -jar avro-tools.jar   idl     ./avfooitem.avdl         ./avfooitem.avsc
	@touch  $@

compile-thrift: ./arena/*.thrift
	@cd arena  &&  thrift    --gen go    -recurse     -out .    thfooitem.thrift
	@touch  $@

compile-protobuf: ./arena/*.proto
	@cd arena  &&  protoc    --go_out=.    pbfooitem.proto
	@touch  $@
