package db

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/ssh"
)

type mySQLSSHDriver struct {
	sshClient *ssh.Client
}

func (e *mySQLSSHDriver) Dial(ctx context.Context, addr string) (net.Conn, error) {
	return e.sshClient.Dial("tcp", addr)
}

var (
	conn    *sqlx.DB
	sshConn *ssh.Client
)

func NewSQLLite(dataSourceName string) error {
	var err error

	conn, err = sqlx.Connect("sqlite3", dataSourceName)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	return nil
}

// New creates a new database connection
func New(dataSourceName string) error {
	var err error

	connType := "mysql"
	if sshConn != nil {
		mysql.RegisterDialContext("mysql+ssh", (&mySQLSSHDriver{sshConn}).Dial)
	}
	conn, err = sqlx.Connect(connType, dataSourceName)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	return nil
}

func NewSSH(host string, port int, user string, pass string, dataSourceName string) error {
	var err error
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshConn, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshConfig)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	return New(dataSourceName)
}

func NewSSHPSK(host string, port int, user string, pskPath string, dataSourceName string) error {
	data, err := os.ReadFile(pskPath)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}
	signer, err := ssh.ParsePrivateKey(data)
	if err != nil {
		return fmt.Errorf("parse private key: %w", err)
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshConn, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), sshConfig)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	return New(dataSourceName)
}
