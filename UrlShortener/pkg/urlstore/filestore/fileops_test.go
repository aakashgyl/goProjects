package filestore

import (
	"os"
	"testing"
)

func TestGetUrlFileStore(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		errExpected bool
	}{
		{
			name:        "valid filename",
			filename:    "urls.txt",
			errExpected: false,
		},
		{
			name:        "invalid filename",
			filename:    "",
			errExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsObj, err := GetUrlFileStoreServiceObj(tt.filename)
			if err != nil && !tt.errExpected || err == nil && tt.errExpected {
				t.Fail()
			}

			if fsObj == nil && !tt.errExpected {
				t.Fail()
			}
		})
	}
}

func TestUrlFileOps_LoadUrls(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		hasData     bool
		errExpected bool
	}{
		{
			name:        "file with data",
			filename:    "testdata/url_file_with_data.txt",
			hasData:     true,
			errExpected: false,
		},
		{
			name:        "empty file",
			filename:    "testdata/url_file_without_data.txt",
			hasData:     false,
			errExpected: false,
		},
		{
			name:        "file with non-compliant data",
			filename:    "testdata/url_file_with_wrong_data.txt",
			hasData:     false,
			errExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsObj, _ := GetUrlFileStoreServiceObj(tt.filename)
			data, err := fsObj.LoadUrls()
			if err != nil && !tt.errExpected || err == nil && tt.errExpected {
				t.Fail()
			}

			if tt.hasData && len(data) == 0 {
				t.Fail()
			}

			if !tt.hasData && len(data) != 0 {
				t.Fail()
			}
		})
	}
}

func TestUrlFileOps_StoreUrls(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		urlMain     string
		urlShort    string
		errExpected bool
	}{
		{
			name:        "proper urls",
			filename:    "testdata/proper_urls.txt",
			urlMain:     "mail.google.com/gmail/u/0",
			urlShort:    "localhost:8888/aaaaaaa",
			errExpected: false,
		},
		{
			name:        "empty urls",
			filename:    "testdata/empty_urls.txt",
			urlMain:     "",
			urlShort:    "",
			errExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsObj, err := GetUrlFileStoreServiceObj(tt.filename)
			err = fsObj.StoreUrls(tt.urlMain, tt.urlShort)
			if err != nil && !tt.errExpected || err == nil && tt.errExpected {
				t.Fail()
			}
			if _, err := os.Stat(tt.filename); err == nil {
				os.Remove(tt.filename)
			}
		})
	}
}
