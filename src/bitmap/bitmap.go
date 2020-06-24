package bitmap

import (
	"bytes"
	"fmt"
)

type Bitmap struct {
	words []uint64
}

func (b Bitmap) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	if word < len(b.words) && b.words[word]&(1<<bit) != 0 {
		return true
	}
	return false
}

func (b *Bitmap) Add(xs ...int) {
	for _, x := range xs {
		word, bit := x/64, uint(x%64)
		for word >= len(b.words) {
			b.words = append(b.words, 0)
		}
		b.words[word] |= (1 << bit)
	}

}

func (b *Bitmap) UnionBitMap(s *Bitmap) {
	for i, k := range s.words {
		if len(b.words) <= i {
			b.words = append(b.words, 0)
		} else {
			b.words[i] |= k
		}
	}
}

func (b *Bitmap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range b.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(',')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (b *Bitmap) Len() int {
	ans := 0
	for _, k := range b.words {
		if k == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if k&(1<<uint(j)) != 0 {
				ans++
			}
		}
	}
	return ans
}
func (b *Bitmap) Remove(x int) {
	word, bit := x/64, x%64
	if len(b.words) <= word {
		return
	}
	b.words[word] -= 1 << uint(bit)
}
func (b *Bitmap) Clear() {
	for i := range b.words {
		b.words[i] = 0
	}
}
func (b *Bitmap) Copy() Bitmap {
	var ans Bitmap
	for i, k := range b.words {
		ans.words = append(ans.words, 0)
		ans.words[i] = k
	}
	return ans
}

func (b *Bitmap) Elems() []int {

	var ans []int
	for i, word := range b.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				ans = append(ans, i*64+j)
			}
		}
	}
	return ans
}
