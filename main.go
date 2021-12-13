package main

import(
	"fmt"
	"net/http"
	"crypto/tls"
	"time"
	"net"
	"sync"
	"os"
	"bufio"
	"flag"
	"log"
	"net/http/httputil"
)

var httpClient = &http.Client{
	Transport: transport,
}

var transport = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: time.Second,
		DualStack: true,
	}).DialContext,
}


func main() {
	var headers string
	var payloads string
	flag.StringVar(&headers, "hf" , "" , "Set the headers file")
	flag.StringVar(&payloads, "p" , "" , "Set the payload file")
	flag.Parse()

	var urls []string

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		urls = append(urls, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n")
	}
	results := make(chan string)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			request(url, headers , payloads)
			defer wg.Done()
		}(url)
	}
	wg.Wait()
	close(results)
}
func request(urls string , headers string , payloads string) {

	file , err := os.Open(headers)

	if err != nil {
		log.Fatal("File could not be read")
	}
	
	defer file.Close()

	time.Sleep(time.Millisecond * 10)

	hScanner := bufio.NewScanner(file)

	var lines []string
	for hScanner.Scan() {
		lines = append(lines, hScanner.Text())
	}

	if err := hScanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n")
	}

	payload_file , err := os.Open(payloads)

	if err !=nil {
		log.Fatal("File could not be read")
	}

	defer payload_file.Close()

	time.Sleep(time.Millisecond * 10)
	pScanner := bufio.NewScanner(payload_file)

	var links []string
	for pScanner.Scan() {
		links = append(links, pScanner.Text())
	}

	if err := pScanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n")
	}

	for _, header := range lines {
		for _, payload := range links {
			req, err:= http.NewRequest("GET" , urls , nil)
			req.Header.Add("User-Agent", "User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36")
			req.Header.Add(header,payload)
			fmt.Printf("[+] Testing: \t %s\n",header)

			if err != nil {
				return
			}
			resp, err := httpClient.Do(req)
			if err != nil {
				return
			}

			res, err := httputil.DumpRequest(req, true)  
 			if err != nil {  
   				log.Fatal(err)  
			}  
			fmt.Print(string(res))
			fmt.Println(resp.StatusCode)
			}
		}
	
}
