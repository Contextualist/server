// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameIndexSubject = "chii_index_related"

// IndexSubject mapped from table <chii_index_related>
type IndexSubject struct {
	ID          uint32  `gorm:"column:idx_rlt_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true"`
	Cat         int8    `gorm:"column:idx_rlt_cat;type:tinyint(3);not null"`
	IndexID     uint32  `gorm:"column:idx_rlt_rid;type:mediumint(8) unsigned;not null;comment:关联目录"`   // 关联目录
	SubjectType uint8   `gorm:"column:idx_rlt_type;type:smallint(6) unsigned;not null;comment:关联条目类型"` // 关联条目类型
	SubjectID   uint32  `gorm:"column:idx_rlt_sid;type:mediumint(8) unsigned;not null;comment:关联条目ID"` // 关联条目ID
	Order       uint32  `gorm:"column:idx_rlt_order;type:mediumint(8) unsigned;not null"`
	Comment     string  `gorm:"column:idx_rlt_comment;type:mediumtext;not null"`
	CreatedTime uint32  `gorm:"column:idx_rlt_dateline;type:int(10) unsigned;not null"`
	Subject     Subject `gorm:"foreignKey:idx_rlt_sid;references:subject_id" json:"subject"`
}

// TableName IndexSubject's table name
func (*IndexSubject) TableName() string {
	return TableNameIndexSubject
}
