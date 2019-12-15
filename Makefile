.PHONY: setup gen clean

clean:
	rm -rf gen/*

setup:
	go install github.com/go-swagger/go-swagger/cmd/swagger
	go install github.com/golang/mock/mockgen

gen: clean
	swagger generate server -t gen -f ./swagger/api.yml --exclude-main -A property -P auth.AuthUser
