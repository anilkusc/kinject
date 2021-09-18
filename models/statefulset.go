package models

type MyStatefulset struct {
	Name      string          `yaml:"name"`
	Namespace string          `yaml:"namespace"`
	Replicas  int             `yaml:"replicas"`
	Env       []MyEnvironment `yaml:"env"`
}
