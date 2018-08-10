package pitreos

import (
	"cloud.google.com/go/storage"
)

type PITR struct {
	chunkSize int64
	threads   int
	Options   *PitreosOptions

	storageBucket *storage.BucketHandle
}

func New(opts *PitreosOptions) *PITR {
	return &PITR{
		chunkSize: 250 * 1024 * 1024,
		threads:   1,
		Options:   opts,
	}
}