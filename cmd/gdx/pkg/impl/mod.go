package impl

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/JiepengTan/godotgo/cmd/gdx/pkg/util"
)

func getLatestCommitHash() (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=format:%H")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error getting latest commit hash: %v", err)
	}
	return out.String(), nil
}
func getCommitTimestamp(commitHash string) (string, error) {
	cmd := exec.Command("git", "show", "-s", "--format=%ci", commitHash)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error getting commit UTC timestamp: %v", err)
	}
	return strings.TrimSpace(out.String()), nil
}
func formatGoVersionTimestamp(utcTimestamp string) (string, error) {
	t, err := time.Parse("2006-01-02 15:04:05 -0700", utcTimestamp)
	if err != nil {
		return "", fmt.Errorf("error parsing time: %v", err)
	}
	return t.UTC().Format("20060102150405"), nil
}
func getGoPackageVersion(commitHash string) (string, error) {
	timestamp, err := getCommitTimestamp(commitHash)
	if err != nil {
		fmt.Println(err)
	}
	timestamp, err = formatGoVersionTimestamp(timestamp)
	if err != nil {
		fmt.Println(err)
	}
	shortHash := commitHash[:12]

	version := fmt.Sprintf("v0.0.0-%s-%s", timestamp, shortHash)
	fmt.Println("Version:", version)
	return version, nil
}

func modifyFile(filePath, tag, modStr string) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer inputFile.Close()

	var outputLines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, tag) {
			line = modStr
		}
		outputLines = append(outputLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range outputLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", filePath, err)
			return
		}
	}
	writer.Flush()
}
func replaceMod(tag string, version string, dirList []string, relDir string, fileList []string) {
	modStr := tag + version
	for _, dir := range dirList {
		modPath := path.Join(relDir, dir, "go.mod")
		modifyFile(modPath, tag, modStr)
	}

	for _, file := range fileList {
		modPath := path.Join(relDir, file)
		modifyFile(modPath, tag, modStr)
	}
}
func directoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func UpdateMod() {
	commitHash, err := getLatestCommitHash()
	if err != nil {
		fmt.Println(err)
		return
	}

	version, err := getGoPackageVersion(commitHash)
	if err != nil {
		fmt.Println(err)
		return
	}

	tag := "github.com/JiepengTan/godotgo "
	relDir := "./spx"

	if exists, err := directoryExists(relDir); err != nil || !exists {
		fmt.Println("Error checking directory ", relDir)
		return
	}

	dirList := []string{}
	fileList := []string{"cmd/gdx/config/project//scriptsgo.mod"}

	replaceMod(tag, version, dirList, relDir, fileList)

	// run go mod tidy
	rawDir, _ := os.Getwd()
	for _, dir := range dirList {
		modPath := path.Join(relDir, dir)
		os.Chdir(modPath)
		if err := util.RunGoModTidy(); err != nil {
			fmt.Println(err)
		}
		os.Chdir(rawDir)
	}
}
