.PHONY:\
benchmark
benchmark:									               \
	benchmark-serialization-performance		               \
	benchmark-deserialization-performance                  \
	benchmark-serialization-deserialization-performance    \
	benchmark-serialization-deserialization-elapsed-time   \
	benchmark-serialization-message-size-footprint         \
	merge-output-images-of-plots

.PHONY:\
benchmark-serialization-performance
benchmark-serialization-performance:
	@$(call benchmark-performance,a-serialization-performance,1)

.PHONY:\
benchmark-deserialization-performance
benchmark-deserialization-performance:
	@$(call benchmark-performance,b-deserialization-performance,1)

.PHONY:\
benchmark-serialization-deserialization-performance
benchmark-serialization-deserialization-performance:
	@$(call benchmark-performance,c-serialization-deserialization-performance,1)

.PHONY:\
benchmark-serialization-deserialization-elapsed-time
benchmark-serialization-deserialization-elapsed-time:
	@$(call benchmark-single-metric,d-serialization-deserialization-elapsed-time,Average Elapsed Time in nsecs\\n(lower is better),ns,1)
	@montage   -mode concatenate    ./arena-results/d-*--result.png    ./arena-results/d-cpu1----category-overall-results.png

.PHONY:\
benchmark-serialization-message-size-footprint
benchmark-serialization-message-size-footprint:
	@$(call benchmark-single-metric,e-serialization-eventual-message-size-footprint,Average Size in Bytes\\n(lower is better),bytes,1)
	@montage   -mode concatenate    ./arena-results/e-*--result.png    ./arena-results/e-cpu1----category-overall-results.png

.PHONY:\
merge-output-images-of-plots
merge-output-images-of-plots: # merge all images into one
	@convert -append       ./arena-results/*-cpu1----category-overall-results.png      ./arena-results/x-cpu1-all-results.png
