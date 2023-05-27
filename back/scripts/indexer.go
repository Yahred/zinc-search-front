package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"sync"

	"github.com/jhillyerd/enmime"
	"zinc-index-back.com/config"
)

type Mail struct {
	MessageID              string `json:"message_id"`
	Date                   string `json:"date"`
	From                   string `json:"from"`
	To                     string `json:"to"`
	Subject                string `json:"subject"`
	MimeVersion            string `json:"mime_version"`
	ContentType            string `json:"content_type"`
	ContentTrasferEncoding string `json:"content_transfer_encoding"`
	XFrom                  string `json:"x_from"`
	XTo                    string `json:"x_to"`
	XCC                    string `json:"x_cc"`
	XBCC                   string `json:"x_bcc"`
	XFolder                string `json:"x_folder"`
	XOrigin                string `json:"x_origin"`
	XFileName              string `json:"x_filename"`
	Body                   string `json:"body"`
}

type BulkRequest struct {
	Index   string `json:"index"`
	Records []Mail `json:"records"`
}

var wg sync.WaitGroup

func main() {
	CreateIndex()

	cpu, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()

	IndexMails()

	runtime.GC()
	mem, err := os.Create("memory.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer mem.Close()
	if err := pprof.WriteHeapProfile(mem); err != nil {
		log.Fatal(err)
	}
}

func CreateIndex() {
	url := config.ZincSearchURL + "/api/index"

	encodedZincSearchCredentials := config.GetZincSearchCredentials()

	headers := http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Basic " + encodedZincSearchCredentials},
	}

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	req.Header = headers

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

}

func IndexMails() {
	chunkSize := 2000
	emails := make([]Mail, 0)

	err := filepath.WalkDir(config.RootDirectoryData, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		file, err := ioutil.ReadFile(path)

		content, err := enmime.ReadEnvelope(bytes.NewReader(file))

		if err != nil {
			log.Fatal(err)
			return nil
		}

		email := Mail{
			MessageID:              content.GetHeader("Message-ID"),
			Date:                   content.GetHeader("Date"),
			From:                   content.GetHeader("From"),
			To:                     content.GetHeader("To"),
			Subject:                content.GetHeader("Subject"),
			MimeVersion:            content.GetHeader("Mime-Version"),
			ContentType:            content.GetHeader("Content-Type"),
			ContentTrasferEncoding: content.GetHeader("Content-Transfer-Encoding"),
			XFrom:                  content.GetHeader("X-From"),
			XTo:                    content.GetHeader("X-To"),
			XCC:                    content.GetHeader("X-CC"),
			XBCC:                   content.GetHeader("X-BCC"),
			XFolder:                content.GetHeader("X-Folder"),
			XOrigin:                content.GetHeader("X-Origin"),
			XFileName:              content.GetHeader("X-FileName"),
			Body:                   content.Text,
		}

		if err != nil {
			log.Fatal(err)
			return nil
		}

		emails = append(emails, email)

		if len(emails) == chunkSize {
			emailsClone := make([]Mail, 0)
			emailsClone = append(emailsClone, emails...)
			emails = make([]Mail, 0)
			go SendToZincBulk(emailsClone)
		}

		return nil
	})

	wg.Wait()

	fmt.Println("Reading files completed")

	if err != nil {
		print("Me detuve")
	}

	fmt.Println("Indexing completed")
}

func SendToZincBulk(emails []Mail) {
	wg.Add(1)

	defer wg.Done()

	url := config.ZincSearchURL + "/api/_bulkv2"
	encodedZincSearchCredentials := config.GetZincSearchCredentials()

	headers := http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Basic " + encodedZincSearchCredentials},
	}

	bulkRequest := BulkRequest{
		Index:   "emails",
		Records: emails,
	}

	bulkRequestJSON, err := json.Marshal(bulkRequest)

	if err != nil {
		log.Fatal(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bulkRequestJSON))

	if err != nil {
		log.Fatal(err)
		return
	}

	req.Header = headers

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()
	fmt.Print(resp)

}

func SendToZincSearch(email []byte) {
	url := config.ZincSearchURL + "/emails/_doc"

	encodedZincSearchCredentials := config.GetZincSearchCredentials()

	headers := http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Basic " + encodedZincSearchCredentials},
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(email))

	if err != nil {
		log.Fatal(err)
		return
	}

	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(resp)
}
