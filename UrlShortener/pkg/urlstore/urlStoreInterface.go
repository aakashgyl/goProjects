package urlstore

// UrlStore can be implemented by different storage types (eg, DB, File etc)
type UrlStore interface {
	StoreUrls(string, string) error
	LoadUrls() (map[string]string, map[string]string, error)
}
