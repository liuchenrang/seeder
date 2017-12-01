package idgen

type IDGen interface{
	generateSegment(bizTag string  ) (uint64, uint64, error)
}