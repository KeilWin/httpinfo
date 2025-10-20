package handlers

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

type IpInfoService interface {
	GetIpInfo(ipAddress string) any
}

type IpLocationNet struct {
	client *http.Client
}

func (p *IpLocationNet) GetIpInfo(ipAddress string) ([]byte, error) {
	response, err := p.client.Get(fmt.Sprintf("https://api.iplocation.net/?ip=%s", ipAddress))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func NewIpLocationNet(newClientFunc func() *http.Client) *IpLocationNet {
	return &IpLocationNet{
		client: newClientFunc(),
	}
}

func CheckIpAddress(ipAddress string) error {
	parsedIp := net.ParseIP(ipAddress)
	if parsedIp == nil {
		return fmt.Errorf("IP address structure is wrong: %s", ipAddress)
	}
	return nil
}

func NewIpHandler(newClientFunc func() *http.Client) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress := r.PathValue("ipAddress")
		if err := CheckIpAddress(ipAddress); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ipLocationNet := NewIpLocationNet(newClientFunc)
		body, err := ipLocationNet.GetIpInfo(ipAddress)
		if err != nil {
			http.Error(w, fmt.Sprintf("Can't get ip info: %s", ipAddress), http.StatusInternalServerError)
			return
		}
		w.Write(body)
		w.Header().Add("content-type", "application/json")
	})
}
