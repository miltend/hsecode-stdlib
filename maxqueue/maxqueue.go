package maxqueue

//
////go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"
//
//import "github.com/cheekybits/genny/generic"
//
//type ValueType generic.Type
//
//type MaxQueue struct {
//	in  []int
//	out []int
//}
//
//func New() *MaxQueue {
//	return &MaxQueue{}
//}
//func (q *MaxQueue) Max() (int, error) {
//
//	return 0, nil
//}
//
//func (q *MaxQueue) Pop() (int, error) {
//	value := q.in[len(q.in)]
//	q.in = q.in[:len(q.in)-1]
//	q.out = append(q.out, value)
//	return value, nil
//}
//
//func (q *MaxQueue) Push(value int) {
//	q.in = append(q.in, value)
//}
