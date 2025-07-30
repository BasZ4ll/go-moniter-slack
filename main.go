package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Config struct {
	SlackWebhook  string
	CheckInterval time.Duration
	PortsToCheck  []int
	ServiceName   []string
	Hostname      string
}

var notifiedPorts = make(map[int]bool)

func main() {
	log.Println("Starting monitoring...")
	config := Config{
		SlackWebhook:  "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXXX",
		CheckInterval: 1 * time.Minute,
		PortsToCheck:  []int{8081, 9001, 9000},
		ServiceName:   []string{"FoodCourt", "TopUp", "Sync Transaction"},
		Hostname:      "TEST-FOODCOURT",
	}
	startMonitoring(config)
}

func startMonitoring(config Config) {
	for {
		for i, port := range config.PortsToCheck {
			serviceName := config.ServiceName[i]
			if isPortOpen("localhost", port) {
				if notifiedPorts[port] {
					log.Printf("Port %d กลับมาใช้งานได้แล้ว", port)
					notifiedPorts[port] = false
					sendRecoveryNotification(config.SlackWebhook, port)
				}
			} else {
				if !notifiedPorts[port] {
					sendSlackNotification(config.SlackWebhook, port, config.Hostname, serviceName)
					notifiedPorts[port] = true
				} else {
					log.Printf("Port %d ยังไม่เปิด (แจ้งเตือนไปแล้ว)", port)
				}
			}
		}
		time.Sleep(config.CheckInterval)
	}
}

func isPortOpen(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func sendSlackNotification(webhook string, port int, hostName string, serviceName string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf(
		"🚨🚨 *Service Alert* 🚨🚨\n"+
			"• Host: `%s`\n"+
			"• Port: `%d`\n"+
			"• Service: `%s`\n"+
			"• Status: `DOWN`\n"+
			"• Time: `%s`\n"+
			"--------------------------",
		hostName, port, serviceName, currentTime,
	)
	payload := fmt.Sprintf(`{"text": "%s"}`, message)

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Printf("สร้าง HTTP request ไม่สำเร็จ: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ส่งแจ้งเตือนล้มเหลว: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("ส่งแจ้งเตือน Slack แล้ว: %s", message)
}

func sendRecoveryNotification(webhook string, port int) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf("✅ Server บน Port `%d` กลับมาใช้งานได้แล้ว! เวลา: %s", port, currentTime)
	payload := fmt.Sprintf(`{"text": "%s"}`, message)

	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Printf("สร้าง HTTP request ไม่สำเร็จ: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("ส่งแจ้งเตือนล้มเหลว: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("ส่งแจ้งเตือนการฟื้นตัวไปยัง Slack แล้ว: %s", message)
}
