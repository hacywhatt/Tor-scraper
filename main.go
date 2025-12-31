package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func main() {

	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)

	if err != nil {

		fmt.Println("Hata: bağlanamadık", err)
		return

	}

	tr := &http.Transport{Dial: dialer.Dial}

	client := &http.Client{

		Transport: tr,
		Timeout:   30 * time.Second,
	}

	fmt.Println("[INFO] IP adresi kontrol ediliyor...")

	ipReq, _ := http.NewRequest("GET", "http://check.torproject.org/api/ip", nil)

	ipReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/110.0.0.0 Safari/537.36")

	resp, err := client.Do(ipReq)

	if err != nil {

		fmt.Println("[ERR] IP kontrolü yapılamadı:", err)

	} else {

		ipBody, _ := io.ReadAll(resp.Body)
		fmt.Println("[INFO] Şu anki IP adresiniz:", strings.TrimSpace(string(ipBody)))

		resp.Body.Close()

	}

	file, err := os.Open("hedef.yaml")

	if err != nil {

		fmt.Println("Hata: hedef.yaml dosyası bulunamadı!", err)
		return

	}

	defer file.Close()

	os.Mkdir("output", 0755)

	reportFile, _ := os.Create("scan_report.log")
	defer reportFile.Close()

	scanner := bufio.NewScanner(file)

	fmt.Print("\n Başlıyoruz\n")

	for scanner.Scan() {
		satir := scanner.Text()
		u, err := url.Parse(satir)
		if err != nil || u.Hostname() == "" {
			continue
		}
		safe := strings.ReplaceAll(u.Hostname(), ".", "_")

		fmt.Printf("[INFO] Scanning: %s ... ", u)

		req, err := http.NewRequest("GET", satir, nil)

		if err != nil {

			fmt.Println("[ERR] İstek oluşturulamadı")
			continue

		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/110.0.0.0 Safari/537.36")

		resp, err := client.Do(req)

		if err != nil {

			fmt.Println("-> [FAIL] TIMEOUT/ERROR")
			logMsg := fmt.Sprintf("[ERR] %s -> %v\n", satir, err)
			reportFile.WriteString(logMsg)

			continue

		}

		fmt.Println("-> [SUCCESS]")

		body, err := io.ReadAll(resp.Body)
		if err == nil {

			filename := fmt.Sprintf("output/%s_%d.html", safe, time.Now().Unix())
			os.WriteFile(filename, body, 0644)

		}

		resp.Body.Close()
		reportFile.WriteString(fmt.Sprintf("[OK] %s -> Active\n", u))
		calisilan_yer, _ := os.Getwd()
		resim_yolu := calisilan_yer + "\\" + safe + ".png"

		var edgeYolu = "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe" //eğer msedge.exe yolu farklıysa bu kısmı değiştirin

		cmd := exec.Command(edgeYolu,
			"--headless",
			"--disable-gpu",
			"--window-size=1920,8000",
			"--proxy-server=socks5://127.0.0.1:9150",
			"--screenshot="+resim_yolu, satir)

		cikti, hata := cmd.CombinedOutput()

		if hata != nil {

			fmt.Printf("Resim çekilemedi! Hata: %v \n\n", hata)

			fmt.Println("Detay:", string(cikti))

		} else {

			fmt.Printf("Resim başarıyla kaydedildi: %v \n\n", resim_yolu)
		}

	}

	fmt.Println("\nTarama tamamlandı.")

}
