package main

type deque struct {
	prep []interface{}
	ap   []interface{}
}

func newDeque() *deque {
	return &deque{}
}

func (d *deque) PushFront(item interface{}) {
	d.ap = append(d.ap, item)
}

func (d *deque) PushBack(item interface{}) {
	d.prep = append(d.prep, item)
}

func (d *deque) Empty() bool {
	return len(d.ap) == 0 && len(d.prep) == 0
}

func (d *deque) PopFront() interface{} {
	lenap := len(d.ap)
	if lenap != 0 {
		v := d.ap[lenap-1]
		d.ap = d.ap[:lenap-1]
		return v
	}
	v := d.prep[0]
	d.prep = d.prep[1:]
	return v
}

func (d *deque) PopBack() interface{} {
	lenprep := len(d.prep)
	if lenprep != 0 {
		v := d.prep[lenprep-1]
		d.prep = d.prep[:lenprep-1]
		return v
	}
	v := d.ap[0]
	d.ap = d.ap[1:]
	return v
}
