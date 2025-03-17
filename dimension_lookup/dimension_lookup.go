package dimension_lookup

type DimLookUpType int

const (
	DimRow DimLookUpType = iota
	DimCol
)

type DimLookUp map[DimLookUpType]map[int]map[int]bool

func NewDimLookUp() DimLookUp {
	dl := make(DimLookUp, 0)
	dl[DimRow] = make(map[int]map[int]bool, 0)
	dl[DimCol] = make(map[int]map[int]bool, 0)

	for i := 0; i < 9; i++ {
		dl[DimRow][i] = make(map[int]bool, 0)
		dl[DimCol][i] = make(map[int]bool, 0)
	}

	return dl
}

func (dl *DimLookUp) Insert(dimType DimLookUpType, key1, key2 int, value bool) {
	(*dl)[dimType][key1][key2] = value
}

func (dl *DimLookUp) IsKeyExists(dimType DimLookUpType, key1, key2 int) bool {
	_, ok := (*dl)[dimType][key1][key2]
	return ok
}

func (dl *DimLookUp) Remove(dimType DimLookUpType, key1, key2 int) {
	if _, ok := (*dl)[dimType][key1]; ok {
		_, ok = (*dl)[dimType][key1][key2]
		if ok {
			delete((*dl)[dimType][key1], key2)
		}
	}
}
