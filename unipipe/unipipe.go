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
	uniconf.SetContexts("jobs.dev.jobs.install")
	uniconf.AddPhase(&uniconf.Phase{
		Name: "config",
		Phases: []*uniconf.Phase{
			{
				Name:     "load",
				Callback: uniconf.Load,
			},
			{
				Name:     "set_contexts",
				Callback: uniconf.ProcessContexts,
			},
			//{
			//	Name:     "print",
			//	Callback: uniconf.PrintConfig,
			//	Args: []interface{}{
			//		"jobs",
			//	},
			//},
			{
				Name:     "flatten",
				Callback: uniconf.FlattenConfig,
			},
			{
				Name:     "process",
				Callback: uniconf.ProcessKeys,
				Args: []interface{}{
					name,
					"jobs",
					[]*uniconf.Processor{
						//{
						//	Callback:    uniconf.InterpolateProcess,
						//	ExcludeKeys: []string{"from"},
						//},
						{
							Callback:    uniconf.FromProcess,
							IncludeKeys: []string{"from"},
						},
					},
				},
			},
		},
	})
	uniconf.Execute()

	job, _ := unitool.CollectInvertedKeyParamsFromJsonPath(u.config, name, "jobs")
	return unitool.MarshallYaml(job)
}