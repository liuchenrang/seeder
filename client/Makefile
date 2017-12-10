clt:
	/usr/local/go/bin/go test -c -i -o client_test seeder/client
	./client_test -test.v -test.run ^TestNewClient$
bench:
	rm -Rf ./client_test
	rm /tmp/iduts.log
	go test  -bench ^BenchmarkLoopsTest5 |tee -a /tmp/idtest5.log &
	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log
	echo total `cat /tmp/iduts.log|grep -v ping|grep id |awk '{print $2}'|wc -l`
#	cat /tmp/iduts.log|grep -v ping|grep id |awk '{print $2}'| sort -k 1 |uniq -c| sort -k 1 -n -r|less


