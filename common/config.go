package common

type Config struct {
	Redis struct {
		Addrs  string
		Passwd string
		Db     int
	} `yaml:"redis"`
	PgDB struct {
		Host   string
		Port   int
		User   string
		Passwd string
		Db     string
	} `yaml:"postgres"`
	Etcd struct {
		Addrs    string
		Resolver string
	} `yaml:"etcd"`
}
