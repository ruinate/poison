// Package utils -----------------------------
// @file      : seaceh.go
// @author    : fzf
// @time      : 2023/6/5 下午1:07
// -------------------------------------------
package utils

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func FindAllFiles(path string) []string {
	file := make([]string, 0)
	if filepath.Ext(path) == ".pcap" || filepath.Ext(path) == ".pcapng" {
		file = append(file, path)
		return file
	}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是 pcap文件  或者pcapng 则添加
		if !info.IsDir() && filepath.Ext(path) == ".pcap" || !info.IsDir() && filepath.Ext(path) == ".pcapng" {
			file = append(file, path)
		}
		return nil
	})
	if err != nil {
		logger.Fatalln(err)
	}
	return file
}
