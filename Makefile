# Set to zsh because I use it and because I like the built in time command output
SHELL := /bin/zsh 

run:
	@time go run day${DAY}/go/main.go

all:
	@for name in day*/go/main.go ; do\
		time go run $${name} ; echo "\n" ; \
	done

test:
	cd day${DAY}/go/ && go test -v

test-all:
	@for name in day*/go/ ; do\
		cd $${name} && echo "Test $${name}" && go test ; cd ../../ ; echo "\n" ; \
	done


