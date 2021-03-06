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
	"github.com/vicanso/cybertect/ent/dnsdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
)

// DNSDetectorResultQuery is the builder for querying DNSDetectorResult entities.
type DNSDetectorResultQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.DNSDetectorResult
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DNSDetectorResultQuery builder.
func (ddrq *DNSDetectorResultQuery) Where(ps ...predicate.DNSDetectorResult) *DNSDetectorResultQuery {
	ddrq.predicates = append(ddrq.predicates, ps...)
	return ddrq
}

// Limit adds a limit step to the query.
func (ddrq *DNSDetectorResultQuery) Limit(limit int) *DNSDetectorResultQuery {
	ddrq.limit = &limit
	return ddrq
}

// Offset adds an offset step to the query.
func (ddrq *DNSDetectorResultQuery) Offset(offset int) *DNSDetectorResultQuery {
	ddrq.offset = &offset
	return ddrq
}

// Order adds an order step to the query.
func (ddrq *DNSDetectorResultQuery) Order(o ...OrderFunc) *DNSDetectorResultQuery {
	ddrq.order = append(ddrq.order, o...)
	return ddrq
}

// First returns the first DNSDetectorResult entity from the query.
// Returns a *NotFoundError when no DNSDetectorResult was found.
func (ddrq *DNSDetectorResultQuery) First(ctx context.Context) (*DNSDetectorResult, error) {
	nodes, err := ddrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dnsdetectorresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ddrq *DNSDetectorResultQuery) FirstX(ctx context.Context) *DNSDetectorResult {
	node, err := ddrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DNSDetectorResult ID from the query.
// Returns a *NotFoundError when no DNSDetectorResult ID was found.
func (ddrq *DNSDetectorResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ddrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dnsdetectorresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ddrq *DNSDetectorResultQuery) FirstIDX(ctx context.Context) int {
	id, err := ddrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DNSDetectorResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DNSDetectorResult entity is not found.
// Returns a *NotFoundError when no DNSDetectorResult entities are found.
func (ddrq *DNSDetectorResultQuery) Only(ctx context.Context) (*DNSDetectorResult, error) {
	nodes, err := ddrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dnsdetectorresult.Label}
	default:
		return nil, &NotSingularError{dnsdetectorresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ddrq *DNSDetectorResultQuery) OnlyX(ctx context.Context) *DNSDetectorResult {
	node, err := ddrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DNSDetectorResult ID in the query.
// Returns a *NotSingularError when exactly one DNSDetectorResult ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ddrq *DNSDetectorResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ddrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = &NotSingularError{dnsdetectorresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ddrq *DNSDetectorResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := ddrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DNSDetectorResults.
func (ddrq *DNSDetectorResultQuery) All(ctx context.Context) ([]*DNSDetectorResult, error) {
	if err := ddrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ddrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ddrq *DNSDetectorResultQuery) AllX(ctx context.Context) []*DNSDetectorResult {
	nodes, err := ddrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DNSDetectorResult IDs.
func (ddrq *DNSDetectorResultQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ddrq.Select(dnsdetectorresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ddrq *DNSDetectorResultQuery) IDsX(ctx context.Context) []int {
	ids, err := ddrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ddrq *DNSDetectorResultQuery) Count(ctx context.Context) (int, error) {
	if err := ddrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ddrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ddrq *DNSDetectorResultQuery) CountX(ctx context.Context) int {
	count, err := ddrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ddrq *DNSDetectorResultQuery) Exist(ctx context.Context) (bool, error) {
	if err := ddrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ddrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ddrq *DNSDetectorResultQuery) ExistX(ctx context.Context) bool {
	exist, err := ddrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DNSDetectorResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ddrq *DNSDetectorResultQuery) Clone() *DNSDetectorResultQuery {
	if ddrq == nil {
		return nil
	}
	return &DNSDetectorResultQuery{
		config:     ddrq.config,
		limit:      ddrq.limit,
		offset:     ddrq.offset,
		order:      append([]OrderFunc{}, ddrq.order...),
		predicates: append([]predicate.DNSDetectorResult{}, ddrq.predicates...),
		// clone intermediate query.
		sql:  ddrq.sql.Clone(),
		path: ddrq.path,
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
//	client.DNSDetectorResult.Query().
//		GroupBy(dnsdetectorresult.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ddrq *DNSDetectorResultQuery) GroupBy(field string, fields ...string) *DNSDetectorResultGroupBy {
	group := &DNSDetectorResultGroupBy{config: ddrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ddrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ddrq.sqlQuery(), nil
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
//	client.DNSDetectorResult.Query().
//		Select(dnsdetectorresult.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (ddrq *DNSDetectorResultQuery) Select(field string, fields ...string) *DNSDetectorResultSelect {
	ddrq.fields = append([]string{field}, fields...)
	return &DNSDetectorResultSelect{DNSDetectorResultQuery: ddrq}
}

func (ddrq *DNSDetectorResultQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ddrq.fields {
		if !dnsdetectorresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ddrq.path != nil {
		prev, err := ddrq.path(ctx)
		if err != nil {
			return err
		}
		ddrq.sql = prev
	}
	return nil
}

func (ddrq *DNSDetectorResultQuery) sqlAll(ctx context.Context) ([]*DNSDetectorResult, error) {
	var (
		nodes = []*DNSDetectorResult{}
		_spec = ddrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DNSDetectorResult{config: ddrq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ddrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ddrq *DNSDetectorResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ddrq.querySpec()
	return sqlgraph.CountNodes(ctx, ddrq.driver, _spec)
}

func (ddrq *DNSDetectorResultQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ddrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ddrq *DNSDetectorResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsdetectorresult.Table,
			Columns: dnsdetectorresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dnsdetectorresult.FieldID,
			},
		},
		From:   ddrq.sql,
		Unique: true,
	}
	if fields := ddrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dnsdetectorresult.FieldID)
		for i := range fields {
			if fields[i] != dnsdetectorresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ddrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ddrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ddrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ddrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, dnsdetectorresult.ValidColumn)
			}
		}
	}
	return _spec
}

func (ddrq *DNSDetectorResultQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ddrq.driver.Dialect())
	t1 := builder.Table(dnsdetectorresult.Table)
	selector := builder.Select(t1.Columns(dnsdetectorresult.Columns...)...).From(t1)
	if ddrq.sql != nil {
		selector = ddrq.sql
		selector.Select(selector.Columns(dnsdetectorresult.Columns...)...)
	}
	for _, p := range ddrq.predicates {
		p(selector)
	}
	for _, p := range ddrq.order {
		p(selector, dnsdetectorresult.ValidColumn)
	}
	if offset := ddrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ddrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DNSDetectorResultGroupBy is the group-by builder for DNSDetectorResult entities.
type DNSDetectorResultGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ddrgb *DNSDetectorResultGroupBy) Aggregate(fns ...AggregateFunc) *DNSDetectorResultGroupBy {
	ddrgb.fns = append(ddrgb.fns, fns...)
	return ddrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ddrgb *DNSDetectorResultGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ddrgb.path(ctx)
	if err != nil {
		return err
	}
	ddrgb.sql = query
	return ddrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ddrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ddrgb *DNSDetectorResultGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ddrgb.fields) > 1 {
		return nil, errors.New("ent: DNSDetectorResultGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ddrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) StringsX(ctx context.Context) []string {
	v, err := ddrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ddrgb *DNSDetectorResultGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ddrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: DNSDetectorResultGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) StringX(ctx context.Context) string {
	v, err := ddrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ddrgb *DNSDetectorResultGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ddrgb.fields) > 1 {
		return nil, errors.New("ent: DNSDetectorResultGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ddrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) IntsX(ctx context.Context) []int {
	v, err := ddrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ddrgb *DNSDetectorResultGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ddrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: DNSDetectorResultGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) IntX(ctx context.Context) int {
	v, err := ddrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ddrgb *DNSDetectorResultGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ddrgb.fields) > 1 {
		return nil, errors.New("ent: DNSDetectorResultGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ddrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ddrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ddrgb *DNSDetectorResultGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ddrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: DNSDetectorResultGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ddrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ddrgb *DNSDetectorResultGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ddrgb.fields) > 1 {
		return nil, errors.New("ent: DNSDetectorResultGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ddrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ddrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ddrgb *DNSDetectorResultGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ddrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: DNSDetectorResultGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ddrgb *DNSDetectorResultGroupBy) BoolX(ctx context.Context) bool {
	v, err := ddrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ddrgb *DNSDetectorResultGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ddrgb.fields {
		if !dnsdetectorresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ddrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ddrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ddrgb *DNSDetectorResultGroupBy) sqlQuery() *sql.Selector {
	selector := ddrgb.sql
	columns := make([]string, 0, len(ddrgb.fields)+len(ddrgb.fns))
	columns = append(columns, ddrgb.fields...)
	for _, fn := range ddrgb.fns {
		columns = append(columns, fn(selector, dnsdetectorresult.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ddrgb.fields...)
}

// DNSDetectorResultSelect is the builder for selecting fields of DNSDetectorResult entities.
type DNSDetectorResultSelect struct {
	*DNSDetectorResultQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ddrs *DNSDetectorResultSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ddrs.prepareQuery(ctx); err != nil {
		return err
	}
	ddrs.sql = ddrs.DNSDetectorResultQuery.sqlQuery()
	return ddrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ddrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ddrs *DNSDetectorResultSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ddrs.fields) > 1 {
		return nil, errors.New("ent: DNSDetectorResultSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ddrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) StringsX(ctx context.Context) []string {
	v, err := ddrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ddrs *DNSDetectorResultSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ddrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: DNSDetectorResultSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) StringX(ctx context.Context) string {
	v, err := ddrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ddrs *DNSDetectorResultSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ddrs.fields) > 1 {
		return nil, errors.New("ent: DNSDetectorResultSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ddrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) IntsX(ctx context.Context) []int {
	v, err := ddrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ddrs *DNSDetectorResultSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ddrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: DNSDetectorResultSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) IntX(ctx context.Context) int {
	v, err := ddrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ddrs *DNSDetectorResultSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ddrs.fields) > 1 {
		return nil, errors.New("ent: DNSDetectorResultSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ddrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ddrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ddrs *DNSDetectorResultSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ddrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: DNSDetectorResultSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) Float64X(ctx context.Context) float64 {
	v, err := ddrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ddrs *DNSDetectorResultSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ddrs.fields) > 1 {
		return nil, errors.New("ent: DNSDetectorResultSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ddrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) BoolsX(ctx context.Context) []bool {
	v, err := ddrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ddrs *DNSDetectorResultSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ddrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{dnsdetectorresult.Label}
	default:
		err = fmt.Errorf("ent: DNSDetectorResultSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ddrs *DNSDetectorResultSelect) BoolX(ctx context.Context) bool {
	v, err := ddrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ddrs *DNSDetectorResultSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ddrs.sqlQuery().Query()
	if err := ddrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ddrs *DNSDetectorResultSelect) sqlQuery() sql.Querier {
	selector := ddrs.sql
	selector.Select(selector.Columns(ddrs.fields...)...)
	return selector
}
