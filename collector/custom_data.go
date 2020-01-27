package collector

type CustomCollector interface {
	Collect()
	Values() map[string]interface{}
}
