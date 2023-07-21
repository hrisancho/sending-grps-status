package server

import (
	"GSS/internal/client"
	"GSS/internal/metrics"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (server *Server) Hi(ctx *fiber.Ctx) (err error) {
	return ctx.SendString("Hi")
}

func (server *Server) addMetrics(ctx *fiber.Ctx) (err error) {
	var userMetricStorage client.UserMetricStorage

	if err := ctx.BodyParser(&userMetricStorage); err != nil {
		return err
	}

	user_uuid := userMetricStorage.UUID
	_, ok := server.UsersMetrics[user_uuid]
	if ok {
		timestamp := userMetricStorage.MetricStorage.Timestamp
		for idx, metric := range server.UsersMetrics[user_uuid] {
			if metric.Timestamp == timestamp {
				if len(userMetricStorage.MetricStorage.MetricMapFloat64) == 1 {
					for k, v := range userMetricStorage.MetricStorage.MetricMapFloat64 {
						server.UsersMetrics[user_uuid][idx].MetricMapFloat64[k] = v
					}
				} else if len(userMetricStorage.MetricStorage.MetricMapUint32) == 1 {
					for k, v := range userMetricStorage.MetricStorage.MetricMapUint32 {
						server.UsersMetrics[user_uuid][idx].MetricMapUint32[k] = v
					}
				} else if len(userMetricStorage.MetricStorage.MetricMapUint64) == 1 {
					for k, v := range userMetricStorage.MetricStorage.MetricMapUint64 {
						server.UsersMetrics[user_uuid][idx].MetricMapUint64[k] = v
					}
				}
				return
			}
		}

		server.UsersMetrics[user_uuid] = append(server.UsersMetrics[user_uuid], userMetricStorage.MetricStorage)
	} else {
		server.UsersMetrics[user_uuid] = []metrics.MetricStorage{userMetricStorage.MetricStorage}
	}

	return
}

func (server *Server) metricsByUUID(ctx *fiber.Ctx) (err error) {
	user_uuid, err := uuid.Parse(ctx.Params("uuid"))
	if err != nil {
		return
	}

	return ctx.JSON(server.UsersMetrics[user_uuid])
}

func (server *Server) getAllMetrics(ctx *fiber.Ctx) (err error) {
	return ctx.JSON(server.UsersMetrics)
}
