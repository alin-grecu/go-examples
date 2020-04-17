package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli" // master is using v2 not v1
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "ns",
				Usage: "Looks up the nameservers for a particular host",
				Action: func(c *cli.Context) error {

					// a simple lookup function
					ns, err := net.LookupNS(c.Args().First())

					if err != nil {
						return err
					}

					// we log the results to our console
					for i := 0; i < len(ns); i++ {
						fmt.Println(ns[i].Host)
					}

					return nil
				},
			},
			{
				Name:  "ip",
				Usage: "Looks up the IP addresses for a particular host",
				Action: func(c *cli.Context) error {
					ip, err := net.LookupIP(c.Args().First())
					if err != nil {
						fmt.Println(err)
					}
					for i := 0; i < len(ip); i++ {
						fmt.Println(ip[i])
					}
					return nil
				},
			},
			{
				Name:  "cname",
				Usage: "Looks up the CNAME for a particular host",
				Action: func(c *cli.Context) error {
					cname, err := net.LookupCNAME(c.Args().First())
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(cname)
					return nil
				},
			},
			{
				Name:  "mx",
				Usage: "Looks up the MX records for a particular host",
				Action: func(c *cli.Context) error {
					mx, err := net.LookupMX(c.Args().First())
					if err != nil {
						fmt.Println(err)
					}
					for i := 0; i < len(mx); i++ {
						fmt.Println(mx[i].Host, mx[i].Pref)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
