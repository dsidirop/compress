#!/usr/bin/env bash

set -o errexit

benchmark_dirname="${1}"
title="${2}"
y_label="${3}"
cpu_count="${4}"
output_files_name_prefix="${benchmark_dirname}-cpu${cpu_count}"

cd  "./arena/${benchmark_dirname}/"

go  clean    -testcache # vital
go  test     -bench=.   -cpu=${cpu_count}  |    tee   "./${output_files_name_prefix}---benchmark-raw-output.dat"

if [[ ${PIPESTATUS[0]} -gt 0 ]]; then
    exit 1
fi



awk \
'/[*][*]/{count++; printf("%d,%s,%s\n", count, $2, $3); }'        \
"./${output_files_name_prefix}---benchmark-raw-output.dat"        \
> "./${output_files_name_prefix}---benchmark-output-parsed.dat"

# messageSizeMax=`                    awk -F','                         'BEGIN{a=0}{ if ($3>0+a) a=$3} END{print a}'     "./${output_files_name_prefix}---benchmark-output-parsed.dat"    `
# messageSizeMaxRoundedUpwards=`      awk -v n="${messageSizeMax}"      'BEGIN{ print int((n+100) / 100 ) * 100 }'                                                 `

tempDir=`mktemp -d -t golang-compression-libs-arena.XXXX`
tempPlotConfigFile="${tempDir}/plot.gp"

title_lowercased_with_dashes=${title// /-}
title_lowercased_with_dashes=${title_lowercased_with_dashes,,}

cp            '../plot.gp'                                                                     "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/${title}\\\\n[${output_files_name_prefix} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
gnuplot \
-e "    file_path='./${output_files_name_prefix}---benchmark-output-parsed.dat'                                                 "    \
-e "    graphic_file_name='../../arena-results/${output_files_name_prefix}--001-${title_lowercased_with_dashes}--result.png'    "    \
-e "    y_label='${y_label}'                                                                                                    "    \
-e "    column_1=1                                                                                                              "    \
-e "    column_2=3                                                                                                              "    \
"${tempPlotConfigFile}"


montage                                                                        \
            -mode concatenate                                                  \
            "../../arena-results/${output_files_name_prefix}--*--result.png"   \
            "../../arena-results/${output_files_name_prefix}----category-overall-results.png"

cp       "./${output_files_name_prefix}---benchmark-output-parsed.dat"   '../../arena-results'

echo     'Plots successfully generated'

rm -rf "${tempDir}"
