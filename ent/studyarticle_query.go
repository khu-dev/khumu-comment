// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/predicate"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"
)

// StudyArticleQuery is the builder for querying StudyArticle entities.
type StudyArticleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.StudyArticle
	// eager-loading edges.
	withComments *CommentQuery
	withAuthor   *KhumuUserQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StudyArticleQuery builder.
func (saq *StudyArticleQuery) Where(ps ...predicate.StudyArticle) *StudyArticleQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit adds a limit step to the query.
func (saq *StudyArticleQuery) Limit(limit int) *StudyArticleQuery {
	saq.limit = &limit
	return saq
}

// Offset adds an offset step to the query.
func (saq *StudyArticleQuery) Offset(offset int) *StudyArticleQuery {
	saq.offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *StudyArticleQuery) Unique(unique bool) *StudyArticleQuery {
	saq.unique = &unique
	return saq
}

// Order adds an order step to the query.
func (saq *StudyArticleQuery) Order(o ...OrderFunc) *StudyArticleQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// QueryComments chains the current query on the "comments" edge.
func (saq *StudyArticleQuery) QueryComments() *CommentQuery {
	query := &CommentQuery{config: saq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studyarticle.Table, studyarticle.FieldID, selector),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, studyarticle.CommentsTable, studyarticle.CommentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAuthor chains the current query on the "author" edge.
func (saq *StudyArticleQuery) QueryAuthor() *KhumuUserQuery {
	query := &KhumuUserQuery{config: saq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studyarticle.Table, studyarticle.FieldID, selector),
			sqlgraph.To(khumuuser.Table, khumuuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, studyarticle.AuthorTable, studyarticle.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StudyArticle entity from the query.
// Returns a *NotFoundError when no StudyArticle was found.
func (saq *StudyArticleQuery) First(ctx context.Context) (*StudyArticle, error) {
	nodes, err := saq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{studyarticle.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *StudyArticleQuery) FirstX(ctx context.Context) *StudyArticle {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StudyArticle ID from the query.
// Returns a *NotFoundError when no StudyArticle ID was found.
func (saq *StudyArticleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = saq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{studyarticle.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *StudyArticleQuery) FirstIDX(ctx context.Context) int {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StudyArticle entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one StudyArticle entity is not found.
// Returns a *NotFoundError when no StudyArticle entities are found.
func (saq *StudyArticleQuery) Only(ctx context.Context) (*StudyArticle, error) {
	nodes, err := saq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{studyarticle.Label}
	default:
		return nil, &NotSingularError{studyarticle.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *StudyArticleQuery) OnlyX(ctx context.Context) *StudyArticle {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StudyArticle ID in the query.
// Returns a *NotSingularError when exactly one StudyArticle ID is not found.
// Returns a *NotFoundError when no entities are found.
func (saq *StudyArticleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = saq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = &NotSingularError{studyarticle.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *StudyArticleQuery) OnlyIDX(ctx context.Context) int {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StudyArticles.
func (saq *StudyArticleQuery) All(ctx context.Context) ([]*StudyArticle, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return saq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (saq *StudyArticleQuery) AllX(ctx context.Context) []*StudyArticle {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StudyArticle IDs.
func (saq *StudyArticleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := saq.Select(studyarticle.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *StudyArticleQuery) IDsX(ctx context.Context) []int {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *StudyArticleQuery) Count(ctx context.Context) (int, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return saq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (saq *StudyArticleQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *StudyArticleQuery) Exist(ctx context.Context) (bool, error) {
	if err := saq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return saq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *StudyArticleQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StudyArticleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *StudyArticleQuery) Clone() *StudyArticleQuery {
	if saq == nil {
		return nil
	}
	return &StudyArticleQuery{
		config:       saq.config,
		limit:        saq.limit,
		offset:       saq.offset,
		order:        append([]OrderFunc{}, saq.order...),
		predicates:   append([]predicate.StudyArticle{}, saq.predicates...),
		withComments: saq.withComments.Clone(),
		withAuthor:   saq.withAuthor.Clone(),
		// clone intermediate query.
		sql:  saq.sql.Clone(),
		path: saq.path,
	}
}

// WithComments tells the query-builder to eager-load the nodes that are connected to
// the "comments" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *StudyArticleQuery) WithComments(opts ...func(*CommentQuery)) *StudyArticleQuery {
	query := &CommentQuery{config: saq.config}
	for _, opt := range opts {
		opt(query)
	}
	saq.withComments = query
	return saq
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *StudyArticleQuery) WithAuthor(opts ...func(*KhumuUserQuery)) *StudyArticleQuery {
	query := &KhumuUserQuery{config: saq.config}
	for _, opt := range opts {
		opt(query)
	}
	saq.withAuthor = query
	return saq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (saq *StudyArticleQuery) GroupBy(field string, fields ...string) *StudyArticleGroupBy {
	group := &StudyArticleGroupBy{config: saq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return saq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (saq *StudyArticleQuery) Select(field string, fields ...string) *StudyArticleSelect {
	saq.fields = append([]string{field}, fields...)
	return &StudyArticleSelect{StudyArticleQuery: saq}
}

func (saq *StudyArticleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range saq.fields {
		if !studyarticle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.sql = prev
	}
	return nil
}

func (saq *StudyArticleQuery) sqlAll(ctx context.Context) ([]*StudyArticle, error) {
	var (
		nodes       = []*StudyArticle{}
		withFKs     = saq.withFKs
		_spec       = saq.querySpec()
		loadedTypes = [2]bool{
			saq.withComments != nil,
			saq.withAuthor != nil,
		}
	)
	if saq.withAuthor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, studyarticle.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &StudyArticle{config: saq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, saq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := saq.withComments; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*StudyArticle)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Comments = []*Comment{}
		}
		query.withFKs = true
		query.Where(predicate.Comment(func(s *sql.Selector) {
			s.Where(sql.InValues(studyarticle.CommentsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.study_article_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "study_article_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "study_article_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Comments = append(node.Edges.Comments, n)
		}
	}

	if query := saq.withAuthor; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*StudyArticle)
		for i := range nodes {
			if nodes[i].author_id == nil {
				continue
			}
			fk := *nodes[i].author_id
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(khumuuser.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "author_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Author = n
			}
		}
	}

	return nodes, nil
}

func (saq *StudyArticleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := saq.querySpec()
	return sqlgraph.CountNodes(ctx, saq.driver, _spec)
}

func (saq *StudyArticleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := saq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (saq *StudyArticleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   studyarticle.Table,
			Columns: studyarticle.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: studyarticle.FieldID,
			},
		},
		From:   saq.sql,
		Unique: true,
	}
	if unique := saq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := saq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studyarticle.FieldID)
		for i := range fields {
			if fields[i] != studyarticle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := saq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := saq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := saq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := saq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (saq *StudyArticleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(studyarticle.Table)
	selector := builder.Select(t1.Columns(studyarticle.Columns...)...).From(t1)
	if saq.sql != nil {
		selector = saq.sql
		selector.Select(selector.Columns(studyarticle.Columns...)...)
	}
	for _, p := range saq.predicates {
		p(selector)
	}
	for _, p := range saq.order {
		p(selector)
	}
	if offset := saq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := saq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StudyArticleGroupBy is the group-by builder for StudyArticle entities.
type StudyArticleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *StudyArticleGroupBy) Aggregate(fns ...AggregateFunc) *StudyArticleGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the group-by query and scans the result into the given value.
func (sagb *StudyArticleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sagb.path(ctx)
	if err != nil {
		return err
	}
	sagb.sql = query
	return sagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sagb *StudyArticleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sagb.fields) > 1 {
		return nil, errors.New("ent: StudyArticleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) StringsX(ctx context.Context) []string {
	v, err := sagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sagb *StudyArticleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = fmt.Errorf("ent: StudyArticleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) StringX(ctx context.Context) string {
	v, err := sagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sagb *StudyArticleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sagb.fields) > 1 {
		return nil, errors.New("ent: StudyArticleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) IntsX(ctx context.Context) []int {
	v, err := sagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sagb *StudyArticleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = fmt.Errorf("ent: StudyArticleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) IntX(ctx context.Context) int {
	v, err := sagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sagb *StudyArticleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sagb.fields) > 1 {
		return nil, errors.New("ent: StudyArticleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sagb *StudyArticleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = fmt.Errorf("ent: StudyArticleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sagb *StudyArticleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sagb.fields) > 1 {
		return nil, errors.New("ent: StudyArticleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sagb *StudyArticleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = fmt.Errorf("ent: StudyArticleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sagb *StudyArticleGroupBy) BoolX(ctx context.Context) bool {
	v, err := sagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sagb *StudyArticleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sagb.fields {
		if !studyarticle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sagb *StudyArticleGroupBy) sqlQuery() *sql.Selector {
	selector := sagb.sql
	columns := make([]string, 0, len(sagb.fields)+len(sagb.fns))
	columns = append(columns, sagb.fields...)
	for _, fn := range sagb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sagb.fields...)
}

// StudyArticleSelect is the builder for selecting fields of StudyArticle entities.
type StudyArticleSelect struct {
	*StudyArticleQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sas *StudyArticleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	sas.sql = sas.StudyArticleQuery.sqlQuery(ctx)
	return sas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sas *StudyArticleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sas *StudyArticleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sas.fields) > 1 {
		return nil, errors.New("ent: StudyArticleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sas *StudyArticleSelect) StringsX(ctx context.Context) []string {
	v, err := sas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sas *StudyArticleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = fmt.Errorf("ent: StudyArticleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sas *StudyArticleSelect) StringX(ctx context.Context) string {
	v, err := sas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sas *StudyArticleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sas.fields) > 1 {
		return nil, errors.New("ent: StudyArticleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sas *StudyArticleSelect) IntsX(ctx context.Context) []int {
	v, err := sas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sas *StudyArticleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = fmt.Errorf("ent: StudyArticleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sas *StudyArticleSelect) IntX(ctx context.Context) int {
	v, err := sas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sas *StudyArticleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sas.fields) > 1 {
		return nil, errors.New("ent: StudyArticleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sas *StudyArticleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sas *StudyArticleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = fmt.Errorf("ent: StudyArticleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sas *StudyArticleSelect) Float64X(ctx context.Context) float64 {
	v, err := sas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sas *StudyArticleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sas.fields) > 1 {
		return nil, errors.New("ent: StudyArticleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sas *StudyArticleSelect) BoolsX(ctx context.Context) []bool {
	v, err := sas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sas *StudyArticleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{studyarticle.Label}
	default:
		err = fmt.Errorf("ent: StudyArticleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sas *StudyArticleSelect) BoolX(ctx context.Context) bool {
	v, err := sas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sas *StudyArticleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sas.sqlQuery().Query()
	if err := sas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sas *StudyArticleSelect) sqlQuery() sql.Querier {
	selector := sas.sql
	selector.Select(selector.Columns(sas.fields...)...)
	return selector
}
