package inventory

import (
	"context"
	"github.com/marvini86/car-parts-shop-service/internal/dto"
	"github.com/marvini86/car-parts-shop-service/internal/grpc"
	pb "github.com/marvini86/car-service-protos/proto/inventory"
	"log"
	"os"
	"time"
)

// CheckAvailability checks the availability of an item
func CheckAvailability(ctx context.Context, codeIntegration string) (itemAvailability dto.ItemAvailabilityDto, err error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	conn, err := grpc.NewGrpcClient(ctxTimeout, os.Getenv("INVENTORY_GRPC_ENDPOINT"))

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
		return
	}

	defer conn.Close()

	client := pb.NewInventoryServiceClient(conn.GetConn())

	res, err := client.CheckAvailability(ctx, &pb.ItemRequest{
		Code: codeIntegration,
	})

	if err != nil {
		log.Fatalf("Error checking availability: %v", err)
		return
	}

	return dto.ItemAvailabilityDto{
		CodeIntegration:   res.Code,
		Name:              res.Name,
		AvailableQuantity: res.AvailableQuantity,
	}, err
}
