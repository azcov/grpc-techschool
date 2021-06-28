package service_test

import (
	"context"
	"testing"

	"github.com/azcov/grpc-techschool/pb"
	"github.com/azcov/grpc-techschool/sample"
	"github.com/azcov/grpc-techschool/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLaptopServer_CreateLaptop(t *testing.T) {
	t.Parallel()
	validlaptop := sample.NewLaptop()

	laptopWithNoID := sample.NewLaptop()
	laptopWithNoID.Id = ""

	laptopWithInvalidID := sample.NewLaptop()
	laptopWithInvalidID.Id = "1"

	laptopWithDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopWithDuplicateID)
	require.Nil(t, err)

	tests := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "succes_with_id",
			laptop: validlaptop,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "succes_no_id",
			laptop: laptopWithNoID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_id",
			laptop: laptopWithInvalidID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "duplicate_id",
			laptop: laptopWithDuplicateID,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tt.laptop,
			}

			server := service.NewLaptopServer(tt.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if tt.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tt.laptop.Id) > 0 {
					require.Equal(t, tt.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.code, st.Code())
			}
		})
	}
}
