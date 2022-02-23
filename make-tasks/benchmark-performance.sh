#!/usr/bin/env bash

benchmark_dirname="${1}"
cpu_count="${2}"
output_files_name_prefix="${benchmark_dirname}-cpu${cpu_count}"

cd "./arena/${benchmark_dirname}/"

go  test  -benchmem  -cpu=${cpu_count}   -bench=.   |    tee "./${output_files_name_prefix}---benchmark-raw-output.dat"

if [[ ${PIPESTATUS[0]} -gt 0 ]]; then
        exit 1
fi




awk                                                                                                                                        \
'/Benchmark_/{count++; gsub(/Benchmark_+.*?_+/, ""); gsub(/[-][0-9]+ /, ""); printf("%d,%s,%s,%s,%s,%s\n", count, $1, $2, $3, $5, $7); }'  \
"./${output_files_name_prefix}---benchmark-raw-output.dat"                                                                                 \
> "./${output_files_name_prefix}---benchmark-output-parsed.dat"

# operationsMax=`                    awk -F','                        'BEGIN{a=0}{ if ($3>0+a) a=$3} END{print a}'     "./${benchmark_dirname}---benchmark-output-parsed.dat"    `
# nanosecondsMax=`                   awk -F','                        'BEGIN{a=0}{ if ($4>0+a) a=$4} END{print a}'     "./${benchmark_dirname}---benchmark-output-parsed.dat"    `
# ramBytesMax=`                      awk -F','                        'BEGIN{a=0}{ if ($5>0+a) a=$5} END{print a}'     "./${benchmark_dirname}---benchmark-output-parsed.dat"    `
# allocationsMax=`                   awk -F','                        'BEGIN{a=0}{ if ($6>0+a) a=$6} END{print a}'     "./${benchmark_dirname}---benchmark-output-parsed.dat"    `

# operationsMaxRoundedUpwards=`      awk -v n="${operationsMax}"      'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `
# nanosecondsMaxRoundedUpwards=`     awk -v n="${nanosecondsMax}"     'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `
# ramBytesMaxRoundedUpwards=`        awk -v n="${ramBytesMax}"        'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `
# allocationsMaxRoundedUpwards=`     awk -v n="${allocationsMax}"     'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `

tempDir=`mktemp -d -t golang-compression-libs-arena.XXXX`
tempPlotConfigFile="${tempDir}/plot.gp"


cp            '../plot.gp'                                                                                                           "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/Avg Elapsed Time (ns) per Operation - Lower is better\\\\n[${benchmark_dirname} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${output_files_name_prefix}---benchmark-output-parsed.dat'                                           "  \
-e "graphic_file_name='../../arena-results/${output_files_name_prefix}--001-time-per-operation--result.png'           "  \
-e "y_label='nanoseconds / operation'                                                                                 "  \
-e "column_1=1                                                                                                        "  \
-e "column_2=4                                                                                                        "  \
"${tempPlotConfigFile}"


cp            '../plot.gp'                                                                                       "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/CPU Operations# - Lower is better\\\\n[${benchmark_dirname} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${output_files_name_prefix}---benchmark-output-parsed.dat'                                           "  \
-e "graphic_file_name='../../arena-results/${output_files_name_prefix}--002-cpu-operations-count--result.png'         "  \
-e "y_label='cpu-ops#'                                                                                                "  \
-e "column_1=1                                                                                                        "  \
-e "column_2=3                                                                                                        "  \
"${tempPlotConfigFile}"

cp            '../plot.gp'                                                                                               "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/RAM Bytes per Operation - Lower is better\\\\n[${benchmark_dirname} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${output_files_name_prefix}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${output_files_name_prefix}--003-ram-bytes-per-operation--result.png'   "  \
-e "y_label='ram-bytes / operation'                                                                                "  \
-e "column_1=1                                                                                                     "  \
-e "column_2=5                                                                                                     "  \
"${tempPlotConfigFile}"

cp            '../plot.gp'                                                                                                 "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/Allocations per Operation - Lower is better\\\\n[${benchmark_dirname} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${output_files_name_prefix}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${output_files_name_prefix}--004-allocations-per-operation--result.png' "  \
-e "y_label='allocations / operation'                                                                              "  \
-e "column_1=1                                                                                                     "  \
-e "column_2=6                                                                                                     "  \
"${tempPlotConfigFile}"


montage                                                                        \
            -mode concatenate                                                  \
            "../../arena-results/${output_files_name_prefix}--*--result.png"   \
            "../../arena-results/${output_files_name_prefix}----category-overall-results.png"


cp                                                                           \
            "./${output_files_name_prefix}---benchmark-output-parsed.dat"    \
            '../../arena-results'


echo "Plots successfully generated"

rm -rf "${tempDir}"
