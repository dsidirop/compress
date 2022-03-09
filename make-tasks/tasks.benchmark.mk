# [snappy           ](https://github.com/klauspost/compress/tree/master/snappy             ) is a drop-in replacement for `github.com/golang/snappy` offering better compression and concurrent streams.
# [huff0            ](https://github.com/klauspost/compress/tree/master/huff0              ) and [FSE](https://github.com/klauspost/compress/tree/master/fse) implementations for raw entropy encoding.
# [gzhttp           ](https://github.com/klauspost/compress/tree/master/gzhttp             ) Provides client and server wrappers for handling gzipped requests efficiently.
# [pgzip            ](https://github.com/klauspost/pgzip                                   ) is a separate package that provides a very fast parallel gzip implementation.
# [fuzz package     ](https://github.com/klauspost/compress-fuzz                           ) for fuzz testing all compressors/decompressors here.

cpucount ?= 1

.PHONY:\
benchmark
benchmark:									               \
	benchmark-serialization-performance		               \
	benchmark-serialization-with-compression-performance   \
	benchmark-deserialization-performance                  \
	benchmark-serialization-deserialization-performance    \
	benchmark-serialization-message-size-footprint         \
	benchmark-serialization-deserialization-with-compression-elapsed-time \
	merge-output-images-of-plots  #		benchmark-serialization-deserialization-elapsed-time

.PHONY:\
benchmark-serialization-performance
benchmark-serialization-performance:    compile-idl
	@$(call benchmark-performance,aa-serialization-performance,$(cpucount))

.PHONY:\
benchmark-deserialization-performance
benchmark-deserialization-performance:    compile-idl
	@$(call benchmark-performance,ab-deserialization-performance,$(cpucount))

.PHONY:\
benchmark-serialization-deserialization-performance
benchmark-serialization-deserialization-performance:    compile-idl
	@$(call benchmark-performance,ac-serialization-deserialization-performance,$(cpucount))

# .PHONY:\
# benchmark-serialization-deserialization-elapsed-time
# benchmark-serialization-deserialization-elapsed-time:    compile-idl
# 	@$(call benchmark-single-metric,ad-serialization-deserialization-elapsed-time,Average Elapsed Time in nsecs - Lower is better,ns,$(cpucount))

.PHONY:\
benchmark-serialization-message-size-footprint
benchmark-serialization-message-size-footprint:    compile-idl
	@$(call benchmark-single-metric,ae-serialization-eventual-message-size-footprint,Eventual Size in Bytes - Lower is better,bytes,$(cpucount))

.PHONY:\
benchmark-serialization-with-compression-performance
benchmark-serialization-with-compression-performance:    compile-idl
	@$(call benchmark-performance,ba-serialization-with-compression-performance,$(cpucount),../plot.serialization-with-compression.gp)

.PHONY:\
benchmark-decompression-deserialization-performance
benchmark-decompression-deserialization-performance:    compile-idl
	@$(call benchmark-performance,bb-decompression-deserialization-performance,$(cpucount))

.PHONY:\
benchmark-serialization-deserialization-with-compression-performance
benchmark-serialization-deserialization-with-compression-performance:    compile-idl
	@$(call benchmark-performance,bc-serialization-deserialization-with-compression-performance,$(cpucount))

.PHONY:\
benchmark-serialization-deserialization-with-compression-elapsed-time
benchmark-serialization-deserialization-with-compression-elapsed-time:    compile-idl
	@$(call benchmark-single-metric,bd-serialization-deserialization-with-compression-elapsed-time,Average Elapsed Time in nsecs - Lower is better,ns,$(cpucount),../plot.serialization-deserialization-with-compression-elapsed-time.gp)



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
