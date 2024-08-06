package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type SClient struct {
	url    string
	client *http.Client
}

func NewSOAPClient(url string, timeout time.Duration) *SClient {
	return &SClient{
		url: url,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (s *SClient) Call(action string, request interface{}, response interface{}) error {
	envelope := struct {
		XMLName   xml.Name `xml:"soap:Envelope"`
		XmlnsXsi  string   `xml:"xmlns:xsi,attr"`
		XmlnsXsd  string   `xml:"xmlns:xsd,attr"`
		XmlnsSoap string   `xml:"xmlns:soap,attr"`
		Body      struct {
			Content interface{} `xml:",any"`
		} `xml:"soap:Body"`
	}{
		XmlnsXsi:  "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsXsd:  "http://www.w3.org/2001/XMLSchema",
		XmlnsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: struct {
			Content interface{} `xml:",any"`
		}{Content: request},
	}

	payload, err := xml.Marshal(envelope)
	if err != nil {
		return fmt.Errorf("error marshalling SOAP request: %w", err)
	}

	req, err := http.NewRequest("POST", s.url, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("SOAPAction", action)

	res, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error executing HTTP request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	log.Println("response: ", string(body))
	if err != nil {
		return fmt.Errorf("error reading HTTP response: %w", err)
	}

	respEnvelope := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			Content interface{} `xml:",any"`
		} `xml:"Body"`
	}{
		Body: struct {
			Content interface{} `xml:",any"`
		}{Content: response},
	}

	err = xml.Unmarshal(body, &respEnvelope)
	log.Println("response: ", respEnvelope)
	if err != nil {
		return fmt.Errorf("error unmarshalling SOAP response: %w", err)
	}

	return nil
}
