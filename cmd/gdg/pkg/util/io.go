package util

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
func SetupFile(force bool, name, embed string, args ...any) error {
	if _, err := os.Stat(name); force || os.IsNotExist(err) {
		if len(args) > 0 {
			embed = fmt.Sprintf(embed, args...)
		}
		if err := os.WriteFile(name, []byte(embed), 0644); err != nil {
			return err
		}
	}
	return nil
}

// Copy a file
func CopyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst, input, 0755); err != nil {
		return err
	}
	return nil
}

func CopyDir(fsys fs.FS, srcDir, dstDir string) error {
	if _, err := os.Stat(dstDir); !os.IsNotExist(err) {
		println("Error: Directory already exists: ", dstDir)
		return nil
	}

	subfs, err := fs.Sub(fsys, srcDir)
	if err != nil {
		println("Error: create sub fs: ", dstDir)
		return err
	}
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		println("Error: creating directory: ", dstDir)
		return err
	}
	return fs.WalkDir(subfs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			println("Error: walking directory: ", srcDir)
			return err
		}

		dstPath := filepath.Join(dstDir, path)
		if d.IsDir() {
			return os.MkdirAll(dstPath, 0755)
		} else {
			srcFile, err := subfs.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			dstFile, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			_, err = io.Copy(dstFile, srcFile)
			return err
		}
	})
}

func DownloadFile(url string, dest string) error {
	println("Downloading file... ", url, "=>", dest)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func Unzip(zipfile, dest string) {
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			panic(err)
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			var dir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				dir = fpath[:lastIndex]
			}

			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(outFile, rc)
			outFile.Close()

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
