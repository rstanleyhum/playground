package table

import (
	"_snippets/golang/crown/crownsql"
	"time"

	"database/sql"
)

// Document stores the Document record in from the Document table in Crown
type Document struct {
	DocumentID                crownsql.UniqueID      `db:"DocumentID"`
	EXTDocumentID             sql.NullString         `db:"EXTDocumentID"`
	EXTSequenceID             sql.NullString         `db:"EXTSequenceID"`
	PatientID                 crownsql.NullUniqueID  `db:"PatientID"`
	AuthorID                  crownsql.NullUniqueID  `db:"AuthorID"`
	IDXDocumentTypeDE         crownsql.DictID        `db:"IDXDocumentTypeDE"`
	DocumentTypeDE            crownsql.DictID        `db:"DocumentTypeDE"`
	PhysicalManifestationDE   crownsql.DictID        `db:"PhysicalManifestationDE"`
	NextAmendmentNumber       int                    `db:"NextAmendmentNumber"`
	EditableChunk             sql.NullString         `db:"EditableChunk"`
	UneditableChunk           sql.NullString         `db:"UneditableChunk"`
	RecordedDTTM              time.Time              `db:"RecordedDTTM"`
	LastEditDTTM              crownsql.NullDatetime  `db:"LastEditDTTM"`
	LastEditID                crownsql.NullUniqueID  `db:"LastEditID"`
	VerifiedDTTM              crownsql.NullDatetime  `db:"VerifiedDTTM"`
	VerifiedID                crownsql.NullUniqueID  `db:"VerifiedID"`
	EIEDTTM                   crownsql.NullDatetime  `db:"EIEDTTM"`
	EIEID                     crownsql.NullUniqueID  `db:"EIEID"`
	PreviousVersionID         crownsql.UniqueID      `db:"PreviousVersionID"`
	NextVersionID             crownsql.UniqueID      `db:"NextVersionID"`
	AccessionNumber           sql.NullString         `db:"AccessionNumber"`
	EncounterID               crownsql.UniqueID      `db:"EncounterID"`
	Transcriptionist          sql.NullString         `db:"Transcriptionist"`
	TranscribedDTTM           crownsql.NullDatetime  `db:"TranscribedDTTM"`
	DictatedDTTM              crownsql.NullDatetime  `db:"DictatedDTTM"`
	VerificationRequiredFLAG  crownsql.NullBool      `db:"VerificationRequiredFLAG"`
	BinaryChunk               crownsql.NullVarbinary `db:"BinaryChunk"`
	Status                    sql.NullString         `db:"Status"`
	ChartLocationDE           crownsql.DictID        `db:"ChartLocationDE"`
	OwnerID                   crownsql.NullUniqueID  `db:"OwnerID"`
	DictationSystemJobNo      sql.NullString         `db:"DictationSystemJobNo"`
	Priority                  crownsql.NullUniqueID  `db:"Priority"`
	TransactionDTTM           crownsql.NullDatetime  `db:"TransactionDTTM"`
	ServiceDTTM               crownsql.NullDatetime  `db:"ServiceDTTM"`
	AuthoredDTTM              crownsql.NullDatetime  `db:"AuthoredDTTM"`
	InterfaceID               sql.NullInt64          `db:"InterfaceID"`
	SiteDE                    crownsql.NullDictID    `db:"SiteDE"`
	OriginatingDocumentID     crownsql.UniqueID      `db:"OriginatingDocumentID"`
	EditableChunkCompressed   crownsql.NullVarbinary `db:"EditableChunkCompressed"`
	UnEditableChunkCompressed crownsql.NullVarbinary `db:"UnEditableChunkCompressed"`
	IsArchived                sql.NullBool           `db:"IsArchived"`
	TextTemplateID            crownsql.NullUniqueID  `db:"TextTemplateID"`
	IsMUNonClinicalSetting    crownsql.Bool          `db:"IsMUNonClinicalSetting"`
}
