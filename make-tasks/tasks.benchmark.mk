# 1    vitalset diversification
# 2    move the repo to private
# 3    benchmark throughput mb/sec


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
	benchmark-serialization-with-compression-eventual-message-size        \
	merge-output-images-of-plots                                          \
	generate-scatter-plot-for-size-vs-time-impl  #		benchmark-serialization-deserialization-elapsed-time


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
	@$(call benchmark-single-metric,ae-serialization-eventual-message-size-footprint,Total  Bytes - Lower is better,bytes,$(cpucount))

.PHONY:\
benchmark-serialization-with-compression-performance
benchmark-serialization-with-compression-performance:    compile-idl
	@$(call benchmark-performance,ba-serialization-with-compression-performance,$(cpucount),../plot.serialization-with-compression.gp,vertical)

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
benchmark-serialization-with-compression-eventual-message-size
benchmark-serialization-with-compression-eventual-message-size:    compile-idl
	@$(call benchmark-single-metric,be-serialization-with-compression-eventual-message-size,Total Bytes - Lower is better,bytes,$(cpucount),../plot.serialization-with-compression-eventual-message-size.gp)

.PHONY:\
merge-output-images-of-plots
merge-output-images-of-plots: # merge all images into one
	@convert     -append       './arena-results/*-cpu$(cpucount)----category-overall-results.png'      './arena-results/x-cpu$(cpucount)-all-results.png'

.PHONY:\
generate-scatter-plot-for-size-vs-time
generate-scatter-plot-for-size-vs-time:                                                      \
    benchmark-serialization-deserialization-with-compression-elapsed-time                    \
	benchmark-serialization-with-compression-eventual-message-size                           \
	generate-scatter-plot-for-size-vs-time-impl
	@$(call generate-scatter-plot-for-size-vs-time)

.PHONY:\
generate-scatter-plot-for-size-vs-time-impl
generate-scatter-plot-for-size-vs-time-impl:
	@$(call generate-scatter-plot-for-size-vs-time)

.PHONY:\
compile-idl
compile-idl:           \
	compile-msgp       \
	compile-avro       \
	compile-thrift     \
	compile-protobuf

compile-msgp: ./arena/fooitem.go   ./arena/curvegenreplyv1.go   ./arena/vitalstemplate.go
	@cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  msgp      --file                           "./$${x}"                      ;  done
	@touch  $@

compile-protobuf: ./arena/*.proto
	@cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  protoc    --go_out=.    "./$${x}"  ;  done
	@touch  $@

compile-avro: ./arena/*.avdl
	@cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  java      -jar       avro-tools.jar   idl  "./$${x}"   "./$${x%.*}.avsc"  ;  done
	@touch  $@

compile-thrift: ./arena/*.thrift
	@cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  thrift    --gen      go    -recurse     -out .    "./$${x}"  ;  done
	@touch  $@
