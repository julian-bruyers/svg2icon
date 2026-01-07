// Package png provides SVG to PNG conversion functionality for icon generation.
//
// This package handles the rasterization of SVG files into PNG format at specific
// pixel dimensions, serving as the foundation for ICO and ICNS icon generation.
package png

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/kanrichan/resvg-go"
)

var (
	contextInstance *resvg.Context
	renderer        *resvg.Renderer
	initOnce        sync.Once
	initErr         error
	initMu          sync.Mutex
)

func initializeRenderer() error {
	initOnce.Do(func() {
		ctx, err := resvg.NewContext(context.Background())
		if err != nil {
			initErr = fmt.Errorf("failed to create resvg context: %w", err)
			return
		}
		contextInstance = ctx

		rend, err := ctx.NewRenderer()
		if err != nil {
			initErr = fmt.Errorf("failed to create resvg renderer: %w", err)
			return
		}
		renderer = rend
	})
	return initErr
}

// SvgToPng converts an SVG file to PNG format at the specified pixel size.
//
// The function rasterizes the SVG using vector graphics processing to produce
// high-quality PNG output suitable for icon generation. The SVG is scaled to
// fit exactly within the specified square dimensions.
//
// Parameters:
//   - svgPath: Path to the source SVG file
//   - pxSize: Output dimensions in pixels (width and height)
//
// Returns the PNG-encoded image data as bytes, or an error if conversion fails.
func SvgToPng(svgPath string, pxSize int) ([]byte, error) {
	err := initializeRenderer()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize resvg renderer: %w", err)
	}

	svgData, err := os.ReadFile(svgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read svg file: %w", err)
	}




	renderer.SetDefaultSize(float32(pxSize), float32(pxSize))

	pngData, err := renderer.Render(svgData)
	if err != nil {
		return nil, fmt.Errorf("resvg render failed: %w", err)
	}

	return pngData, nil
}
