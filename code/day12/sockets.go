package main

import (
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// we're ignoring errors for simplicity
		conn, _, _, _ := ws.UpgradeHTTP(r, w)
		log.Print("Hello Go Web Examples, you're doing great!")

		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Printf("Error in wsutil.ReadClientData: %s", err)
				}

				err = wsutil.WriteServerMessage(conn, op, msg)
				if err != nil {
					log.Printf("Error in ws.WriteServerMessage: %s", err)
				}

				return // avoid an infinite loop of errors
			}
		}()
	}))
}
