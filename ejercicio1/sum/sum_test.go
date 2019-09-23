package sum

import (
	"fmt"
	"testing"
)

func TestGetsum(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{15, 60},
	}

	for _, c := range cases {
		got := Getsum(c.in)
		if got != c.want {
			t.Errorf("Getsum(%q) == %q	, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkGetsum(b *testing.B) {
	p := 15
	want := 60

	// run the Getsum function b.N times
	for n := 0; n < b.N; n++ {
		b.Run("suma", func(b *testing.B) {
			got := Getsum(p)
			if got != want {
				b.Errorf("got %.2d want %.2d", got, want)
			}
		})
		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					Getsum(p)
				}
			},
		)
	}
}

func ExampleGetsum() {
	fmt.Println(Getsum(15))
	//output: 60
}
 