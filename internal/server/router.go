package server

func (server Server) SetupRoutes() {
	server.App.Get("/", server.Hi)
	server.App.Post("/", server.addMetrics)

	server.App.Get("/metrics", server.getAllMetrics)
	server.App.Get("/metrics/:uuid", server.metricsByUUID)
}
