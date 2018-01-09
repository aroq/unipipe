package unipipe

import (
	"github.com/aroq/uniconf/uniconf"
	"github.com/aroq/uniconf/unitool"
)

type Unipipe struct {
	config map[string]interface{}
}

var u *Unipipe

// init initializes unipipe.
func init() {
	u = New()
	u.config = uniconf.Config().(map[string]interface{})
}

// New returns an initialized Unipipe instance.
func New() *Unipipe {
	u := new(Unipipe)
	return u
}

// Job returns a processed Job config by name.
func Job(name string) string { return u.job(name) }
func (u *Unipipe) job(name string) string {
	uniconf.Process(u.config["jobs"], "jobs", "config")
	job, _ := unitool.CollectInvertedKeyParamsFromJsonPath(u.config, name, "jobs")
	return unitool.MarshallYaml(job)
}