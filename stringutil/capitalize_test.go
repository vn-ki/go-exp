package stringutil

import "testing"

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test capitalize string",
			args: "abcdEFGH",
			want: "ABCDEFGH",
		},
		{
			name: "Test capitalize numbers",
			args: "1234",
			want: "1234",
		},
		{
			name: "Test capitalize symbols",
			args: ",.;'[]",
			want: ",.;'[]",
		},
		{
			name: "Test capitalize mix",
			args: "aB;0",
			want: "AB;0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Capitalize(tt.args); got != tt.want {
				t.Errorf("Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_capitalize(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name    string
		args    rune
		want    rune
		wantErr bool
	}{
		{
			name:    "Test capitalize character",
			args:    'a',
			want:    65,
			wantErr: false,
		},
		{
			name:    "Test capitalize numbers",
			args:    '1',
			want:    -1,
			wantErr: true,
		},
		{
			name:    "Test capitalize symbols",
			args:    ',',
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := capitalize(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("capitalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
