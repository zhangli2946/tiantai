package _6

import "bytes"


func convert(s string, numRows int) string {
    var (
        stat = newState(numRows)
    )
    for i, los := 0, len(s); i < los; i += 1 {
        stat.Accept(s[i])
    }
    return stat.String()
}

func newState(rows int) *State {
    as := make([][]byte, rows)
    return &State{as: as, r: rows - 1, n: 1}
}

type State struct {
    as   [][]byte
    r, n int
    curr int
}

func (s State) String() string {
    w := bytes.Buffer{}
    for i := range s.as {
        w.Write(s.as[i])
    }
    return w.String()
}

func (s *State) Accept(c byte) {
    s.as[s.curr] = append(s.as[s.curr], c)
    switch {
    case s.r == 0:
        s.n = 0
    case s.curr == 0:
        s.n = 1
    case s.curr == s.r:
        s.n = -1
    }
    s.curr += s.n
}
