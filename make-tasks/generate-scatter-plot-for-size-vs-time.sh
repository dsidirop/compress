#!/usr/bin/env bash

set -eu

cpu_count="${1}"
top_N_results="21"

imgfpath___scatter_plot_full='./arena-results/x-size-vs-time-scatterplot-full.png'
imgfpath___scatter_plot_topdogs='./arena-results/x-size-vs-time-scatterplot-topdogs.png'
imgfpath___scatter_plot_cherrypicked='./arena-results/x-size-vs-time-scatterplot-cherrypicked.png'

csvfpath___scatter_plot_data='./arena-results/x-final-scatterplot.csv'
csvfpath___scatter_plot_data_topdogs='./arena-results/x-final-scatterplot-topdogs.csv'
csvfpath___scatter_plot_data_cherrypicked='./arena-results/x-final-scatterplot-cherrypicked.csv'

csvfpath___serialization_with_compression_eventual_message_size="./arena-results/be-serialization-with-compression-eventual-message-size-cpu${cpu_count}---benchmark-output-parsed---sorted.csv"
csvfpath___serialization_deserialization_with_compression_elapsed_time="./arena-results/bd-serialization-deserialization-with-compression-elapsed-time-cpu${cpu_count}---benchmark-output-parsed---sorted.csv"

mkdir -p ./arena-results/


# notice the formula we use on the last column   totalscore="10 * $4 + 1 * $7"   which is essentially   totalscore="(3 * size-score%)^2 + speed-score%^2"
# this means that we consider compression to be an order of magnitude more important than speed
echo "Combo,Total Size (in Bytes | Lower is Better),Size-Score% (Lower is Better),Size% Compared to JSON-Uncompressed (Lower is Better),Total Time for Compression+Decompression (in ns | Lower is Better),Time-Score% (Lower is Better),Time% Compared to JSON-Uncompressed (Lower is Better),Total Score ( ((10 * size-score%)^2 + speed-score%^2)/10000 | Lower is Better)" > "${csvfpath___scatter_plot_data}"
join                                                                                                                                           \
       -t,                                                                                                                                     \
       -12                                                                                                                                     \
       -22                                                                                                                                     \
       <(sort     -k2    -t,    "${csvfpath___serialization_with_compression_eventual_message_size}"         )                                 \
       <(sort     -k2    -t,    "${csvfpath___serialization_deserialization_with_compression_elapsed_time}"  )                                 \
       | awk      -F','         '{ printf("%s,%s,%s,%s,%s,%s,%s,%s\n", $1, $3, $4, $5, $7, $8, $9, ((100 * $4 * $4) + ($8 * $8)) / 10000 ) }'  \
       | sort     -t','         -nk8                                                                                                           \
       >> "${csvfpath___scatter_plot_data}"

awk   "NR>=0 && NR<=${top_N_results}"   "${csvfpath___scatter_plot_data}"   >   "${csvfpath___scatter_plot_data_topdogs}"

cat "${csvfpath___scatter_plot_data}"                       \
| grep -Ei 'JSON|CBOR|ProtoBuf'                             \
| grep -Ei 'S2|Zlib|ZStandard|Uncompressed'                 \
| grep -Ei 'DefaultCompression|BestSpeed|Uncompressed'      \
> "${csvfpath___scatter_plot_data_cherrypicked}"

gnuplot                                                                                      \
           -e "          file_path='${csvfpath___scatter_plot_data}'                    "    \
           -e "  graphic_file_name='${imgfpath___scatter_plot_full}'                    "    \
           "./arena/plot.size-vs-time-full.gp"

gnuplot                                                                                      \
           -e "          file_path='${csvfpath___scatter_plot_data_topdogs}'            "    \
           -e "  graphic_file_name='${imgfpath___scatter_plot_topdogs}'                 "    \
           "./arena/plot.size-vs-time-topdogs.gp"

gnuplot                                                                                      \
           -e "          file_path='${csvfpath___scatter_plot_data_cherrypicked}'       "    \
           -e "  graphic_file_name='${imgfpath___scatter_plot_cherrypicked}'            "    \
           "./arena/plot.size-vs-time-topdogs.gp"


echo "** Successfully generated data-file '${csvfpath___scatter_plot_data}'!"
echo "** Successfully generated data-file '${csvfpath___scatter_plot_data_topdogs}'!"
echo "** Successfully generated data-file '${csvfpath___scatter_plot_data_cherrypicked}'!"
echo "** Successfully generated scatterplots '${imgfpath___scatter_plot_full}'!"
echo "** Successfully generated scatterplots '${imgfpath___scatter_plot_topdogs}'!"
echo "** Successfully generated scatterplots '${imgfpath___scatter_plot_cherrypicked}'!"
