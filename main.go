package main

import (
	"context"
	"fmt"
	"log"

	simulationpb "git.risenlighten.com/lasvsim/process_task/api/cosim/v1"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	run()
}

func run() {
	// 加载服务器证书
	certificate, err := credentials.NewClientTLSFromFile("qianxing-grpc.risenlighten.com_public.crt", "qianxing-grpc.risenlighten.com")
	if err != nil {
		log.Fatalf("无法加载服务器证书: %v", err)
	}
	// Connect to the gRPC server
	conn, err := grpc.Dial("qianxing-grpc.risenlighten.com:443", grpc.WithTransportCredentials(certificate))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create the CosimStub
	client := simulationpb.NewCosimClient(conn)

	// 添加授权信息
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjcsIm9pZCI6MTAxLCJuYW1lIjoi572X5b-X5rSqIiwiaWRlbnRpdHkiOiJub3JtYWwiLCJwZXJtaXNzaW9ucyI6WyJ0YXNrLnRhc2sucHVibGljLkFDVElPTl9WSUVXIiwidGFzay50YXNrLnB1YmxpYy5BQ1RJT05fQ09QWSIsInRhc2sudGFzay5wdWJsaWMuQUNUSU9OX1JFUExBWSIsInRhc2sudGFzay5wdWJsaWMuQUNUSU9OX1JFUE9SVCIsInRhc2sudGFzay5wcml2YXRlLkFDVElPTl9WSUVXIiwidGFzay50YXNrLnByaXZhdGUuQUNUSU9OX0FERCIsInRhc2sudGFzay5wcml2YXRlLkFDVElPTl9DT1BZIiwidGFzay50YXNrLnByaXZhdGUuQUNUSU9OX0RFTEVURSIsInRhc2sudGFzay5wcml2YXRlLkFDVElPTl9SRVBMQVkiLCJ0YXNrLnRhc2sucHJpdmF0ZS5BQ1RJT05fUkVQT1JUIiwidGFzay50YXNrLnBlcnNvbmFsLkFDVElPTl9WSUVXIiwidGFzay50YXNrLnBlcnNvbmFsLkFDVElPTl9ERUxFVEUiLCJ0YXNrLnRhc2sucGVyc29uYWwuQUNUSU9OX1JFUExBWSIsInRhc2sudGFzay5wZXJzb25hbC5BQ1RJT05fUkVQT1JUIiwicmVzb3VyY2UudmVoaWNsZS5wdWJsaWMuQUNUSU9OX1ZJRVciLCJyZXNvdXJjZS52ZWhpY2xlLnB1YmxpYy5BQ1RJT05fVVNFIiwicmVzb3VyY2UudmVoaWNsZS5wcml2YXRlLkFDVElPTl9WSUVXIiwicmVzb3VyY2UudmVoaWNsZS5wcml2YXRlLkFDVElPTl9BREQiLCJyZXNvdXJjZS52ZWhpY2xlLnByaXZhdGUuQUNUSU9OX1VQREFURSIsInJlc291cmNlLnZlaGljbGUucHJpdmF0ZS5BQ1RJT05fREVMRVRFIiwicmVzb3VyY2UudmVoaWNsZS5wcml2YXRlLkFDVElPTl9VU0UiLCJyZXNvdXJjZS52ZWhpY2xlLnBlcnNvbmFsLkFDVElPTl9WSUVXIiwicmVzb3VyY2UudmVoaWNsZS5wZXJzb25hbC5BQ1RJT05fVVBEQVRFIiwicmVzb3VyY2UudmVoaWNsZS5wZXJzb25hbC5BQ1RJT05fREVMRVRFIiwicmVzb3VyY2Uuc2Vuc29yLnB1YmxpYy5BQ1RJT05fVklFVyIsInJlc291cmNlLnNlbnNvci5wdWJsaWMuQUNUSU9OX1VTRSIsInJlc291cmNlLnNlbnNvci5wcml2YXRlLkFDVElPTl9WSUVXIiwicmVzb3VyY2Uuc2Vuc29yLnByaXZhdGUuQUNUSU9OX0FERCIsInJlc291cmNlLnNlbnNvci5wcml2YXRlLkFDVElPTl9VUERBVEUiLCJyZXNvdXJjZS5zZW5zb3IucHJpdmF0ZS5BQ1RJT05fREVMRVRFIiwicmVzb3VyY2Uuc2Vuc29yLnByaXZhdGUuQUNUSU9OX1VTRSIsInJlc291cmNlLnNlbnNvci5wZXJzb25hbC5BQ1RJT05fVklFVyIsInJlc291cmNlLnNlbnNvci5wZXJzb25hbC5BQ1RJT05fVVBEQVRFIiwicmVzb3VyY2Uuc2Vuc29yLnBlcnNvbmFsLkFDVElPTl9ERUxFVEUiLCJyZXNvdXJjZS5tYXAucHVibGljLkFDVElPTl9WSUVXIiwicmVzb3VyY2UubWFwLnB1YmxpYy5BQ1RJT05fVVNFIiwicmVzb3VyY2UubWFwLnByaXZhdGUuQUNUSU9OX1ZJRVciLCJyZXNvdXJjZS5tYXAucHJpdmF0ZS5BQ1RJT05fQUREIiwicmVzb3VyY2UubWFwLnByaXZhdGUuQUNUSU9OX1VQREFURSIsInJlc291cmNlLm1hcC5wcml2YXRlLkFDVElPTl9ERUxFVEUiLCJyZXNvdXJjZS5tYXAucHJpdmF0ZS5BQ1RJT05fVVNFIiwicmVzb3VyY2UubWFwLnBlcnNvbmFsLkFDVElPTl9WSUVXIiwicmVzb3VyY2UubWFwLnBlcnNvbmFsLkFDVElPTl9VUERBVEUiLCJyZXNvdXJjZS5tYXAucGVyc29uYWwuQUNUSU9OX0RFTEVURSIsInJlc291cmNlLnNjZW5hcmlvLnB1YmxpYy5BQ1RJT05fVklFVyIsInJlc291cmNlLnNjZW5hcmlvLnB1YmxpYy5BQ1RJT05fVVNFIiwicmVzb3VyY2Uuc2NlbmFyaW8ucHJpdmF0ZS5BQ1RJT05fVklFVyIsInJlc291cmNlLnNjZW5hcmlvLnByaXZhdGUuQUNUSU9OX0NPUFkiLCJyZXNvdXJjZS5zY2VuYXJpby5wcml2YXRlLkFDVElPTl9VUERBVEUiLCJyZXNvdXJjZS5zY2VuYXJpby5wcml2YXRlLkFDVElPTl9ERUxFVEUiLCJyZXNvdXJjZS5zY2VuYXJpby5wcml2YXRlLkFDVElPTl9VU0UiLCJyZXNvdXJjZS5zY2VuYXJpby5wcml2YXRlLkFDVElPTl9BREQiLCJyZXNvdXJjZS5zY2VuYXJpby5wZXJzb25hbC5BQ1RJT05fVklFVyIsInJlc291cmNlLnNjZW5hcmlvLnBlcnNvbmFsLkFDVElPTl9VUERBVEUiLCJyZXNvdXJjZS5zY2VuYXJpby5wZXJzb25hbC5BQ1RJT05fREVMRVRFIiwicmVzb3VyY2UudHJhZmZpY19mbG93X2NvbmZpZy5wdWJsaWMuQUNUSU9OX1ZJRVciLCJyZXNvdXJjZS50cmFmZmljX2Zsb3dfY29uZmlnLnB1YmxpYy5BQ1RJT05fVVNFIiwicmVzb3VyY2UudHJhZmZpY19mbG93X2NvbmZpZy5wcml2YXRlLkFDVElPTl9WSUVXIiwicmVzb3VyY2UudHJhZmZpY19mbG93X2NvbmZpZy5wcml2YXRlLkFDVElPTl9VUERBVEUiLCJyZXNvdXJjZS50cmFmZmljX2Zsb3dfY29uZmlnLnByaXZhdGUuQUNUSU9OX0RFTEVURSIsInJlc291cmNlLnRyYWZmaWNfZmxvd19jb25maWcucHJpdmF0ZS5BQ1RJT05fVVNFIiwicmVzb3VyY2UudHJhZmZpY19mbG93X2NvbmZpZy5wcml2YXRlLkFDVElPTl9BREQiLCJyZXNvdXJjZS50cmFmZmljX2Zsb3dfY29uZmlnLnBlcnNvbmFsLkFDVElPTl9WSUVXIiwicmVzb3VyY2UudHJhZmZpY19mbG93X2NvbmZpZy5wZXJzb25hbC5BQ1RJT05fVVBEQVRFIiwicmVzb3VyY2UudHJhZmZpY19mbG93X2NvbmZpZy5wZXJzb25hbC5BQ1RJT05fREVMRVRFIiwicmVzb3VyY2UubWFwLnByaXZhdGUuQUNUSU9OX0RPV05MT0FEIiwicmVzb3VyY2UubWFwLnB1YmxpYy5BQ1RJT05fRE9XTkxPQUQiXSwiaXNzIjoidXNlciIsInN1YiI6Ikxhc1ZTaW0iLCJleHAiOjE2OTExMjQ3MzEsIm5iZiI6MTY5MTAzODMzMSwiaWF0IjoxNjkxMDM4MzMxLCJqdGkiOiI3In0.s8Uqp4p7bNCxkSF9M9PUn5fQdSgfpoU2ajDN59Qt7aY"
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer "+token)

	// Start the simulation
	startResp, err := client.Start(ctx, &simulationpb.StartSimulationReq{
		TaskId:   3958,
		RecordId: 1536,
	})
	if checkError(err) {
		return
	}
	fmt.Println(">>>")
	for {
		// Get the vehicle information
		_, err := client.GetVehicle(ctx, &simulationpb.GetVehicleReq{
			SimulationId: startResp.SimulationId,
			VehicleId:    "ego",
		})
		if checkError(err) {
			break
		}

		stepResult, err := client.NextStep(ctx, &simulationpb.NextStepReq{
			SimulationId: startResp.SimulationId,
		})
		if checkError(err) {
			break
		}

		// Check for simulation completion
		if stepResult.State.Progress < 0 || stepResult.State.Progress >= 100 {
			fmt.Printf("Simulation ended, status: %s\n", stepResult.State.Msg)
			break
		}
	}

	fmt.Printf("id: %d\n", startResp.SimulationId)

	// Stop the simulation
	client.Stop(ctx, &simulationpb.StopSimulationReq{
		SimulationId: startResp.SimulationId,
	})

	// Get the results
	_, err = client.GetResults(ctx, &simulationpb.GetResultsReq{
		SimulationId: startResp.SimulationId,
	})
	if err != nil {
		fmt.Printf("Error getting results: %v\n", err)
	} else {
		fmt.Println("仿真完成")
		// Print results here if needed: fmt.Printf("Simulation ended, results: %v\n", result.Results)
	}
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err)
	}

	return err != nil
}
