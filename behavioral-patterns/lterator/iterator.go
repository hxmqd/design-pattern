// 模式用于顺序访问集合对象的元素，不需要知道集合对象的底层表示

package lterator

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Container interface {
	GetIterator() Iterator
}

type NameRepository struct {
	Names []string
}

func NewNameRepository(names []string) *NameRepository{
	return &NameRepository{
		Names: names,
	}
}

func (n *NameRepository) GetIterator() *NameIterator{
	data := make([]interface{}, 0)
	for i := range n.Names {
		data = append(data, n.Names[i])
	}
	return &NameIterator{data: data, index:0}
}

type NameIterator struct {
	data []interface{}
	index int
}

func (n *NameIterator) HasNext() bool{
	if n.index < len(n.data){
		return true
	}
	return false
}

func (n *NameIterator) Next() interface{}{
	if n.HasNext() {
		 result := n.data[n.index]
		 n.index+=1
		 return result
	}
	return nil
}






