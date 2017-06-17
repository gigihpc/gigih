package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

func New(size int) Queue {
  q := [size]Queue{}
	return q
}

func (q Queue)Push(key interface{})  {
  q = append(q,key);
}

func (q Queue)Pop() interface{}  {
  return q[len(q)-1]
}

func (q Queue)Contains(key interface{}) bool  {
  isFill := false
  for index := 0; index < Len(); index++ {
    if q[index] == key {
      isFill = true
    }
  }
  return isFill
  }

func (q Queue)Len() int  {
  return len(q)
}

func (q Queue)Keys() []interface{}  {
  return q[0]
}
