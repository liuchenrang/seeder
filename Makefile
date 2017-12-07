build:
	cd thrift && $(MAKE) build
run:
	go build -o sbin/seeder &&  ./sbin/seeder
debug:
	go build -gcflags "-N -l" -o sbin/seeder && sudo gdb ./sbin/seeder
