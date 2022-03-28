define benchmark-performance =
             dos2unix   --quiet     './make-tasks/benchmark-performance.sh'    \
        &&   chmod      +x          './make-tasks/benchmark-performance.sh'    \
        &&                          './make-tasks/benchmark-performance.sh'    '$(1)'  '$(2)'  '$(3)'  '$(4)'
endef
