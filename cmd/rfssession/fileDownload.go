package rfssession

import (
	"context"
	"fmt"
	"i9pkgs/i9types"
	"log"
	"os"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func downloadFile(command string, cmdArgs []string, serverWorkPath string, connStream *websocket.Conn) {
	if cmdArgsLen := len(cmdArgs); cmdArgsLen != 2 {
		fmt.Printf("error: download: %d arguments provided, 2 required\n", cmdArgsLen)
		return
	}

	filename := cmdArgs[0]
	destination := cmdArgs[1]

	sendData := map[string]any{
		"workPath": serverWorkPath,
		"command":  command,
		"cmdArgs":  []string{filename},
	}

	if w_err := wsjson.Write(context.Background(), connStream, sendData); w_err != nil {
		log.Println(fmt.Errorf("rfssession: %s: write error: %s", command, w_err))
		return
	}

	var recvData i9types.WSResp

	if r_err := wsjson.Read(context.Background(), connStream, &recvData); r_err != nil {
		log.Println(fmt.Errorf("rfssession: %s: read error: %s", command, r_err))
		return
	}

	if recvData.Status == "f" {
		fmt.Printf("error: %s: %s\n", command, recvData.Error)
		return
	}

	if err := os.WriteFile(destination, recvData.Body.([]byte), 0644); err != nil {
		fmt.Printf("%s: %s\n", command, err)
	}
}
