package idgen

type IDGen interface{
	GenerateSegment(bizTag string  ) (uint64, uint64, error)
	UpdateStep(bizTag string) (int64, error)
}