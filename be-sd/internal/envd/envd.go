package envd

import "syscall"

// BackendServices keeps information about back-end services.
type BackendServices struct {
	SD  string `json:"sd"`
	API string `json:"api"`
}

// FrontendServices keeps information about front-end services.
type FrontendServices struct {
	Web string `json:"web"`
}

// InfraServices keeps information about infrastructure services.
type InfraServices struct {
	Mongo         string `json:"mongo"`
	Elasticsearch string `json:"elasticsearch"`
	Kibana        string `json:"kibana"`
	Grafana       string `json:"grafana"`
}

// Services keeps information about group of service components.
type Services struct {
	Backend  BackendServices  `json:"backend"`
	Frontend FrontendServices `json:"frontend"`
	Infra    InfraServices    `json:"infra"`
}

// SD represents final service information distinguished by internal and external services.
type SD struct {
	Port     string   `json:"-"`
	Internal Services `json:"internal"`
	Public   Services `json:"public"`
}

func internalServices() (Services, error) {
	evBackendSD := "KN_INT_BE_SD"
	evBackendAPI := "KN_INT_BE_API"
	evFrontendWeb := "KN_INT_FE_WEB"
	evInfraMongo := "KN_INT_INFRA_MONGO"
	evInfraElasticsearch := "KN_INT_INFRA_ELASTICSEARCH"
	evInfraKibana := "KN_INT_INFRA_KIBANA"
	evInfraGrafana := "KN_INT_INFRA_GRAFANA"

	sd, _ := syscall.Getenv(evBackendSD)
	api, _ := syscall.Getenv(evBackendAPI)
	web, _ := syscall.Getenv(evFrontendWeb)
	mdb, _ := syscall.Getenv(evInfraMongo)
	es, _ := syscall.Getenv(evInfraElasticsearch)
	kibana, _ := syscall.Getenv(evInfraKibana)
	grafana, _ := syscall.Getenv(evInfraGrafana)

	return Services{
		Backend: BackendServices{
			SD:  sd,
			API: api,
		},
		Frontend: FrontendServices{
			Web: web,
		},
		Infra: InfraServices{
			Mongo:         mdb,
			Elasticsearch: es,
			Kibana:        kibana,
			Grafana:       grafana,
		},
	}, nil
}

func publicServices() (Services, error) {
	evBackendSD := "KN_PUB_BE_SD"
	evBackendAPI := "KN_PUB_BE_API"
	evFrontendWeb := "KN_PUB_FE_WEB"
	evInfraMongoDB := "KN_PUB_INFRA_MONGO"
	evInfraElasticsearch := "KN_PUB_INFRA_ELASTICSEARCH"
	evInfraKibana := "KN_PUB_INFRA_KIBANA"
	evInfraGrafana := "KN_PUB_INFRA_GRAFANA"

	sd, _ := syscall.Getenv(evBackendSD)
	api, _ := syscall.Getenv(evBackendAPI)
	web, _ := syscall.Getenv(evFrontendWeb)
	mdb, _ := syscall.Getenv(evInfraMongoDB)
	es, _ := syscall.Getenv(evInfraElasticsearch)
	kibana, _ := syscall.Getenv(evInfraKibana)
	grafana, _ := syscall.Getenv(evInfraGrafana)

	return Services{
		Backend: BackendServices{
			SD:  sd,
			API: api,
		},
		Frontend: FrontendServices{
			Web: web,
		},
		Infra: InfraServices{
			Mongo:         mdb,
			Elasticsearch: es,
			Kibana:        kibana,
			Grafana:       grafana,
		},
	}, nil
}

// NewSD harvests specific env variables and returns a new instance of SD.
func NewSD() (*SD, error) {
	evSDPort := "KN_BE_SD_PORT"
	sdPort, _ := syscall.Getenv(evSDPort)

	is, _ := internalServices()
	ps, _ := publicServices()

	sd := SD{
		Port:     sdPort,
		Internal: is,
		Public:   ps,
	}

	return &sd, nil
}
