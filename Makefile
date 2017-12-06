build:
	cd thrift && $(MAKE) build
run:
	go build -o sbin/seeder &&  ./sbin/seeder
debug:
	go build -gcflags "-N -l" -o sbin/seeder && sudo gdb ./sbin/seeder
clt:
	cd ./client
	/usr/local/go/bin/go test -c -i -o client_test seeder/client 
	./client_test -test.v -test.run ^TestNewClient$
bench:
	cd ./client
	/usr/local/go/bin/go test -c -i -o client_test seeder/client 
	./client_test -bench=.  ^BenchmarkId$
