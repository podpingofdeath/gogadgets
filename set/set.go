// 实现一个任意类型的集合结构

package set

// 定义一个空结构体
type nullT struct{}

// 集合中的元素需要实现Element接口
type Element interface {
	Hash() string //如果两个Element的Hash()相同，则认为它们是相同的
}

// Set 集合类型
type Set struct {
	//非线程安全类型
	Elements map[Element]nullT
	Keys     map[string]nullT
}

// NewSet 创建一个新的集合，返回集合的地址
func NewSet() *Set {
	var s Set
	s.Elements = make(map[Element]nullT, 8)
	s.Keys = make(map[string]nullT, 8)
	return &s
}

// Copy 复制一个新的集合
func Copy(other *Set) *Set {
	s3 := NewSet()
	for key := range other.Keys {
		s3.Keys[key] = nullT{}
	}
	for e := range other.Elements {
		s3.Elements[e] = nullT{}
	}
	return s3
}

// Size 返回集合中元素的个数
func (s *Set) Size() int {
	return len(s.Keys)
}

// Exist 判断集合中某个元素是否存在
func (s *Set) Exist(element Element) bool {
	if _, ok := s.Keys[element.Hash()]; ok {
		return true
	}
	return false
}

// Add 向集合中添加元素，返回集合本身，方便链式调用
func (s *Set) Add(element ...Element) *Set {
	for _, v := range element {
		if !s.Exist(v) {
			s.Elements[v] = nullT{}
			s.Keys[v.Hash()] = nullT{}
		}
	}
	return s
}

// Remove 从集合中删除某个元素
func (s *Set) Remove(element Element) *Set {
	delete(s.Keys, element.Hash())
	for k := range s.Elements {
		if k.Hash() == element.Hash() {
			delete(s.Elements, k)
			break
		}
	}
	return s
}

// Clean 清空集合中的所有元素
func (s *Set) Clean() *Set {
	for k := range s.Keys {
		delete(s.Keys, k)
	}
	for k := range s.Elements {
		delete(s.Elements, k)
	}
	return s
}

// Union 与另一集合做并集，会改变自身
func (s *Set) Union(other *Set) *Set {
	var key string
	for e := range other.Elements {
		key = e.Hash()
		if _, ok := s.Keys[key]; !ok {
			s.Elements[e] = nullT{}
			s.Keys[key] = nullT{}
		}
	}
	return s
}

// Join 与另一集合做交集，会改变自身，效率低，建议用下面的Join函数
func (s *Set) Join(other *Set) *Set {
	if s.Size() == 0 || other.Size() == 0 {
		return s.Clean()
	}

	for e := range s.Elements {
		if !other.Exist(e) {
			delete(s.Elements, e)
			delete(s.Keys, e.Hash())
		}
	}
	return s
}

// ToSlice 将集合转为切片
func (s *Set) ToSlice() []Element {
	es := make([]Element, 0, s.Size())
	for e := range s.Elements {
		es = append(es, e)
	}
	return es
}

// Join 多个Set取交集，产生新的集合
func Join(ss ...*Set) *Set {
	s3 := NewSet()
	if ss == nil || len(ss) == 0 {
		return s3
	}

	s1 := ss[0]
	var sign bool
	for e := range s1.Elements {
		sign = true
		for _, s := range ss[1:] {
			if !s.Exist(e) {
				sign = false
				break
			}
		}
		if sign == true {
			s3.Elements[e] = nullT{}
			s3.Keys[e.Hash()] = nullT{}
		}
	}

	return s3
}

// Union 多个Set取并集，不改变原来的集合
func Union(ss ...*Set) *Set {
	s3 := NewSet()
	if ss == nil || len(ss) == 0 {
		return s3
	}

	for _, s := range ss {
		for e := range s.Elements {
			if !s3.Exist(e) {
				s3.Elements[e] = nullT{}
				s3.Keys[e.Hash()] = nullT{}
			}
		}
	}
	return s3
}
