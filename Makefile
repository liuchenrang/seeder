release:
	go build  -o sbin/seeder
	rm -Rf sbin/log4go.xml
	rm -Rf sbin/seeder.yaml
	cp log4go.xml sbin/
	cp seeder.yaml sbin/

test:
	go build -gcflags "-N -l"  -o sbin/seeder  &&   ./sbin/seeder  2>&1 | tee /tmp/seeder.log

race:
	go build -gcflags "-N -l"  -race -o sbin/seeder  &&   ./sbin/seeder  2>&1 | tee /tmp/seeder.log

run:
	go build -o sbin/seeder  &&   ./sbin/seeder  2>&1 | tee /tmp/seeder.log

build:
	cd thrift && $(MAKE) build
