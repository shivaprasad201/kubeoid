package cmd

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func buildImage(buildContextDir, tagName string) {

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	HandleError(err, ":unable to init client")

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	scanDir(buildContextDir, tw)

	tarReader := bytes.NewReader(buf.Bytes())

	imageBuildResponse, err := cli.ImageBuild(
		ctx,
		tarReader,
		types.ImageBuildOptions{
			Context:    tarReader,
			Dockerfile: "Dockerfile",
			Tags:       []string{tagName},
			Remove:     true})
	HandleError(err, " :unable to build docker image")

	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	HandleError(err, " :unable to read image build response")
}

func scanDir(dirPath string, tw *tar.Writer) {
	dir, err := os.Open(dirPath)
	HandleError(err, " :unable open the given path for reading")
	defer dir.Close()

	fis, err := dir.Readdir(0)
	HandleError(err, " :unable to read given directory")

	for _, fi := range fis {
		curPath := dirPath + "/" + fi.Name()
		if fi.IsDir() {
			scanDir(curPath, tw)
		} else {
			tarWrite(curPath, tw, fi)
		}
	}
}

func tarWrite(path string, tw *tar.Writer, fi os.FileInfo) {
	fr, err := os.Open(path)
	HandleError(err, " :unable to read path for writing tarball")
	defer fr.Close()

	h := new(tar.Header)
	h.Name = path
	h.Size = fi.Size()
	h.Mode = int64(fi.Mode())
	h.ModTime = fi.ModTime()

	err = tw.WriteHeader(h)
	HandleError(err, " :unable to write tarball header")

	_, err = io.Copy(tw, fr)
	HandleError(err, " :unable to copy files to tarball")
}
