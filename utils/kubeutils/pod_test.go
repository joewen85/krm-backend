package kubeutils

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestNewPod(t *testing.T) {
	type args struct {
		kubecofig string
		item      *corev1.Pod
	}
	tests := []struct {
		name string
		args args
		want *Pod
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPod(tt.args.kubecofig, tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPod() = %v, want %v", got, tt.want)
			}
		})
	}
}
