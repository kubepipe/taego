package mk8s

import (
	"reflect"
	"testing"

	"k8s.io/client-go/kubernetes"
)

func TestGetInClusterClientset(t *testing.T) {
	tests := []struct {
		name string
		want *kubernetes.Clientset
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := GetInClusterClientset(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetInClusterClientset() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
