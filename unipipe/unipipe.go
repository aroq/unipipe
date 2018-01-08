package unipipe

import (
	"github.com/aroq/uniconf/uniconf"
	"github.com/aroq/uniconf/unitool"
)

type Unipipe struct {
	uniconf.Uniconf
}

var u *Unipipe

// Init initializes unipipe.
func init() {
	u = New()
}

// New returns an initialized Uniconf instance.
func New() *Unipipe {
	u := new(Unipipe)
	return u
}

func Job(name string) string { return u.job(name) }
func (u *Unipipe) job(name string) string {
	config := uniconf.Config().(map[string]interface{})
	uniconf.Process(config["jobs"], "jobs", "config")
	job, _ := unitool.CollectInvertedKeyParamsFromJsonPath(config, name, "jobs")
	return unitool.MarshallYaml(job)
}