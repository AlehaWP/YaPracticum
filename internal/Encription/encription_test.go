package encription

import (
	"testing"
)

func Test_generateRandom(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"Тест 1",
			args{
				16,
			},
			16,
			false,
		},
		{
			"Тест 2",
			args{
				5,
			},
			5,
			false,
		},
		{
			"Тест 3",
			args{
				0,
			},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateRandom(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateRandom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("generateRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncriptStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"Тест 1",
			args{
				"askjdh",
			},
			32,
			false,
		},
		{
			"Тест 2",
			args{
				"",
			},
			0,
			true,
		},
		{
			"Тест 3",
			args{
				"gkshlakjsdfh;ajkfh;alskdj;alksdfasjkdfhkjsdfb,asjdbfajdflakjdfnaksndf",
			},
			32,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncriptStr(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncriptStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("EncriptStr() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
