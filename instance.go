package wmiv2

type miInstance struct {
	properties *DynamicStruct
}

func (inst *miInstance) GetProperties() *DynamicStruct {
	return inst.properties
}
