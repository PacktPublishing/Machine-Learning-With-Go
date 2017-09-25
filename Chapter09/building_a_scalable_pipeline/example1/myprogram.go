package main

import (
	"log"

	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm using the IP of our
	// Kubernetes cluster. Here we will use  localhost
	// to mimic the sceneario when you have k8s running
	// locally and/or you are forwarding the Pachyderm
	// port to your localhost.. By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a data repository called "training."
	if err := c.CreateRepo("training"); err != nil {
		log.Fatal(err)
	}

	// Create a data repository called "attributes."
	if err := c.CreateRepo("attributes"); err != nil {
		log.Fatal(err)
	}

	// Now, we will list all the current data repositories
	// on the Pachyderm cluster as a sanity check. We
	// should now have two data repositories.
	repos, err := c.ListRepo(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Check that the number of repos is what we expect.
	if len(repos) != 2 {
		log.Fatal("Unexpected number of data repositories")
	}

	// Check that the name of the repo is what we expect.
	if repos[0].Repo.Name != "attributes" || repos[1].Repo.Name != "training" {
		log.Fatal("Unexpected data repository name")
	}
}
