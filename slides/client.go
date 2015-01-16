package main

// 'docker' package
import "github.com/fsouza/go-dockerclient"

func main() {
    endpoint := "unix:///var/run/docker.sock"
    client, _ := docker.NewClient(endpoint)
    //...
}
