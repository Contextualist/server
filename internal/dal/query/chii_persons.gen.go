// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/bangumi/server/internal/dal/dao"
)

func newPerson(db *gorm.DB) person {
	_person := person{}

	_person.personDo.UseDB(db)
	_person.personDo.UseModel(&dao.Person{})

	tableName := _person.personDo.TableName()
	_person.ALL = field.NewField(tableName, "*")
	_person.ID = field.NewUint32(tableName, "prsn_id")
	_person.Name = field.NewString(tableName, "prsn_name")
	_person.Type = field.NewUint8(tableName, "prsn_type")
	_person.Infobox = field.NewString(tableName, "prsn_infobox")
	_person.Producer = field.NewBool(tableName, "prsn_producer")
	_person.Mangaka = field.NewBool(tableName, "prsn_mangaka")
	_person.Artist = field.NewBool(tableName, "prsn_artist")
	_person.Seiyu = field.NewBool(tableName, "prsn_seiyu")
	_person.Writer = field.NewBool(tableName, "prsn_writer")
	_person.Illustrator = field.NewBool(tableName, "prsn_illustrator")
	_person.Actor = field.NewBool(tableName, "prsn_actor")
	_person.Summary = field.NewString(tableName, "prsn_summary")
	_person.Img = field.NewString(tableName, "prsn_img")
	_person.ImgAnidb = field.NewString(tableName, "prsn_img_anidb")
	_person.Comment = field.NewUint32(tableName, "prsn_comment")
	_person.Collects = field.NewUint32(tableName, "prsn_collects")
	_person.Dateline = field.NewUint32(tableName, "prsn_dateline")
	_person.Lastpost = field.NewUint32(tableName, "prsn_lastpost")
	_person.Lock = field.NewInt8(tableName, "prsn_lock")
	_person.AnidbID = field.NewUint32(tableName, "prsn_anidb_id")
	_person.Ban = field.NewUint8(tableName, "prsn_ban")
	_person.Redirect = field.NewUint32(tableName, "prsn_redirect")
	_person.Nsfw = field.NewBool(tableName, "prsn_nsfw")
	_person.Fields = personHasOneFields{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Fields", "dao.PersonField"),
	}

	_person.fillFieldMap()

	return _person
}

type person struct {
	personDo personDo

	ALL         field.Field
	ID          field.Uint32
	Name        field.String
	Type        field.Uint8
	Infobox     field.String
	Producer    field.Bool
	Mangaka     field.Bool
	Artist      field.Bool
	Seiyu       field.Bool
	Writer      field.Bool
	Illustrator field.Bool
	Actor       field.Bool
	Summary     field.String
	Img         field.String
	ImgAnidb    field.String
	Comment     field.Uint32
	Collects    field.Uint32
	Dateline    field.Uint32
	Lastpost    field.Uint32
	Lock        field.Int8
	AnidbID     field.Uint32
	Ban         field.Uint8
	Redirect    field.Uint32
	Nsfw        field.Bool
	Fields      personHasOneFields

	fieldMap map[string]field.Expr
}

func (p person) Table(newTableName string) *person {
	p.personDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p person) As(alias string) *person {
	p.personDo.DO = *(p.personDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *person) updateTableName(table string) *person {
	p.ALL = field.NewField(table, "*")
	p.ID = field.NewUint32(table, "prsn_id")
	p.Name = field.NewString(table, "prsn_name")
	p.Type = field.NewUint8(table, "prsn_type")
	p.Infobox = field.NewString(table, "prsn_infobox")
	p.Producer = field.NewBool(table, "prsn_producer")
	p.Mangaka = field.NewBool(table, "prsn_mangaka")
	p.Artist = field.NewBool(table, "prsn_artist")
	p.Seiyu = field.NewBool(table, "prsn_seiyu")
	p.Writer = field.NewBool(table, "prsn_writer")
	p.Illustrator = field.NewBool(table, "prsn_illustrator")
	p.Actor = field.NewBool(table, "prsn_actor")
	p.Summary = field.NewString(table, "prsn_summary")
	p.Img = field.NewString(table, "prsn_img")
	p.ImgAnidb = field.NewString(table, "prsn_img_anidb")
	p.Comment = field.NewUint32(table, "prsn_comment")
	p.Collects = field.NewUint32(table, "prsn_collects")
	p.Dateline = field.NewUint32(table, "prsn_dateline")
	p.Lastpost = field.NewUint32(table, "prsn_lastpost")
	p.Lock = field.NewInt8(table, "prsn_lock")
	p.AnidbID = field.NewUint32(table, "prsn_anidb_id")
	p.Ban = field.NewUint8(table, "prsn_ban")
	p.Redirect = field.NewUint32(table, "prsn_redirect")
	p.Nsfw = field.NewBool(table, "prsn_nsfw")

	p.fillFieldMap()

	return p
}

func (p *person) WithContext(ctx context.Context) *personDo { return p.personDo.WithContext(ctx) }

func (p person) TableName() string { return p.personDo.TableName() }

func (p person) Alias() string { return p.personDo.Alias() }

func (p *person) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *person) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 24)
	p.fieldMap["prsn_id"] = p.ID
	p.fieldMap["prsn_name"] = p.Name
	p.fieldMap["prsn_type"] = p.Type
	p.fieldMap["prsn_infobox"] = p.Infobox
	p.fieldMap["prsn_producer"] = p.Producer
	p.fieldMap["prsn_mangaka"] = p.Mangaka
	p.fieldMap["prsn_artist"] = p.Artist
	p.fieldMap["prsn_seiyu"] = p.Seiyu
	p.fieldMap["prsn_writer"] = p.Writer
	p.fieldMap["prsn_illustrator"] = p.Illustrator
	p.fieldMap["prsn_actor"] = p.Actor
	p.fieldMap["prsn_summary"] = p.Summary
	p.fieldMap["prsn_img"] = p.Img
	p.fieldMap["prsn_img_anidb"] = p.ImgAnidb
	p.fieldMap["prsn_comment"] = p.Comment
	p.fieldMap["prsn_collects"] = p.Collects
	p.fieldMap["prsn_dateline"] = p.Dateline
	p.fieldMap["prsn_lastpost"] = p.Lastpost
	p.fieldMap["prsn_lock"] = p.Lock
	p.fieldMap["prsn_anidb_id"] = p.AnidbID
	p.fieldMap["prsn_ban"] = p.Ban
	p.fieldMap["prsn_redirect"] = p.Redirect
	p.fieldMap["prsn_nsfw"] = p.Nsfw

}

func (p person) clone(db *gorm.DB) person {
	p.personDo.ReplaceDB(db)
	return p
}

type personHasOneFields struct {
	db *gorm.DB

	field.RelationField
}

func (a personHasOneFields) Where(conds ...field.Expr) *personHasOneFields {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a personHasOneFields) WithContext(ctx context.Context) *personHasOneFields {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a personHasOneFields) Model(m *dao.Person) *personHasOneFieldsTx {
	return &personHasOneFieldsTx{a.db.Model(m).Association(a.Name())}
}

type personHasOneFieldsTx struct{ tx *gorm.Association }

func (a personHasOneFieldsTx) Find() (result *dao.PersonField, err error) {
	return result, a.tx.Find(&result)
}

func (a personHasOneFieldsTx) Append(values ...*dao.PersonField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a personHasOneFieldsTx) Replace(values ...*dao.PersonField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a personHasOneFieldsTx) Delete(values ...*dao.PersonField) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a personHasOneFieldsTx) Clear() error {
	return a.tx.Clear()
}

func (a personHasOneFieldsTx) Count() int64 {
	return a.tx.Count()
}

type personDo struct{ gen.DO }

//Get a person from database.
//
//where(prsn_id=@id)
func (p personDo) Get(id uint32) (result *dao.Person, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("prsn_id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = p.UnderlyingDB().Where(generateSQL.String(), params).Take(&result)
	} else {
		executeSQL = p.UnderlyingDB().Where(generateSQL.String()).Take(&result)
	}
	err = executeSQL.Error
	return
}

func (p personDo) Debug() *personDo {
	return p.withDO(p.DO.Debug())
}

func (p personDo) WithContext(ctx context.Context) *personDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personDo) Clauses(conds ...clause.Expression) *personDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personDo) Returning(value interface{}, columns ...string) *personDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personDo) Not(conds ...gen.Condition) *personDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personDo) Or(conds ...gen.Condition) *personDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personDo) Select(conds ...field.Expr) *personDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personDo) Where(conds ...gen.Condition) *personDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *personDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p personDo) Order(conds ...field.Expr) *personDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personDo) Distinct(cols ...field.Expr) *personDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personDo) Omit(cols ...field.Expr) *personDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personDo) Join(table schema.Tabler, on ...field.Expr) *personDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personDo) LeftJoin(table schema.Tabler, on ...field.Expr) *personDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personDo) RightJoin(table schema.Tabler, on ...field.Expr) *personDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personDo) Group(cols ...field.Expr) *personDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personDo) Having(conds ...gen.Condition) *personDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personDo) Limit(limit int) *personDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personDo) Offset(offset int) *personDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *personDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personDo) Unscoped() *personDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personDo) Create(values ...*dao.Person) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personDo) CreateInBatches(values []*dao.Person, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personDo) Save(values ...*dao.Person) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personDo) First() (*dao.Person, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Person), nil
	}
}

func (p personDo) Take() (*dao.Person, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Person), nil
	}
}

func (p personDo) Last() (*dao.Person, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Person), nil
	}
}

func (p personDo) Find() ([]*dao.Person, error) {
	result, err := p.DO.Find()
	return result.([]*dao.Person), err
}

func (p personDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.Person, err error) {
	buf := make([]*dao.Person, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personDo) FindInBatches(result *[]*dao.Person, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personDo) Attrs(attrs ...field.AssignExpr) *personDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personDo) Assign(attrs ...field.AssignExpr) *personDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personDo) Joins(fields ...field.RelationField) *personDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personDo) Preload(fields ...field.RelationField) *personDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personDo) FirstOrInit() (*dao.Person, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Person), nil
	}
}

func (p personDo) FirstOrCreate() (*dao.Person, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Person), nil
	}
}

func (p personDo) FindByPage(offset int, limit int) (result []*dao.Person, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p personDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p *personDo) withDO(do gen.Dao) *personDo {
	p.DO = *do.(*gen.DO)
	return p
}
