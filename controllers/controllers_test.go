package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewInfo(t *testing.T) {
	type args struct {
		c             *gin.Context
		info          *Info
		returnDataMsg string
	}
	tests := []struct {
		name            string
		args            args
		wantKuberconfig string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotKuberconfig := NewInfo(tt.args.c, tt.args.info, tt.args.returnDataMsg); gotKuberconfig != tt.wantKuberconfig {
				t.Errorf("NewInfo() = %v, want %v", gotKuberconfig, tt.wantKuberconfig)
			}
		})
	}
}
