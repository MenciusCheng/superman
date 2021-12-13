package daenerys

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSimpleHandler_Gen(t *testing.T) {
	type fields struct {
		HandlerName string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			fields: fields{HandlerName: "CommodityDetail"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleHandler{
				HandlerName: tt.fields.HandlerName,
			}
			wr := &bytes.Buffer{}
			err := s.Gen(wr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(wr.String())
		})
	}
}
