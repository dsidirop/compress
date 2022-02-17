benchmark:
	@  go test -bench=.        | tee ./arenaresults/out.dat                                           \
                                                                                                      \
	&& awk                                                                                            \
	              '/Benchmark/{count++; gsub(/BenchmarkTest/,""); printf("%d,%s,%s,%s\n",count,$$1,$$2,$$3); }'    \
				    ./arenaresults/out.dat                                                            \
				  > ./arenaresults/final.dat                                                          \
                                                                                                      \
	&& gnuplot                                                                                        \
            -e "file_path='./arenaresults/final.dat'                   "                              \
            -e "graphic_file_name='./arenaresults/operations.png'      "                              \
            -e "y_label='number of operations'                         "                              \
            -e "y_range_min='000000000''                               "                              \
            -e "y_range_max='400000000'                                "                              \
            -e "column_1=1                                             "                              \
            -e "column_2=3                                             "                              \
            ./arenaresults/performance.gp                                                             \
                                                                                                      \
	&& gnuplot                                                                                        \
            -e "file_path='./arenaresults/final.dat'                     "                            \
            -e "graphic_file_name='./arenaresults/time_operations.png'   "                            \
            -e "y_label='each operation in nanoseconds'                  "                            \
            -e "y_range_min='000''                                       "                            \
            -e "y_range_max='45000'                                      "                            \
            -e "column_1=1                                               "                            \
            -e "column_2=4                                               "                            \
            ./arenaresults/performance.gp                                                             \
                                                                                                      \
	&& rm                                                                                             \
            -f ./arenaresults/out.dat                                                                 \
			   ./arenaresults/final.dat                                                               \
                                                                                                      \
    && echo "'graphic/operations.png' and 'graphic/time_operations.png' graphics were generated."
