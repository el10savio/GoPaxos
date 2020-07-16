package proposer

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Prepare starts a Paxos round sending
// a prepare request to all the Paxos
// peers including itself
func Prepare(key string, value string) error {
	if key == "" {
		return errors.New("empty key provided")
	}

	if value == "" {
		return errors.New("empty value provided")
	}

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
			Accept(key, value, uuid)
			break
		}
	}

	if acceptedPeersCount < majorityPeersCount {
		return errors.New("majority consensus not obtained")
	}

	return nil
}

// Accept starts the accept phase sending
// an accept request to all the Paxos
// peers including itself
func Accept(key string, value string, uuid string) error {
	if key == "" {
		return errors.New("empty key provided")
	}

	if value == "" {
		return errors.New("empty value provided")
	}

	if uuid == "" {
		return errors.New("empty uuid provided")
	}

	peers := GetPeerList()

	if len(peers) == 0 {
		return errors.New("nil peers present")
	}

	acceptedPeersCount := 0
	responses := make([]bool, 0)
	majorityPeersCount := len(peers)/2 + 1

	for _, peer := range peers {
		response, err := SendAcceptRequest(peer, uuid)
		if err != nil {
			log.Fatalln("error in sending accept request to peer", peer, ":", err)
		}

		if response == http.StatusOK {
			acceptedPeersCount++
			responses = append(responses, true)
		}

		// Break when majorityPeersCount reached
		if acceptedPeersCount >= majorityPeersCount {
			Learn(key, value)
			break
		}
	}

	if acceptedPeersCount < majorityPeersCount {
		return errors.New("majority consensus not obtained")
	}

	return nil
}

// Learn in the final phase of the Paxos round
// telling all the nodes in the cluster to
// save the agreed upon data
func Learn(key string, value string) error {
	if key == "" {
		return errors.New("empty key provided")
	}

	if value == "" {
		return errors.New("empty value provided")
	}

	peers := GetPeerList()

	if len(peers) == 0 {
		return errors.New("nil peers present")
	}

	for _, peer := range peers {
		_, err := SendLearnRequest(peer, key, value)
		if err != nil {
			log.Fatalln("error in sending learn request to peer", peer, ":", err)
		}
	}

	return nil
}

// SendPrepareRequest sends the HTTP Prepare GET request to a given peer
func SendPrepareRequest(peer string, uuid string) (int, error) {
	if peer == "" {
		return 0, errors.New("empty peer provided")
	}

	if uuid == "" {
		return 0, errors.New("empty uuid provided")
	}

	url := fmt.Sprintf("http://%s.%s/prepare/%s", peer, GetNetwork(), uuid)

	return SendRequest(url)
}

// SendAcceptRequest sends the HTTP Accept GET request to a given peer
func SendAcceptRequest(peer string, uuid string) (int, error) {
	if peer == "" {
		return 0, errors.New("empty peer provided")
	}

	if uuid == "" {
		return 0, errors.New("empty uuid provided")
	}

	url := fmt.Sprintf("http://%s.%s/accept/%s", peer, GetNetwork(), uuid)

	return SendRequest(url)
}

// SendLearnRequest sends the HTTP Learn GET request to a given peer
func SendLearnRequest(peer string, key string, value string) (int, error) {
	if peer == "" {
		return 0, errors.New("empty peer provided")
	}

	if key == "" {
		return 0, errors.New("empty key provided")
	}

	if value == "" {
		return 0, errors.New("empty value provided")
	}

	url := fmt.Sprintf("http://%s.%s/learn/%s/%s", peer, GetNetwork(), key, value)

	return SendRequest(url)
}
