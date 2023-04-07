package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type InstalledPackage struct {
	// ID               primitive.ObjectID `bson:"_id,omitempty"`
	PackageName      string `bson:"package-name,omitempty" json:"package-name" binding:"required"`
	InstalledVersion string `bson:"installed-version,omitempty" json:"installed-version" binding:"required"`
	CandidateVersion string `bson:"candidate-version,omitempty" json:"candidate-version" binding:"required"`
}

type CollectedPackage struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	PackageName            string             `bson:"package-name,omitempty" json:"package-name" binding:"required"`
	InstalledVersion       string             `bson:"installed-version,omitempty" json:"installed-version" binding:"required"`
	CandidateVersion       string             `bson:"candidate-version,omitempty" json:"candidate-version" binding:"required"`
	InstalledVersionStatus string             `bson:"installed-version-status,omitempty" json:"installes-version-status" binding:"required"`
	InstalledVersionCves   []string           `bson:"installed-version-cves,omitempty" json:"installed-version-cves" binding:"required"`
	LatestVersion          string             `bson:"latest-version,omitempty" json:"latest-version" binding:"required"`
	LatestReleaseNotes     []string           `bson:"latest-release-notes,omitempty" json:"latest-release-notes" binding:"required" `
}

var cves = []string{"cve", "cve"}
var releases = []string{"release1", "link"}

var CollectedPackageInstance = CollectedPackage{PackageName: "xyz", InstalledVersion: "slsl", CandidateVersion: "slsl1", InstalledVersionStatus: "ls", InstalledVersionCves: cves, LatestVersion: "inifinite", LatestReleaseNotes: releases}
