define build_docker_image =
           GOPRIVATE="bitbucket.org/laerdalmedical"                   \
        && DOCKER_BUILDKIT=1                                          \
        && COMPOSE_DOCKER_CLI_BUILD=1                                 \
        && docker                                                     \
                build                                                 \
                --ssh=default                                         \
                --tag="$(1)"                                          \
                --file="$(2)"                                         \
                --progress=plain                                      \
                .
endef
