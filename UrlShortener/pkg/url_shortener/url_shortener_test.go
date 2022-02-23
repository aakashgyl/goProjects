package url_shortener

//import "testing"
//
//func TestGetShortenedUrl(t *testing.T) {
//	tests := []struct {
//		name        string
//		url         string
//		errExpected bool
//	}{
//		{
//			name:        "valid url",
//			url:         "mail.google.com/gmail/u/0",
//			errExpected: false,
//		},
//		{
//			name:        "invalid url",
//			url:         "",
//			errExpected: true,
//		},
//	}
//
//	CreateUrlServiceObject()
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			_, err := CreateShortUrl(tt.url)
//			if err != nil && !tt.errExpected || err == nil && tt.errExpected {
//				t.Fail()
//			}
//		})
//	}
//}
//
//func TestGetSameUrl(t *testing.T) {
//	CreateUrlServiceObject()
//	url := "mail.google.com/gmail/u/0"
//	shortUrl1, _ := CreateShortUrl(url)
//	shortUrl2, _ := CreateShortUrl(url)
//
//	if shortUrl1 != shortUrl2 {
//		t.Errorf("Short URLs should be same for a particular input url")
//	}
//}
