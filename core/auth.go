package core

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"strings"
)

var tlsConfig *tls.Config

func init() {
	certPool := x509.NewCertPool()
	body, _ := os.ReadFile("riotgames.pem")
	certPool.AppendCertsFromPEM(body)
	tlsConfig = &tls.Config{RootCAs: certPool}
}

type GameAuth struct {
	Pid         int32
	Port, Token string
}

func (ga *GameAuth) BasicHeader() string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte("riot:"+ga.Token))
}

func NewGameAuth(p *process.Process) (*GameAuth, error) {
	arguments, err := p.CmdlineSlice()
	if err != nil {
		return nil, fmt.Errorf("error parsing client arguments: %w", err)
	}
	ga := &GameAuth{Pid: p.Pid}
	for _, arg := range arguments {
		if strings.HasPrefix(arg, `"--app-port=`) {
			ga.Port = strings.TrimSuffix(strings.TrimPrefix(arg, `"--app-port=`), `"`)
		} else if strings.HasPrefix(arg, `"--remoting-auth-token`) {
			ga.Token = strings.TrimSuffix(strings.TrimPrefix(arg, `"--remoting-auth-token=`), `"`)
		}
	}
	return ga, nil
}

func NewGameAuthWithToken(port, token string) *GameAuth {
	return &GameAuth{Token: token, Port: port}
}
