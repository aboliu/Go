package main

import (
	"testing"
)

func TestGetMachineGame(t *testing.T) {
	type args struct {
		MID string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "查無此人", args: args{MID: "000"}, want: 0},
		{name: "totalBalanceInGame小於0", args: args{MID: "123"}, want: 0},
		{name: "totalBalanceInGame大於0", args: args{MID: "123456"}, want: 1},
		{name: "BetBase為空字串", args: args{MID: "1234567"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMachineGame(tt.args.MID); got != tt.want {
				t.Errorf("GetMachineGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
