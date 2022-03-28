define generate-scatter-plot-for-size-vs-time =
             dos2unix   --quiet     './make-tasks/generate-scatter-plot-for-size-vs-time.sh'    \
        &&   chmod      +x          './make-tasks/generate-scatter-plot-for-size-vs-time.sh'    \
        &&                          './make-tasks/generate-scatter-plot-for-size-vs-time.sh'   '$(1)'
endef
