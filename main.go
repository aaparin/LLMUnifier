package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	RootDir       string   `yaml:"root_dir"`
	CommentSymbol string   `yaml:"comment_symbol"`
	Paths         []string `yaml:"paths"`
	OutputFile    string   `yaml:"output_file"`
}

func main() {
	config, err := loadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	if config.CommentSymbol == "" {
		config.CommentSymbol = "#"
	}

	if config.OutputFile == "" {
		config.OutputFile = "combined.txt"
	}

	content := ""
	for _, path := range config.Paths {
		fullPath := filepath.Join(config.RootDir, path)
		if strings.HasPrefix(path, "!") {
			continue // Исключаем пути, начинающиеся с '!'
		}

		if strings.Contains(path, "*") {
			err := filepath.Walk(filepath.Dir(fullPath), func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				if shouldIncludeFile(path, config.Paths, config.RootDir) && strings.HasSuffix(path, filepath.Ext(fullPath)) {
					fileContent, err := readFile(path)
					if err != nil {
						fmt.Printf("Error reading file %s: %v\n", path, err)
						return nil
					}
					relPath, _ := filepath.Rel(config.RootDir, path)
					content += fmt.Sprintf("%s %s\n%s\n\n", config.CommentSymbol, relPath, fileContent)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Error walking directory %s: %v\n", fullPath, err)
			}
		} else {
			if shouldIncludeFile(fullPath, config.Paths, config.RootDir) {
				fileContent, err := readFile(fullPath)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", fullPath, err)
					continue
				}
				relPath, _ := filepath.Rel(config.RootDir, fullPath)
				content += fmt.Sprintf("%s %s\n%s\n\n", config.CommentSymbol, relPath, fileContent)
			}
		}
	}

	err = writeToFile(config.OutputFile, content)
	if err != nil {
		fmt.Printf("Error writing combined file: %v\n", err)
		return
	}

	fmt.Printf("Files combined successfully into %s!\n", config.OutputFile)
}

func loadConfig(filename string) (Config, error) {
	var config Config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func readFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func writeToFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

func shouldIncludeFile(file string, paths []string, rootDir string) bool {
	relPath, _ := filepath.Rel(rootDir, file)
	for _, path := range paths {
		if strings.HasPrefix(path, "!") {
			excludePath := strings.TrimPrefix(path, "!")
			if strings.HasPrefix(relPath, excludePath) {
				return false
			}
		}
	}
	return true
}
