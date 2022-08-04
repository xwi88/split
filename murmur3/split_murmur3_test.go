package murmur3

import (
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		data string
		mod  int
		seed int
	}
	tests := []struct {
		name  string
		args  args
		wantB int
	}{
		{name: "xwi88_100_0", args: args{data: "xwi88", mod: 100, seed: 0x00}, wantB: 31},
		{name: "xwi88_100_1", args: args{data: "xwi88", mod: 100, seed: 0x01}, wantB: 80},
		{name: "xwi88_1000_0", args: args{data: "xwi88", mod: 1000, seed: 0x00}, wantB: 531},
		{name: "xwi88_1000_1", args: args{data: "xwi88", mod: 1000, seed: 0x01}, wantB: 680},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := Split(tt.args.data, tt.args.mod, tt.args.seed); gotB != tt.wantB {
				t.Errorf("Split() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	data := "xwi88"
	mod := 100
	seed := 0x00

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Split(data, mod, seed+i)
	}
}
