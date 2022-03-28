.PHONY:\
clean
clean:
	@rm -rf ./arena-results/*
	@rm -f ./compile-*
	@find ./arena -name "*.dat" -o -name "*.csv" -o -name "*.png" -type f  -exec  rm {} \;
