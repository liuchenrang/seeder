build: 
	thrift --out ./src/packages/thrift  --gen go ./src/resource/tutorial.thrift
