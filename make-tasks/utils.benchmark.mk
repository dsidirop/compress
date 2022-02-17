define benchmark =
       cd ./arena/$(1)/                                                                               \
                                                                                                      \
    && go test -bench=.   |    tee ./$(1)__benchmark_raw_output.dat                                   \
                                                                                                      \
    && awk                                                                                            \
              '/BenchmarkTest/{count++; gsub(/BenchmarkTest___.*?___/, ""); printf("%d,%s,%s,%s\n", count, $$1, $$2, $$3); }'    \
                    ./$(1)__benchmark_raw_output.dat                                                  \
                  > ./$(1)__benchmark_output_parsed.dat                                               \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)__benchmark_output_parsed.dat'             "                         \
            -e "graphic_file_name='../../arenaresults/$(1)__operations.png' "                         \
            -e "y_label='number of operations'                              "                         \
            -e "y_range_min='0000000''                                      "                         \
            -e "y_range_max='1000000'                                       "                         \
            -e "column_1=1                                                  "                         \
            -e "column_2=3                                                  "                         \
            ../performance.gp                                                                         \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)__benchmark_output_parsed.dat'                   "                   \
            -e "graphic_file_name='../../arenaresults/$(1)__time_operations.png'  "                   \
            -e "y_label='each operation in nanoseconds'                           "                   \
            -e "y_range_min='00000''                                              "                   \
            -e "y_range_max='45000'                                               "                   \
            -e "column_1=1                                                        "                   \
            -e "column_2=4                                                        "                   \
            ../performance.gp                                                                         \
                                                                                                      \
                                                                                                      \
    && cp    ./$(1)__benchmark_output_parsed.dat    ../../arenaresults                                \
                                                                                                      \
    && echo "'arenaresults/$(1)__operations.png' and 'arenaresults/$(1)__time_operations.png' graphics were generated."
endef