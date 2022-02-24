.PHONY:\
clean
clean:
	@rm -rf ./arena-results/*
	@rm -f ./compile-*
	@find ./arena -name "*.dat" -type f -delete
