package server

func (server Server) SetupRoutes() {
	server.App.Get("/", server.getAllMetrics)
	server.App.Post("/", server.addMetrics)

	server.App.Get("/:uuid", server.metricsByUUID)
}
