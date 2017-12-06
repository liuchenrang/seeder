clt:

	/usr/local/go/bin/go test -c -i -o client_test seeder/client 
	./client_test -test.v -test.run ^TestNewClient$
bench:

	/usr/local/go/bin/go test -c -i -o client_test seeder/client 
	go test  -bench BenchmarkLoopsParallel|tee -a /tmp/id.log
#	cat /tmp/id.log|grep -v ping|grep id |awk '{print $2}'| sort -k 1 |uniq -c


