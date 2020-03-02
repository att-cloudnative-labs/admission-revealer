package main

import (
	"encoding/json"
	"log"
	"net/http"

	"k8s.io/api/admission/v1"
)

func main() {
	http.HandleFunc("/webhook", handleAdmissionRequest)
	err := http.ListenAndServeTLS(":8443", "/etc/certs/cert.pem", "/etc/certs/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func handleAdmissionRequest(w http.ResponseWriter, req *http.Request) {
	var admissionRev v1.AdmissionReview
	err := json.NewDecoder(req.Body).Decode(&admissionRev)
	if err != nil {
		log.Printf("error reading admission request body: %s", err.Error())
		return
	}

	requestJson, err := json.Marshal(admissionRev)
	if err != nil {
		log.Panicf("error marshalling admission request: %s", err.Error())
		return
	}
	log.Printf("%s", requestJson)

	w.Header().Set("Content-Type", "application/json")
	admissionResp := v1.AdmissionResponse{UID: admissionRev.Request.UID, Allowed: true}

	if err := json.NewEncoder(w).Encode(admissionResp); err != nil {
		log.Printf("error writing admission response: %s", err.Error())
		return
	}
}
