release:
	go build  -o sbin/seeder
	rm -Rf sbin/log4go.xml
	rm -Rf sbin/seeder.yaml
	cp log4go.xml sbin/
	cp seeder.yaml sbin/

d:
	go build -gcflags "-N -l"  -race -o sbin/seeder  &&   ./sbin/seeder -start 2>&1 | tee /tmp/seeder.log
#	go build -gcflags "-N -l"  -o sbin/seeder  &&   ./sbin/seeder -start 2>&1 | tee /tmp/seeder.log
seeder:
	go build -gcflags "-N -l"  -o sbin/seeder


build:
	cd thrift && $(MAKE) build
