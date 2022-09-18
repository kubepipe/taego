package util

import "testing"

func TestGetLocalIp(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := GetLocalIp(); got != tt.want {
			t.Errorf("%q. GetLocalIp() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_getLocalIP(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := getLocalIP()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. getLocalIP() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. getLocalIP() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
