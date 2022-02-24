# https://github.com/glycerine/golang-thrift-minimal-example/blob/master/serialize.go

cpucount ?= 1

.PHONY:\
benchmark
benchmark:									               \
	benchmark-serialization-performance		               \
	benchmark-deserialization-performance                  \
	benchmark-serialization-deserialization-performance    \
	benchmark-serialization-message-size-footprint         \
	merge-output-images-of-plots  #		benchmark-serialization-deserialization-elapsed-time


.PHONY:\
benchmark-serialization-performance
benchmark-serialization-performance:
	@$(call benchmark-performance,a-serialization-performance,$(cpucount))

.PHONY:\
benchmark-deserialization-performance
benchmark-deserialization-performance:
	@$(call benchmark-performance,b-deserialization-performance,$(cpucount))

.PHONY:\
benchmark-serialization-deserialization-performance
benchmark-serialization-deserialization-performance:
	@$(call benchmark-performance,c-serialization-deserialization-performance,$(cpucount))

# .PHONY:\
# benchmark-serialization-deserialization-elapsed-time
# benchmark-serialization-deserialization-elapsed-time:
# 	@$(call benchmark-single-metric,d-serialization-deserialization-elapsed-time,Average Elapsed Time in nsecs - Lower is better,ns,$(cpucount))

.PHONY:\
benchmark-serialization-message-size-footprint
benchmark-serialization-message-size-footprint:
	@$(call benchmark-single-metric,e-serialization-eventual-message-size-footprint,Eventual Size in Bytes - Lower is better,bytes,$(cpucount))

.PHONY:\
merge-output-images-of-plots
merge-output-images-of-plots: # merge all images into one
	@convert -append       ./arena-results/*-cpu$(cpucount)----category-overall-results.png      ./arena-results/x-cpu$(cpucount)-all-results.png
