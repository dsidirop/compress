# 1    run amd64 and arm64 tests on the dev-cluster
# 2    make cbor work in c/c++
# 3    calculate the impact in terms of cost-savings for each compression - will need to update levo-emulator-cli + grafana charts too probably

cpucount ?= 1

.PHONY:\
benchmark
benchmark:                        \
	compile-idl                   \
	benchmark-raw


.PHONY:\
benchmark-raw
benchmark-raw:									                                  \
    benchmark-serialization-performance                                           \
    benchmark-deserialization-performance                                         \
    benchmark-serialization-deserialization-performance                           \
    benchmark-serialization-message-size-footprint                                \
    benchmark-serialization-with-compression-performance                          \
    benchmark-decompression-deserialization-performance                           \
    benchmark-serialization-deserialization-with-compression-performance          \
    benchmark-serialization-deserialization-with-compression-elapsed-time         \
    benchmark-serialization-with-compression-eventual-message-size                \
	merge-output-images-of-plots                                                  \
	generate-scatter-plot-for-size-vs-time-impl  #		benchmark-serialization-deserialization-elapsed-time


.PHONY:\
benchmark-serialization-performance
benchmark-serialization-performance:
	@$(call benchmark-performance,aa-serialization-performance,$(cpucount))

.PHONY:\
benchmark-deserialization-performance
benchmark-deserialization-performance:
	@$(call benchmark-performance,ab-deserialization-performance,$(cpucount))

.PHONY:\
benchmark-serialization-deserialization-performance
benchmark-serialization-deserialization-performance:
	@$(call benchmark-performance,ac-serialization-deserialization-performance,$(cpucount))

# .PHONY:\
# benchmark-serialization-deserialization-elapsed-time
# benchmark-serialization-deserialization-elapsed-time:
# 	@$(call benchmark-single-metric,ad-serialization-deserialization-elapsed-time,Average Elapsed Time in nsecs - Lower is better,ns,$(cpucount))

.PHONY:\
benchmark-serialization-message-size-footprint
benchmark-serialization-message-size-footprint:
	@$(call benchmark-single-metric,ae-serialization-eventual-message-size-footprint,Total  Bytes - Lower is better,bytes,$(cpucount))

.PHONY:\
benchmark-serialization-with-compression-performance
benchmark-serialization-with-compression-performance:
	@$(call benchmark-performance,ba-serialization-with-compression-performance,$(cpucount),../plot.serialization-with-compression.gp,vertical)

.PHONY:\
benchmark-decompression-deserialization-performance
benchmark-decompression-deserialization-performance:
	@$(call benchmark-performance,bb-decompression-deserialization-performance,$(cpucount),../plot.serialization-with-compression.gp,vertical)

.PHONY:\
benchmark-serialization-deserialization-with-compression-performance
benchmark-serialization-deserialization-with-compression-performance:
	@$(call benchmark-performance,bc-serialization-deserialization-with-compression-performance,$(cpucount),../plot.serialization-with-compression.gp,vertical)

.PHONY:\
benchmark-serialization-deserialization-with-compression-elapsed-time
benchmark-serialization-deserialization-with-compression-elapsed-time:
	@$(call benchmark-single-metric,bd-serialization-deserialization-with-compression-elapsed-time,Average Elapsed Time in nsecs - Lower is better,ns,$(cpucount),../plot.serialization-deserialization-with-compression-elapsed-time.gp)

.PHONY:\
benchmark-serialization-with-compression-eventual-message-size
benchmark-serialization-with-compression-eventual-message-size:
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

.PHONY:\
generate-scatter-plot-for-size-vs-time-impl
generate-scatter-plot-for-size-vs-time-impl:
	@$(call generate-scatter-plot-for-size-vs-time,$(cpucount))

.PHONY:\
compile-idl
compile-idl:           \
	compile-msgp       \
	compile-avro       \
	compile-thrift     \
	compile-protobuf   \
	compile-easyjson

compile-msgp: ./arena/fooitem.go   ./arena/curvegenreplyv1.go   ./arena/vitalstemplate.go    ./arena/simeventregistereventcmd.go
	( which msgp                                            &&    cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  msgp      --file                           "./$${x}"                      ;  done ) || exit 0
	@touch  $@

compile-avro: ./arena/*.avdl   # java must be jdk8    jdk14+ doesn't work for some reason
	( which java    &&   test -f "./arena/avro-tools.jar"   &&    cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  java      -jar       avro-tools.jar   idl  "./$${x}"   "./$${x%.*}.avsc"  ;  done ) || exit 0
	@touch  $@

compile-thrift: ./arena/*.thrift
	( which thrift                                          &&    cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  thrift    --gen      go    -recurse     -out .    "./$${x}"               ;  done ) || exit 0
	@touch  $@

compile-protobuf: ./arena/*.proto
	( which protoc                                          &&    cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  protoc    --go_out=.    "./$${x}"                                         ;  done ) || exit 0
	@touch  $@

compile-easyjson: ./arena/fooitem.go   ./arena/curvegenreplyv1.go   ./arena/vitalstemplate.go    ./arena/simeventregistereventcmd.go
	( which easyjson                                        &&    cd arena    &&    for file in $^ ; do    x=$$(basename "$${file}");  easyjson   -all      "./$${x}"                                            ;  done ) || exit 0
	@touch  $@
