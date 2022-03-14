#!/usr/bin/env bash

imgfpath___scatter_plot_data='./arena-results/x-size-vs-time-scatterplot.png'

csvfpath___scatter_plot_data='./arena-results/x-final-scatterplot.csv'
csvfpath___serialization_with_compression_eventual_message_size='./arena/be-serialization-with-compression-eventual-message-size/be-serialization-with-compression-eventual-message-size-cpu1---benchmark-output-parsed---sorted.dat'
csvfpath___serialization_deserialization_with_compression_elapsed_time='./arena/bd-serialization-deserialization-with-compression-elapsed-time/bd-serialization-deserialization-with-compression-elapsed-time-cpu1---benchmark-output-parsed---sorted.dat'


# notice the formula we use on the last column   totalscore="10 * $4 + 1 * $7"   which is essentially   totalscore="(10 * size-score%) + (1 * speed-score%)"
# this means that we consider compression to be an order of magnitude more important than speed
echo "Combo,Total Size (in Bytes),Size-Score%,Total Time for Compression+Decompression (in ns),Time-Score%,Total Score" > "${csvfpath___scatter_plot_data}"
join                                                                                                             \
       -t,                                                                                                       \
       -12                                                                                                       \
       -22                                                                                                       \
       <(sort     -k2    -t,    "${csvfpath___serialization_with_compression_eventual_message_size}"         )   \
       <(sort     -k2    -t,    "${csvfpath___serialization_deserialization_with_compression_elapsed_time}"  )   \
       | awk      -F','         '{ printf("%s,%s,%s,%s,%s,%s\n", $1, $3, $4, $6, $7, (10 * $4) + (1 * $7)) }'    \
       | sort     -t','         -nk6                                                                             \
       >> "${csvfpath___scatter_plot_data}"

echo "** Successfully generated data-file '${csvfpath___scatter_plot_data}'!"



gnuplot                                                                                   \
           -e "          file_path='${csvfpath___scatter_plot_data}'                 "    \
           -e "  graphic_file_name='${imgfpath___scatter_plot_data}'                 "    \
           "./arena/plot.size-vs-time.gp"

echo "** Successfully generated scatterplot image '${imgfpath___scatter_plot_data}'!"
