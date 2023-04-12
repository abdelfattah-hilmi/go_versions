package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type InstalledPackage struct {
	// ID               primitive.ObjectID `bson:"_id,omitempty"`
	PackageName      string `bson:"package-name,omitempty" json:"package-name" binding:"required"`
	InstalledVersion string `bson:"installed-version,omitempty" json:"installed-version" binding:"required"`
	CandidateVersion string `bson:"candidate-version,omitempty" json:"candidate-version" binding:"required"`
}

type Repology struct {
	LatestVersion string `bson:"latest-version,omitempty" json:"latest-version" binding:"required"`
	Cves          []Cve  `bson:"cves,omitempty" json:"cves" binding:"required"`
}

type Cve struct {
	CveId            string `bson:"cve-id,omitempty" json:"cve-id" binding:"required"`
	CveLink          string `bson:"cve-link,omitempty" json:"cve-link" binding:"required"`
	AffectedVersions string `bson:"affected-versions,omitempty" json:"affected-versions" binding:"required"`
}
type CollectedPackage struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	PackageName      string             `bson:"package-name,omitempty" json:"package-name" binding:"required"`
	InstalledVersion string             `bson:"installed-version,omitempty" json:"installed-version" binding:"required"`
	CandidateVersion string             `bson:"candidate-version,omitempty" json:"candidate-version" binding:"required"`
	// InstalledVersionStatus string             `bson:"installed-version-status,omitempty" json:"installes-version-status" binding:"required"`
	// InstalledVersionCves []string `bson:"installed-version-cves,omitempty" json:"installed-version-cves" binding:"required"`
	Cves               []Cve    `bson:"cves,omitempty" json:"cves" binding:"required"`
	LatestVersion      string   `bson:"latest-version,omitempty" json:"latest-version" binding:"required"`
	LatestReleaseNotes []string `bson:"latest-release-notes,omitempty" json:"latest-release-notes" binding:"required" `
}
