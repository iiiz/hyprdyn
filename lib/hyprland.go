package hyprdyn

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/user"
	"path/filepath"
)

type HyprlandRequestClient struct {
	connection *net.UnixAddr
}

const (
	requestSocket string = ".socket.sock"
	MaxBufferSize int    = 8192
)

var hyprlandClient *HyprlandRequestClient

func InitHyprlandClient() {
	hyprlandClient = NewHyprlandClient()
}

func getHyprlandSocket() (*string, error) {
	signature, signatureIsSet := os.LookupEnv("HYPRLAND_INSTANCE_SIGNATURE")
	if !signatureIsSet || signature == "" {
		return nil, errors.New("Instance signature not set, are you sure Hyprland is running?")
	}

	runtimeDir, runtimeDirIsSet := os.LookupEnv("XDG_RUNTIME_DIR")

	if !runtimeDirIsSet || runtimeDir == "" {
		u, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("error while getting the current user: %w", err)
		}

		runtimeDir = filepath.Join("/run/user", u.Uid)
	}

	socket := filepath.Join(runtimeDir, "hypr", signature, requestSocket)

	return &socket, nil
}

func NewHyprlandClient() *HyprlandRequestClient {
	socket, err := getHyprlandSocket()
	Check(err)

	return &HyprlandRequestClient{
		connection: &net.UnixAddr{
			Net:  "unix",
			Name: *socket,
		},
	}
}

func (c *HyprlandRequestClient) sendCommand(command string, args *string) (response []byte, err error) {
	requestBuf := bytes.NewBuffer(nil)

	requestBuf.Write([]byte{'j', '/'})
	requestBuf.WriteString(command)

	if args != nil {
		requestBuf.WriteByte(' ')
		requestBuf.WriteString(*args)
	}

	requestBytes := requestBuf.Bytes()
	if len(requestBytes) > MaxBufferSize {
		return nil, fmt.Errorf(
			"Request violates max buffer size (%d): %s",
			MaxBufferSize,
			requestBytes,
		)
	}

	connection, err := net.DialUnix("unix", nil, c.connection)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to socket: %w", err)
	}

	writer := bufio.NewWriter(connection)
	_, err = writer.Write(requestBytes)
	if err != nil {
		return nil, fmt.Errorf("Error writing request: %w", err)
	}

	err = writer.Flush()
	if err != nil {
		return nil, fmt.Errorf("Error on socket flush: %w", err)
	}

	responseBuf, err := io.ReadAll(connection)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %w", err)
	}

	if closeErr := connection.Close(); closeErr != nil {
		err = errors.Join(err, fmt.Errorf("Error on close: %w", closeErr))
	}

	return responseBuf, err
}

func UnmarshalHyprlandResponse[T any](response []byte, d *T) (T, error) {
	if len(response) == 0 {
		return *d, errors.New("empty response")
	}

	err := json.Unmarshal(response, &d)
	if err != nil {
		return *d, fmt.Errorf(
			"Error unmarshaling: %w, response: %s",
			err,
			response,
		)
	}
	return *d, nil
}
