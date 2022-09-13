package vm

type ImageType string

const (
	QCOW2 = ImageType("qcow2")
)

type Image struct {
	Name string
	Path string
	Type ImageType
	Size int
}

func (i *Image) GetName() string {
	return i.Name
}

func (i *Image) GetSize() int {
	return i.Size
}

func (i *Image) GetType() ImageType {
	return i.Type
}
