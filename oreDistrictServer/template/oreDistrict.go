package template

type OreDistrict struct {
	Id    uint32
	Total uint32
}

type OreDistrictTemplate struct {
	data map[uint32]*OreDistrict
}

func (i *OreDistrictTemplate) load() {
	i.data = make(map[uint32]*OreDistrict)
	rf := readRf(OreDistrict{})
	for k := 0; k < rf.NumRecord(); k++ {
		oreDistrict := rf.Record(k).(*OreDistrict)
		i.data[oreDistrict.Id] = oreDistrict
	}
}

// GetOreDistrict 获得指定矿洞
func (i *OreDistrictTemplate) GetOreDistrict(id uint32) *OreDistrict {
	if ret, ok := i.data[id]; ok {
		return ret
	}
	return nil
}

// GetAll 获取所有矿洞
func (i *OreDistrictTemplate) GetAll() []*OreDistrict {
	var ret []*OreDistrict
	for _, data := range i.data {
		ret = append(ret, data)
	}
	return ret
}
