

benchmark:				\
	benchmark-serialization		\
	benchmark-deserialization




benchmark-serialization:
	@$(call benchmark,serialization)


benchmark-deserialization:
	@$(call benchmark,deserialization)
