package kodoku

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid(4, 4)
	t.Log(g.String())
}

func TestGrid_Clone(t *testing.T) {
	g := NewGrid(4, 4)

	g.Fill(0, 0, T1)
	g.Fill(0, 1, T2)
	g.Fill(0, 2, T3)
	g.Fill(0, 3, T4)
	g.Fill(1, 0, T1)
	g.Fill(1, 1, T2)
	g.Fill(1, 2, T3)

	g.FillAndLock(1, 3, T4)
	g.FillAndLock(2, 0, T1)
	g.FillAndLock(2, 1, T2)
	g.FillAndLock(2, 2, T3)
	g.FillAndLock(2, 3, T4)
	g.FillAndLock(3, 0, T1)
	g.FillAndLock(3, 1, T2)
	g.FillAndLock(3, 2, T3)
	g.FillAndLock(3, 3, T4)

	g1 := g.Clone()

	g1.Fill(0, 0, T2)
	g1.Fill(0, 1, T1)
	g1.Fill(0, 2, T4)
	g1.Fill(0, 3, T3)
	g1.Fill(1, 0, T2)
	g1.Fill(1, 1, T4)
	g1.Fill(1, 2, T1)

	t.Log(fmt.Sprintf("%b, %v", g.rows, g.lock))
	t.Log(fmt.Sprintf("%b, %v", g1.rows, g1.lock))

	t.Log(fmt.Sprintf("%v", reflect.DeepEqual(g, g1)))
	t.Log(fmt.Sprintf("%v", reflect.DeepEqual(g.rows, g1.rows)))
	t.Log(fmt.Sprintf("%v", reflect.DeepEqual(g.lock, g1.lock)))
}

func TestGrid_Fill(t *testing.T) {
	g := NewGrid(4, 4)
	g.Fill(0, 0, T1)
	g.Fill(0, 1, T2)
	g.Fill(0, 2, T3)
	g.Fill(0, 3, T4)
	g.Fill(1, 0, T1)
	g.Fill(1, 1, T2)
	g.Fill(1, 2, T3)
	g.Fill(1, 3, T4)
	g.Fill(2, 0, T1)
	g.Fill(2, 1, T2)
	g.Fill(2, 2, T3)
	g.Fill(2, 3, T4)
	g.Fill(3, 0, T1)
	g.Fill(3, 1, T2)
	g.Fill(3, 2, T3)
	g.Fill(3, 3, T4)
	t.Log(g.String())
	t.Log(fmt.Sprintf("%b, %v", g.rows, g.lock))
	t.Log(fmt.Sprintf("%b, %b", g.rows[0][1], g.rows[2][2]))
	t.Log(fmt.Sprintf("%b", g.rows[0][1]^g.rows[2][2]))
}

func TestGrid_FillFromCSV(t *testing.T) {
	g := NewGrid(4, 4)
	g.FillFromCSV("1,1,1,1\n1,1,1,1\n1,1,1,1\n1,1,1,1")
	t.Log(g.String())
}
