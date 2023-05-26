package domain

// At the moment Catalog is the only addition of this module, but in other services it is possible that there is more than
// one, for that reason a generic solution.

var aggregates = newAggregateRegistry()

func newAggregateRegistry() *aggregateRegistry {
	return &aggregateRegistry{
		Catalog: "catalog.Catalog",
		//Other:   "catalog.Other",
		//Another: "catalog.Another",
	}
}

type aggregateRegistry struct {
	Catalog string
	//Other   string
	//Another string
}
