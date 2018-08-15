package tool

import (
	"testing"
)

func TestMul(t *testing.T) {
	type args struct {
		v1    string
		v2    string
		place int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"5100", "0.1", 2}, "510.00"},
		{"case2", args{"5100", "0.001", 2}, "5.10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.v1, tt.args.v2, tt.args.place); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	type args struct {
		v1    string
		v2    string
		place int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"10", "3", 2}, "3.33"},
		{"case2", args{"5100", "1000", 2}, "5.10"},
		{"case3", args{"0.2", "0.3", 2}, "0.66"},
		{"case4", args{"5100", "10", 2}, "510.00"},
		{"case5", args{"5100", "10", 4}, "510.0000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.args.v1, tt.args.v2, tt.args.place); got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		v1    string
		v2    string
		place int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"1", "1", 2}, "0.00"},
		{"case2", args{"1", "2", 2}, "-1.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.v1, tt.args.v2, tt.args.place); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNewBalance(t *testing.T) {
	type args struct {
		v1    string
		v2    string
		v3    string
		place int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{v1: "1.0", v2: "2.0", v3: "3.0", place: 2}, "2.00"},
		{"case1", args{v1: "1.0", v2: "3.0", v3: "2.0", place: 2}, "0.00"},
		{"case1", args{v1: "1.0", v2: "5.0", v3: "2.0", place: 2}, "-2.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNewBalance(tt.args.v1, tt.args.v2, tt.args.v3, tt.args.place); got != tt.want {
				t.Errorf("GetNewBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWinAmount(t *testing.T) {
	type args struct {
		v1    string
		v2    string
		v3    string
		place int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{"100.0", "0", "20", 2}, "0.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWinAmount(tt.args.v1, tt.args.v2, tt.args.v3, tt.args.place); got != tt.want {
				t.Errorf("GetWinAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
