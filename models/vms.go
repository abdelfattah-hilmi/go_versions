package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vm struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Ip                string             `bson:"ip,omitempty" json:"ip" binding:"required"`
	Distro            []string           `bson:"distro,omitempty" json:"distro" binding:"required"`
	DateOfExecution   string             `bson:"date-of-execution,omitempty" json:"date-of-execution" binding:"required"`
	DateOfCollection  string             `bson:"date-of-collection,omitempty" json:"date-of-collection" `
	InstalledPackages []InstalledPackage `bson:"installed-packages,omitempty" json:"installed-packages" `
	CollectedPackages []CollectedPackage `bson:"collected-packages,omitempty" json:"collected-packages" `
}
