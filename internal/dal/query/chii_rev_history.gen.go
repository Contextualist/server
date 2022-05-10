// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/bangumi/server/internal/dal/dao"
)

func newRevisionHistory(db *gorm.DB) revisionHistory {
	_revisionHistory := revisionHistory{}

	_revisionHistory.revisionHistoryDo.UseDB(db)
	_revisionHistory.revisionHistoryDo.UseModel(&dao.RevisionHistory{})

	tableName := _revisionHistory.revisionHistoryDo.TableName()
	_revisionHistory.ALL = field.NewField(tableName, "*")
	_revisionHistory.ID = field.NewUint32(tableName, "rev_id")
	_revisionHistory.Type = field.NewUint8(tableName, "rev_type")
	_revisionHistory.Mid = field.NewUint32(tableName, "rev_mid")
	_revisionHistory.TextID = field.NewUint32(tableName, "rev_text_id")
	_revisionHistory.CreatedAt = field.NewUint32(tableName, "rev_dateline")
	_revisionHistory.CreatorID = field.NewUint32(tableName, "rev_creator")
	_revisionHistory.Summary = field.NewString(tableName, "rev_edit_summary")

	_revisionHistory.fillFieldMap()

	return _revisionHistory
}

type revisionHistory struct {
	revisionHistoryDo revisionHistoryDo

	ALL       field.Field
	ID        field.Uint32
	Type      field.Uint8
	Mid       field.Uint32
	TextID    field.Uint32
	CreatedAt field.Uint32
	CreatorID field.Uint32
	Summary   field.String

	fieldMap map[string]field.Expr
}

func (r revisionHistory) Table(newTableName string) *revisionHistory {
	r.revisionHistoryDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r revisionHistory) As(alias string) *revisionHistory {
	r.revisionHistoryDo.DO = *(r.revisionHistoryDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *revisionHistory) updateTableName(table string) *revisionHistory {
	r.ALL = field.NewField(table, "*")
	r.ID = field.NewUint32(table, "rev_id")
	r.Type = field.NewUint8(table, "rev_type")
	r.Mid = field.NewUint32(table, "rev_mid")
	r.TextID = field.NewUint32(table, "rev_text_id")
	r.CreatedAt = field.NewUint32(table, "rev_dateline")
	r.CreatorID = field.NewUint32(table, "rev_creator")
	r.Summary = field.NewString(table, "rev_edit_summary")

	r.fillFieldMap()

	return r
}

func (r *revisionHistory) WithContext(ctx context.Context) *revisionHistoryDo {
	return r.revisionHistoryDo.WithContext(ctx)
}

func (r revisionHistory) TableName() string { return r.revisionHistoryDo.TableName() }

func (r revisionHistory) Alias() string { return r.revisionHistoryDo.Alias() }

func (r *revisionHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *revisionHistory) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 7)
	r.fieldMap["rev_id"] = r.ID
	r.fieldMap["rev_type"] = r.Type
	r.fieldMap["rev_mid"] = r.Mid
	r.fieldMap["rev_text_id"] = r.TextID
	r.fieldMap["rev_dateline"] = r.CreatedAt
	r.fieldMap["rev_creator"] = r.CreatorID
	r.fieldMap["rev_edit_summary"] = r.Summary
}

func (r revisionHistory) clone(db *gorm.DB) revisionHistory {
	r.revisionHistoryDo.ReplaceDB(db)
	return r
}

type revisionHistoryDo struct{ gen.DO }

func (r revisionHistoryDo) Debug() *revisionHistoryDo {
	return r.withDO(r.DO.Debug())
}

func (r revisionHistoryDo) WithContext(ctx context.Context) *revisionHistoryDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r revisionHistoryDo) Clauses(conds ...clause.Expression) *revisionHistoryDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r revisionHistoryDo) Returning(value interface{}, columns ...string) *revisionHistoryDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r revisionHistoryDo) Not(conds ...gen.Condition) *revisionHistoryDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r revisionHistoryDo) Or(conds ...gen.Condition) *revisionHistoryDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r revisionHistoryDo) Select(conds ...field.Expr) *revisionHistoryDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r revisionHistoryDo) Where(conds ...gen.Condition) *revisionHistoryDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r revisionHistoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *revisionHistoryDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r revisionHistoryDo) Order(conds ...field.Expr) *revisionHistoryDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r revisionHistoryDo) Distinct(cols ...field.Expr) *revisionHistoryDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r revisionHistoryDo) Omit(cols ...field.Expr) *revisionHistoryDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r revisionHistoryDo) Join(table schema.Tabler, on ...field.Expr) *revisionHistoryDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r revisionHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *revisionHistoryDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r revisionHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *revisionHistoryDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r revisionHistoryDo) Group(cols ...field.Expr) *revisionHistoryDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r revisionHistoryDo) Having(conds ...gen.Condition) *revisionHistoryDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r revisionHistoryDo) Limit(limit int) *revisionHistoryDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r revisionHistoryDo) Offset(offset int) *revisionHistoryDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r revisionHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *revisionHistoryDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r revisionHistoryDo) Unscoped() *revisionHistoryDo {
	return r.withDO(r.DO.Unscoped())
}

func (r revisionHistoryDo) Create(values ...*dao.RevisionHistory) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r revisionHistoryDo) CreateInBatches(values []*dao.RevisionHistory, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r revisionHistoryDo) Save(values ...*dao.RevisionHistory) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r revisionHistoryDo) First() (*dao.RevisionHistory, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionHistory), nil
	}
}

func (r revisionHistoryDo) Take() (*dao.RevisionHistory, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionHistory), nil
	}
}

func (r revisionHistoryDo) Last() (*dao.RevisionHistory, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionHistory), nil
	}
}

func (r revisionHistoryDo) Find() ([]*dao.RevisionHistory, error) {
	result, err := r.DO.Find()
	return result.([]*dao.RevisionHistory), err
}

func (r revisionHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.RevisionHistory, err error) {
	buf := make([]*dao.RevisionHistory, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r revisionHistoryDo) FindInBatches(result *[]*dao.RevisionHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r revisionHistoryDo) Attrs(attrs ...field.AssignExpr) *revisionHistoryDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r revisionHistoryDo) Assign(attrs ...field.AssignExpr) *revisionHistoryDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r revisionHistoryDo) Joins(fields ...field.RelationField) *revisionHistoryDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r revisionHistoryDo) Preload(fields ...field.RelationField) *revisionHistoryDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r revisionHistoryDo) FirstOrInit() (*dao.RevisionHistory, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionHistory), nil
	}
}

func (r revisionHistoryDo) FirstOrCreate() (*dao.RevisionHistory, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.RevisionHistory), nil
	}
}

func (r revisionHistoryDo) FindByPage(offset int, limit int) (result []*dao.RevisionHistory, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r revisionHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r *revisionHistoryDo) withDO(do gen.Dao) *revisionHistoryDo {
	r.DO = *do.(*gen.DO)
	return r
}
