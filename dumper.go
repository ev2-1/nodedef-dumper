package main

import (
	"github.com/HimbeerserverDE/mt-multiserver-proxy"
	"github.com/anon55555/mt"
	pt "github.com/ev2-1/mt-multiserver-playerTools"

	"encoding/json"

	"log"
	"os"
	"time"
)

func init() {
	pt.RegisterPlayerListUpdateHandler(&pt.PlayerListUpdateHandler{
		Join: func(name string) {
			cc := pt.GetPlayerByName(name)

			cc.SendCmd(&mt.ToCltChatMsg{
				Type:      mt.SysMsg,
				Text:      "Dumper is installed, please disable in production\nNOTE: as its not able to do more, it dumps the MULTIPLEXED media pool, which makes it more or less useless with multiblie servers configured\nNOTE: use with one server at a time (disable all except one)\nNOTE: the plugin will OVERWRITE definitions if already present",
				Timestamp: time.Now().Unix(),
			})
		},
	})

	proxy.RegisterPacketHandler(&proxy.PacketHandler{
		SrvHandler: func(sc *proxy.ServerConn, pkt *mt.Pkt) bool {
			switch cmd := pkt.Cmd.(type) {
			case *mt.ToCltNodeDefs:
				sc.Log("[dumper] defs")

				// save:
				go saveNodeDefs(sc.Client().ServerName(), cmd.Defs)
			}

			return false
		},
	})
}

func saveNodeDefs(srv string, defs []mt.NodeDef) {
	os.Remove(proxy.Path("dumped_def.json"))
	f, err := os.OpenFile(proxy.Path("dumped_def.json"), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)

	err = encoder.Encode(defs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully dumped %d nodedefinitions for server %s", len(defs), srv)
}
