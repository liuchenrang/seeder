package idgen

type IDGen interface {
	GenerateSegment(bizTag string) (currentId uint64, cacheSteop uint64, step uint64, e error)
	Find(bizTag string) (currentId uint64, cacheStep uint64, step uint64, e error)
}
