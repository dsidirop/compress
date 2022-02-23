#!/usr/bin/env bash

cd "./arena/${1}/"

go  test  -benchmem  -cpu=1   -bench=.   |    tee "./${1}---benchmark-raw-output.dat"

if [[ ${PIPESTATUS[0]} -gt 0 ]]; then
        exit 1
fi

awk \
'/Benchmark_/{count++; gsub(/Benchmark_+.*?_+/, ""); gsub(/[-][0-9]+ /, ""); printf("%d,%s,%s,%s,%s,%s\n", count, $1, $2, $3, $5, $7); }' \
"./${1}---benchmark-raw-output.dat" \
> "./${1}---benchmark-output-parsed.dat"

# operationsMax=`                    awk -F','                        'BEGIN{a=0}{ if ($3>0+a) a=$3} END{print a}'     "./${1}---benchmark-output-parsed.dat"    `
# nanosecondsMax=`                   awk -F','                        'BEGIN{a=0}{ if ($4>0+a) a=$4} END{print a}'     "./${1}---benchmark-output-parsed.dat"    `
# ramBytesMax=`                      awk -F','                        'BEGIN{a=0}{ if ($5>0+a) a=$5} END{print a}'     "./${1}---benchmark-output-parsed.dat"    `
# allocationsMax=`                   awk -F','                        'BEGIN{a=0}{ if ($6>0+a) a=$6} END{print a}'     "./${1}---benchmark-output-parsed.dat"    `

# operationsMaxRoundedUpwards=`      awk -v n="${operationsMax}"      'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `
# nanosecondsMaxRoundedUpwards=`     awk -v n="${nanosecondsMax}"     'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `
# ramBytesMaxRoundedUpwards=`        awk -v n="${ramBytesMax}"        'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `
# allocationsMaxRoundedUpwards=`     awk -v n="${allocationsMax}"     'BEGIN{ print int((n+100) / 100 ) * 100 }'                            `

tempDir=`mktemp -d -t golang-compression-libs-arena.XXXX`
tempPlotConfigFile="${tempDir}/plot.gp"

cp            '../plot.gp'                                                     "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${1}] CPU Operations#\\\\n(lower is better)/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${1}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${1}--cpu-operations-count--result.png'          "  \
-e "y_label='cpu-ops#'                                                                      "  \
-e "column_1=1                                                                              "  \
-e "column_2=3                                                                              "  \
"${tempPlotConfigFile}"

cp            '../plot.gp'                                                             "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${1}] Time (ns) per Operation\\\\n(lower is better)/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${1}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${1}--time-per-operation--result.png'            "  \
-e "y_label='nanoseconds / operation'                                                       "  \
-e "column_1=1                                                                              "  \
-e "column_2=4                                                                              "  \
"${tempPlotConfigFile}"

cp            '../plot.gp'                                                             "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${1}] RAM Bytes per Operation\\\\n(lower is better)/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${1}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${1}--ram-bytes-per-operation--result.png'       "  \
-e "y_label='ram-bytes / operation'                                                         "  \
-e "column_1=1                                                                              "  \
-e "column_2=5                                                                              "  \
"${tempPlotConfigFile}"

cp            '../plot.gp'                                                               "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${1}] Allocations per Operation\\\\n(lower is better)/g"   "${tempPlotConfigFile}"
gnuplot \
-e "file_path='./${1}---benchmark-output-parsed.dat'                                        "  \
-e "graphic_file_name='../../arena-results/${1}--allocations-per-operation--result.png'     "  \
-e "y_label='allocations / operation'                                                       "  \
-e "column_1=1                                                                              "  \
-e "column_2=6                                                                              "  \
"${tempPlotConfigFile}"

cp   "./${1}---benchmark-output-parsed.dat"  '../../arena-results'

echo "Plots successfully generated";

rm -rf "${tempDir}"
