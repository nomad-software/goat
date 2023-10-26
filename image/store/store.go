package store

import (
	"embed"
	"encoding/base64"
	"path/filepath"

	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/image/gif"
	"github.com/nomad-software/goat/image/png"
	"github.com/nomad-software/goat/log"
)

// Store is an image store of embedded images.
type Store struct {
	fs embed.FS
}

// New creates a new image store.
func New(fs embed.FS) *Store {
	st := &Store{
		fs: fs,
	}

	return st
}

// GetImage gets an image from the store.
// Only Png and Gif formats are supported.
func (s *Store) GetImage(name string) image.Image {
	b, err := s.fs.ReadFile(name)
	if err != nil {
		log.Error(err)
		panic("cannot read file")
	}

	str := base64.StdEncoding.EncodeToString(b)
	ext := filepath.Ext(name)

	switch ext {
	case "gif":
		return gif.New(str)
	case "png":
		return png.New(str)
	default:
		panic("image extension not recognised")
	}
}
