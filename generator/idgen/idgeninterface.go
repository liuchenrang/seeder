package idgen

type IDGen interface{
	GetId(bizTag string , step int ) int
}