package mexc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSignRequest(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		params string
		secret string
		want   string
	}{
		{
			name:   "ok",
			params: "symbol=BTCUSDT&side=BUY&type=LIMIT&quantity=1&price=11&recvWindow=5000&timestamp=1644489390087",
			secret: "45d0b3c26f2644f19bfb98b07741b2f5",
			want:   "fd3e4e8543c5188531eb7279d68ae7d26a573d0fc5ab0d18eb692451654d837a",
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := signRequest(tc.params, tc.secret)
			require.Equal(t, tc.want, got)
		})
	}
}
