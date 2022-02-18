.PHONY:\
benchmark
benchmark:									\
	benchmark-serialization-performance		\
	benchmark-deserialization-performance   \
	benchmark-serialization-message-size-footprint

.PHONY:\
benchmark-serialization-performance
benchmark-serialization-performance:
	@$(call benchmark-performance,a-serialization-performance)

.PHONY:\
benchmark-deserialization-performance
benchmark-deserialization-performance:
	@$(call benchmark-performance,b-deserialization-performance)

.PHONY:\
benchmark-serialization-message-size-footprint
benchmark-serialization-message-size-footprint:
	@$(call benchmark-serialization-eventual-message-size-footprint,c-serialization-eventual-message-size-footprint)
