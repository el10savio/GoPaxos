package proposer

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

// Prepare ...
func Prepare(value string) error {
	uuid := GenerateUUID()
	peers := GetPeerList()

	if len(peers) == 0 {
		return errors.New("nil peers present")
	}

	acceptedPeersCount := 0
	responses := make([]bool, 0)
	majorityPeersCount := len(peers)/2 + 1

	for _, peer := range peers {
		response, err := SendPrepareRequest(peer, uuid)
		if err != nil {
			log.Fatalln("error in sending prepare request to peer", peer, ":", err)
		}

		if response == http.StatusOK {
			acceptedPeersCount++
			responses = append(responses, true)
		}

		// Break when majorityPeersCount reached
		if acceptedPeersCount >= majorityPeersCount {
			Accept(value, uuid)
			break
		}
	}

	if acceptedPeersCount < majorityPeersCount {
		return errors.New("majority consensus not obtained")
	}

	return nil
}

// Accept ...
func Accept(value string, uuid string) error {
	peers := GetPeerList()

	if len(peers) == 0 {
		return errors.New("nil peers present")
	}

	acceptedPeersCount := 0
	responses := make([]bool, 0)
	majorityPeersCount := len(peers)/2 + 1

	for _, peer := range peers {
		response, err := SendAcceptRequest(peer, value, uuid)
		if err != nil {
			log.Fatalln("error in sending accept request to peer", peer, ":", err)
		}

		if response == http.StatusOK {
			acceptedPeersCount++
			responses = append(responses, true)
		}

		// Break when majorityPeersCount reached
		if acceptedPeersCount >= majorityPeersCount {
			Learn(value, uuid)
			break
		}
	}

	if acceptedPeersCount < majorityPeersCount {
		return errors.New("majority consensus not obtained")
	}

	return nil
}

// Learn ...
func Learn(value string, uuid string) error {
	peers := GetPeerList()

	if len(peers) == 0 {
		return errors.New("nil peers present")
	}

	for _, peer := range peers {
		_, err := SendLearnRequest(peer, value)
		if err != nil {
			log.Fatalln("error in sending learn request to peer", peer, ":", err)
		}
	}

	return nil
}

// SendPrepareRequest ...
func SendPrepareRequest(peer string, uuid string) (int, error) {
	if peer == "" {
		return 0, errors.New("empty peer provided")
	}

	if uuid == "" {
		return 0, errors.New("empty uuid provided")
	}

	url := "http://" + strings.TrimSpace(peer) + "." + strings.TrimSpace(GetNetwork()) + "/prepare/" + strings.TrimSpace(uuid)

	return SendRequest(url)
}

// SendAcceptRequest ...
func SendAcceptRequest(peer string, value string, uuid string) (int, error) {
	if peer == "" {
		return 0, errors.New("empty peer provided")
	}

	if value == "" {
		return 0, errors.New("empty value provided")
	}

	if uuid == "" {
		return 0, errors.New("empty uuid provided")
	}

	url := "http://" + strings.TrimSpace(peer) + "." + strings.TrimSpace(GetNetwork()) + "/accept/" + strings.TrimSpace(uuid) + "/" + strings.TrimSpace(value)

	return SendRequest(url)
}

// SendLearnRequest ...
func SendLearnRequest(peer string, value string) (int, error) {
	if peer == "" {
		return 0, errors.New("empty peer provided")
	}

	if value == "" {
		return 0, errors.New("empty value provided")
	}

	url := "http://" + strings.TrimSpace(peer) + "." + strings.TrimSpace(GetNetwork()) + "/learn/" + strings.TrimSpace(value)

	return SendRequest(url)
}
