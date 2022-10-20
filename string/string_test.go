package string

import "testing"

func TestEncrypt(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name:       "test",
			args:       args{input: []byte("123456")},
			wantOutput: "313233343536d41d8cd98f00b204e9800998ecf8427e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Encrypt(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("Encrypt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestGenerateRandNumStr(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{n: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandNumStr(tt.args.n)
			if len(got) != tt.args.n {
				t.Errorf("GenerateRandNumStr() = %v, want %v", got, tt.args.n)
			}
			for _, v := range got {
				if v < '0' || v > '9' {
					t.Errorf("GenerateRandNumStr() = %v, want %v", got, tt.args.n)
				}
			}
			t.Logf("GenerateRandNumStr() = %v", got)
		})
	}
}

func TestGenerateRandLetterStr(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{n: 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandLetterStr(tt.args.n)
			if len(got) != tt.args.n {
				t.Errorf("GenerateRandLetterStr() = %v, want %v", got, tt.args.n)
			}
			for _, v := range got {
				if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') {
					t.Errorf("GenerateRandLetterStr() = %v, want %v", got, tt.args.n)
				}
			}
			t.Logf("GenerateRandLetterStr() = %v", got)
		})
	}
}
