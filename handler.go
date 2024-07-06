package main

import (
	"net/http"

	"encoding/json"

	"github.com/boltdb/bolt"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also

// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if path, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return

func YAMLHandler(yamldata []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsToUrls []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}
	if err := yaml.Unmarshal(yamldata, &pathsToUrls); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathsToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r, pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

// JSONHandler will parse the provided JSON and then return

func JSONHandler(jsondata []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsToUrls []struct {
		Path string `json:"path"`
		URL  string `json:"url"`
	}
	if err := json.Unmarshal(jsondata, &pathsToUrls); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, pathtourl := range pathsToUrls {
			if pathtourl.Path == r.URL.Path {
				http.Redirect(w, r, pathtourl.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

// BOLTHandler will use the provided BoltDB and then return

func BOLTHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("pathstourls"))
			if bucket != nil {
				cursor := bucket.Cursor()
				for path, url := cursor.First(); path != nil; path, url = cursor.Next() {
					if string(path) == r.URL.Path {
						http.Redirect(w, r, string(url), http.StatusFound)
						return nil
					}
				}
			}
			return nil
		}); err != nil {
			panic(err)
		}
		fallback.ServeHTTP(w, r)
	}
}
