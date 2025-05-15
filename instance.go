package wmiv2

import "github.com/walliba/go-wmiv2/internal/mi"

type miInstance struct {
	raw *mi.Instance
}

func (inst *miInstance) GetProperties() any {
	return nil
}
