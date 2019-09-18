package list

import "testing"

func checkList(t *testing.T, l *List, values []int) {
	i := 0
	for item := l.First();item != nil; item = item.Next() {
		ev := item.Value.(int)
		if ev != values[i] {
			t.Errorf("expected %d at %d, got %d",values[i], i , ev)
		}
		i++
	}
}

func TestReverse(t *testing.T) {
	list := New()
	list.Append(34)
	list.Append(12)
	list.Append(65)
	list.Prepend(45)
	list.Prepend(43)
	list.Append(56)

	list.Reverse()
	values := []int{56, 65, 12, 34, 45, 43}
	checkList(t, list, values)
	
	list.Reverse()
	values = []int{43, 45, 34, 12, 65, 56}	
	checkList(t, list, values)	
}

func TestAddValues(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append(45)
	list.Append(67)
	list.Append(744)
	list.Append(-64)
	list.Append(3)
	checkList(t, list, []int{1, 45, 67, 744, -64, 3})

	list = New()
	list.Prepend(1)
	list.Prepend(45)
	list.Prepend(67)
	list.Prepend(744)
	list.Prepend(-64)
	list.Prepend(3)
	checkList(t, list, []int{3, -64, 744, 67, 45, 1})
}
