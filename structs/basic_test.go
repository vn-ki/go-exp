package structs

import (
	"testing"
)

func TestWallet_Balance(t *testing.T) {
	tests := []struct {
		name   string
		wallet Wallet
		want   int64
	}{
		{
			name:   "Test balance",
			wallet: Wallet{100, 300},
			want:   100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wallet.Balance(); got != tt.want {
				t.Errorf("Wallet.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallet_AddMoney(t *testing.T) {
	type fields struct {
		money    int64
		MaxMoney int64
	}
	type args struct {
		money int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		expect  int64
	}{
		{
			name:    "Test add money normal",
			fields:  fields{200, 400},
			args:    args{100},
			wantErr: false,
			expect:  300,
		},
		{
			name:    "Test add money overflow",
			fields:  fields{200, 400},
			args:    args{400},
			wantErr: true,
			expect:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wallet{
				Money:    tt.fields.money,
				MaxMoney: tt.fields.MaxMoney,
			}
			if err := w.AddMoney(tt.args.money); (err != nil) != tt.wantErr {
				t.Errorf("Wallet.AddMoney() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWallet_TakeMoney(t *testing.T) {
	type fields struct {
		money    int64
		MaxMoney int64
	}
	type args struct {
		money int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		expect  int64
	}{
		{
			name:    "Test add money normal",
			fields:  fields{200, 400},
			args:    args{100},
			wantErr: false,
			expect:  100,
		},
		{
			name:    "Test add money underflow",
			fields:  fields{200, 400},
			args:    args{300},
			wantErr: true,
			expect:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wallet{
				Money:    tt.fields.money,
				MaxMoney: tt.fields.MaxMoney,
			}
			if err := w.TakeMoney(tt.args.money); (err != nil) != tt.wantErr {
				t.Errorf("Wallet.TakeMoney() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
