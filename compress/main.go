package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	threshold int64 = 1024 * 1024
)

func main() {
	dir := "./"
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if path == dir || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".jpg") || strings.HasSuffix(info.Name(), ".png") {
			if info.Size() < threshold {
				return nil
			}
			fmt.Printf("processing path: %s ,size: %fM\n", path, float64(info.Size())/float64(1024*1024))
			// 复制源文件
			newpath := strings.TrimSuffix(path, filepath.Ext(path)) + "_copy" + filepath.Ext(path)

			// 压缩
			// yum install -y ImageMagick
			cmd := exec.Command("convert", path, "-quality", "10%", newpath)
			_, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
				return err
			}

			// 删除源文件
			err = os.Remove(path)
			if err != nil {
				return fmt.Errorf("failed to remove file: %s", path)
			}

			// 更新复制的文件名
			err = os.Rename(newpath, path)
			if err != nil {
				return fmt.Errorf("failed to rename file: %s", path)
			}
		}
		return nil
	})

}
