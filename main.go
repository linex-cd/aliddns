package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	log "github.com/sirupsen/logrus"
)

type CONFIG struct {
	AccessKeyId  string `json:"accessKeyId"`
	AccessSecret string `json:"accessSecret"`
	RecordID     string `json:"RecordID"`
}

var Config CONFIG

func main() {
	file, _ := os.Open("./config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}

	go SetDDNSService()
	for {
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func SetDDNSService() {
	var WanIP string
	var RecordIP string
	var RecordRR string

	RecordIP, RecordRR = GetAliRecordIP() // 服务器启动时，从阿里云获取一次

	for {
		WanIP = GetWanIPStr()
		log.Info("Get WAN IP: ", WanIP)
		if WanIP != "" && WanIP != RecordIP {
			log.Info("Wan IP changed. Will change the record IP.")
			err := SetDDNS(RecordRR, WanIP)
			if err == nil {
				RecordIP = WanIP
			}

		} else {
			//log.Info("Wan IP hold.")
		}
		time.Sleep(time.Duration(60) * time.Second)
	}
}

func SetDDNS(RecordRR string, wanIP string) (err error) {
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", Config.AccessKeyId, Config.AccessSecret)

	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"

	request.RecordId = Config.RecordID
	request.RR = RecordRR
	request.Type = "A"
	request.Value = wanIP //GetWanIPStr() //"118.123.37.212"
	request.Lang = "en"
	request.UserClientIp = wanIP // "118.123.37.211"
	request.TTL = "600"
	request.Priority = "1"
	request.Line = "default"

	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error(), response)
		return err
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}

func GetAliRecordIP() (recordIP string, recordRR string) {
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", Config.AccessKeyId, Config.AccessSecret)

	request := alidns.CreateDescribeDomainRecordInfoRequest()
	request.Scheme = "https"

	request.RecordId = Config.RecordID
	request.Lang = "en"
	request.UserClientIp = ""

	response, err := client.DescribeDomainRecordInfo(request)
	if err != nil {
		fmt.Print(err.Error())
		return "", ""
	}
	log.Info("Record IP: ", response.Value)
	return response.Value, response.RR
}

func GetWanIPStr() (wanip string) {
	cmd := exec.Command("wsl", "curl", "ident.me")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Error("error: ", err)
		return ""
	}
	//fmt.Printf("in all caps: %q\n", out.String())

	wanip = out.String()
	if wanip != "" {
		//log.Info("Get WAN IP ok: ", wanip)
	} else {
		log.Warn("Get WAN IP failed")
	}
	return wanip
}
