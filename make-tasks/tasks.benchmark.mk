benchmark:									\
	benchmark-serialization-performance		\
	benchmark-deserialization-performance


benchmark-serialization-performance:
	@$(call benchmark,serializationperf)


benchmark-deserialization-performance:
	@$(call benchmark,deserializationperf)
