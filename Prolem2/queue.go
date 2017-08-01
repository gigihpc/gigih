package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}
type Queue_Val struct {
	buf               []interface{}
	head, tail, count int
}

func New(size int) *Queue_Val {
	return &Queue_Val{
		buf: make([]interface{}, 0),
	}
}

func (q *Queue_Val) Push(key interface{}) {
	q.buf = append(q.buf, key)
	if q.count > 0 && (q.head == q.tail) {
		q.count++
		//q.buf = append(q.buf, key)
		q.tail++
	} else if q.count == 5 && q.tail >= 5 {
		q.head++
		q.tail++
	} else {
		//q.buf = append(q.buf, key)
		q.count++
		q.tail++
	}
}

func (q *Queue_Val) Pop() interface{} {
	tmp := q.buf[q.head]
	if q.count > 0 {
		q.count--
		q.head++
	}
	return tmp
}

func (q *Queue_Val) Contains(key interface{}) bool {
	check := false
	for i := 0; i < len(q.buf); i++ {
		if q.buf[i] == key {
			check = true
		}
	}
	return check
}

func (q *Queue_Val) Len() int {
	return q.count
}

func (q *Queue_Val) Keys() []interface{} {
	return q.buf
}
