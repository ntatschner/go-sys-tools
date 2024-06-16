package ports

import (
	"github.com/ntatschner/go-sys-tools/internal/core/domain"
)

type BackgroundPort interface {
	GetBackground(path string) (string, error)
}