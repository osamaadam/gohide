package hide

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"

	"github.com/pkg/errors"
)

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func isLinux() bool {
	return runtime.GOOS == "linux"
}

func Hide(files ...string) error {
	if isWindows() {
		return hideWindows(false, files...)
	} else if isLinux() {
		return hideLinux(files...)
	} else {
		return errors.WithStack(errors.Errorf("%s is not a supported os", runtime.GOOS))
	}
}

func Unhide(files ...string) error {
	if isWindows() {
		return hideWindows(true, files...)
	} else if isLinux() {
		return unhideLinux(files...)
	} else {
		return errors.WithStack(errors.Errorf("%s is not a supported os", runtime.GOOS))
	}
}

func unhideLinux(files ...string) error {
	// TODO
	return nil
}

func hideLinux(files ...string) error {
	for _, file := range files {
		filename := filepath.Base(file)
		filedir := filepath.Dir(file)

		if !strings.HasPrefix(filename, ".") {
			hiddenFileName := fmt.Sprintf(".%s", filename)
			hiddenFilePath := filepath.Join(filedir, hiddenFileName)

			if err := os.Rename(file, hiddenFilePath); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

func hideWindows(unhide bool, files ...string) error {
	for _, file := range files {
		filename, err := syscall.UTF16PtrFromString(file)
		if err != nil {
			return errors.WithStack(err)
		}
		var attr uint32 = syscall.FILE_ATTRIBUTE_HIDDEN

		if unhide {
			attr = syscall.FILE_ATTRIBUTE_NORMAL
		}

		if err := syscall.SetFileAttributes(filename, attr); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
