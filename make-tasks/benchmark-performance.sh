#!/usr/bin/env bash

benchmark_dirname="${1}"

cd "./arena/${benchmark_dirname}/"

go  test  -benchmem  -cpu=1   -bench=.   |    tee "./${benchmark_dirname}---benchmark-raw-output.dat"

if [[ ${PIPESTATUS[0]} -gt 0 ]]; then
        exit 1
fi



awk                                                                                                                                        \
'/Benchmark_/{count++; gsub(/Benchmark_+.*?_+/, ""); gsub(/[-][0-9]+ /, ""); printf("%d,%s,%s,%s,%s,%s\n", count, $1, $2, $3, $5, $7); }'  \
"./${benchmark_dirname}---benchmark-raw-output.dat"                                                                                        \
> "./${benchmark_dirname}---benchmark-output-parsed.dat"

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

cp            '../plot.gp'                                                                     "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${benchmark_dirname}] CPU Operations#\\\\n(lower is better)/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${benchmark_dirname}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${benchmark_dirname}--cpu-operations-count--result.png'          "  \
-e "y_label='cpu-ops#'                                                                                      "  \
-e "column_1=1                                                                                              "  \
-e "column_2=3                                                                                              "  \
"${tempPlotConfigFile}"

cp            '../plot.gp'                                                                     "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${benchmark_dirname}] Time (ns) per Operation\\\\n(lower is better)/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${benchmark_dirname}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${benchmark_dirname}--time-per-operation--result.png'            "  \
-e "y_label='nanoseconds / operation'                                                                       "  \
-e "column_1=1                                                                                              "  \
-e "column_2=4                                                                                              "  \
"${tempPlotConfigFile}"

cp            '../plot.gp'                                                                     "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${benchmark_dirname}] RAM Bytes per Operation\\\\n(lower is better)/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${benchmark_dirname}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${benchmark_dirname}--ram-bytes-per-operation--result.png'       "  \
-e "y_label='ram-bytes / operation'                                                                         "  \
-e "column_1=1                                                                                              "  \
-e "column_2=5                                                                                              "  \
"${tempPlotConfigFile}"

cp            '../plot.gp'                                                                       "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${benchmark_dirname}] Allocations per Operation\\\\n(lower is better)/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${benchmark_dirname}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${benchmark_dirname}--allocations-per-operation--result.png'     "  \
-e "y_label='allocations / operation'                                                                       "  \
-e "column_1=1                                                                                              "  \
-e "column_2=6                                                                                              "  \
"${tempPlotConfigFile}"


montage                                                                 \
            -mode concatenate                                           \
            "../../arena-results/${benchmark_dirname}--*--result.png"   \
            "../../arena-results/${benchmark_dirname}---category-overall-results.png"


cp                                                                    \
            "./${benchmark_dirname}---benchmark-output-parsed.dat"    \
            '../../arena-results'


echo "Plots successfully generated"

rm -rf "${tempDir}"
