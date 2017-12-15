build:
	@echo "\n--------------> build thrift <--------------\n"
	mkdir -p packages
	rm -rf packages/*
	thrift  --gen go -out packages generator.thrift
	thrift  --gen go -out packages inthrift.thrift
