package dingtalk

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/EmptyZeroRain/proper_dingtalk_bot/internal/security"

	mockmessage "github.com/EmptyZeroRain/proper_dingtalk_bot/test/mocks/message"
	"github.com/golang/mock/gomock"
)

func TestNewClient(t *testing.T) {
	type args struct {
		accessToken string
		secret      string
		domain      string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "",
			args: args{
				accessToken: "123456",
				secret:      "111111",
				domain:      "oapi.dingtalk.com",
			},
			want: &Client{
				AccessToken: "123456",
				Secret:      "111111",
				Domain:      "oapi.dingtalk.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.accessToken, tt.args.secret, tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Send(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	message := mockmessage.NewMockMessage(ctrl)

	t.Run("message return error", func(t *testing.T) {
		c := &Client{}

		message.EXPECT().ToByte().Return([]byte{}, errors.New("test"))

		if _, _, err := c.Send(message); err == nil {
			t.Error("send error")
		}
	})

	t.Run("security.URL return error", func(t *testing.T) {
		c := &Client{
			AccessToken: "test-access-token",
			Secret:      "test-secret",
		}

		message.EXPECT().ToByte().Return([]byte{}, nil)
		monkey.Patch(security.URL, func(domain, accessToken, secret string) (string, error) {
			return "", errors.New("URL error")
		})

		if _, _, err := c.Send(message); err == nil {
			t.Error("send error")
		}
	})

	t.Run("http.NewRequest return error", func(t *testing.T) {
		c := &Client{
			AccessToken: "test-access-token",
			Secret:      "test-secret",
		}

		message.EXPECT().ToByte().Return([]byte{}, nil)
		monkey.Patch(security.URL, func(domain, accessToken, secret string) (string, error) {
			return "https://oapi.dingtalk.com/robot/send?access_token=ewfewfwfwefwafew", nil
		})

		monkey.Patch(http.NewRequest, func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("NewRequest error")
		})

		if _, _, err := c.Send(message); err == nil {
			t.Error("send error")
		}
	})
}
