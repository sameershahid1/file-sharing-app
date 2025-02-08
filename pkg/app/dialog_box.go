package app

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) OpenFileDialogBox() (string, error) {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "All Files",
				Pattern:     "*",
			},
		},
	})

	if err != nil {
		return "", err
	}

	return selection, nil
}
