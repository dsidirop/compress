define benchmark-serialization-eventual-message-size-footprint =
       cd  './arena/$(1)/'                                                            \
                                                                                      \
    && go  test   -bench=.   |    tee   './$(1)--benchmark-raw-output.dat'            \
                                                                                      \
    && awk                                                                            \
              '/[*][*]/{count++; printf("%d,%s,%s\n", count, $$4, $$5); }'            \
                './$(1)--benchmark-raw-output.dat'                                    \
              > './$(1)--benchmark-output-parsed.dat'                                 \
                                                                                                                     \
    && gnuplot                                                                                                       \
            -e "    file_path='./$(1)--benchmark-output-parsed.dat'                                      "           \
            -e "    graphic_file_name='../../arena-results/$(1)--eventual-message-size-footprint.png'    "           \
            -e "    y_label='bytes'                                                                      "           \
            -e "    y_range_min='0000''                                                                  "           \
            -e "    y_range_max='2000'                                                                   "           \
            -e "    column_1=1                                                                           "           \
            -e "    column_2=3                                                                           "           \
            '../plot.gp'                                                                                             \
                                                                                      \
    && cp       './$(1)--benchmark-output-parsed.dat'   '../../arena-results'         \
                                                                                      \
    && echo     'Plots successfully generated'

endef
