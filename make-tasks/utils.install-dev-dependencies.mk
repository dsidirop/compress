define install-dev-dependencies =
             dos2unix   --quiet     './make-tasks/install-dev-dependencies.sh'    \
        &&   chmod      +x          './make-tasks/install-dev-dependencies.sh'    \
        &&                          './make-tasks/install-dev-dependencies.sh'    '$(1)'
endef
