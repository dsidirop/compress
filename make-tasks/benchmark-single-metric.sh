#!/usr/bin/env bash

set -eu

benchmark_dirname="${1}"
title="${2}"
y_label="${3}"
cpu_count="${4}"
gnuplot_config_file="${5:-../plot.gp}"
output_files_name_prefix="${benchmark_dirname}-cpu${cpu_count}"

mkdir -p ./arena-results/             # order
cd  "./arena/${benchmark_dirname}/"   # order

GOAMD64=${GOAMD64:-v3}  # nice cpu optimization   gives 5%-to-30% boost in some compressors   snappy seems to be unaffected though

go  clean    -testcache # vital
go  test     -bench=.   -cpu=${cpu_count}  |    tee   "./${output_files_name_prefix}---benchmark-raw-output.csv"

if [[ ${PIPESTATUS[0]} -gt 0 ]]; then
    exit 1
fi



awk                                                                         \
     '/[*][*]/{count++; printf("%d,%s,%s\n", count, $2, $3); }'             \
     "./${output_files_name_prefix}---benchmark-raw-output.csv"             \
  >  "./${output_files_name_prefix}---benchmark-output-parsed.csv"

jsonPerformance=`  awk -F','    'BEGIN{ jsonPerformance=0; }   { if ($2 == "JSON" || $2 == "JSON+Uncompressed") { jsonPerformance=$3; } }   END{ print jsonPerformance; }'     "./${output_files_name_prefix}---benchmark-output-parsed.csv"    `

tempDir=`mktemp -d -t golang-compression-libs-arena.XXXX`
tempPlotConfigFile="${tempDir}/plot.gp"

title_lowercased_with_dashes=${title// /-}
title_lowercased_with_dashes=${title_lowercased_with_dashes,,}

cp            "${gnuplot_config_file}"                                                         "${tempPlotConfigFile}"
sed    -i     "s/___TITLE___/${title}\\\\n[${output_files_name_prefix} cpu#=${cpu_count}]/g"   "${tempPlotConfigFile}"
cat           "./${output_files_name_prefix}---benchmark-output-parsed.csv"                                                                                                                                                        \
  |           sort   -t','    -nk3                                                                                                                                                                                                 \
  |           awk    -F','                                                 '//{count++; printf("%d,%s,%s\n", count, $2, $3); }'                                                                                                    \
  |           awk    -F','    -v jsonPerformance="${jsonPerformance}"      'BEGIN { minvalue=0 } { if (minvalue==0) { minvalue=$3 }; printf("%s,%s,%s,%04.0f%%,%04.0f%%\n", $1, $2, $3, (100*$3/minvalue), (100*$3/jsonPerformance)); }'   \
  >           "./${output_files_name_prefix}---benchmark-output-parsed---sorted.csv"

gnuplot \
-e "    file_path='./${output_files_name_prefix}---benchmark-output-parsed---sorted.csv'                                        "    \
-e "    graphic_file_name='../../arena-results/${output_files_name_prefix}--001-${title_lowercased_with_dashes}--result.png'    "    \
-e "    y_label='${y_label}'                                                                                                    "    \
-e "    column_1=1                                                                                                              "    \
-e "    column_2=5                                                                                                              "    \
"${tempPlotConfigFile}"


overall_results_fname="${output_files_name_prefix}----category-overall-results.png"
overall_results_filepath=`realpath  "../../arena-results/${overall_results_fname}"`


montage                                                                        \
            -mode concatenate                                                  \
            "../../arena-results/${output_files_name_prefix}--*--result.png"   \
            "${overall_results_filepath}"


cp                                                                                    \
            "./${output_files_name_prefix}---benchmark-output-parsed---sorted.csv"    \
            "../../arena-results"


echo
echo        "** Csv '${output_files_name_prefix}---benchmark-output-parsed---sorted.csv' successfully generated!"
echo        "** Plot '${overall_results_fname}' successfully generated!"
echo

rm -rf "${tempDir}"
