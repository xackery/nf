package main

import (
	"io"
	"net"
	"os"
	"strconv"

	"github.com/xackery/nf/db"
	"github.com/xackery/nf/gateway"
	"github.com/xackery/nf/insecure"
	ranksv1 "github.com/xackery/nf/proto/aas/ranks/v1"
	aasv1 "github.com/xackery/nf/proto/aas/v1"
	accountsv1 "github.com/xackery/nf/proto/accounts/v1"
	adventuresv1 "github.com/xackery/nf/proto/adventures/v1"
	charactersv1 "github.com/xackery/nf/proto/characters/v1"
	"github.com/xackery/nf/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	dbType := "mysql"
	dsn := os.Getenv("EQ_DSN")
	if dsn == "" {
		dsn = "root:root@tcp(localhost:3306)/eq"
	}
	pskPath := os.Getenv("EQ_PSK_PATH")
	if pskPath != "" {
		dbType = "mysql+ssh"
	}
	if os.Getenv("EQ_SQLITE") != "" {
		dbType = "sqlite3"
	}

	if os.Getenv("EQ_CONN") != "" {
		dbType = os.Getenv("EQ_CONN")
	}

	switch dbType {
	case "sqlite3":
		err = db.NewSQLLite(os.Getenv("EQ_SQLITE"))
		if err != nil {
			log.Fatalln("Failed to connect to database via mysqlite:", err)
		}
	case "mysql+ssh":
		port := 22
		portStr := os.Getenv("EQ_PSK_PORT")
		if portStr != "" {
			port, err = strconv.Atoi(portStr)
			if err != nil {
				log.Fatalln("Failed to parse EQ_PSK_PORT:", err)
			}
		}
		err = db.NewSSHPSK(os.Getenv("EQ_PSK_HOST"), port, os.Getenv("EQ_PSK_USER"), pskPath, dsn)
		if err != nil {
			log.Fatalln("Failed to connect to database via psk:", err)
		}
	default:
		err = db.New(dsn)
		if err != nil {
			log.Fatalln("Failed to connect to database via mysql:", err)
		}
	}
	log.Info("Connected to database via " + dbType)

	s := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	aasv1.RegisterAaServiceServer(s, server.New())
	ranksv1.RegisterRankServiceServer(s, server.New())
	accountsv1.RegisterAccountServiceServer(s, server.New())
	adventuresv1.RegisterAdventureServiceServer(s, server.New())
	charactersv1.RegisterCharacterServiceServer(s, server.New())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
