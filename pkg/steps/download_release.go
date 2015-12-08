package steps

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mfojtik/oinc/pkg/log"
	"github.com/mfojtik/oinc/pkg/util"
)

const (
	ReleaseURL = "https://api.github.com/repos/openshift/origin/releases/latest"
)

type Asset struct {
	URL string `json:"browser_download_url"`
}

type Release struct {
	Assets []Asset `json:"assets"`
}

type DownloadReleaseStep struct {
	ParallelStep
}

func (*DownloadReleaseStep) String() string { return "download-release" }

func (*DownloadReleaseStep) Execute() error {
	log.Info("Creating %q directory ...", BaseDir)
	if err := os.MkdirAll(BaseDir, 0770); err != nil {
		return err
	}
	log.Info("Downloading Origin release archive ...")
	resp, err := http.Get(ReleaseURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(strings.NewReader(string(body)))
	var rel Release
	for {
		if err := dec.Decode(&rel); err == io.EOF {
			break
		} else if err != nil {
			log.Error("Failed to decode GitHub release JSON: %v", err)
		}
	}
	downloadUrl := ""
	for _, asset := range rel.Assets {
		if util.IsDarwin() {
			if strings.Contains(asset.URL, "darwin-amd64") {
				downloadUrl = asset.URL
			}
		} else {
			if strings.Contains(asset.URL, "linux-amd64") {
				downloadUrl = asset.URL
			}
		}
	}
	if err := downloadTar(downloadUrl, filepath.Join(BaseDir, "release.tar.gz")); err != nil {
		return err
	}
	//defer os.Remove(filepath.Join(BaseDir, "release.tar.gz"))
	log.Debug("Extracting release tar to %q", filepath.Join(BaseDir, "bin"))
	if err := os.MkdirAll(filepath.Join(BaseDir, "bin"), 0700); err != nil {
		return err
	}
	_, err = util.RunCommand("tar", "-x", "-z", "-f", filepath.Join(BaseDir, "release.tar.gz"), "-C", filepath.Join(BaseDir, "bin"))
	if err != nil {
		log.Info("OpenShift commands installed to %q", filepath.Join(BaseDir, "bin"))
	}
	return err
}

func downloadTar(url, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
