package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func runPlaybook() {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}

func runAdhoc() {
	ansibleConnectionOptions := &options.AnsibleConnectionOptions{}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		ModuleName: "ping",
	}

	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		//StdoutCallback:    "oneline",
	}

	log.Println("Command: ", adhoc)

	err := adhoc.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello World")
	runPlaybook()
	// runAdhoc()
}
