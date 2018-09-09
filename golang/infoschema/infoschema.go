package infoschema

import (
	"database/sql"
	"time"
)

// CheckConstraint holds an entry for the Check_Constraints table
type CheckConstraint struct {
	ConstraintCatalog string `db: "CONSTRAINT_CATALOG"`
	ConstraintSchema  string `db: "CONSTRAINT_SCHEMA"`
	ConstraintName    string `db: "CONSTRAINT_NAME"`
	CheckClause       string `db: "CHECK_CLAUSE"`
}

// ColumnDomainUsage holds an entry for the Column_Domain_Usage table
type ColumnDomainUsage struct {
	DomainCatalog string `db: "DOMAIN_CATALOG"`
	DomainSchema  string `db: "DOMAIN_SCHEMA"`
	DomainName    string `db: "DOMAIN_NAME"`
	TableCatalog  string `db: "TABLE_CATALOG"`
	TableSchema   string `db: "TABLE_SCHEMA"`
	TableName     string `db: "TABLE_NAME"`
	ColumnName    string `db: "COLUMN_NAME"`
}

// ColumnPrivilege holds an entry for the Column_Privileges table
type ColumnPrivilege struct {
	Grantor       string `db: "GRANTOR"`
	Grantee       string `db: "GRANTEE"`
	TableCatalog  string `db: "TABLE_CATALOG"`
	TableSchema   string `db: "TABLE_SCHEMA"`
	TableName     string `db: "TABLE_NAME"`
	ColumnName    string `db: "COLUMN_NAME"`
	PrivilegeType string `db: "PRIVILEGE_TYPE"`
	IsGrantable   string `db: "IS_GRANTABLE"`
}

// Column holds an entry for the Columns table
type Column struct {
	TableCatalog           string         `db: "TABLE_CATALOG"`
	TableSchema            string         `db: "TABLE_SCHEMA"`
	TableName              string         `db: "TABLE_NAME"`
	ColumnName             string         `db: "COLUMN_NAME"`
	OrdinalPosition        int            `db: "ORDINAL_POSITION"`
	ColumnDefault          string         `db: "COLUMN_DEFAULT"`
	IsNullable             string         `db: "IS_NULLABLE"`
	DataType               string         `db: "DATA_TYPE"`
	CharacterMaximumLength sql.NullInt64  `db: "CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   sql.NullInt64  `db: "CHARACTER_OCTET_LENGTH"`
	NumericPrecision       sql.NullInt64  `db: "NUMERIC_PRECISION"`
	NumericPrecisionRadix  sql.NullInt64  `db: "NUMERIC_PRECISION_RADIX"`
	NumericScale           sql.NullInt64  `db: "NUMERIC_SCALE"`
	DateTimePrec           sql.NullInt64  `db: "DATE_TIME_PREC"`
	CharacterSetCatalog    sql.NullString `db: "CHARACTER_SET_CATALOG"`
	CharacterSetSchema     sql.NullString `db: "CHARACTER_SET_SCHEMA"`
	CharacterSetName       sql.NullString `db: "CHARACTER_SET_NAME"`
	CollationCatalog       sql.NullString `db: "COLLATION_CATALOG"`
	CollationSchema        sql.NullString `db: "COLLATION_SCHEMA"`
	CollationName          sql.NullString `db: "COLLATION_NAME"`
	DomainCatalog          sql.NullString `db: "DOMAIN_CATALOG"`
	DomainSchema           sql.NullString `db: "DOMAIN_SCHEMA"`
	DomainName             sql.NullString `db: "DOMAIN_NAME"`
}

// ConstraintColumnUsage holds an entry for the Constraint_Column_Usage table
type ConstraintColumnUsage struct {
	TableCatalog      string `db: "TABLE_CATALOG"`
	TableSchema       string `db: "TABLE_SCHEMA"`
	TableName         string `db: "TABLE_NAME"`
	ColumnName        string `db: "COLUMN_NAME"`
	ConstraintCatalog string `db: "CONSTRAINT_CATALOG"`
	ConstraintSchema  string `db: "CONSTRAINT_SCHEMA"`
	ConstraintName    string `db: "CONSTRAINT_NAME"`
}

// ConstraintTableUsage holds an entry for the Constraint_Table_Usage table
type ConstraintTableUsage struct {
	TableCatalog      string `db: "TABLE_CATALOG"`
	TableSchema       string `db: "TABLE_SCHEMA"`
	TableName         string `db: "TABLE_NAME"`
	ConstraintCatalog string `db: "CONSTRAINT_CATALOG"`
	ConstraintSchema  string `db: "CONSTRAINT_SCHEMA"`
	ConstraintName    string `db: "CONSTRAINT_NAME"`
}

// DomainConstraint holds an entry for the Domain_Constraints table
type DomainConstraint struct {
	ConstraintCatalog string `db: "CONSTRAINT_CATALOG"`
	ConstraintSchema  string `db: "CONSTRAINT_SCHEMA"`
	ConstraintName    string `db: "CONSTRAINT_NAME"`
	DomainCatalog     string `db: "DOMAIN_CATALOG"`
	DomainSchema      string `db: "DOMAIN_SCHEMA"`
	DomainName        string `db: "DOMAIN_NAME"`
	IsDeferrable      string `db: "IS_DEFERRABLE"`
	InitiallyDeferred string `db: "INITIALLY_DEFERRED"`
}

// Domain holds an entry for the Domains table
type Domain struct {
	DomainCatalog          string         `db: "DOMAIN_CATALOG"`
	DomainSchema           string         `db: "DOMAIN_SCHEMA"`
	DomainName             string         `db: "DOMAIN_NAME"`
	DataType               string         `db: "DATA_TYPE"`
	CharacterMaximumLength sql.NullInt64  `db: "CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   sql.NullInt64  `db: "CHARACTER_OCTET_LENGTH"`
	CollationCatalog       sql.NullString `db: "COLLATION_CATALOG"`
	CollationSchema        sql.NullString `db: "COLLATION_SCHEMA"`
	CollationName          sql.NullString `db: "COLLATION_NAME"`
	CharacterSetCatalog    sql.NullString `db: "CHARACTER_SET_CATALOG"`
	CharacterSetSchema     sql.NullString `db: "CHARACTER_SET_SCHEMA"`
	CharacterSetName       sql.NullString `db: "CHARACTER_SET_NAME"`
	NumericPrecision       sql.NullInt64  `db: "NUMERIC_PRECISION"`
	NumericPrecisionRadix  sql.NullInt64  `db: "NUMERIC_PRECISION_RADIX"`
	NumericScale           sql.NullInt64  `db: "NUMERIC_SCALE"`
	DatetimePrecision      sql.NullInt64  `db: "DATETIME_PRECISION"`
	DomainDefault          string         `db: "DOMAIN_DEFAULT"`
}

// KeyColumnUsage holds an entry for the Key_Column_Usage table
type KeyColumnUsage struct {
	ConstraintCatalog string `db: "CONSTRAINT_CATALOG"`
	ConstraintSchema  string `db: "CONSTRAINT_SCHEMA"`
	ConstraintName    string `db: "CONSTRAINT_NAME"`
	TableCatalog      string `db: "TABLE_CATALOG"`
	TableSchema       string `db: "TABLE_SCHEMA"`
	TableName         string `db: "TABLE_NAME"`
	ColumnName        string `db: "COLUMN_NAME"`
	OrdinalPosition   int    `db: "ORDINAL_POSITION"`
}

// Parameter holds an entry for the Parameters table
type Parameter struct {
	SpecificCatalog        string         `db: "SPECIFIC_CATALOG"`
	SpecificSchema         string         `db: "SPECIFIC_SCHEMA"`
	SpecificName           string         `db: "SPECIFIC_NAME"`
	OridinalPosition       int            `db: "ORIDINAL_POSITION"`
	ParameterMode          string         `db: "PARAMETER_MODE"`
	IsResult               string         `db: "IS_RESULT"`
	AsLocator              string         `db: "AS_LOCATOR"`
	ParameterName          sql.NullString `db: "PARAMETER_NAME"`
	DataType               string         `db: "DATA_TYPE"`
	CharacterMaximumLength sql.NullInt64  `db: "CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   sql.NullInt64  `db: "CHARACTER_OCTET_LENGTH"`
	CollationCatalog       sql.NullString `db: "COLLATION_CATALOG"`
	CollationSchema        sql.NullString `db: "COLLATION_SCHEMA"`
	CollationName          sql.NullString `db: "COLLATION_NAME"`
	CharacterSetCatalog    sql.NullString `db: "CHARACTER_SET_CATALOG"`
	CharacterSetSchema     sql.NullString `db: "CHARACTER_SET_SCHEMA"`
	CharacterSetName       sql.NullString `db: "CHARACTER_SET_NAME"`
	NumericPrecision       sql.NullInt64  `db: "NUMERIC_PRECISION"`
	NumericPrecisionRadix  sql.NullInt64  `db: "NUMERIC_PRECISION_RADIX"`
	NumericScale           sql.NullInt64  `db: "NUMERIC_SCALE"`
	DatetimePrecision      sql.NullInt64  `db: "DATETIME_PRECISION"`
	IntervalType           sql.NullString `db: "INTERVAL_TYPE"`
	IntervalPrecision      sql.NullInt64  `db: "INTERVAL_PRECISION"`
	UserDefinedTypeCatalog sql.NullString `db: "USER_DEFINED_TYPE_CATALOG"`
	UserDefinedTypeSchema  sql.NullString `db: "USER_DEFINED_TYPE_SCHEMA"`
	UserDefinedTypeName    sql.NullString `db: "USER_DEFINED_TYPE_NAME"`
	ScopeCatalog           sql.NullString `db: "SCOPE_CATALOG"`
	ScopeSchema            sql.NullString `db: "SCOPE_SCHEMA"`
	ScopeName              sql.NullString `db: "SCOPE_NAME"`
}

// ReferentialConstraint holds an entry for the Referential_Constraint table
type ReferentialConstraint struct {
	ConstraintCatalog       string         `db: "CONSTRAINT_CATALOG"`
	ConstraintSchema        string         `db: "CONSTRAINT_SCHEMA"`
	ConstraintName          string         `db: "CONSTRAINT_NAME"`
	UniqueConstraintCatalog string         `db: "UNIQUE_CONSTRAINT_CATALOG"`
	UniqueConstraintSchema  string         `db: "UNIQUE_CONSTRAINT_SCHEMA"`
	UniqueConstraintName    string         `db: "UNIQUE_CONSTRAINT_NAME"`
	MatchOption             sql.NullString `db: "MATCH_OPTION"`
	UpdateRule              string         `db: "UPDATE_RULE"`
	DeleteRule              string         `db: "DELETE_RULE"`
}

// RoutineColumn holds an entry for the Routine_Columns table
type RoutineColumn struct {
	TableCatalog           string         `db: "TABLE_CATALOG"`
	TableSchema            string         `db: "TABLE_SCHEMA"`
	TableName              string         `db: "TABLE_NAME"`
	ColumnName             string         `db: "COLUMN_NAME"`
	OrdinalPosition        int            `db: "ORDINAL_POSITION"`
	ColumnDefault          string         `db: "COLUMN_DEFAULT"`
	IsNullable             string         `db: "IS_NULLABLE"`
	DataType               string         `db: "DATA_TYPE"`
	CharacterMaximumLength sql.NullInt64  `db: "CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   sql.NullInt64  `db: "CHARACTER_OCTET_LENGTH"`
	NumericPrecision       sql.NullInt64  `db: "NUMERIC_PRECISION"`
	NumericPrecisionRadix  sql.NullInt64  `db: "NUMERIC_PRECISION_RADIX"`
	NumericScale           sql.NullInt64  `db: "NUMERIC_SCALE"`
	DatetimePrecision      sql.NullInt64  `db: "DATETIME_PRECISION"`
	CharacterSetCatalog    sql.NullString `db: "CHARACTER_SET_CATALOG"`
	CharacterSetSchema     sql.NullString `db: "CHARACTER_SET_SCHEMA"`
	CharacterSetName       sql.NullString `db: "CHARACTER_SET_NAME"`
	CollationCatalog       sql.NullString `db: "COLLATION_CATALOG"`
	CollationSchema        sql.NullString `db: "COLLATION_SCHEMA"`
	CollationName          sql.NullString `db: "COLLATION_NAME"`
	DomainCatalog          sql.NullString `db: "DOMAIN_CATALOG"`
	DomainSchema           sql.NullString `db: "DOMAIN_SCHEMA"`
	DomainName             sql.NullString `db: "DOMAIN_NAME"`
}

// Routine holds an entry for the Routines table
type Routine struct {
	SpecificCatalog        string         `db: "SPECIFIC_CATALOG"`
	SpecificSchema         string         `db: "SPECIFIC_SCHEMA"`
	SpecificName           string         `db: "SPECIFIC_NAME"`
	RoutineCatalog         string         `db: "ROUTINE_CATALOG"`
	RoutineSchema          string         `db: "ROUTINE_SCHEMA"`
	RoutineName            string         `db: "ROUTINE_NAME"`
	RoutineType            string         `db: "ROUTINE_TYPE"`
	ModuleCatalog          sql.NullString `db: "MODULE_CATALOG"`
	ModuleSchema           sql.NullString `db: "MODULE_SCHEMA"`
	ModuleName             sql.NullString `db: "MODULE_NAME"`
	UdtCatalog             sql.NullString `db: "UDT_CATALOG"`
	UdtSchema              sql.NullString `db: "UDT_SCHEMA"`
	UdtName                sql.NullString `db: "UDT_NAME"`
	DataType               string         `db: "DATA_TYPE"`
	CharacterMaximumLength int            `db: "CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   int            `db: "CHARACTER_OCTET_LENGTH"`
	CollationCatalog       sql.NullString `db: "COLLATION_CATALOG"`
	CollationSchema        sql.NullString `db: "COLLATION_SCHEMA"`
	CollationName          sql.NullString `db: "COLLATION_NAME"`
	CharacterSetCatalog    sql.NullString `db: "CHARACTER_SET_CATALOG"`
	CharacterSetSchema     sql.NullString `db: "CHARACTER_SET_SCHEMA"`
	CharacterSetName       sql.NullString `db: "CHARACTER_SET_NAME"`
	NumericPrecision       sql.NullInt64  `db: "NUMERIC_PRECISION"`
	NumericPrecisionRadix  sql.NullInt64  `db: "NUMERIC_PRECISION_RADIX"`
	NumericScale           sql.NullInt64  `db: "NUMERIC_SCALE"`
	DatetimePrecision      sql.NullInt64  `db: "DATETIME_PRECISION"`
	IntervalType           sql.NullString `db: "INTERVAL_TYPE"`
	IntervalPrecision      sql.NullInt64  `db: "INTERVAL_PRECISION"`
	TypeUdtCatalog         sql.NullString `db: "TYPE_UDT_CATALOG"`
	TypeUdtSchema          sql.NullString `db: "TYPE_UDT_SCHEMA"`
	TypeUdtName            sql.NullString `db: "TYPE_UDT_NAME"`
	ScopeCatalog           sql.NullString `db: "SCOPE_CATALOG"`
	ScopeSchema            sql.NullString `db: "SCOPE_SCHEMA"`
	ScopeName              sql.NullString `db: "SCOPE_NAME"`
	MaximumCardinality     sql.NullInt64  `db: "MAXIMUM_CARDINALITY"`
	DtdIdentifier          sql.NullString `db: "DTD_IDENTIFIER"`
	RoutineBody            string         `db: "ROUTINE_BODY"`
	RoutineDefinition      sql.NullString `db: "ROUTINE_DEFINITION"`
	ExternalName           sql.NullString `db: "EXTERNAL_NAME"`
	ExternalLanguage       sql.NullString `db: "EXTERNAL_LANGUAGE"`
	ParameterStyle         sql.NullString `db: "PARAMETER_STYLE"`
	IsDeterministic        string         `db: "IS_DETERMINISTIC"`
	SQLDataAccess          string         `db: "SQL_DATA_ACCESS"`
	IsNullCall             string         `db: "IS_NULL_CALL"`
	SQLPath                sql.NullString `db: "SQL_PATH"`
	SchemaLevelRoutine     string         `db: "SCHEMA_LEVEL_ROUTINE"`
	MaxDynamicResultsSet   int            `db: "MAX_DYNAMIC_RESULTS_SET"`
	IsUserDefinedCast      int            `db: "IS_USER_DEFINED_CAST"`
	IsImplicitlyInvocable  string         `db: "IS_IMPLICITLY_INVOCABLE"`
	Created                time.Time      `db: "CREATED"`
	Altered                time.Time      `db: "ALTERED"`
}

// Schema holds an entry for the Schemata table
type Schema struct {
	CatalogName                string         `db: "CATALOG_NAME"`
	SchemaName                 string         `db: "SCHEMA_NAME"`
	SchemaOwner                string         `db: "SCHEMA_OWNER"`
	DefaultCharacterSetCatalog sql.NullString `db: "DEFAULT_CHARACTER_SET_CATALOG"`
	DefaultCharacterSetSchema  sql.NullString `db: "DEFAULT_CHARACTER_SET_SCHEMA"`
	DefaultCharacterSetName    string         `db: "DEFAULT_CHARACTER_SET_NAME"`
}

// TableConstraint holds an entry for the Table_Constraints table
type TableConstraint struct {
	ConstraintCatalog string `db: "CONSTRAINT_CATALOG"`
	ConstraintSchema  string `db: "CONSTRAINT_SCHEMA"`
	ConstraintName    string `db: "CONSTRAINT_NAME"`
	TableCatalog      string `db: "TABLE_CATALOG"`
	TableSchema       string `db: "TABLE_SCHEMA"`
	TableName         string `db: "TABLE_NAME"`
	ConstraintType    string `db: "CONSTRAINT_TYPE"`
	IsDeferrable      string `db: "IS_DEFERRABLE"`
	InitiallyDeferred string `db: "INITIALLY_DEFERRED"`
}

// TablePrivilege holds an entry for the Table_Privilege table
type TablePrivilege struct {
	Grantor       string `db: "GRANTOR"`
	Grantee       string `db: "GRANTEE"`
	TableCatalog  string `db: "TABLE_CATALOG"`
	TableSchema   string `db: "TABLE_SCHEMA"`
	TableName     string `db: "TABLE_NAME"`
	PrivilegeType string `db: "PRIVILEGE_TYPE"`
	IsGrantable   string `db: "IS_GRANTABLE"`
}

// Table holds an entry for the Tables table
type Table struct {
	TableCatalog string `db: "TABLE_CATALOG"`
	TableSchema  string `db: "TABLE_SCHEMA"`
	TableName    string `db: "TABLE_NAME"`
	TableType    string `db: "TABLE_TYPE"`
}

// ViewColumnUsage holds an entry for the View_Column_Usage table
type ViewColumnUsage struct {
	ViewCatalog  string `db: "VIEW_CATALOG"`
	ViewSchema   string `db: "VIEW_SCHEMA"`
	ViewName     string `db: "VIEW_NAME"`
	TableCatalog string `db: "TABLE_CATALOG"`
	TableSchema  string `db: "TABLE_SCHEMA"`
	TableName    string `db: "TABLE_NAME"`
	ColumnName   string `db: "COLUMN_NAME"`
}

// ViewTableUsage holds an entry for the View_Table_Usage table
type ViewTableUsage struct {
	ViewCatalog  string `db: "VIEW_CATALOG"`
	ViewSchema   string `db: "VIEW_SCHEMA"`
	ViewName     string `db: "VIEW_NAME"`
	TableCatalog string `db: "TABLE_CATALOG"`
	TableSchema  string `db: "TABLE_SCHEMA"`
	TableName    string `db: "TABLE_NAME"`
}

// View holds an entry for the Views table
type View struct {
	TableCatalog   string         `db: "TABLE_CATALOG"`
	TableSchema    string         `db: "TABLE_SCHEMA"`
	TableName      string         `db: "TABLE_NAME"`
	ViewDefinition sql.NullString `db: "VIEW_DEFINITION"`
	CheckOption    string         `db: "CHECK_OPTION"`
	IsUpdatable    string         `db: "IS_UPDATABLE"`
}
