package blacklist

import (
	"reflect"
	"testing"
)

func TestClient_SearchBlacklistFull(t *testing.T) {
	type fields struct {
		url    string
		apiKey string
	}
	type args struct {
		bvn string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBlr []BlacklistLoanRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				url:    tt.fields.url,
				apiKey: tt.fields.apiKey,
			}
			gotBlr, err := c.SearchBlacklistFull(tt.args.bvn)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchBlacklistFull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBlr, tt.wantBlr) {
				t.Errorf("SearchBlacklistFull() gotBlr = %v, want %v", gotBlr, tt.wantBlr)
			}
		})
	}
}
