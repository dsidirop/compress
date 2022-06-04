#!/usr/bin/env bash

set -o errexit

benchmark_dirname="${1}"
cpu_count="${2}"
gnuplot_config_file="${3:-../plot.gp}"
resulting_image_merge_mode="${4}"

output_files_name_prefix="${benchmark_dirname}-cpu${cpu_count}"

cd "./arena/${benchmark_dirname}/"

go  clean    -testcache  # vital
go  test     -benchmem     -cpu=${cpu_count}   -bench=.   |    tee "./${output_files_name_prefix}---benchmark-raw-output.dat"

if [[ ${PIPESTATUS[0]} -gt 0 ]]; then
        exit 1
fi




awk                                                                                                                                            \
    '/Benchmark_/{count++; gsub(/Benchmark_+.*?_+/, ""); gsub(/[-][0-9]+ /, ""); printf("%d,%s,%s,%s,%s,%s\n", count, $1, $2, $3, $5, $7); }'  \
    "./${output_files_name_prefix}---benchmark-raw-output.dat"                                                                                 \
  > "./${output_files_name_prefix}---benchmark-output-parsed.dat"

tempDir=`mktemp -d -t golang-compression-libs-arena.XXXX`
tempPlotConfigFile="${tempDir}/plot.gp"

cp            "${gnuplot_config_file}"                                                                                               "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/Avg Elapsed Time (ns) per Operation - Lower is better\\\\n[${benchmark_dirname} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
cat           "./${output_files_name_prefix}---benchmark-output-parsed.dat"          \
  |           sort   -t','    -nk4                                                   \
  |           awk    -F','    '//{count++; printf("%d,%s,%s\n", count, $2, $4); }'   \
  |           awk    -F','    'BEGIN { minvalue=0 } { if (minvalue==0) { minvalue=$3 }; printf("%s,%s,%s,%.0f%%\n", $1, $2, $3, (100*$3/minvalue)); }'   \
  >           "./${output_files_name_prefix}---benchmark-output-parsed---sorted-by-time-per-operation.dat"


gnuplot \
  -e "file_path='./${output_files_name_prefix}---benchmark-output-parsed---sorted-by-time-per-operation.dat'            "  \
  -e "graphic_file_name='../../arena-results/${output_files_name_prefix}--001-time-per-operation--result.png'           "  \
  -e "y_label='nanoseconds / operation'                                                                                 "  \
  -e "column_1=1                                                                                                        "  \
  -e "column_2=3                                                                                                        "  \
  "${tempPlotConfigFile}"


cp            "${gnuplot_config_file}"                                                                           "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/CPU Operations# - Lower is better\\\\n[${benchmark_dirname} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
cat           "./${output_files_name_prefix}---benchmark-output-parsed.dat"          \
  |           sort   -t','    -nk3                                                   \
  |           awk    -F','    '//{count++; printf("%d,%s,%s\n", count, $2, $3); }'   \
  |           awk    -F','    'BEGIN { minvalue=0 } { if (minvalue==0) { minvalue=$3 }; printf("%s,%s,%s,%.0f%%\n", $1, $2, $3, (100*$3/minvalue)); }'   \
  >           "./${output_files_name_prefix}---benchmark-output-parsed---sorted-by-cpu-operations-count.dat"
gnuplot \
  -e "file_path='./${output_files_name_prefix}---benchmark-output-parsed---sorted-by-cpu-operations-count.dat'          "  \
  -e "graphic_file_name='../../arena-results/${output_files_name_prefix}--002-cpu-operations-count--result.png'         "  \
  -e "y_label='cpu-ops#'                                                                                                "  \
  -e "column_1=1                                                                                                        "  \
  -e "column_2=3                                                                                                        "  \
  "${tempPlotConfigFile}"

cp            "${gnuplot_config_file}"                                                                                   "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/RAM Bytes per Operation - Lower is better\\\\n[${benchmark_dirname} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
cat           "./${output_files_name_prefix}---benchmark-output-parsed.dat"          \
  |           sort   -t','    -nk5                                                   \
  |           awk    -F','    '//{count++; printf("%d,%s,%s\n", count, $2, $5); }'   \
  |           awk    -F','    'BEGIN { minvalue=0 } { if (minvalue==0) { minvalue=$3 }; printf("%s,%s,%s,%.0f%%\n", $1, $2, $3, (100*$3/minvalue)); }'   \
  >           "./${output_files_name_prefix}---benchmark-output-parsed---sorted-by-ram-bytes-per-operation.dat"
gnuplot \
  -e "file_path='./${output_files_name_prefix}---benchmark-output-parsed---sorted-by-ram-bytes-per-operation.dat'    "  \
  -e "graphic_file_name='../../arena-results/${output_files_name_prefix}--003-ram-bytes-per-operation--result.png'   "  \
  -e "y_label='ram-bytes / operation'                                                                                "  \
  -e "column_1=1                                                                                                     "  \
  -e "column_2=3                                                                                                     "  \
  "${tempPlotConfigFile}"

cp            "${gnuplot_config_file}"                                                                                     "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/Allocations per Operation - Lower is better\\\\n[${benchmark_dirname} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
cat           "./${output_files_name_prefix}---benchmark-output-parsed.dat"          \
  |           sort   -t','    -nk6                                                   \
  |           awk    -F','    '//{count++; printf("%d,%s,%s\n", count, $2, $6); }'   \
  |           awk    -F','    'BEGIN { minvalue=0 } { if (minvalue==0) { minvalue=$3 }; printf("%s,%s,%s,%.0f%%\n", $1, $2, $3, (100*$3/minvalue)); }'   \
  >           "./${output_files_name_prefix}---benchmark-output-parsed---sorted-by-allocations-per-operation.dat"
gnuplot \
  -e "file_path='./${output_files_name_prefix}---benchmark-output-parsed---sorted-by-allocations-per-operation.dat'  "  \
  -e "graphic_file_name='../../arena-results/${output_files_name_prefix}--004-allocations-per-operation--result.png' "  \
  -e "y_label='allocations / operation'                                                                              "  \
  -e "column_1=1                                                                                                     "  \
  -e "column_2=3                                                                                                     "  \
  "${tempPlotConfigFile}"


overall_results_fname="${output_files_name_prefix}----category-overall-results.png"
overall_results_filepath=`realpath  "../../arena-results/${overall_results_fname}"`

if [[ "${resulting_image_merge_mode}" == "vertical" ]]; then
  convert                                                                      \
            -append                                                            \
            "../../arena-results/${output_files_name_prefix}--*--result.png"   \
            "${overall_results_filepath}"

else # horizontal
  convert                                                                      \
            +append                                                            \
            "../../arena-results/${output_files_name_prefix}--*--result.png"   \
            "${overall_results_filepath}"

fi


cp                                                                                    \
            ./${output_files_name_prefix}---benchmark-output-parsed---sorted-*.dat    \
            "../../arena-results"


echo        "Plot '${overall_results_fname}' successfully generated"

rm -rf "${tempDir}"
