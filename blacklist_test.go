package blacklist

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var cc *Client
var ee error

func TestMain(m *testing.M) {

	// Write code here to run before tests
	ee = godotenv.Load("vars.env")
	if ee != nil {
		log.Fatalf("authentication error: %v", ee)
	}

	cc = NewBlackListClient(os.Getenv("BLACKLIST_BASE_URL"), os.Getenv("BLACKLIST_API_KEY"))

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func TestClient_SearchBlacklistFull(t *testing.T) {

	tests := []struct {
		cl      *Client
		name    string
		bvn     string
		wantErr bool
	}{
		// TODO: Add test cases.
		{cc,
			"valid BVN",
			os.Getenv("VALID_BVN_NUMBER"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.cl.SearchBlacklistFull(tt.bvn)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchBlacklistFull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotBlr, tt.wantBlr) {
			//	t.Errorf("SearchBlacklistFull() gotBlr = %v, want %v", gotBlr, tt.wantBlr)
			//}
		})
	}
}

func TestClient_SearchCRCFull(t *testing.T) {

	tests := []struct {
		cl      *Client
		name    string
		bvn     string
		wantErr bool
	}{
		// TODO: Add test cases.
		{cc,
			"valid BVN",
			os.Getenv("VALID_BVN_NUMBER"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cd, err := tt.cl.SearchCRCFull(tt.bvn)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchBlacklistFull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if cd != nil {
				t.Logf("crc result: %+v\n", cd)
			}
			//if !reflect.DeepEqual(gotBlr, tt.wantBlr) {
			//	t.Errorf("SearchBlacklistFull() gotBlr = %v, want %v", gotBlr, tt.wantBlr)
			//}
		})
	}
}
