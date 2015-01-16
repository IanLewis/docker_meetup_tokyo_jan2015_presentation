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
        fmt.Println("ID: ", c.ID)
        fmt.Println("Name: ", c.Names)
        fmt.Println("Created: ", c.Created)
    }
}
