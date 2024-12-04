SHELL := /bin/zsh

run:
	time go run day${DAY}/go/main.go

all:
	i = 1 ; while [[ $$i -le 4 ]] ; do \
		time go run day$$i/go/main.go ; echo "\n" ; \
		((i = i + 1)) ; \
	done
