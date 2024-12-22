package ginx

import (
	"reflect"
	"testing"

	"github.com/DaHuangQwQ/ginx/openapi"
	"github.com/gin-gonic/gin"
)

func TestServer_marshalSpec(t *testing.T) {
	type fields struct {
		Engine  *gin.Engine
		OpenAPI openapi.OpenAPI
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Engine:  tt.fields.Engine,
				OpenAPI: tt.fields.OpenAPI,
			}
			got, err := s.marshalSpec()
			if (err != nil) != tt.wantErr {
				t.Errorf("marshalSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marshalSpec() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_saveOpenAPIToFile(t *testing.T) {
	type fields struct {
		Engine  *gin.Engine
		OpenAPI openapi.OpenAPI
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Engine:  tt.fields.Engine,
				OpenAPI: tt.fields.OpenAPI,
			}
			if err := s.saveOpenAPIToFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("saveOpenAPIToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_registerOpenAPIRoutes(t *testing.T) {
	type fields struct {
		Engine  *gin.Engine
		OpenAPI openapi.OpenAPI
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Engine:  tt.fields.Engine,
				OpenAPI: tt.fields.OpenAPI,
			}
			s.registerOpenAPIRoutes(tt.args.path)
		})
	}
}
