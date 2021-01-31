// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/cybertect/ent/httpdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
)

// HTTPDetectorResultQuery is the builder for querying HTTPDetectorResult entities.
type HTTPDetectorResultQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.HTTPDetectorResult
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HTTPDetectorResultQuery builder.
func (hdrq *HTTPDetectorResultQuery) Where(ps ...predicate.HTTPDetectorResult) *HTTPDetectorResultQuery {
	hdrq.predicates = append(hdrq.predicates, ps...)
	return hdrq
}

// Limit adds a limit step to the query.
func (hdrq *HTTPDetectorResultQuery) Limit(limit int) *HTTPDetectorResultQuery {
	hdrq.limit = &limit
	return hdrq
}

// Offset adds an offset step to the query.
func (hdrq *HTTPDetectorResultQuery) Offset(offset int) *HTTPDetectorResultQuery {
	hdrq.offset = &offset
	return hdrq
}

// Order adds an order step to the query.
func (hdrq *HTTPDetectorResultQuery) Order(o ...OrderFunc) *HTTPDetectorResultQuery {
	hdrq.order = append(hdrq.order, o...)
	return hdrq
}

// First returns the first HTTPDetectorResult entity from the query.
// Returns a *NotFoundError when no HTTPDetectorResult was found.
func (hdrq *HTTPDetectorResultQuery) First(ctx context.Context) (*HTTPDetectorResult, error) {
	nodes, err := hdrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{httpdetectorresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hdrq *HTTPDetectorResultQuery) FirstX(ctx context.Context) *HTTPDetectorResult {
	node, err := hdrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HTTPDetectorResult ID from the query.
// Returns a *NotFoundError when no HTTPDetectorResult ID was found.
func (hdrq *HTTPDetectorResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hdrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{httpdetectorresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hdrq *HTTPDetectorResultQuery) FirstIDX(ctx context.Context) int {
	id, err := hdrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HTTPDetectorResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one HTTPDetectorResult entity is not found.
// Returns a *NotFoundError when no HTTPDetectorResult entities are found.
func (hdrq *HTTPDetectorResultQuery) Only(ctx context.Context) (*HTTPDetectorResult, error) {
	nodes, err := hdrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{httpdetectorresult.Label}
	default:
		return nil, &NotSingularError{httpdetectorresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hdrq *HTTPDetectorResultQuery) OnlyX(ctx context.Context) *HTTPDetectorResult {
	node, err := hdrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HTTPDetectorResult ID in the query.
// Returns a *NotSingularError when exactly one HTTPDetectorResult ID is not found.
// Returns a *NotFoundError when no entities are found.
func (hdrq *HTTPDetectorResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = hdrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = &NotSingularError{httpdetectorresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hdrq *HTTPDetectorResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := hdrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HTTPDetectorResults.
func (hdrq *HTTPDetectorResultQuery) All(ctx context.Context) ([]*HTTPDetectorResult, error) {
	if err := hdrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hdrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hdrq *HTTPDetectorResultQuery) AllX(ctx context.Context) []*HTTPDetectorResult {
	nodes, err := hdrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HTTPDetectorResult IDs.
func (hdrq *HTTPDetectorResultQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := hdrq.Select(httpdetectorresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hdrq *HTTPDetectorResultQuery) IDsX(ctx context.Context) []int {
	ids, err := hdrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hdrq *HTTPDetectorResultQuery) Count(ctx context.Context) (int, error) {
	if err := hdrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hdrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hdrq *HTTPDetectorResultQuery) CountX(ctx context.Context) int {
	count, err := hdrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hdrq *HTTPDetectorResultQuery) Exist(ctx context.Context) (bool, error) {
	if err := hdrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hdrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hdrq *HTTPDetectorResultQuery) ExistX(ctx context.Context) bool {
	exist, err := hdrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HTTPDetectorResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hdrq *HTTPDetectorResultQuery) Clone() *HTTPDetectorResultQuery {
	if hdrq == nil {
		return nil
	}
	return &HTTPDetectorResultQuery{
		config:     hdrq.config,
		limit:      hdrq.limit,
		offset:     hdrq.offset,
		order:      append([]OrderFunc{}, hdrq.order...),
		predicates: append([]predicate.HTTPDetectorResult{}, hdrq.predicates...),
		// clone intermediate query.
		sql:  hdrq.sql.Clone(),
		path: hdrq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.HTTPDetectorResult.Query().
//		GroupBy(httpdetectorresult.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (hdrq *HTTPDetectorResultQuery) GroupBy(field string, fields ...string) *HTTPDetectorResultGroupBy {
	group := &HTTPDetectorResultGroupBy{config: hdrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hdrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hdrq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
//	}
//
//	client.HTTPDetectorResult.Query().
//		Select(httpdetectorresult.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (hdrq *HTTPDetectorResultQuery) Select(field string, fields ...string) *HTTPDetectorResultSelect {
	hdrq.fields = append([]string{field}, fields...)
	return &HTTPDetectorResultSelect{HTTPDetectorResultQuery: hdrq}
}

func (hdrq *HTTPDetectorResultQuery) prepareQuery(ctx context.Context) error {
	for _, f := range hdrq.fields {
		if !httpdetectorresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hdrq.path != nil {
		prev, err := hdrq.path(ctx)
		if err != nil {
			return err
		}
		hdrq.sql = prev
	}
	return nil
}

func (hdrq *HTTPDetectorResultQuery) sqlAll(ctx context.Context) ([]*HTTPDetectorResult, error) {
	var (
		nodes = []*HTTPDetectorResult{}
		_spec = hdrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &HTTPDetectorResult{config: hdrq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, hdrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (hdrq *HTTPDetectorResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hdrq.querySpec()
	return sqlgraph.CountNodes(ctx, hdrq.driver, _spec)
}

func (hdrq *HTTPDetectorResultQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hdrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (hdrq *HTTPDetectorResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   httpdetectorresult.Table,
			Columns: httpdetectorresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: httpdetectorresult.FieldID,
			},
		},
		From:   hdrq.sql,
		Unique: true,
	}
	if fields := hdrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, httpdetectorresult.FieldID)
		for i := range fields {
			if fields[i] != httpdetectorresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hdrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hdrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hdrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hdrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, httpdetectorresult.ValidColumn)
			}
		}
	}
	return _spec
}

func (hdrq *HTTPDetectorResultQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(hdrq.driver.Dialect())
	t1 := builder.Table(httpdetectorresult.Table)
	selector := builder.Select(t1.Columns(httpdetectorresult.Columns...)...).From(t1)
	if hdrq.sql != nil {
		selector = hdrq.sql
		selector.Select(selector.Columns(httpdetectorresult.Columns...)...)
	}
	for _, p := range hdrq.predicates {
		p(selector)
	}
	for _, p := range hdrq.order {
		p(selector, httpdetectorresult.ValidColumn)
	}
	if offset := hdrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hdrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HTTPDetectorResultGroupBy is the group-by builder for HTTPDetectorResult entities.
type HTTPDetectorResultGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hdrgb *HTTPDetectorResultGroupBy) Aggregate(fns ...AggregateFunc) *HTTPDetectorResultGroupBy {
	hdrgb.fns = append(hdrgb.fns, fns...)
	return hdrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (hdrgb *HTTPDetectorResultGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := hdrgb.path(ctx)
	if err != nil {
		return err
	}
	hdrgb.sql = query
	return hdrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := hdrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (hdrgb *HTTPDetectorResultGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(hdrgb.fields) > 1 {
		return nil, errors.New("ent: HTTPDetectorResultGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := hdrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) StringsX(ctx context.Context) []string {
	v, err := hdrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hdrgb *HTTPDetectorResultGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hdrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: HTTPDetectorResultGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) StringX(ctx context.Context) string {
	v, err := hdrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (hdrgb *HTTPDetectorResultGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(hdrgb.fields) > 1 {
		return nil, errors.New("ent: HTTPDetectorResultGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := hdrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) IntsX(ctx context.Context) []int {
	v, err := hdrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hdrgb *HTTPDetectorResultGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hdrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: HTTPDetectorResultGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) IntX(ctx context.Context) int {
	v, err := hdrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (hdrgb *HTTPDetectorResultGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(hdrgb.fields) > 1 {
		return nil, errors.New("ent: HTTPDetectorResultGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := hdrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := hdrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hdrgb *HTTPDetectorResultGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hdrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: HTTPDetectorResultGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) Float64X(ctx context.Context) float64 {
	v, err := hdrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (hdrgb *HTTPDetectorResultGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(hdrgb.fields) > 1 {
		return nil, errors.New("ent: HTTPDetectorResultGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := hdrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := hdrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (hdrgb *HTTPDetectorResultGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hdrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: HTTPDetectorResultGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hdrgb *HTTPDetectorResultGroupBy) BoolX(ctx context.Context) bool {
	v, err := hdrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hdrgb *HTTPDetectorResultGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range hdrgb.fields {
		if !httpdetectorresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := hdrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hdrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hdrgb *HTTPDetectorResultGroupBy) sqlQuery() *sql.Selector {
	selector := hdrgb.sql
	columns := make([]string, 0, len(hdrgb.fields)+len(hdrgb.fns))
	columns = append(columns, hdrgb.fields...)
	for _, fn := range hdrgb.fns {
		columns = append(columns, fn(selector, httpdetectorresult.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(hdrgb.fields...)
}

// HTTPDetectorResultSelect is the builder for selecting fields of HTTPDetectorResult entities.
type HTTPDetectorResultSelect struct {
	*HTTPDetectorResultQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (hdrs *HTTPDetectorResultSelect) Scan(ctx context.Context, v interface{}) error {
	if err := hdrs.prepareQuery(ctx); err != nil {
		return err
	}
	hdrs.sql = hdrs.HTTPDetectorResultQuery.sqlQuery()
	return hdrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) ScanX(ctx context.Context, v interface{}) {
	if err := hdrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (hdrs *HTTPDetectorResultSelect) Strings(ctx context.Context) ([]string, error) {
	if len(hdrs.fields) > 1 {
		return nil, errors.New("ent: HTTPDetectorResultSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := hdrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) StringsX(ctx context.Context) []string {
	v, err := hdrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (hdrs *HTTPDetectorResultSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = hdrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: HTTPDetectorResultSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) StringX(ctx context.Context) string {
	v, err := hdrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (hdrs *HTTPDetectorResultSelect) Ints(ctx context.Context) ([]int, error) {
	if len(hdrs.fields) > 1 {
		return nil, errors.New("ent: HTTPDetectorResultSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := hdrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) IntsX(ctx context.Context) []int {
	v, err := hdrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (hdrs *HTTPDetectorResultSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = hdrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: HTTPDetectorResultSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) IntX(ctx context.Context) int {
	v, err := hdrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (hdrs *HTTPDetectorResultSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(hdrs.fields) > 1 {
		return nil, errors.New("ent: HTTPDetectorResultSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := hdrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) Float64sX(ctx context.Context) []float64 {
	v, err := hdrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (hdrs *HTTPDetectorResultSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = hdrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: HTTPDetectorResultSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) Float64X(ctx context.Context) float64 {
	v, err := hdrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (hdrs *HTTPDetectorResultSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(hdrs.fields) > 1 {
		return nil, errors.New("ent: HTTPDetectorResultSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := hdrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) BoolsX(ctx context.Context) []bool {
	v, err := hdrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (hdrs *HTTPDetectorResultSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = hdrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{httpdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: HTTPDetectorResultSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (hdrs *HTTPDetectorResultSelect) BoolX(ctx context.Context) bool {
	v, err := hdrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hdrs *HTTPDetectorResultSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hdrs.sqlQuery().Query()
	if err := hdrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hdrs *HTTPDetectorResultSelect) sqlQuery() sql.Querier {
	selector := hdrs.sql
	selector.Select(selector.Columns(hdrs.fields...)...)
	return selector
}
