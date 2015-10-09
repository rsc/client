package keybase

import (
	"encoding/base64"
	"fmt"
	"net"
	"sync"

	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/service"
)

var con net.Conn
var startOnce sync.Once

// ServerURI should match run mode environment.
func Init(homeDir string, runMode string, serverURI string) {
	startOnce.Do(func() {
		libkb.G.Init()
		usage := libkb.Usage{
			Config:    true,
			API:       true,
			KbKeyring: true,
		}
		config := libkb.AppConfig{HomeDir: homeDir, RunMode: libkb.DevelRunMode, Debug: true, LocalRPCDebug: "Acsvip", ServerURI: serverURI}
		libkb.G.Configure(config, usage)
		(service.NewService(false)).StartLoopbackServer(libkb.G)
		Reset()
	})
}

// Takes base64 encoded msgpack rpc payload
func WriteB64(str string) bool {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Base64 decode error:", err, str)
	}
	n, err := con.Write(data)
	if err != nil {
		fmt.Println("Write error: ", err)
		return false
	}
	if n != len(data) {
		fmt.Println("Did not write all the data")
		return false
	}
	return true
}

// Blocking read, returns base64 encoded msgpack rpc payload
// bufferSize must be divisible by 3 to ensure that we don't split
// our b64 encode across a payload boundary if we go over our buffer
// size
const targetBufferSize = 50 * 1024
const bufferSize = targetBufferSize - (targetBufferSize % 3)

func ReadB64() string {
	data := make([]byte, bufferSize)

	n, err := con.Read(data)
	if n > 0 && err == nil {
		str := base64.StdEncoding.EncodeToString(data[0:n])
		return str
	}

	if err != nil {
		fmt.Println("read error:", err)
		// attempt to fix the connection
		Reset()
	}

	return ""
}

func Reset() {
	if con != nil {
		con.Close()
	}

	var err error
	libkb.G.SocketWrapper = nil
	con, _, err = libkb.G.GetSocket()

	if err != nil {
		fmt.Println("loopback socker error:", err)
	}
}

// So We can export these libkb types
type AndroidSecretRetriever libkb.SecretRetriever
type AndroidSecretStorer libkb.SecretStorer
type AndroidSecretStore libkb.SecretStore

// We have to duplicate the interface defined in libkb.AndroidKeyStore
// Otherwise we get an undefined param error when we use this as an argument
// in an exported func
type AndroidKeyStore interface {
	RetrieveSecret(username string) ([]byte, error)
	StoreSecret(username string, secret []byte) error
	ClearSecret(username string) error
	GetUsersWithStoredSecretsMsgPack() ([]byte, error)
	SetupKeyStore(username string) error
}

func SetGlobalAndroidKeyStore(s AndroidKeyStore) {
	// TODO: Gross! can we fix this?
	libkb.G.SetAndroidKeyStore(libkb.AndroidKeyStore(s))
}
