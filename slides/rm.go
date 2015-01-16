package main

import (
    "fmt"
    "github.com/fsouza/go-dockerclient"  // 'docker' package
)

func main() {
    endpoint := "unix:///var/run/docker.sock"
    client, _ := docker.NewClient(endpoint)

    clist, _ := client.ListContainers(docker.ListContainersOptions{All: false})
    for _, c := range clist {
        client.StopContainer(c.ID, 30)
        client.RemoveContainer(docker.RemoveContainerOptions{
            ID: c.ID,
            RemoveVolumes: true,
        })
        fmt.Println("ID: ", c.ID)
    }
}
