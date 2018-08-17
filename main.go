package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
	"crypto/tls"
	"io"
	"sync"
	"os"
)
const AGENT string = "friendly-link-checker/1.0"

func main() {
  aFlag := flag.Bool("a", false, "return all status codes")
  flag.Parse()

  inarg := flag.Args()

  inUrl := inarg[0]
  if inUrl == "" {
    fmt.Println("Usage: ./deadlinks [-a] <url>")
    os.Exit(-1)
  }

  // validate input url
  startURL, err := url.Parse(inUrl)
  if err != nil {
    log.Fatal("Bad URL", err)
  }

  // handle empty scheme
  if !startURL.IsAbs() {
    startURL.Scheme = "https"
  }

  // get start page
  base := startURL.String()
  startbody, err := fetchBody(base)
  if err != nil {
    log.Fatal("problem with fetchBody", err)
  }

  links := collectLinks(startbody)

  type Link struct {
    uri string
    scode int
    status string
  }

  queue := make(chan Link, len(links))
  var wg sync.WaitGroup
  wg.Add(len(links))

  for _, l := range links {
    go func(ln string) {

      defer wg.Done()
      curlin := Link{
	uri: fixUrl(ln, base),
      }
      fmt.Println("fetching... ", curlin.uri)
      s, c, err := fetchStatus(curlin.uri)

      if err != nil {
	fmt.Println("Errors found: ", err)
      }
      curlin.status = s
      curlin.scode = c
      queue <- curlin
    }(l)


  }
  wg.Wait()
  close(queue)

  for qlink := range queue {
    if *aFlag || qlink.scode != 200 {
      fmt.Printf("[%d]\t%s\t%s\n", qlink.scode, qlink.uri, qlink.status)
    }
  }


  fmt.Println("OK")

}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

func fetchBody (uri string) (io.ReadCloser, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	timeout := time.Duration(30 * time.Second)
	client := http.Client{Transport: transport, Timeout: timeout}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	// politeness
	req.Header.Set("User-Agent", AGENT)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func fetchStatus (uri string) (string, int, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	timeout := time.Duration(60 * time.Second)
	client := http.Client{Transport: transport, Timeout: timeout}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", 0, err
	}

	// politeness
	req.Header.Set("User-Agent", AGENT)
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}

	return resp.Status, resp.StatusCode, nil
}