package client

import (
	"GSS/internal/client/config"
	"GSS/internal/metrics"
	pb "GSS/proto/grpc"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func StateFunCli(config config.Config) {
	conn, err := grpc.Dial(config.GrpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewStatusClient(conn)

	ctx := context.Background()
	//Объявление сообщения
	meseg := StateRequest()
	//Отправка изображения

	response, err := cli.StateFun(ctx, meseg)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: ", response)
}

func StateRequest() (mes *pb.StateRequest) {
	mStorage, err := metrics.Get()
	if err != nil {
		log.Print(err)
	}
	//ссылки
	meseg := &pb.StateRequest{
		NextGC:      mStorage.MetricMapUint32["NextGC"],
		NumForcedGC: mStorage.MetricMapUint32["NumForcedGC"],

		BuckHashSys:  mStorage.MetricMapUint64["BuckHashSys"],
		Frees:        mStorage.MetricMapUint64["Frees"],
		GCSys:        mStorage.MetricMapUint64["GCSys"],
		HeapAlloc:    mStorage.MetricMapUint64["HeapAlloc"],
		HeapIdle:     mStorage.MetricMapUint64["HeapIdle"],
		HeapInuse:    mStorage.MetricMapUint64["HeapInuse"],
		HeapObjects:  mStorage.MetricMapUint64["HeapObjects"],
		HeapReleased: mStorage.MetricMapUint64["HeapReleased"],
		HeapSys:      mStorage.MetricMapUint64["HeapSys"],
		LastGC:       mStorage.MetricMapUint64["LastGC"],
		Lookups:      mStorage.MetricMapUint64["Lookups"],
		MCacheInuse:  mStorage.MetricMapUint64["MCacheInuse"],
		MCacheSys:    mStorage.MetricMapUint64["MCacheSys"],
		MSpanInuse:   mStorage.MetricMapUint64["MSpanInuse"],
		MSpanSys:     mStorage.MetricMapUint64["MSpanSys"],
		Mallocs:      mStorage.MetricMapUint64["Mallocs"],
		NumGC:        mStorage.MetricMapUint64["NumGC"],
		OtherSys:     mStorage.MetricMapUint64["OtherSys"],
		PauseTotalNs: mStorage.MetricMapUint64["PauseTotalNs"],
		StackInuse:   mStorage.MetricMapUint64["StackInuse"],
		StackSys:     mStorage.MetricMapUint64["StackSys"],
		Alloc:        mStorage.MetricMapUint64["Alloc"],
		Sys:          mStorage.MetricMapUint64["Sys"],
		TotalAlloc:   mStorage.MetricMapUint64["TotalAlloc"],
		RandomValue:  mStorage.MetricMapUint64["RandomValue"],
		TotalMemory:  mStorage.MetricMapUint64["TotalMemory"],
		FreeMemory:   mStorage.MetricMapUint64["FreeMemory"],

		GCCPUFraction: float32(mStorage.MetricMapFloat64["GCCPUFraction"]),
	}

	return meseg
}
