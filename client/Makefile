clt:
	/usr/local/go/bin/go test -test.v -test.run ^TestNewClient$$
bench:
	rm -Rf ./client_test
	rm -Rf /tmp/iduts.log
#	go test  -bench ^BenchmarkLoopsTest5 |tee -a /tmp/idtest5.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
	go test -parallel 500  -bench ^BenchmarkLoopsUts$$ -run ^$$ |tee -a /tmp/iduts.log
#	echo total `cat /tmp/iduts.log|grep -v ping|grep id |awk '{print $2}'|wc -l`
	cat /tmp/iduts.log|grep -v ping|grep id |awk '{print $2}'| sort -k 1 |uniq -c| sort -k 1 -n -r|less

benchfix:
	rm -Rf ./client_test
	rm -Rf /tmp/iduts.log
#	go test  -bench ^BenchmarkLoopsTest5 |tee -a /tmp/idtest5.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
	go test   -test.benchtime 5s -test.bench ^BenchmarkLoopsMultiTag$$ -run ^$$ |tee  /tmp/iduts.log &
#	echo total `cat /tmp/iduts.log|grep -v ping|grep id |awk '{print $2}'|wc -l`


benchmusnow:
	rm -Rf ./client_test
	rm -Rf /tmp/iduts.log
	go test   -test.benchtime 5s -test.bench ^BenchmarkSnowMultiTag$$ -run ^$$ |tee  /tmp/iduts.log
	cat /tmp/iduts.log|grep -v ping|grep MultiTag |awk '{print $3}'| sort -k 1 |uniq -c| sort -k 1 -n -r|less
benchsnow:
	rm -Rf ./client_test
	rm -Rf /tmp/iduts.log
	go test   -test.benchtime 5s -test.bench ^BenchmarkLoopsSnow$$ -run ^$$ |tee  /tmp/iduts.log
	cat /tmp/iduts.log|grep -v ping|grep id |awk '{print $2}'| sort -k 1 |uniq -c| sort -k 1 -n -r|less
benin:
	rm -Rf ./client_test
	rm -Rf /tmp/iduts.log
#	go test  -bench ^BenchmarkLoopsTest5 |tee -a /tmp/idtest5.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
#	go test  -bench ^BenchmarkLoopsUts$ |tee -a /tmp/iduts.log &
	go test -parallel 100   -bench ^BenchmarkInThrift$$ |tee -a /tmp/iduts.log &
#	echo total `cat /tmp/iduts.log|grep -v ping|grep id |awk '{print $2}'|wc -l`
#	cat /tmp/iduts.log|grep -v ping|grep id |awk '{print $2}'| sort -k 1 |uniq -c| sort -k 1 -n -r|less


