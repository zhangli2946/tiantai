package _31

import (
    "github.com/go-playground/assert/v2"
    "sort"
    "testing"
)

func TestCheck(t *testing.T) {
    {
        p := []int{4, 3, 2, 1}
        t.Log(check(p))
    }
    {
        p := []int{4, 3, 1, 2}
        t.Log(check(p))
    }
    {
        p := []int{3, 1, 2}
        t.Log(check(p))
    }
    {
        p := []int{2, 3, 1}
        t.Log(check(p))
    }
    {
        p := []int{4, 2, 3, 1}
        t.Log(check(p))
    }
}

func TestSort(t *testing.T) {
    p := []int{4, 3, 2, 1}
    sort.Ints(p)
    t.Log(p)
}

func TestSolve(t *testing.T) {
    {
        p := []int{9, 8, 7, 6, 5, 4, 3, 2}
        nextPermutation(p)
        assert.Equal(t, p, []int{2, 3, 4, 5, 6, 7, 8, 9})
    }
    {
        p := []int{2, 3, 1}
        nextPermutation(p)
        assert.Equal(t, p, []int{3, 1, 2})
    }
    {
        p := []int{9, 8, 7, 3, 1, 4, 2}
        nextPermutation(p)
        assert.Equal(t, p, []int{9, 8, 7, 3, 2, 1, 4})
    }
    {
        p := []int{9, 8, 7, 3, 1, 4, 1}
        nextPermutation(p)
        assert.Equal(t, p, []int{9, 8, 7, 3, 4, 1, 1})
    }

    {
        p := []int{2, 3, 2}
        nextPermutation(p)
        assert.Equal(t, p, []int{3, 2, 2})
    }

    {
        p := []int{5, 4, 7, 5, 3, 2}
        nextPermutation(p)
        assert.Equal(t, p, []int{5, 5, 2, 3, 4, 7})
    }

}
