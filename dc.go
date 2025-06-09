package main

import (
	"time"
)

type DublinCore struct {
	// An entity responsible for making contributions to the resource
	// Examples of a contributor include a person, an organization, or a service.
	// Typically, the name of a contributor should be used to indicate the entity.
	Contributor []string `xml:"contributor" json:"contributor"`
	// The spatial or temporal topic of the resource, the spatial applicability of the resource,
	// or the jurisdiction under which the resource is relevant.
	Coverage string `xml:"coverage" json:"coverage"`
	// An entity primarily responsible for making the resource.
	// Examples of a creator include a person, an organization, or a
	// service. Typically, the name of a creator should be used to indicate the entity.
	Creator []string `xml:"creator" json:"creator"`
	// A point or period of time associated with an event in the life cycle of the resource.
	Date time.Time `xml:"date" json:"date"`

	// An account of the resource.
	Description []string `xml:"description" json:"description"`

	// The size or duration of the resource.
	// Recommended practice is to specify the file size in megabytes and duration in ISO 8601 format.
	// For images media type
	Format string `xml:"format" json:"format"`

	// A related resource that is substantially the same as the pre-existing described resource, but in another format.
	HasFormat []string `xml:"hasFormat" json:"has_format"`

	// A related resource that is included either physically or logically in the described resource.
	HasPart []string `xml:"hasPart" json:"has_part"`

	// An unambiguous reference to the resource within a given context.
	Identifier string `xml:"identifier" json:"identifier"`

	// A pre-existing related resource that is substantially the same as the described resource, but in another format.
	IsFormatOf []string `xml:"isFormatOf" json:"is_format_of"`

	// A related resource in which the described resource is physically or logically included.
	IsPartOf []string `xml:"isPartOf" json:"is_part_of"`

	// A language of the resource.
	Language []string `xml:"language" json:"language"`
	// An entity responsible for making the resource available
	// Examples of a publisher include a person, an organization, or a
	// service. Typically, the name of a publisher should be used to indicate the entity.

	// Date on which the resource was changed.
	Modified time.Time `xml:"modified" json:"modified"`

	Publisher []string `xml:"language" json:"publisher"`
	// A related resource.
	// Recommended best practice is to identify the related resource
	// by means of a string conforming to a formal identification system.
	Relation []string `xml:"language" json:"relation"`
	// Information about rights held in and over the resource.
	// typically, rights information includes a statement about various property
	// rights associated with the resource, including intellectual property rights.
	Rights []string `xml:"rights" json:"rights"`
	// A related resource from which the described resource is derived.
	// The described resource may be derived from the related resource in whole or in part.
	// Recommended best practice is to identify the related resource by means of a string
	// conforming to a formal identification system.
	Source string `xml:"source" json:"source"`
	// The topic of the resource.
	// Typically, the subject will be represented using keywords, key phrases, or
	// classification codes. Recommended best practice is to use a controlled vocabulary.
	// To describe the spatial or temporal topic of the resource, use the dc:coverage element.
	Subject []string `xml:"subject" json:"subject"`
	// A name given to the resource.
	// Typically, a title will be a name by which the resource is formally known.
	Title []string `xml:"title" json:"title"`

	// The nature or genre of the resource.
	Type []string `xml:"type" json:"type"`

	// A list of subunits of the resource
	TableOfContents []string `xml:"tableOfContents" json:"table_of_contents"`
}
