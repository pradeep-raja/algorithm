package list

import "testing"

func checkList(t *testing.T, l *List, values []int) {
	i := 0
	for item := l.First(); item != nil; item = item.Next() {
		ev := item.Value.(int)
		if ev != values[i] {
			t.Errorf("expected %d at %d, got %d", values[i], i, ev)
		}
		i++
	}
}

func compareList(t *testing.T, list1, list2 *List) {
	if list1.Length() != list2.Length() {
		t.Errorf("expected %d, got %d", list1.Length(), list2.Length())
		return
	}

	item2 := list2.First()
	for item1 := list1.First(); item1 != nil; item1 = item1.Next() {
		item1Val := item1.Value.(int)
		item2Val := item2.Value.(int)
		if item1Val != item2Val {
			t.Errorf("List item value not matched %d != %d", item1Val, item2Val)
		}

		if item1 == item2 {
			t.Errorf("Both items are pointing to same memory")
		}
		item2 = item2.Next()
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

func TestSubList(t *testing.T) {
	list := New()
	list.Append(1)
	list.Append(45)
	list.Append(67)
	list.Append(744)
	list.Append(-64)
	list.Append(3)
	newList := list.SubList(0, -1)

	compareList(t, list, newList)

	newList = list.SubList(1, 2)
	testList := New()
	testList.Append(45)
	testList.Append(67)
	compareList(t, newList, testList)

	newList = list.SubList(2, 1)
	testList = New()
	testList.Append(67)
	testList.Append(45)
	compareList(t, newList, testList)

	newList = list.SubList(-1, 0)
	testList = list.SubList(0, -1)
	testList.Reverse()
	compareList(t, newList, testList)

	newList = list.SubList(0, 8)
	if newList != nil {
		t.Errorf("Non nil list created from invalid range")
	}

	newList = list.SubList(8, -1)
	if newList != nil {
		t.Errorf("Non nil list created from invalid range")
	}

	newList = list.SubList(-2, 0)
	if newList != nil {
		t.Errorf("Non nil list created from invalid range")
	}

	newList = list.SubList(-2, -4)
	if newList != nil {
		t.Errorf("Non nil list created from invalid range")
	}
}
