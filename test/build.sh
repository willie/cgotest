rm -f a.out cgotest.a cgotest.h

go version

# build the library
CGO_ENABLED=1 go build -buildmode=c-archive -mod=vendor ../cgotest.go ../cgotest_bridge.go

# test the library
gcc main.c cgotest.a -framework CoreFoundation -framework Security -lpthread
./a.out
