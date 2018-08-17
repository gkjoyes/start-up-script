package jobs

import (
	"os/exec"

	"github.com/g-kutty/start-up-script/logger"
)

// Execute runs each jobs.
func (j InputJobs) Execute() error {
	for _, job := range j.Jobs {
		// construct commands for jobs.
		cmds := []string{"-n", " ", "-r", job.Name, "-e"}
		if job.Path != "" {
			cmds = append(cmds, "cd "+job.Path)
		}

		// append commands and execute.
		cmds = append(cmds, job.Commands...)
		if err := exec.Command("guake", cmds...).Run(); err != nil {
			return err
		}

		// log job details.
		logger.Info().Message(job.Name, logger.FormattedMessage(cmds...)).Log()
	}
	return nil
}
