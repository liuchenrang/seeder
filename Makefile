build:
	cd thrift && $(MAKE) build
run:
	rm -Rf ./logs/*
	go build -o sbin/seeder &&  ./sbin/seeder -start | tee /tmp/seeder.log
debug:
	go build -gcflags "-N -l" -o sbin/seeder  && sudo gdb ./sbin/seeder -start | tee /tmp/seeder.log
seeder:
	go build -gcflags "-N -l" -o sbin/seeder
