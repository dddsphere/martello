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

type (
	CatalogStatus string
)

var status = newStatusRegistry()

func newStatusRegistry() *statusRegistry {
	return &statusRegistry{
		Unknown:            "unknown",
		Initialized:        "initialized",
		WaitingForApproval: "waiting-for-approval",
		Approved:           "approved",
		Active:             "status",
		Paused:             "paused",
		Deprecated:         "deprecated",
		Archived:           "archived",
	}
}

type statusRegistry struct {
	Unknown            CatalogStatus
	Initialized        CatalogStatus
	WaitingForApproval CatalogStatus
	Approved           CatalogStatus
	Active             CatalogStatus
	Paused             CatalogStatus
	Deprecated         CatalogStatus
	Archived           CatalogStatus
}

func (cs CatalogStatus) Name() string {
	return cs.String()
}

func (cs CatalogStatus) String() string {
	return string(cs)
}
