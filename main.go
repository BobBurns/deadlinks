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

  links, texts := collectLinks(startbody)

  type Link struct {
    uri string
    linkText string
    scode int
    status string
  }
  // Fetch 10 links at a time to ease file descriptor usage

  start := 0
  done := false
  n := len(links)
  broken := 0
  limit := 10
  for {
    if n - start < 10 {
      limit = n - start
    }
    fmt.Printf(".")
    queue := make(chan Link, limit)
    var wg sync.WaitGroup
    wg.Add(limit)

    for i := start; i<start+10; i++ {
      if i == n {
	done = true
	break
      }
      go func(ln string, txt string) {

	defer wg.Done()
	curlin := Link{
	  uri: fixUrl(ln, base),
	  linkText: txt,
	}
	if curlin.uri == "" {
	  return
	}
	//fmt.Println("fetching... ", curlin.uri)
	resp, err := fetchStatus(curlin.uri)

	if err != nil {
//	  fmt.Println("Errors found: ", err)
	  curlin.status = err.Error()
	} else {
	  curlin.status = resp.Status
	  curlin.scode = resp.StatusCode
	  // help mem leaks
	  resp.Body.Close()
	}
	queue <- curlin
      }(links[i], texts[i])


    }

    wg.Wait()
    close(queue)

    fmt.Println()
    for qlink := range queue {
      if *aFlag || qlink.scode != 200 {
	fmt.Printf("[%d]\t%s\t%s [%s]\n", qlink.scode, qlink.uri, qlink.status, qlink.linkText)
	if  qlink.scode != 200 {
	  broken++
	}
      }
    }
    if done == true {
      break
    }
    start += 10
  }

  fmt.Println()
  fmt.Println("========================")
  fmt.Printf("Fetched %d links\n ", n)
  fmt.Printf("%d links broken\n", broken)
  fmt.Printf("%d links OK\n", n-broken)

}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil || uri.Scheme == "mailto" {
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

func fetchStatus (uri string) (*http.Response , error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	timeout := time.Duration(30 * time.Second)
	client := http.Client{Transport: transport, Timeout: timeout}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		//return "", 0, err
		return nil, err
	}

	// politeness
	req.Header.Set("User-Agent", AGENT)
	resp, err := client.Do(req)
	if err != nil {
		//return "", 0, err
		return nil, err
	}

	//return resp.Status, resp.StatusCode, nil
	return resp, nil
}
