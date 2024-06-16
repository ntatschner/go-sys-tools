package adapters

import "C"

import (
	"github.com/ntatschner/go-sys-tools/internal/core/domain"
	"github.com/ntatschner/go-sys-tools/internal/core/ports/primary"
)

type BackgroundAdapter struct {
	BackgroundPort primary.BackgroundPort
}

func (b BackgroundAdapter) GetBackground(path string) (domain.Background, error) {
	return b.BackgroundPort.GetBackground(path)
}
