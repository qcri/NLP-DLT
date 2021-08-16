package main

import (
	"errors"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func (cc *Chaincode) verifyOrg(stub shim.ChaincodeStubInterface, id string)(bool, error) {
	bytes, err := stub.GetState(id)
	if bytes != nil {
		log.Errorf("[%s][%s][verifyOrg] The identity already exists", CHANNEL_ENV, IDREGISTRY)
		return true, err
	}
	if err != nil {
		log.Errorf("[%s][%s][verifyOrg] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
		return false, err
	}
	log.Info("[%s][%s][verifyOrg] The must be registered", CHANNEL_ENV, IDREGISTER)
	return false, err
}

func (cc *Chaincode) recordOrg(stub shim.ChaincodeStubInterface, organization Organization, id string)(error) {
	idBytes, err := json.Marshal(organization)
	if err != nil {
		log.Errorf("[%s][%s][recordOrg] Error parsing: %v", CHANNEL_ENV, ERRORParsingOrg, err.Error())
		return errors.New(ERRORParsingID + err.Error())
	}

	err = stub.PutState(id, idBytes) // PuState of Client (Organization) Identity and Organtization struct
	if err != nil {
		log.Errorf("[%s][%s][recordOrg] Error storing: %v", CHANNEL_ENV, ERRORStoringOrg, err.Error())
		return errors.New(ERRORStoringIdentity + err.Error())
	}

	store := make(map[string]Organization)
	store["org"] = organization
	err = stub.PutState(store["org"].mno_name, []byte(id)) // PuState of Client (Organization) Name and Organtization Identity
	if err != nil {
		log.Errorf("[%s][%s][recordOrg] Error storing: %v", CHANNEL_ENV, ERRORStoringOrg, err.Error())
		return errors.New(ERRORStoringIdentity + err.Error())
	}
	return nil
}

func (cc *Chaincode) recoverOrgId(stub shim.ChaincodeStubInterface, organization Organization)(string, error) {
	store := make(map[string]Organization)
	store["org"] = organization
	id_org_bytes, err := stub.GetState(store["org"].mno_name)
	id_org := string([]byte(id_org_bytes))
	if err != nil {
		log.Errorf("[%s][%s][recoverOrgId] Error recovering: %v", CHANNEL_ENV, ERRORRecoveringOrg, err.Error())
		return "", err
	}
	if id_org_bytes == nil {
		log.Errorf("[%s][%s][recoverOrgId] The identity not exists", CHANNEL_ENV, IDREGISTRY)
		return "", err
	}
	return id_org, nil
}