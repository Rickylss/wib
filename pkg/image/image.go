package image

type Image interface {
	GetName() string
	SetSize() int
	GetSize() int
}

type Qcow2Image struct {
	Name string
	Size int64
}

type RawImage struct {
	Name string
	Size int64
}

func (qi *Qcow2Image) GetName() string {
	return ""
}
func (qi *Qcow2Image) SetSize() int {
	return 0
}
func (qi *Qcow2Image) GetSize() int {
	return 0
}
