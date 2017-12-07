clt:
	/usr/local/go/bin/go test -c -i -o client_test seeder/client
	./client_test -test.v -test.run ^TestNewClient$
bench:
	rm -Rf ./client_test
	go test  -bench ^BenchmarkLoopsParallel$ |tee /tmp/id.log
	cat /tmp/id.log|grep -v ping|grep id |awk '{print $2}'| sort -k 1 |uniq -c


