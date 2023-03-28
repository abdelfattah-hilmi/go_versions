package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type InstalledPackage struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	PackageName      string             `bson:"package-name,omitempty"`
	InstalledVersion string             `bson:"installed-version,omitempty"`
	CandidateVersion string             `bson:"title,omitempty"`
}

type CollectedPackage struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	PackageName            string             `bson:"package-name,omitempty"`
	InstalledVersion       string             `bson:"installes-version,omitempty"`
	CandidateVersion       string             `bson:"candidate-version,omitempty"`
	InstalledVersionStatus string             `bson:"installed-version-status,omitempty"`
	InstalledVersionCves   []string           `bson:"installed-version-cves,omitempty"`
	LatestVersion          string             `bson:"latest-version,omitempty"`
	LatestReleaseNotes     []string           `bson:"latest-release-notes,omitempty"`
}

var cves = []string{"cve", "cve"}
var releases = []string{"release1", "link"}

var CollectedPackageInstance = CollectedPackage{PackageName: "xyz", InstalledVersion: "slsl", CandidateVersion: "slsl1", InstalledVersionStatus: "ls", InstalledVersionCves: cves, LatestVersion: "inifinite", LatestReleaseNotes: releases}
