package models

type MyDeployment struct {
	Name      string          `yaml:"name"`
	Namespace string          `yaml:"namespace"`
	Replicas  int             `yaml:"replicas"`
	Env       []MyEnvironment `yaml:"env"`
}
