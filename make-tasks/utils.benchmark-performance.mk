define benchmark-performance =
       cd ./arena/$(1)/                                                                               \
                                                                                                      \
    && go  test  -benchmem  -bench=.   |    tee ./$(1)--benchmark-raw-output.dat                      \
                                                                                                      \
    && awk                                                                                            \
              '/Benchmark_/{count++; gsub(/Benchmark_+.*?_+/, ""); gsub(/[-][0-9]+ /, ""); printf("%d,%s,%s,%s,%s,%s\n", count, $$1, $$2, $$3, $$5, $$7); }'    \
                    "./$(1)--benchmark-raw-output.dat"                                                \
                  > "./$(1)--benchmark-output-parsed.dat"                                             \
                                                                                                      \
    && operationsMax=`                    awk -F','  'BEGIN{a=0}{ if ($$3>0+a) a=$$3} END{print a}'     "./$(1)--benchmark-output-parsed.dat"     `  \
    && nanosecondsMax=`                   awk -F','  'BEGIN{a=0}{ if ($$4>0+a) a=$$4} END{print a}'     "./$(1)--benchmark-output-parsed.dat"     `  \
    && bytesMax=`                         awk -F','  'BEGIN{a=0}{ if ($$5>0+a) a=$$5} END{print a}'     "./$(1)--benchmark-output-parsed.dat"     `  \
    && allocationsMax=`                   awk -F','  'BEGIN{a=0}{ if ($$6>0+a) a=$$6} END{print a}'     "./$(1)--benchmark-output-parsed.dat"     `  \
                                                                                                                                                     \
    && operationsMaxRoundedUpwards=`      awk -v n="$${operationsMax}"      'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `  \
    && nanosecondsMaxRoundedUpwards=`     awk -v n="$${nanosecondsMax}"     'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `  \
    && bytesMaxRoundedUpwards=`           awk -v n="$${bytesMax}"           'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `  \
    && allocationsMaxRoundedUpwards=`     awk -v n="$${allocationsMax}"     'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `  \
                                                                                                                                                     \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)--benchmark-output-parsed.dat'                    "                  \
            -e "graphic_file_name='../../arena-results/$(1)--operations-count.png' "                  \
            -e "y_label='operations#'                                              "                  \
            -e "y_range_min='0000000''                                             "                  \
            -e "y_range_max='$${operationsMaxRoundedUpwards}'                      "                  \
            -e "column_1=1                                                         "                  \
            -e "column_2=3                                                         "                  \
            ../plot.gp                                                                                \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)--benchmark-output-parsed.dat'                      "                \
            -e "graphic_file_name='../../arena-results/$(1)--time-per-operation.png' "                \
            -e "y_label='nanoseconds / operation'                                    "                \
            -e "y_range_min='00000''                                                 "                \
            -e "y_range_max='$${nanosecondsMaxRoundedUpwards}'                       "                \
            -e "column_1=1                                                           "                \
            -e "column_2=4                                                           "                \
            ../plot.gp                                                                                \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)--benchmark-output-parsed.dat'                        "              \
            -e "graphic_file_name='../../arena-results/$(1)--bytes-per-operations.png' "              \
            -e "y_label='bytes / operation'                                            "              \
            -e "y_range_min='0000''                                                    "              \
            -e "y_range_max='$${bytesMaxRoundedUpwards}'                               "              \
            -e "column_1=1                                                             "              \
            -e "column_2=5                                                             "              \
            ../plot.gp                                                                                \
                                                                                                      \
    && gnuplot                                                                                        \
            -e "file_path='./$(1)--benchmark-output-parsed.dat'                              "        \
            -e "graphic_file_name='../../arena-results/$(1)--allocations-per-operations.png' "        \
            -e "y_label='allocations / operation'                                            "        \
            -e "y_range_min='000''                                                           "        \
            -e "y_range_max='$${allocationsMaxRoundedUpwards}'                               "        \
            -e "column_1=1                                                                   "        \
            -e "column_2=6                                                                   "        \
            ../plot.gp                                                                                \
                                                                                                      \
    && cp    ./$(1)--benchmark-output-parsed.dat    ../../arena-results                               \
                                                                                                      \
    && echo "Plots successfully generated"
endef