package app

type ObjectMeta struct {
	Name            string `json:"name,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

type Service struct {
	ObjectMeta `json:"metadata,omitempty"`
}

type Pod struct {
	ObjectMeta `json:"metadata,omitempty"`
}

type Deployment struct {
	ObjectMeta `json:"metadata,omitempty"`
}

//Config represent configuration written in .kube/config file
type Config struct {
	Clusters       []ClusterWrap `yaml:"clusters"`
	Contexts       []ContextWrap `yaml:"contexts"`
	CurrentContext string        `yaml:"current-context"`
}

type Cluster struct {
	Server string `yaml:"server"`
}

type ClusterWrap struct {
	Name    string  `yaml:"name"`
	Cluster Cluster `yaml:"cluster"`
}

type Context struct {
	Cluster   string `yaml:"cluster"`
	Namespace string `yaml:"namespace"`
}

type ContextWrap struct {
	Name    string  `json:"name"`
	Context Context `yaml:"context"`
}
