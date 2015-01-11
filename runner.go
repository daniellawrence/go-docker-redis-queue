package main

import (
	"fmt"
	"github.com/adeven/redismq"
	"github.com/samalba/dockerclient"
)

func main() {
	testQueue := redismq.CreateQueue("localhost", "6379", "", 9, "goqueue")

	consumer, err := testQueue.AddConsumer("testconsumer")
	
	if err != nil {
		panic(err)
	}
	containerConfig := &dockerclient.ContainerConfig{Image: "ubuntu:12.04",
		Cmd: []string{"date"},
	}
	hostConfig := &dockerclient.HostConfig{PublishAllPorts: true} 
	logconfig := &dockerclient.LogOptions{Follow: true, Stdout: false, Stderr: false}

	fmt.Println("Reading all messages on redis://localhost:6379/goqueue")

	for {
		p, err := consumer.Get()

		if err != nil {
			fmt.Println(err)
			continue
		}

		docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)


		// Create a container
		containerId, err := docker.CreateContainer(containerConfig, "")
		if err != nil {
			fmt.Println(err)
		}


		// Start the container
		err = docker.StartContainer(containerId, hostConfig)
		if err != nil {
			fmt.Println(err)
		}

			
		// read continer logs
		output, err := docker.ContainerLogs(containerId, logconfig)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("OUT: %s\n", output)

		// pop the job off the queue
		err = p.Ack()
		if err != nil {
			fmt.Println(err)
		}
	}
	
}
