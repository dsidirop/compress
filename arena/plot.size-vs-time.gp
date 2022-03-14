set terminal png size 10000,10000

# allows grid lines to be drawn on the plot
set grid x,y

# setting the graphic file name to be saved
set output graphic_file_name

# the graphic's main title
set title "Overall Size in Bytes   vS   Overall Time for Compression+Decompression"

# since the input file is a CSV file, we need to tell gnuplot that data fields are separated by comma
set datafile separator comma

# disable key box
set key off

# offset
set offset 1,1,1,1

# label for x axis
set xlabel "Size in Bytes"

# label for y axis
set ylabel "Overall time for Compression+Decompression"


plot file_path                                                             \
                using       2:4:(stringcolumn(1))                          \
                with        labels                                         \
                            point                                          \
                            pointtype   7                                  \
                            pointsize   2                                  \
                            offset char 0.1,0.1                            \
                            notitle

