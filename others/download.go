package others

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func Download() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/denoland/deno/releases/latest", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GITHUB"))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	// Extract the download URL for the desired asset
	assets := result["assets"].([]interface{})
	var downloadURL string
	for _, asset := range assets {
		assetMap := asset.(map[string]interface{})
		if assetMap["name"].(string) == "deno-x86_64-unknown-linux-gnu.zip" { // Change this to the desired asset name
			downloadURL = assetMap["browser_download_url"].(string)
			break
		}
	}

	// Download the release asset
	resp, err = http.Get(downloadURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create the output file
	out, err := os.Create("latest_release.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Copy the response body to the output file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latest release downloaded successfully.")

	// Check if the "deno" binary exists in the destination directory
	if _, err := os.Stat("deno"); os.IsNotExist(err) {
		// Unzip the downloaded file
		err = unzip("latest_release.zip", ".")
		if err != nil {
			log.Fatal(err)
		}

		// Change permissions of the binary called "deno"
		err = exec.Command("chmod", "+x", "deno").Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Permissions changed successfully.")
	} else {
		fmt.Println("Binary already exists, skipping download and unzip.")
	}
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
