package rpc

import (
	"context"
	"testing"

	"github.com/test-go/testify/require"
)

func TestRPCError(t *testing.T) {
	if testEnv == "mock" {
		testConfig := beforeEach(t)
		_, err := testConfig.provider.ChainID(context.Background())
		require.NoError(t, err)
		require.Nil(t, err)

		_, err = testConfig.provider.Events(context.Background(), EventsInput{ResultPageRequest: ResultPageRequest{ChunkSize: 0}})
		require.Error(t, err)
		rpcErr := err.(*RPCError)
		require.Equal(t, rpcErr.Code, InternalError)
		require.NotNil(t, rpcErr.Message, "Internal Error")
		require.NotNil(t, rpcErr.Data, "-ChuckSize error message-")
	}
}
