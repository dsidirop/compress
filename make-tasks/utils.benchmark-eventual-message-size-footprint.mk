define benchmark-serialization-eventual-message-size-footprint =
       cd  './arena/$(1)/'                                                                                                                            \
                                                                                                                                                      \
    && go  test   -bench=.   |    tee   './$(1)---benchmark-raw-output.dat'                                                                           \
                                                                                                                                                      \
    && awk                                                                                                                                            \
              '/[*][*]/{count++; printf("%d,%s,%s\n", count, $$4, $$5); }'                                                                            \
                './$(1)---benchmark-raw-output.dat'                                                                                                   \
              > './$(1)---benchmark-output-parsed.dat'                                                                                                \
                                                                                                                                                      \
    && messageSizeMax=`                    awk -F','  'BEGIN{a=0}{ if ($$3>0+a) a=$$3} END{print a}'     "./$(1)---benchmark-output-parsed.dat"    `  \
                                                                                                                                                      \
    && messageSizeMaxRoundedUpwards=`      awk -v n="$${messageSizeMax}"      'BEGIN{ print int((n+100) / 100 ) * 100 }'                           `  \
                                                                                                                                                      \
    && tempDir=`mktemp -d -t golang-compression-libs-arena.XXXX`                                                                                      \
                                                                                                                                                      \
    && cp            '../plot.gp'                               "$${tempDir}/plot.gp"                                                                 \
    && sed    -i     's/___TITLE___/Eventual Size in Bytes/g'   "$${tempDir}/plot.gp"                                                                 \
    && gnuplot                                                                                                                                        \
            -e "    file_path='./$(1)---benchmark-output-parsed.dat'                                     "                                            \
            -e "    graphic_file_name='../../arena-results/$(1)--eventual-message-size-footprint.png'    "                                            \
            -e "    y_label='bytes'                                                                      "                                            \
            -e "    y_range_min='0000''                                                                  "                                            \
            -e "    y_range_max='$${messageSizeMaxRoundedUpwards}'                                       "                                            \
            -e "    column_1=1                                                                           "                                            \
            -e "    column_2=3                                                                           "                                            \
            "$${tempDir}/plot.gp"                                                                                                                     \
                                                                                                                                                      \
    && cp       './$(1)---benchmark-output-parsed.dat'   '../../arena-results'                                                                        \
                                                                                                                                                      \
    && echo     'Plots successfully generated';                                                                                                       \
                                                                                                                                                      \
    rm -rf "$${tempDir}"
endef
