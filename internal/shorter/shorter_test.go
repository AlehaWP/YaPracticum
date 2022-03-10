package shorter

import "testing"

func TestMakeShortner(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Тест 1",
			args{
				"some_string",
			},
			"31ee76261d87fed8cb9d4c465c48158c",
		},
		{
			"Тест 2",
			args{
				"",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeShortner(tt.args.s); got != tt.want {
				t.Errorf("MakeShortner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMakeShortner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeShortner("123")
	}
}
