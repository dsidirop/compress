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
	@$(call benchmark-performance,a-serialization-performance)
	@montage   -mode concatenate   -tile x1   ./arena-results/a-*--result.png    ./arena-results/a--category-result.png

.PHONY:\
benchmark-deserialization-performance
benchmark-deserialization-performance:
	@$(call benchmark-performance,b-deserialization-performance)
	@montage   -mode concatenate   -tile x1   ./arena-results/b-*--result.png    ./arena-results/b--category-result.png

.PHONY:\
benchmark-serialization-deserialization-performance
benchmark-serialization-deserialization-performance:
	@$(call benchmark-performance,c-serialization-deserialization-performance)
	@montage   -mode concatenate   -tile x1   ./arena-results/c-*--result.png    ./arena-results/c--category-result.png

.PHONY:\
benchmark-serialization-deserialization-elapsed-time
benchmark-serialization-deserialization-elapsed-time:
	@$(call benchmark-single-metric,d-serialization-deserialization-elapsed-time,Time in Nanoseconds)
	@montage   -mode concatenate   -tile x1   ./arena-results/d-*--result.png    ./arena-results/d--category-result.png

.PHONY:\
benchmark-serialization-message-size-footprint
benchmark-serialization-message-size-footprint:
	@$(call benchmark-single-metric,e-serialization-eventual-message-size-footprint, Size in Bytes)
	@montage   -mode concatenate   -tile x1   ./arena-results/e-*--result.png    ./arena-results/e--category-result.png

.PHONY:\
merge-output-images-of-plots
merge-output-images-of-plots:
	@montage   -mode concatenate   -tile x100   ./arena-results/?--category-result.png    ./arena-results/x-all-results.png
