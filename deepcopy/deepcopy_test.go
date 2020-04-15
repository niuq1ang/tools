package deepcopy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	type Task struct {
		S  string
		SI *string
		I  int64
		II *int64
		F  float64
		FI *float64
		B  bool
		BI *bool
		A  []*Task
		IT interface{}
	}

	t1 := new(Task)
	var s string = ""
	t1.S = s
	t1.SI = &s
	var i int64 = 0
	t1.I = i
	t1.II = &i
	var f float64 = 0
	t1.F = f
	t1.FI = &f
	var b bool = false
	t1.B = b
	t1.BI = &b
	var a []*Task
	t1.A = a
	t1.IT = 100

	t2 := Copy(t1).(*Task)
	t.Log(fmt.Sprintf("t1:%+v", t1))
	t.Log(fmt.Sprintf("t2:%+v", t2))
	if reflect.DeepEqual(t1, t2) {
		t.Log("compare default value ok")
	} else {
		t.Fatalf("not eqaul")
	}

	t3 := new(Task)
	s = "adc"
	t3.S = s
	t3.SI = &s
	i = 100
	t3.I = i
	t3.II = &i
	f = 1000
	t3.F = f
	t3.FI = &f
	b = true
	t3.B = b
	t3.BI = &b
	a = []*Task{new(Task)}
	t3.A = a
	t3.IT = new(Task)
	t4 := Copy(t3).(*Task)
	t.Log(fmt.Sprintf("t3:%+v", t3))
	t.Log(fmt.Sprintf("t4:%+v", t4))
	if reflect.DeepEqual(t3, t4) {
		t.Log("compare common value ok")
	} else {
		t.Fatalf("not eqaul")
	}
}
