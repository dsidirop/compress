define benchmark-single-metric =
             dos2unix   --quiet     './make-tasks/benchmark-single-metric.sh'    \
        &&   chmod      +x          './make-tasks/benchmark-single-metric.sh'    \
        &&                          './make-tasks/benchmark-single-metric.sh'    '$(1)'   '$(2)'   '$(3)'   '$(4)'
endef
