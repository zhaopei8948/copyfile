package main

import (
	"bytes"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
	"path"
)

func handleMessage(message []byte, suffix, messageDir string) {
	buff := bytes.Buffer{}
	uid, _ := uuid.NewV4()
	nowt := time.Now()
	strtime := nowt.Format("20060102150405")
	nano := nowt.UnixNano() % nowt.Unix()

	buff.WriteString(strtime)
	buff.WriteString(strconv.FormatInt(nano, 10))
	strtime = buff.String()

	buff.Reset()
	buff.WriteString(messageDir)
	buff.WriteByte(os.PathSeparator)
	// buff.WriteString("/")
	buff.WriteString(uid.String())
	buff.WriteString("_")
	buff.WriteString(strtime)
	buff.WriteString(".writing")
	fileName := buff.String()

	buff.Reset()
	buff.WriteString(messageDir)
	buff.WriteByte(os.PathSeparator)
	// buff.WriteString("/")
	buff.WriteString(uid.String())
	buff.WriteString("_")
	buff.WriteString(strtime)
	buff.WriteString(suffix)
	finalFileName := buff.String()

	// log.Printf("filename = [%s]\n", fileName)

	if _, err := os.Stat(messageDir); os.IsNotExist(err) {
		os.MkdirAll(messageDir, 0755)
	}

	err := ioutil.WriteFile(fileName, message, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = os.Rename(fileName, finalFileName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("success create file %s\n", finalFileName)
	}
}

func readFileToMemory(srcFile string) []byte {
	content, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File contents: %s", content)
	return content
}

func main() {
	arguments := os.Args
	if len(arguments) < 4 {
		log.Printf("usage: %s <source file> <target directory> <count>", arguments[0])
		return
	}

	suffix := path.Ext(arguments[1])
	content := readFileToMemory(arguments[1])
	count, err := strconv.Atoi(arguments[3])
	if err != nil {
		log.Println("count error", err)
		return
	}

	if count <= 0 {
		log.Printf("count cannot be less than zero")
		return
	}

	for i := 0; i < count; i++ {
		handleMessage(content, suffix, arguments[2])
	}
}
