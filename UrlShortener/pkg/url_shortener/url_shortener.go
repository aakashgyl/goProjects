package url_shortener

import (
	"errors"
	"fmt"
	"github.com/aakashgyl/UrlShortener/pkg/urlstore"
	"github.com/aakashgyl/UrlShortener/pkg/urlstore/filestore"
)

const (
	charSet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charSetLen = len(charSet)
)

var (
	fileName      = "url.txt" // file used for persistent storage of processed URLs
	urlsCount     = 0         // counter to generate unique URLs
	shortUrlIdLen = 7         // length of short URL generated
)

type UrlService interface {
	CreateShortenedUrl(string) (string, error)
	GetMainUrl(string) (string, error)
}

type Url struct {
	Storage        urlstore.UrlStore // generic store object to store urls and load on app start
	MainToShortMap map[string]string // map[main_url]: short_url
	ShortToMainMap map[string]string // map[short_url]: main_url
}

// CreateUrlServiceObject initializes URL operator to process URLs
func CreateUrlServiceObject() (UrlService, error) {
	var err error
	urlObj := &Url{}

	// initializing urlOperator with file storage
	urlObj.Storage, err = filestore.GetUrlFileStoreServiceObj(fileName)
	if err != nil {
		return urlObj, err
	}

	urlObj.MainToShortMap, urlObj.ShortToMainMap, err = urlObj.Storage.LoadUrls()
	if err != nil {
		return urlObj, err
	}

	urlsCount = len(urlObj.MainToShortMap)
	return urlObj, nil
}

func (u *Url) GetMainUrl(shortUrl string) (string, error) {
	if mainUrl, ok := u.ShortToMainMap[shortUrl]; ok {
		return mainUrl, nil
	}
	return "", errors.New("Short URL not found")
}

// CreateShortenedUrl gets short url if old URL is passed, else generates a new short url
func (u *Url) CreateShortenedUrl(url string) (string, error) {
	if err := validateUrl(url); err != nil {
		return "", err
	}

	if shorturl, ok := u.MainToShortMap[url]; ok {
		return shorturl, nil
	}

	shorturl := generateShortUrl()
	u.MainToShortMap[url] = shorturl
	u.ShortToMainMap[shorturl] = url

	fmt.Printf("Original URL: %s, Shortened URL: %s\n", url, shorturl)
	err := u.Storage.StoreUrls(url, shorturl)
	if err != nil {
		return "", err
	}

	return shorturl, nil
}

func validateUrl(url string) error {
	if url == "" {
		return errors.New("Empty URL")
	}
	return nil
}

func generateShortUrl() string {
	shorturl := make([]byte, shortUrlIdLen)
	count := urlsCount

	for i := 0; i < shortUrlIdLen; i++ {
		char := charSet[count%charSetLen]
		shorturl[shortUrlIdLen-i-1] = char
		count = count / charSetLen
	}

	urlsCount++
	return string(shorturl)
}
