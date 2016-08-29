package spooler

import (
	"bufio"
	"github.com/Marneus68/gvp/config"
	"github.com/Marneus68/gvp/ps2pdf"
	"github.com/Marneus68/utils"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

func Start(con *config.Config) {
	// Create the output directory
	err := os.MkdirAll(utils.SubstituteHomeDir(con.OutDir), 0777)
	if err != nil {
		if os.IsExist(err) {
			log.Println(err.Error())
		} else {
			log.Panic(err.Error())
		}
	}
	ln, err := net.Listen("tcp", con.Port)
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go spool(con, conn, ps2pdf.Convert, func(dest string) {
			if con.Mail == true {
				log.Println(dest)
				// TODO: Send mail
			}
		})
	}
}

func spool(
	con *config.Config,
	conn net.Conn,
	psFun func(string, string, func(string)),
	mailFun func(string),
) {
	defer func() {
		log.Println("Connection closed.")
		conn.Close()
	}()
	log.Println("New incomming connection...")
	err := conn.SetReadDeadline(time.Now().Add(time.Duration(con.Timeout) * time.Second))
	if err != nil {
		log.Println(err.Error())
	}
	ts := time.Now().Format(time.UnixDate)
	scanner := bufio.NewScanner(conn)
	content := []byte("")
	for scanner.Scan() {
		content = append(content, []byte(scanner.Text()+"\n")...)
	}

	out := utils.SubstituteHomeDir(con.OutDir) + string(filepath.Separator) + ts
	pdl := out + ".pdl"

	ioutil.WriteFile(
		pdl,
		[]byte(content),
		0777,
	)

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}

	log.Println("Created temporary file in " + pdl)
	if psFun != nil {
		psFun(pdl, utils.SubstituteHomeDir(con.OutDir), mailFun)
	}
}
