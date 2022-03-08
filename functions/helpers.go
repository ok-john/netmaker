package functions

import (
	"encoding/json"
	"strings"

	"github.com/gravitl/netmaker/database"
	"github.com/gravitl/netmaker/models"
)

const (
	DNS_CHAR_SET = "abcdefghijklmnopqrstuvwxyz1234567889-."
)

// NetworkExists - check if network exists
func NetworkExists(name string) (bool, error) {

	var network string
	var err error
	if network, err = database.FetchRecord(database.NETWORKS_TABLE_NAME, name); err != nil {
		return false, err
	}
	return len(network) > 0, nil
}

// NameInDNSCharSet - name in dns char set
func NameInDNSCharSet(name string) bool {
	for _, char := range name {
		if !strings.Contains(DNS_CHAR_SET, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

// NameInNodeCharSet - name in node char set
func NameInNodeCharSet(name string) bool {
	for _, char := range name {
		if !strings.Contains(DNS_CHAR_SET, strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}

// RemoveDeletedNode - remove deleted node
func RemoveDeletedNode(nodeid string) bool {
	return database.DeleteRecord(database.DELETED_NODES_TABLE_NAME, nodeid) == nil
}

// GetAllExtClients - get all ext clients
func GetAllExtClients() ([]models.ExtClient, error) {
	var extclients []models.ExtClient
	collection, err := database.FetchRecords(database.EXT_CLIENT_TABLE_NAME)

	if err != nil {
		return extclients, err
	}

	for _, value := range collection {
		var extclient models.ExtClient
		err := json.Unmarshal([]byte(value), &extclient)
		if err != nil {
			return []models.ExtClient{}, err
		}
		// add node to our array
		extclients = append(extclients, extclient)
	}

	return extclients, nil
}

// DeleteKey - deletes a key
func DeleteKey(network models.Network, i int) {

	network.AccessKeys = append(network.AccessKeys[:i],
		network.AccessKeys[i+1:]...)

	if networkData, err := json.Marshal(&network); err != nil {
		return
	} else {
		database.Insert(network.NetID, string(networkData), database.NETWORKS_TABLE_NAME)
	}
}
