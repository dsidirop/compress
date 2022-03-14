##
# gnuplot script to generate a performance graphic - it expects the following parameters:
#
# file_path - path to the file from which the data will be read
# graphic_file_name - the graphic file name to be saved 
# y_label - the desired label for y axis
# y_range_min - minimum range for values in y axis
# y_range_max - maximum range for values in y axis
# column_1 - the first column to be used in plot command
# column_2 - the second column to be used in plot command
##

# graphic will be saved as 800x600 png image file
set terminal png size 12000,800

# allows grid lines to be drawn on the plot
set grid x,y

# setting the graphic file name to be saved
set output graphic_file_name

# the graphic's main title
set title "___TITLE___"

# since the input file is a CSV file, we need to tell gnuplot that data fields are separated by comma
# set datafile separator ","
set datafile separator comma

# disable key box
set key off

# label for y axis
set ylabel y_label

# range for values in x/y axes
set xrange[0.1:245]
set yrange[0:]

# to avoid displaying large numbers in exponential format
set format y "%.0f"

# vertical label for x values 
set xtics rotate by 65 right

# set boxplots
set style       fill    solid       0.1
set boxwidth    0.2     relative

# plot graphic for each line of input file
plot file_path                                                                      \
        using column_1:column_2:($0+1):xtic(2)      with    boxes   lc      var,    \
    ''  using column_1:column_2:column_2            with    labels  offset  0,0.5
