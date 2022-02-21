.PHONY:\
benchmark
benchmark:									               \
	benchmark-serialization-performance		               \
	benchmark-deserialization-performance                  \
	benchmark-serialization-deserialization-performance    \
	benchmark-serialization-message-size-footprint         \
	merge-output-images-of-plots

.PHONY:\
benchmark-serialization-performance
benchmark-serialization-performance:
	@$(call benchmark-performance,a-serialization-performance)

.PHONY:\
benchmark-deserialization-performance
benchmark-deserialization-performance:
	@$(call benchmark-performance,b-deserialization-performance)

.PHONY:\
benchmark-serialization-deserialization-performance
benchmark-serialization-deserialization-performance:
	@$(call benchmark-performance,c-serialization-deserialization-performance)

.PHONY:\
benchmark-serialization-message-size-footprint
benchmark-serialization-message-size-footprint:
	@$(call benchmark-serialization-eventual-message-size-footprint,d-serialization-eventual-message-size-footprint)

.PHONY:\
merge-output-images-of-plots
merge-output-images-of-plots:
	@montage   -mode concatenate     ./arena-results/*--result.png    ./arena-results/x-all-results.png
