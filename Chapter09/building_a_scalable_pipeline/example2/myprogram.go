package main

import (
	"log"
	"os"

	"github.com/pachyderm/pachyderm/src/client"
)

func main() {

	// Connect to Pachyderm on our localhost.  By default
	// Pachyderm will be exposed on port 30650.
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Start a commit in our "attributes" data repo on the "master" branch.
	commit, err := c.StartCommit("attributes", "master")
	if err != nil {
		log.Fatal(err)
	}

	// Open one of the attributes JSON files.
	f, err := os.Open("1.json")
	if err != nil {
		log.Fatal(err)
	}

	// Put a file containing the attributes into the data repository.
	if _, err := c.PutFile("attributes", commit.ID, "1.json", f); err != nil {
		log.Fatal(err)
	}

	// Finish the commit.
	if err := c.FinishCommit("attributes", commit.ID); err != nil {
		log.Fatal(err)
	}

	// Start a commit in our "training" data repo on the "master" branch.
	commit, err = c.StartCommit("training", "master")
	if err != nil {
		log.Fatal(err)
	}

	// Open up the training data set.
	f, err = os.Open("diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Put a file containing the training data set into the data repository.
	if _, err := c.PutFile("training", commit.ID, "diabetes.csv", f); err != nil {
		log.Fatal(err)
	}

	// Finish the commit.
	if err := c.FinishCommit("training", commit.ID); err != nil {
		log.Fatal(err)
	}
}
