#!/usr/bin/env bash

benchmark_dirname="${1}"
title="${2}"
y_label="${3}"

cd  "./arena/${benchmark_dirname}/"

go  test   -bench=.   |    tee   "./${benchmark_dirname}---benchmark-raw-output.dat"

if [[ ${PIPESTATUS[0]} -gt 0 ]]; then
    exit 1
fi

awk \
'/[*][*]/{count++; printf("%d,%s,%s\n", count, $2, $3); }' \
"./${benchmark_dirname}---benchmark-raw-output.dat"        \
> "./${benchmark_dirname}---benchmark-output-parsed.dat"

# messageSizeMax=`                    awk -F','                         'BEGIN{a=0}{ if ($3>0+a) a=$3} END{print a}'     "./${benchmark_dirname}---benchmark-output-parsed.dat"    `
# messageSizeMaxRoundedUpwards=`      awk -v n="${messageSizeMax}"      'BEGIN{ print int((n+100) / 100 ) * 100 }'                                                 `

tempDir=`mktemp -d -t golang-compression-libs-arena.XXXX`
tempPlotConfigFile="${tempDir}/plot.gp"

cp            '../plot.gp'                                        "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${benchmark_dirname}] ${title}/g"   "${tempPlotConfigFile}"
gnuplot \
-e "    file_path='./${benchmark_dirname}---benchmark-output-parsed.dat'                                             "    \
-e "    graphic_file_name='../../arena-results/${benchmark_dirname}--eventual-message-size-footprint--result.png'    "    \
-e "    y_label='${y_label}'                                                                                         "    \
-e "    column_1=1                                                                                                   "    \
-e "    column_2=3                                                                                                   "    \
"${tempPlotConfigFile}"

cp       "./${benchmark_dirname}---benchmark-output-parsed.dat"   '../../arena-results'

echo     'Plots successfully generated';

rm -rf "${tempDir}"
