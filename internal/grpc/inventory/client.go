package inventory

import (
	"context"
	"github.com/marvini86/car-parts-shop-service/internal/dto"
	pb "github.com/marvini86/car-service-protos/proto/inventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

// CheckAvailability checks the availability of an item
func CheckAvailability(ctx context.Context, codeIntegration string) (itemAvailability dto.ItemAvailabilityDto, err error) {
	conn, err := grpc.NewClient(os.Getenv("INVENTORY_GRPC_ENDPOINT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
		return
	}

	defer conn.Close()
	client := pb.NewInventoryServiceClient(conn)

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
