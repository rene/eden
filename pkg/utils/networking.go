package utils

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/lf-edge/eden/pkg/defaults"
	log "github.com/sirupsen/logrus"
)

// IFInfo stores information about net address and subnet
type IFInfo struct {
	Subnet        *net.IPNet
	FirstAddress  net.IP
	SecondAddress net.IP
}

func getSubnetByInd(ind int) (*net.IPNet, error) {
	if ind < 0 || ind > 255 {
		return nil, fmt.Errorf("error in index %d", ind)
	}
	_, curNet, err := net.ParseCIDR(fmt.Sprintf("192.168.%d.1/24", ind))
	return curNet, err
}

func getIPByInd(ind int) ([]net.IP, error) {
	if ind < 0 || ind > 255 {
		return nil, fmt.Errorf("error in index %d", ind)
	}
	IP := net.ParseIP(fmt.Sprintf("192.168.%d.10", ind))
	if IP == nil {
		return nil, fmt.Errorf("error in ParseIP for index %d", ind)
	}
	ips := []net.IP{IP}
	IP2 := net.ParseIP(fmt.Sprintf("192.168.%d.11", ind))
	if IP2 == nil {
		return nil, fmt.Errorf("error in ParseIP for index %d", ind)
	}
	ips = append(ips, IP2)
	return ips, nil
}

// GetSubnetsNotUsed prepare map with subnets and ip not used by any interface of host
func GetSubnetsNotUsed(count int) ([]IFInfo, error) {
	var result []IFInfo
	curSubnetInd := 0
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for ; len(result) < count; curSubnetInd++ {
		curNet, err := getSubnetByInd(curSubnetInd)
		if err != nil {
			return nil, fmt.Errorf("error in GetSubnetsNotUsed: %s", err)
		}
		contains := false
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					if curNet.Contains(ipnet.IP) {
						contains = true
						break
					}
				}
			}
		}
		if !contains {
			ips, err := getIPByInd(curSubnetInd)
			if err != nil {
				return nil, fmt.Errorf("error in getIPByInd: %s", err)
			}
			result = append(result, IFInfo{
				Subnet:        curNet,
				FirstAddress:  ips[0],
				SecondAddress: ips[1],
			})
		}
	}
	return result, nil
}

// GetIPForDockerAccess is service function to obtain IP for adam access
// The function is filter out docker bridge
func GetIPForDockerAccess() (ipv4, ipv6 net.IP, err error) {
	networks, err := GetDockerNetworks()
	if err != nil {
		log.Errorf("GetDockerNetworks: %s", err)
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
out:
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			for _, el := range networks {
				if el.Contains(ipnet.IP) {
					continue out
				}
			}
			if ipv4 == nil && ipnet.IP.To4() != nil {
				ipv4 = ipnet.IP.To4()
			}
			if ipv6 == nil && ipnet.IP.To4() == nil {
				ipv6 = ipnet.IP.To16()
			}
			if ipv4 != nil && ipv6 != nil {
				break
			}
		}
	}
	if ipv4 == nil && ipv6 == nil {
		return ipv4, ipv6, errors.New("no IP found")
	}
	return ipv4, ipv6, nil
}

// ResolveURL concatenate parts of url
func ResolveURL(b, p string) (string, error) {
	u, err := url.Parse(p)
	if err != nil {
		return "", err
	}
	base, err := url.Parse(b)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(u).String(), nil
}

// ipToBigInt converts net.IP to *big.Int
func ipToBigInt(ip net.IP) *big.Int {
	return big.NewInt(0).SetBytes(ip)
}

// bigIntToIP converts *big.Int to net.IP of the given length
func bigIntToIP(i *big.Int, length int) net.IP {
	b := i.Bytes()
	if len(b) < length {
		padded := make([]byte, length)
		copy(padded[length-len(b):], b)
		b = padded
	}
	return b
}

// GetNetworkIPs returns the first, second, and last usable IPs from a *net.IPNet.
func GetNetworkIPs(subnet string) (gateway, dhcpStart, dhcpEnd net.IP, err error) {
	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, nil, nil,
			fmt.Errorf("failed to parse network subnet %s: %w", subnet, err)
	}
	networkIP := ipNet.IP
	prefixLen, bits := ipNet.Mask.Size()

	if bits-prefixLen < 2 {
		return nil, nil, nil,
			fmt.Errorf("subnet too small: need at least 2 host bits (got /%d)", prefixLen)
	}

	ipLen := len(networkIP)
	networkInt := ipToBigInt(networkIP.Mask(ipNet.Mask))

	// First usable IP (gateway): network IP + 1
	firstUsable := big.NewInt(0).Add(networkInt, big.NewInt(1))

	// Second usable IP (DHCP start): gateway + 1
	secondUsable := big.NewInt(0).Add(firstUsable, big.NewInt(1))

	// Last usable IP (DHCP end): broadcast - 1 (IPv4) or max IP in range - 1 (IPv6)
	broadcastInt := big.NewInt(0).Add(networkInt, big.NewInt(0).Lsh(
		big.NewInt(1), uint(bits-prefixLen)))
	lastUsable := big.NewInt(0).Sub(broadcastInt, big.NewInt(2))

	gateway = bigIntToIP(firstUsable, ipLen)
	dhcpStart = bigIntToIP(secondUsable, ipLen)
	dhcpEnd = bigIntToIP(lastUsable, ipLen)
	return gateway, dhcpStart, dhcpEnd, nil
}

// GetFileSizeURL returns file size for url
func GetFileSizeURL(url string) int64 {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	return int64(size)
}

// RepeatableAttempt do request several times waiting for nil error and expected status code
func RepeatableAttempt(client *http.Client, req *http.Request) (response *http.Response, err error) {
	maxRepeat := defaults.DefaultRepeatCount
	delayTime := defaults.DefaultRepeatTimeout

	for i := 0; i < maxRepeat; i++ {
		timer := time.AfterFunc(2*delayTime, func() {
			i = 0
		})
		resp, err := client.Do(req)
		wrongCode := false
		if err == nil {
			// we should check the status code of the response and try again if needed
			if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusAccepted {
				return resp, nil
			}
			wrongCode = true
			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Debugf("bad status: %s", resp.Status)
			} else {
				log.Debugf("bad status (%s) in response (%s)", resp.Status, string(buf))
			}
		}
		log.Debugf("error %s URL %s: %v", req.Method, req.RequestURI, err)
		timer.Stop()
		if wrongCode {
			log.Infof("Received unexpected StatusCode(%s): repeat request (%d) of (%d)",
				http.StatusText(resp.StatusCode), i, maxRepeat)
		} else {
			log.Infof("Attempt to re-establish connection (%d) of (%d)", i, maxRepeat)
		}
		time.Sleep(delayTime)
	}
	return nil, fmt.Errorf("all connection attempts failed")
}

// UploadFile send file in form
func UploadFile(client *http.Client, url, filePath, prefix string) (result *http.Response, err error) {
	body, writer := io.Pipe()

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	mwriter := multipart.NewWriter(writer)
	req.Header.Add("Content-Type", mwriter.FormDataContentType())

	errchan := make(chan error)

	fileName := filepath.Base(filePath)
	if prefix != "" {
		fileName = fmt.Sprintf("%s/%s", prefix, fileName)
	}

	go func() {
		defer writer.Close()
		defer mwriter.Close()
		w, err := mwriter.CreateFormFile("file", fileName)
		if err != nil {
			errchan <- err
			return
		}
		in, err := os.Open(filePath)
		if err != nil {
			errchan <- err
			return
		}
		defer in.Close()

		counter := &writeCounter{step: 10 * 1024 * 1024, message: "Uploading..."}
		if written, err := io.Copy(w, io.TeeReader(in, counter)); err != nil {
			errchan <- fmt.Errorf("error copying %s (%d bytes written): %v", filePath, written, err)
			return
		}
		fmt.Printf("\n")

		if err := mwriter.Close(); err != nil {
			errchan <- err
			return
		}
		log.Info("Waiting for SHA256 calculation")
	}()
	respchan := make(chan *http.Response)
	go func() {
		resp, err := client.Do(req)
		if err != nil {
			errchan <- err
		} else {
			respchan <- resp
		}
	}()
	var merr error
	var resp *http.Response
	select {
	case merr = <-errchan:
		return nil, fmt.Errorf("http/multipart error: %v", merr)
	case resp = <-respchan:
		return resp, nil
	}
}

// FindUnusedPort : find port number not currently used by the host.
func FindUnusedPort() (uint16, error) {
	// We let the kernel to find the port for us.
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return uint16(l.Addr().(*net.TCPAddr).Port), nil
}
