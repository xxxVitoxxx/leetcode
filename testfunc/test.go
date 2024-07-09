package testfunc

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

// ElementsMatch assert that the specified listA(slice) is equal to specified
// listB(slice) ignoring the order of the elements. if there are duplicate
// element, the number of appearances of each of them in both lists should match.
//
// ElementsMatch([3, 1, 3, 2], [2, 3, 1, 3])
func ElementsMatch(listA, listB interface{}) bool {
	if isEmpty(listA) && isEmpty(listB) {
		return true
	}

	if !isList(listA) || !isList(listB) {
		return false
	}

	extraA, extraB := diffList(listA, listB)
	if len(extraA) == 0 && len(extraB) == 0 {
		return true
	}

	fmt.Println(formatListDiff(listA, listB, extraA, extraB))

	return false
}

// isEmpty gets whether the specified object is considered empty or not.
func isEmpty(object interface{}) bool {
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)
	switch objValue.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return isEmpty(deref)
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

// isList check that the list is array or slice
func isList(list interface{}) bool {
	kind := reflect.TypeOf(list).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return false
	}
	return true
}

// diffLists diffs two arrays/slices and return slices of elements that are only in A and only in B.
// if some element is present multiple times, each instance is counted separately (e.g. if something
// is 2x in A and 5x in B, it will be 0x in extraA and 3x in extraB). the order of items in both lists is ignored.
func diffList(listA, listB interface{}) (extraA, extraB []interface{}) {
	aValue, bValue := reflect.ValueOf(listA), reflect.ValueOf(listB)
	aLen, bLen := aValue.Len(), bValue.Len()

	visited := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		element := aValue.Index(i).Interface()
		found := false
		for j := 0; j < bLen; j++ {
			if visited[j] {
				continue
			}
			if objectsAreEqual(bValue.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}

		if !found {
			extraA = append(extraA, element)
		}
	}

	for j := 0; j < bLen; j++ {
		if visited[j] {
			continue
		}
		extraB = append(extraB, bValue.Index(j).Interface())
	}

	return
}

// objectsAreEqual determines if two objects are considered equal.
//
// this function does no assertion of any kind.
func objectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

func formatListDiff(listA, listB interface{}, extraA, extraB []interface{}) string {
	var msg bytes.Buffer

	msg.WriteString("elements differ")
	if len(extraA) > 0 {
		msg.WriteString("\n\nextra elements in list A: \n")
		msg.WriteString(spew.Config.Sdump(extraA))
	}
	if len(extraB) > 0 {
		msg.WriteString("\n\nextra elements in list B: \n")
		msg.WriteString(spew.Config.Sdump(extraB))
	}
	msg.WriteString("\n\nlistA:\n")
	msg.WriteString(spew.Config.Sdump(listA))
	msg.WriteString("\n\nlistB:\n")
	msg.WriteString(spew.Config.Sdump(listB))

	return msg.String()
}
