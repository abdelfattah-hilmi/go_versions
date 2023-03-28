package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Vm struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Ip                string             `bson:"ip,omitempty"`
	Distro            string             `bson:"distro,omitempty"`
	DateOfExecution   string             `bson:"date-of-execution,omitempty"`
	DateOfCollection  string             `bson:"date-of-collection,omitempty"`
	InstalledPackages []InstalledPackage `bson:"installed-packages,omitempty"`
	CollectedPackages []CollectedPackage `bson:"collected-packages,omitempty"`
}
