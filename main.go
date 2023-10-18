package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	user = flag.Bool("user", false, "Install in the current directory if set")
)

func main() {
	flag.Parse()

	version, err := readVersionFile()
	if err != nil {
		fmt.Println("Error reading .copilot-version:", err)
		os.Exit(1)
	}

	err = downloadAndInstallCopilot(version, *user)
	if err != nil {
		fmt.Println("Error downloading or installing copilot:", err)
		os.Exit(1)
	}
}

func readVersionFile() (string, error) {
	data, err := os.ReadFile(".copilot-version")
	if err != nil {
		return "", err
	}

	dataStr := string(data)

	return strings.TrimSpace(dataStr), nil
}

func downloadAndInstallCopilot(version string, local bool) error {
	url := fmt.Sprintf("https://github.com/aws/copilot-cli/releases/download/v%s/copilot-%s-%s", version, runtime.GOOS, runtime.GOARCH)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var installPath string
	if local {
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}
		installPath = fmt.Sprintf("%s/copilot", currentDir)
	} else {
		installPath = "/usr/local/bin/copilot"
	}

	out, err := os.Create(installPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	err = os.Chmod(installPath, 0755)
	if err != nil {
		return err
	}

	var cmd *exec.Cmd
	if local {
		cmd = exec.Command("./copilot", "-v")
	} else {
		cmd = exec.Command("copilot", "-v")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to verify copilot installation: %w\nOutput: %s", err, output)
	}

	fmt.Println(string(output))

	return nil
}
