package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// InputJobs struct.
type InputJobs struct {
	Jobs []Job `json:"jobs"`
}

// Job represent job structure.
type Job struct {
	Name         string   `json:"name"`
	Commands     []string `json:"commands"`
	Path         string   `json:"path"`
	GitPull      bool     `json:"git_pull"`
	SourceBranch string   `json:"source_branch"`
	DestBranch   string   `json:"dest_branch"`
}

// Reader struct.
type Reader struct {
	Path string `json:"Path"`
}

// NewReader Create new reader and return.
func NewReader(loc string) *Reader {
	return &Reader{
		Path: loc,
	}
}

// Read input file and process intput data.
func (r *Reader) Read() (*InputJobs, error) {

	// read json input file.
	jobs, err := ioutil.ReadFile(r.Path)
	if err != nil {
		return nil, err
	}

	// unmarshal into jobs struct.
	var data InputJobs
	err = json.Unmarshal(jobs, &data)
	if err != nil {
		return nil, err
	}

	// validate each jobs.
	err = validateJobs(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// check each jobs has valid input params.
func validateJobs(jobs *InputJobs) error {
	for _, job := range jobs.Jobs {
		if strings.TrimSpace(job.Path) != "" {
			if _, err := os.Stat(job.Path); err != nil {
				return fmt.Errorf("Job[%s] has invalid path [%s]", job.Name, job.Path)
			}
		}
	}
	return nil
}
