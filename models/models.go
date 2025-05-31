package models

// Model corresponds to rebac.models
type Model struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"uniqueIndex;size:255"`
	Description string
	Types       []Type `gorm:"foreignKey:ModelID"`
}

// Type corresponds to rebac.types
type Type struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255"`
	ModelID   uint
	Model     Model      `gorm:"constraint:OnDelete:CASCADE;"`
	Relations []Relation `gorm:"foreignKey:TypeID"`
}

// Relation corresponds to rebac.relations
type Relation struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"size:255"`
	TypeID uint
	Type   Type `gorm:"constraint:OnDelete:CASCADE;"`
}

// DirectRelation corresponds to rebac.direct_relations
type DirectRelation struct {
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	RelationID         uint
	Relation           Relation `gorm:"foreignKey:RelationID;constraint:OnDelete:CASCADE;"`
	UserTypeID         uint
	UserType           Type `gorm:"foreignKey:UserTypeID;constraint:OnDelete:CASCADE;"`
	RequiredRelationID uint
	RequiredRelation   Relation `gorm:"foreignKey:RequiredRelationID"`
	ModelID            uint
	Model              Model `gorm:"foreignKey:ModelID"`
}

// ImpliedRelation corresponds to rebac.implied_relations
type ImpliedRelation struct {
	ID                  uint `gorm:"primaryKey;autoIncrement"`
	RelationID          uint
	Relation            Relation `gorm:"foreignKey:RelationID;constraint:OnDelete:CASCADE;"`
	PrivilegeRelationID uint
	PrivilegeRelation   Relation `gorm:"foreignKey:PrivilegeRelationID;constraint:OnDelete:CASCADE;"`
	LinkingRelationID   uint
	LinkingRelation     Relation `gorm:"foreignKey:LinkingRelationID;constraint:OnDelete:CASCADE;"`
	ModelID             uint
	Model               Model `gorm:"foreignKey:ModelID"`
}

// Policy corresponds to rebac.policies
type Policy struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	UserTypeID       uint
	UserType         Type `gorm:"foreignKey:UserTypeID"`
	UserIdentifier   string
	RelationID       uint
	Relation         Relation `gorm:"foreignKey:RelationID"`
	ObjectTypeID     uint
	ObjectType       Type `gorm:"foreignKey:ObjectTypeID"`
	ObjectIdentifier string
	ModelID          uint
	Model            Model `gorm:"foreignKey:ModelID"`
}
