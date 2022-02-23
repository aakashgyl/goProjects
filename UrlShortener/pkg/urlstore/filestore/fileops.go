package filestore

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aakashgyl/UrlShortener/pkg/urlstore"
)

const URL_SEPARATOR = "#URL-SEPARATOR#"

type urlFileStorage struct {
	File string
}

// GetUrlFileStoreServiceObj implements UrlStore to persist URL data in file
func GetUrlFileStoreServiceObj(filename string) (urlstore.UrlStore, error) {
	if filename == "" {
		return nil, errors.New("empty filename")
	}
	return urlFileStorage{File: filename}, nil
}

func (u urlFileStorage) StoreUrls(url string, shorturl string) error {
	if url == "" || shorturl == "" {
		return errors.New("urls cannot be empty")
	}

	file, err := os.OpenFile(u.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(url + URL_SEPARATOR + shorturl + "\n")
	if err != nil {
		return err
	}

	return nil
}

func (u urlFileStorage) LoadUrls() (map[string]string, map[string]string, error) {
	urlMap := make(map[string]string)
	idToMainUrlMap := make(map[string]string)

	if _, err := os.Stat(u.File); err != nil {
		fmt.Println("Error -> ", err.Error())
		return urlMap, idToMainUrlMap, err
	}

	file, err := os.OpenFile(u.File, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return urlMap, idToMainUrlMap, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			return urlMap, idToMainUrlMap, nil
		}

		if err != nil {
			return urlMap, idToMainUrlMap, err
		}

		urlLine := strings.Split(string(line), URL_SEPARATOR)

		if len(urlLine) != 2 {
			return urlMap, idToMainUrlMap, fmt.Errorf("invalid data (%s) in file", line)
		}
		urlMap[urlLine[0]] = urlLine[1]
		idToMainUrlMap[urlLine[1]] = urlLine[0]
	}
}
