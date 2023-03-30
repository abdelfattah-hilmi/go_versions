package main

import (
	"context"
	"time"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/go-co-op/gocron"
)

func runPlaybook() {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"playbook.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}

func runCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Weeks().Do(
		func() {
			runPlaybook()
		})
	s.StartBlocking()
}

func main() {

	runPlaybook()
	//! runCronJobs() This runs once a week
}
