package mixer

import (
	"context"
	"fmt"

	driver "github.com/antha-lang/antha/driver/antha_driver_v1"
	"google.golang.org/grpc"
)

type BaseMixer struct {
	connection       string
	expectedSubTypes []string
}

func NewBaseMixer(connection string, subTypes ...string) *BaseMixer {
	return &BaseMixer{
		connection:       connection,
		expectedSubTypes: subTypes,
	}
}

func (bm *BaseMixer) ConnectInit() (*grpc.ClientConn, error) {
	if bm.connection == "" {
		return nil, nil

	} else {
		conn, err := grpc.Dial(bm.connection, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		c := driver.NewDriverClient(conn)
		ctx := context.Background()
		if reply, err := c.DriverType(ctx, &driver.TypeRequest{}); err != nil {
			return nil, err
		} else if typ := reply.GetType(); typ != "antha.mixer.v1.Mixer" {
			return nil, fmt.Errorf("Expected to find a mixer driver at %s but instead found: %s", bm.connection, typ)
		} else if subtypes := reply.GetSubtypes(); len(subtypes) != len(bm.expectedSubTypes) {
			return nil, fmt.Errorf("Expected to find a %v mixer driver at %s but instead found: %v", bm.expectedSubTypes, bm.connection, subtypes)
		} else {
			for idx, est := range bm.expectedSubTypes {
				if subtypes[idx] != est {
					return nil, fmt.Errorf("Expected to find a %v mixer driver at %s but instead found: %v", bm.expectedSubTypes, bm.connection, subtypes)
				}
			}
			return conn, nil
		}
	}
}