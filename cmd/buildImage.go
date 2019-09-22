package cmd

import (
	"bufio"
	"io"
	"log"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

//BuildImage builds the container image
func BuildImage(dockerBuildCtxDir, tagName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(300)*time.Second)
	defer cancel()
	dockerFileTarReader := strings.NewReader(dockerBuildCtxDir)

	// buildArgs := make(map[string]*string)
	// add any build args if you want to
	// buildArgs["ENV"] = os.Getenv("GO_ENV")

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	resp, err := cli.ImageBuild(
		ctx,
		dockerFileTarReader,
		types.ImageBuildOptions{
			Dockerfile: "Dockerfile",
			Tags:       []string{tagName},
			NoCache:    true,
			Remove:     true,
			// BuildArgs:  buildArgs,
		}) //cli is the docker client instance created from the engine-api
	if err != nil {
		log.Println(dockerFileTarReader)
		log.Println(err, " :unable to build docker image")
		return err
	}
	return writeToLog(resp.Body)
}

//writes from the build response to the log
func writeToLog(reader io.ReadCloser) error {
	defer reader.Close()
	rd := bufio.NewReader(reader)
	for {
		n, _, err := rd.ReadLine()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		log.Println(string(n))
	}
	return nil
}
