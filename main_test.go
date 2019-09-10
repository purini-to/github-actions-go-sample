package main

import "testing"

func TestGetMessage(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: `Should get message is "Hello Actions!"`, want: "Hello Actions!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessage(); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Should no errors run"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
