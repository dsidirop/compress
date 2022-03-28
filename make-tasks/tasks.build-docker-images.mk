.PHONY:\
build-docker-image-amd64
build-docker-image-amd64:
	@$(call build_docker_image,magellanrepo.azurecr.io/tests/amd64/magellan-serialization-and-compression-benchmarks-testbed:localdev,Dockerfile.amd64)


.PHONY:\
build-docker-image-arm64
build-docker-image-arm64:
	$(call build_docker_image,magellanrepo.azurecr.io/tests/arm64/magellan-serialization-and-compression-benchmarks-testbed:localdev,Dockerfile.arm64)
