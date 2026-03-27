package tree

import "testing"

func TestTree(t *testing.T) {

	tr := NewTree[int]()

	tr.Add(5)
	tr.Add(2)
	tr.Add(8)
	tr.Add(1)

	result := tr.Values()
	expected := []int{1, 2, 5, 8}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("ожидали %v, получили %v", expected, result)
		}
	}
}
