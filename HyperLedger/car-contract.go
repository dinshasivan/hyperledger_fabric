package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CarContract provides functions for managing a Car asset
type CarContract struct {
	contractapi.Contract
}

// Car represents the structure of a car asset
type Car struct {
	AssetType        string `json:"AssetType"`
	CarId            string `json:"CarId"`
	Color            string `json:"Color"`
	DateOfManufacture string `json:"DateOfManufacture"`
	OwnedBy          string `json:"OwnedBy"`
	Make             string `json:"Make"`
	Model            string `json:"Model"`
	Status           string `json:"Status"`
}

// CarExists returns true when a car with given ID exists in the world state
func (c *CarContract) CarExists(ctx contractapi.TransactionContextInterface, carID string) (bool, error) {
	data, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return data != nil, nil
}

// CreateCar creates a new instance of Car
func (c *CarContract) CreateCar(ctx contractapi.TransactionContextInterface, carID string, make string, model string, color string, manufacturerName string, dateOfManufacture string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID != "Org1MSP" {
		return "", fmt.Errorf("User under MSPID %v can't perform this action", clientOrgID)
	}

	exists, err := c.CarExists(ctx, carID)
	if err != nil {
		return "", fmt.Errorf("Could not read from world state: %v", err)
	}
	if exists {
		return "", fmt.Errorf("The asset %s already exists", carID)
	}

	car := Car{
		AssetType:        "car",
		CarId:            carID,
		Color:            color,
		DateOfManufacture: dateOfManufacture,
		Make:             make,
		Model:            model,
		OwnedBy:          manufacturerName,
		Status:           "In Factory",
	}

	bytes, err := json.Marshal(car)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal car object: %v", err)
	}

	err = ctx.GetStub().PutState(carID, bytes)
	if err != nil {
		return "", fmt.Errorf("Failed to put car in world state: %v", err)
	}

	return fmt.Sprintf("Successfully added car %s", carID), nil
}

// ReadCar retrieves an instance of Car from the world state
func (c *CarContract) ReadCar(ctx contractapi.TransactionContextInterface, carID string) (*Car, error) {
	exists, err := c.CarExists(ctx, carID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state: %v", err)
	}
	if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", carID)
	}

	bytes, err := ctx.GetStub().GetState(carID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get car from world state: %v", err)
	}

	var car Car
	err = json.Unmarshal(bytes, &car)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Car: %v", err)
	}

	return &car, nil
}

// DeleteCar removes the instance of Car from the world state
func (c *CarContract) DeleteCar(ctx contractapi.TransactionContextInterface, carID string) (string, error) {

clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
if err != nil {
return "", err
}
if clientOrgID == "Org1MSP" {

exists, err := c.CarExists(ctx, carID)
if err != nil {
return "", fmt.Errorf("Could not read from world state. %s", err)
} else if !exists {
return "", fmt.Errorf("The asset %s does not exist", carID)
}

err = ctx.GetStub().DelState(carID)
if err != nil {
return "", err
} else {
return fmt.Sprintf("Car with id %v is deleted from the world state.", carID), nil
}

} else {
return "", fmt.Errorf("User under following MSP:%v cannot able to perform this action", clientOrgID)
}
}

