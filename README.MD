#结构名与类图名对应
函数与类图接口名称保持一致
结构文件 必须包含 New结构名的 函数
每个文件必须有个文件名_test.go的测试
包名和结构一致时, 文件的名字和包名一样
驼峰命名的结果一律都是小写字母命名的文件

id 都用 uint64

申请具体TagSegment 可以多个协成申请吗 
    不可以, 会造成DBGEN多次刷新
segment可以并发访问吗?
    可以, id是 基于原子的    

/seeder/servers/  临时节点
/seeder/workers/    临时顺序节点产生id
/seeder/data/ip-port/time -> time()
/seeder/data/ip-port/id -> 123
