#!/usr/bin/env bash

cd  "./arena/${1}/"

go  test   -bench=.   |    tee   "./${1}---benchmark-raw-output.dat"

if [[ ${PIPESTATUS[0]} -gt 0 ]]; then
    exit 1
fi

awk \
'/[*][*]/{count++; printf("%d,%s,%s\n", count, $2, $3); }' \
"./${1}---benchmark-raw-output.dat"                        \
> "./${1}---benchmark-output-parsed.dat"

messageSizeMax=`                    awk -F','                         'BEGIN{a=0}{ if ($3>0+a) a=$3} END{print a}'     "./${1}---benchmark-output-parsed.dat"    `
messageSizeMaxRoundedUpwards=`      awk -v n="${messageSizeMax}"      'BEGIN{ print int((n+100) / 100 ) * 100 }'                                                 `

tempDir=`mktemp -d -t golang-compression-libs-arena.XXXX`
tempPlotConfigFile="${tempDir}/plot.gp"

cp            '../plot.gp'                    "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/[${1}] ${2}/g"   "${tempPlotConfigFile}"
gnuplot \
-e "    file_path='./${1}---benchmark-output-parsed.dat'                                             "    \
-e "    graphic_file_name='../../arena-results/${1}--eventual-message-size-footprint--result.png'    "    \
-e "    y_label='${3}'                                                                               "    \
-e "    y_range_min='0000''                                                                          "    \
-e "    y_range_max='${messageSizeMaxRoundedUpwards}'                                                "    \
-e "    column_1=1                                                                                   "    \
-e "    column_2=3                                                                                   "    \
"${tempPlotConfigFile}"

cp       "./${1}---benchmark-output-parsed.dat"   '../../arena-results'

echo     'Plots successfully generated';

rm -rf "${tempDir}"
