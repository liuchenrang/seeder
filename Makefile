build:
	cd thrift && $(MAKE) build

d:
	go build -gcflags "-N -l"  -race -o sbin/seeder  &&   ./sbin/seeder -start 2>&1 | tee /tmp/seeder.log
#	go build -gcflags "-N -l"  -o sbin/seeder  &&   ./sbin/seeder -start 2>&1 | tee /tmp/seeder.log
seeder:
	go build -gcflags "-N -l"  -o sbin/seeder
release:
	go build  -o sbin/seeder
