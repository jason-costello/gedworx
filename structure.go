package gedworx

import (
	"time"
)

// Level Markers:
//  0 means 'must be a record'
//  n means 'level inherited from rule instantiation'
//  +1, +2, etc, indicate nesting withing nearest preceding structure with lesser level

// Four cardinality markers are used: {0:1}, {1:1}, {0:M}, and {1:M}.

type Dataset struct {
	Header  Header
	Records []Record
	Trailer Trailer
}

type Header struct {
	Gedcom Gedcom
	Schema []Schema
	Source HeadSource
	// Destination - An identifier for the system expected to receive this document. See HEAD.SOUR for guidance on choosing identifiers.
	Destination *string
	// Date -The DateExact that this document was created.
	Date *DateExact
	// Submitter - A contributor of information in the substructure. This is metadata about the structure itself, not data about its subject.
	// @<XREF:SUBM>@
	// @<XREF:tag>@ means a pointer to a structure with this cross-reference template; @VOID@ is also permitted.
	Submitter *SubmitterRecord
	Copyright *string
	Language  *Language
	// This is a placeholder for providing a default PLAC.FORM, and must not have a payload.
	Place *Place
	Notes []Note
}

type Place struct {
	City    *string
	County  *string
	State   *string
	Country *string
}

// HeadSource An identifier for the product producing this dataset.
//A registration process for these identifiers existed for a time, but no longer
//does. If an existing identifier is known, it should be used. Otherwise,
//a URI owned by the product should be used instead.
type HeadSource struct {
	// Version g7:VERS
	// An identifier that represents the version level assigned to the associated product. It is defined and changed by the creators of the product.
	Version *string
	// Name g7:NAME
	// The name of the superstructure’s subject, represented as a simple string.
	Name        *string
	Corporation *Corporation
	// Data g7:Head-sour-data
	// The electronic data source or digital repository from which this dataset was exported. The payload is the name of that source, with substructures providing additional details about the source (not the export).
	Data *HeadSourceData
}

type SubmitterRecord struct {
	Name           string
	Address        []Address
	Phone          []Phone
	Email          []Email
	Fax            []Fax
	WWW            []WWW
	MultimediaLink []Object
	Languages      []Language
	Identifiers    []Identifier
	Notes          []Note
	ChangeDates    []ChangeDate
	CreationDates  []CreationDate
}

// Language datatype represents a human language or family of related languages, as defined in BCP 46. It consists of a sequence of language subtags separated by hyphens, where language subtags are registered by the IANA.
// The ABNF grammar for language tags is given in BCP 47, section 2.1, production Language-Tag.
// The URI for the Language datatype is xsd:Language.
type Language struct {
	Type           string
	Subtag         string
	Description    string
	Added          string
	SuppressScript *string
	Scope          *string
}

type Identifier struct {
	// Reference is a user-generated identifier for a structure.
	Reference Reference
	// UniqueID is a globally-unique identifier for a structure.
	UniqueID string
	// ExternalID is an identifier maintained by an external authority that applies to the subject of the structure.
	ExternalID ExternalID
}

// ExternalID - An identifier for the subject of the superstructure. The identifier is maintained by some external authority; the authority owning the identifier is provided in the TYPE substructure; see EXID.TYPE for more.
//Depending on the maintaining authority, an EXID may be a unique identifier for the subject, an identifier for 1 of several views of the subject, or an identifier for the externally-maintained copy of the same information as is contained in this structure. However, unlike UID and REFN, EXID does not identify a structure; structures with the same EXID may have originated independently rather than by edits from the same starting point.
type ExternalID struct {
	Text string
	Type *string
}

// Reference - A user-defined number or text that the submitter uses to identify the superstructure. For instance, it may be a record number within the submitter’s automated or manual system, or it may be a page and position number on a pedigree chart.
//This is metadata about the structure itself, not data about its subject. Multiple structures describing different aspects of the same subject must not have the same REFN value.
type Reference struct {
	Text string
	Type *string
}

type CreationDate DateExact
type ChangeDate struct {
	Date  DateExact
	Notes []Note
}

type Note struct {
	Text      string
	MediaType []string
	Language  []Language
}
type NoteTranslation struct {
	Name string
}
type Object struct {
	Crop  *Crop
	Title *string
}
type Crop struct {
	Top    *int
	Left   *int
	Height *int
	Width  *int
}
type HeadSourceData struct {
	// Date - The principal date of the subject of the superstructure. The payload is a DateExact.
	Date        *DateExact
	Corporation *Corporation
}

// DateExact is used for timestamps and other fully-known dates.
type DateExact time.Time

// Corporation g7:CORP
// The name of the business, corporation, or person that produced or commissioned the product.
type Corporation struct {
	Address []Address
	Phone   []Phone
	Email   []Email
	Fax     []Fax
	WWW     []WWW
}

type Address struct {
	// AddressOne - The first line of the address, used for indexing. This is the value of the line corresponding to the ADDR tag line in the address structure
	AddressOne string
	// AddressTwo - The first line of the address, used for indexing. This is the value of the line corresponding to the ADDR tag line in the address structure
	AddressTwo string
	// AddressThree - The third line of the address, used for indexing. This is the value of the second CONT line subordinate to the ADDR tag in the address structure. See ADDRESS_STRUCTURE for more.
	AddressThree *string
	// The name of the city used in the address.
	City *string
	// A geographical division of a larger jurisdictional area, such as a state within the United States of America
	State *string
	// A code used by a postal service to identify an area to facilitate mail handling.
	Post *string
	// The name of the country that pertains to the associated address
	Country *string
}

// Phone - A telephone number. Telephone numbers have many regional variations
//and can contain non-digit characters. Users should be encouraged to use
//internationalized telephone numbers rather than local versions.
//As a starting point for this recommendation, there are international
//standards that use a “‘+’” shorthand for the international prefix
//(for example, in place of “011” in the US or “00” in the UK).
//Examples are +1 (555) 555-1234 (US) or +44 20 1234 1234 (UK).
type Phone string

// Email An electronic mail address, as defined by any relevant standard such as RFC 3696, RFC 5321, or RFC 5322.
//If an invalid email address is present upon import, it should be preserved as-is on export.
//The version 5.5.1 specification contained a typo where this tag was sometimes written EMAI and sometimes written EMAIL. EMAIL should be used in version 7.0 and later.
type Email string

// Fax - A fax telephone number appropriate for sending data facsimiles. See PHON for more.
type Fax string

// WWW - A URL or other locator for a World Wide Web page, as defined by any relevant standard such as whatwg/url, RFC 3986, RFC 3987, and so forth.
//
//If an invalid or no longer existing web address is present upon import, it should be preserved as-is on export.
type WWW string

// Schema - A container for storing meta-information about the extension tags
// used in this document.
type Schema struct {
	ExtensionTags []ExtensionTag
}

type ExtensionTag struct{}

// Gedcom - GEDC (GEDCOM) g7:GEDC
// A container for information about the entire document.
// It is recommended that applications write GEDC
// with its required subrecord VERS as the first substructure of HEAD.
type Gedcom struct {
	// Version: VERS (Version) g7:GEDC-VERS
	// The version number of the official specification that this document’s
	// data conforms to. This must include the major and minor version
	// (for example, “7.0”); it may include the patch as
	// well (for example, “7.0.1”), but doing so is not required.
	// See A Guide to Version Numbers for more.
	Version string
}

type Record struct{}

type Trailer struct{}
