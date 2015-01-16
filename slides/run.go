package main

import (
    "fmt"
    "github.com/fsouza/go-dockerclient"  // 'docker' package
)

func main() {
    endpoint := "unix:///var/run/docker.sock"
    client, _ := docker.NewClient(endpoint)

    c, _ := client.CreateContainer(docker.CreateContainerOptions{
        Name: "my-nginx",
        Config: &docker.Config{
            Image: "nginx",
        },
    })

    client.StartContainer(c.ID, &docker.HostConfig{
        PortBindings: map[docker.Port][]docker.PortBinding{
            "80/tcp": []docker.PortBinding{
                docker.PortBinding{HostPort: "8000"},
            },
        },
        Binds: []string{"/opt/workspace/src/github.com/IanMLewis/docker_meetup_slides:/usr/share/nginx/html:ro"},
    })

    fmt.Println("ID: ", c.ID)
    fmt.Println("Name: ", c.Name)
    fmt.Println("Created: ", c.Created)
}
