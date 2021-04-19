container=$$(docker images --filter "dangling=true" -q)
FILES = $(shell ls)
all:
	echo $(FILES)

.PHONY: clean

clean:
	@docker rmi -f $(container)

test: 
	@go test -coverprofile fmtcoverage.html fmt  