package util

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []int64
}

func (s *IntSet) Add(n int) {
	word, bit := n / 64, uint(n % 64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Has(n int) bool {
	word, bit := n / 64, uint(n % n)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

func (s *IntSet) Remove(n int) bool {
	has := s.Has(n)
	if !has {
		return false
	}
	word, bit := n / 64, uint(n % 64)
	s.words[word] = s.words[word] - (1 << bit)
	return true
}

func (s *IntSet) UnionSet(t *IntSet) {
	for key, val := range t.words {
		if key < len(s.words) {
			s.words[key] |= val
		} else {
			s.words = append(s.words, val)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for key, val := range s.words {
		if val == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if val & (1 << uint(i)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(',')
				}
				fmt.Fprintf(&buf, "%d", 64*key+i)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var num int = 0
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			if word & (1 << uint(i)) != 0 {
				num++
			}
		}
	}
	return num
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	var t *IntSet = new(IntSet)
	for _, word := range s.words {
		t.words = append(t.words, word)
	}
	return t
}
