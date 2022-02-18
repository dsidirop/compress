define benchmark =
       cd ./arena/$(1)/                                                                               \
                                                                                                      \
    && go  test  -benchmem  -bench=.   |    tee ./$(1)__benchmark_raw_output.dat                      \
                                                                                                      \
    && awk                                                                                            \
              '/Benchmark_/{count++; gsub(/Benchmark_+.*?_+/, ""); gsub(/[-][0-9]+ /, ""); printf("%d,%s,%s,%s,%s,%s\n", count, $$1, $$2, $$3, $$5, $$7); }'    \
                    ./$(1)__benchmark_raw_output.dat                                                  \
                  > ./$(1)__benchmark_output_parsed.dat                                               \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)__benchmark_output_parsed.dat'                   "                   \
            -e "graphic_file_name='../../arenaresults/$(1)__operations_count.png' "                   \
            -e "y_label='operations#'                                             "                   \
            -e "y_range_min='0000000''                                            "                   \
            -e "y_range_max='1000000'                                             "                   \
            -e "column_1=1                                                        "                   \
            -e "column_2=3                                                        "                   \
            ../performance.gp                                                                         \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)__benchmark_output_parsed.dat'                      "                \
            -e "graphic_file_name='../../arenaresults/$(1)__time_per_operation.png'  "                \
            -e "y_label='nanoseconds / operation'                                    "                \
            -e "y_range_min='00000''                                                 "                \
            -e "y_range_max='45000'                                                  "                \
            -e "column_1=1                                                           "                \
            -e "column_2=4                                                           "                \
            ../performance.gp                                                                         \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)__benchmark_output_parsed.dat'                        "              \
            -e "graphic_file_name='../../arenaresults/$(1)__bytes_per_operations.png'  "              \
            -e "y_label='bytes / operation'                                            "              \
            -e "y_range_min='0000''                                                    "              \
            -e "y_range_max='2000'                                                     "              \
            -e "column_1=1                                                             "              \
            -e "column_2=5                                                             "              \
            ../performance.gp                                                                         \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)__benchmark_output_parsed.dat'                              "        \
            -e "graphic_file_name='../../arenaresults/$(1)__allocations_per_operations.png'  "        \
            -e "y_label='allocations / operation'                                            "        \
            -e "y_range_min='000''                                                           "        \
            -e "y_range_max='40'                                                             "        \
            -e "column_1=1                                                                   "        \
            -e "column_2=6                                                                   "        \
            ../performance.gp                                                                         \
                                                                                                      \
    && cp    ./$(1)__benchmark_output_parsed.dat    ../../arenaresults                                \
                                                                                                      \
    && echo "Plots successfully generated"
endef