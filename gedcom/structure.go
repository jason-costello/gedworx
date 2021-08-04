package gedcom
// https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#IDENTIFIER_STRUCTURE
import (
	"time"
)

// Level Markers:
//  0 means 'must be a record'
//  n means 'level inherited from rule instantiation'
//  +1, +2, etc, indicate nesting withing nearest preceding structure with lesser level

// Four cardinality markers are used: {0:1}, {1:1}, {0:M}, and {1:M}.

type Dataset struct {
	// Header - {1:1} Dataset must include one and only one header
	Header Header
	// Records - {0:M} Dataset {0:M} records
	Records []Record
	// Trailer - {1:1} Dataset must include one and only one trailer
	Trailer Trailer
}

type Header struct {
	// Gedcom - {1:1}  https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#GEDC
	Gedcom Gedcom
	// Schema - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#SCHMA
	Schema []Schema
	// HeadSource - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#HEAD-SOUR
	Source HeadSource
	// Destination - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#DEST
	Destination *string
	// Date - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#HEAD-DATE
	Date *DateExact
	// Submitter - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#SUBM
	Submitter *SubmitterRecord
	// Copyright - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#COPR
	Copyright *string
	// Language - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#HEAD-LANG
	Language  *Language
	// Place -  {0:1}  https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#HEAD-PLAC
	// https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#HEAD-PLACHEAD-PLAC-FORM
	Place *Form
	// Notes - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#NOTE_STRUCTURE
	Notes []Note
}


// Form follows the iana media-types document
// https://www.iana.org/assignments/media-types/media-types.xhtml
type Form struct{
	Name string
	Template string
	Reference string
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
	// Version - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#VERS
	Version *string
	// Name - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#NAME
	Name  *string
	// Corporation - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#NAME
	Corporation *Corporation
	// Data - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#Head-sour-data
	Data *HeadSourceData
}

type SubmitterRecord struct {
	// Name - {1:1} - g7:NAME
	Name           string
	// Address - {0:1}
	Address        *Address
	// Phone - {0:M} g7:PHON
	Phone          []Phone
	// Email - {0:M} g7:EMAIL
	Email          []Email
	// Fax - {0:M} g7:FAX
	Fax            []Fax
	// WWW - {0:M} g7:WWW
	WWW            []WWW
	// MultimediaLink - {0:M} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#MULTIMEDIA_LINK
	MultimediaLink []Object
	// Languages - {0:M} g7:SUBM-LANG
	Languages      []Language
	// Identifiers - {0:M} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#IDENTIFIER_STRUCTURE
	Identifiers    []Identifier
	// Notes - {0:M} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#NOTE_STRUCTURE
	Notes          []Note
	// ChangeDates - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#CHANGE_DATE
	ChangeDates    *ChangeDate
	// CreationDates - {0:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#CREATION_DATE
	CreationDate  *CreationDate
}

// Language  g7:LANG
// https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry
type Language struct {
	Type           string
	Subtag         string
	Description    string
	Added          string
	SuppressScript *string
	Scope          *string
}

type Identifier struct {
	// Reference - {1:1} g7:REFN
	Reference Reference
	// UniqueID - {1:1} g7:UID
	UniqueID string
	// ExternalID - {1:1} g7:EXID
	ExternalID ExternalID
}

// ExternalID - g7:EXID
type ExternalID struct {

	Text string
	// Type - {0:1} g7:EXID-TYPE
	Type *string
}

// Reference - g7:REFN
type Reference struct {
	Text string
	// Type - {0:1} g7 - REFN-TYPE
	Type *string
}
// CreationDate https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#CREATION_DATE
type CreationDate DateExact

type ChangeDate struct {
	// Date - {1:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#date
	Date DateExact
	// Notes - {0:M} - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#NOTE_STRUCTURE
	Notes []Note
}

// Note - g7:NOTE
type Note struct {
	Text string
	// MIME - {0:1} g7:MIME
	MIME []string
	// Language - {0:1} g7:LANG
	Language  *Language
	// Translation - {0:1} g7:NOTE-TRAN
	Translation NoteTranslation
}

// SharedNote - {1:1} g7:SNOTE
type SharedNote *Note

// NoteTranslation - g7:NOTE-TRAN
type NoteTranslation struct {
	// Language is the language the translation is in
	Language *string
	// MIME is the MIME type of the translation if translating between types
	// example is going from text/html to text/plain
	MIME *string
	// Translated text
	Translation string
}

// Object - @<XREF:OBJE>@g7:OBJE
type Object struct {
	// Crop {0:1} g7:GROP
	Crop  *Crop
	Title *string
}
// Crop - g7:CROP
type Crop struct {
	// Top - {0:1} g7:TOP
	Top    *int
	// Left - {0:1} g7:LEFT
	Left   *int
	// Height - {0:1} g7:HEIGHT
	Height *int
	// Width - {0:1} g7:WIDTH
	Width  *int
}
type HeadSourceData struct {
	// Date - {0:1} g7:HEAD-DATE
	Date        *DateExact
	// Corporation - {0:1} g7:CORP
	Corporation *Corporation
}

// DateExact - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#date
type DateExact time.Time

// Corporation - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#CORP
type Corporation struct {
	// Address - {0:1} g7:ADDR
	Address []Address
	// Phone - {0:M}  g7:PHONE
	Phone   []Phone
	// Email - {0:M}  g7:EMAIL
	Email   []Email
	// Fax - {0:M}  g7:FAX
	Fax     []Fax
	// WWW - {0:M} g7:WWW
	WWW     []WWW
}

// Address - g7:ADDR
type Address struct {
	// AddressOne - {0:1} g7:ADR1
	AddressOne string
	// AddressTwo - {0:1} g7:ADR2
	AddressTwo string
	// AddressThree - {0:1} g7:ADR3
	AddressThree *string
	// City - {0:1} g7:CITY
	City *string
	// State - {0:1} g7:STAE
	State *string
	// Post - {0:1} g7:POST
	Post *string
	// Country - {0:1} g7:CTRY
	Country *string
}



// Phone - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#PHON
type Phone string

// Email - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#EMAIL
type Email string

// Fax - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#FAX
type Fax string

// WWW - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#WWW
type WWW string

// Schema - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#SCHMA
type Schema struct {
	// ExtensionTags - {0:M} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#TAG
	ExtensionTags []ExtensionTag
}

// ExtensionTag - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#TAG
type ExtensionTag struct{}

// Gedcom - https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#GEDC
type Gedcom struct {
	// Version:  {1:1} https://gedcom.io/specifications/FamilySearchGEDCOMv7.html#GEDC-VERS
	Version string
}

// type Record struct{}

type Trailer struct{}
