package api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestNewUserApi(t *testing.T) {
	tests := []struct {
		name string
		want *UserAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserApi(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAPI_GetBal(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := &UserAPI{}
			ua.GetBal(tt.args.c)
		})
	}
}

func TestUserAPI_GetOrder(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := &UserAPI{}
			ua.GetOrder(tt.args.c)
		})
	}
}

func TestUserAPI_GetWithD(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := &UserAPI{}
			ua.GetWithD(tt.args.c)
		})
	}
}

func TestUserAPI_Login(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := &UserAPI{}
			ua.Login(tt.args.c)
		})
	}
}

func TestUserAPI_PostOrder(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := &UserAPI{}
			ua.PostOrder(tt.args.c)
		})
	}
}

func TestUserAPI_PostWithD(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := &UserAPI{}
			ua.PostWithD(tt.args.c)
		})
	}
}

func TestUserAPI_Regis(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := &UserAPI{}
			ua.Regis(tt.args.c)
		})
	}
}
