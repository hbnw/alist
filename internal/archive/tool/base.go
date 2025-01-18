package tool

import (
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/stream"
	"io"
)

type Tool interface {
	AcceptedExtensions() []string
	GetMeta(ss *stream.SeekableStream, args model.ArchiveArgs) (model.ArchiveMeta, error)
	List(ss *stream.SeekableStream, args model.ArchiveInnerArgs) ([]model.Obj, error)
	Extract(ss *stream.SeekableStream, args model.ArchiveInnerArgs) (io.ReadCloser, int64, error)
	Decompress(ss *stream.SeekableStream, outputPath string, args model.ArchiveInnerArgs, up model.UpdateProgress) error
}