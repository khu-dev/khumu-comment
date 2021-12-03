// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/khu-dev/khumu-comment/ent/migrate"

	"github.com/khu-dev/khumu-comment/ent/article"
	"github.com/khu-dev/khumu-comment/ent/comment"
	"github.com/khu-dev/khumu-comment/ent/khumuuser"
	"github.com/khu-dev/khumu-comment/ent/likecomment"
	"github.com/khu-dev/khumu-comment/ent/studyarticle"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Article is the client for interacting with the Article builders.
	Article *ArticleClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// KhumuUser is the client for interacting with the KhumuUser builders.
	KhumuUser *KhumuUserClient
	// LikeComment is the client for interacting with the LikeComment builders.
	LikeComment *LikeCommentClient
	// StudyArticle is the client for interacting with the StudyArticle builders.
	StudyArticle *StudyArticleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Article = NewArticleClient(c.config)
	c.Comment = NewCommentClient(c.config)
	c.KhumuUser = NewKhumuUserClient(c.config)
	c.LikeComment = NewLikeCommentClient(c.config)
	c.StudyArticle = NewStudyArticleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Article:      NewArticleClient(cfg),
		Comment:      NewCommentClient(cfg),
		KhumuUser:    NewKhumuUserClient(cfg),
		LikeComment:  NewLikeCommentClient(cfg),
		StudyArticle: NewStudyArticleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:       cfg,
		Article:      NewArticleClient(cfg),
		Comment:      NewCommentClient(cfg),
		KhumuUser:    NewKhumuUserClient(cfg),
		LikeComment:  NewLikeCommentClient(cfg),
		StudyArticle: NewStudyArticleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Article.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Article.Use(hooks...)
	c.Comment.Use(hooks...)
	c.KhumuUser.Use(hooks...)
	c.LikeComment.Use(hooks...)
	c.StudyArticle.Use(hooks...)
}

// ArticleClient is a client for the Article schema.
type ArticleClient struct {
	config
}

// NewArticleClient returns a client for the Article from the given config.
func NewArticleClient(c config) *ArticleClient {
	return &ArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `article.Hooks(f(g(h())))`.
func (c *ArticleClient) Use(hooks ...Hook) {
	c.hooks.Article = append(c.hooks.Article, hooks...)
}

// Create returns a create builder for Article.
func (c *ArticleClient) Create() *ArticleCreate {
	mutation := newArticleMutation(c.config, OpCreate)
	return &ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Article entities.
func (c *ArticleClient) CreateBulk(builders ...*ArticleCreate) *ArticleCreateBulk {
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Article.
func (c *ArticleClient) Update() *ArticleUpdate {
	mutation := newArticleMutation(c.config, OpUpdate)
	return &ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleClient) UpdateOne(a *Article) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticle(a))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleClient) UpdateOneID(id int) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticleID(id))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Article.
func (c *ArticleClient) Delete() *ArticleDelete {
	mutation := newArticleMutation(c.config, OpDelete)
	return &ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ArticleClient) DeleteOne(a *Article) *ArticleDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ArticleClient) DeleteOneID(id int) *ArticleDeleteOne {
	builder := c.Delete().Where(article.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleDeleteOne{builder}
}

// Query returns a query builder for Article.
func (c *ArticleClient) Query() *ArticleQuery {
	return &ArticleQuery{
		config: c.config,
	}
}

// Get returns a Article entity by its id.
func (c *ArticleClient) Get(ctx context.Context, id int) (*Article, error) {
	return c.Query().Where(article.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleClient) GetX(ctx context.Context, id int) *Article {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComments queries the comments edge of a Article.
func (c *ArticleClient) QueryComments(a *Article) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, article.CommentsTable, article.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a Article.
func (c *ArticleClient) QueryAuthor(a *Article) *KhumuUserQuery {
	query := &KhumuUserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(khumuuser.Table, khumuuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, article.AuthorTable, article.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArticleClient) Hooks() []Hook {
	return c.hooks.Article
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `comment.Hooks(f(g(h())))`.
func (c *CommentClient) Use(hooks ...Hook) {
	c.hooks.Comment = append(c.hooks.Comment, hooks...)
}

// Create returns a create builder for Comment.
func (c *CommentClient) Create() *CommentCreate {
	mutation := newCommentMutation(c.config, OpCreate)
	return &CommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Comment entities.
func (c *CommentClient) CreateBulk(builders ...*CommentCreate) *CommentCreateBulk {
	return &CommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	mutation := newCommentMutation(c.config, OpUpdate)
	return &CommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withComment(co))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id int) *CommentUpdateOne {
	mutation := newCommentMutation(c.config, OpUpdateOne, withCommentID(id))
	return &CommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	mutation := newCommentMutation(c.config, OpDelete)
	return &CommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CommentClient) DeleteOneID(id int) *CommentDeleteOne {
	builder := c.Delete().Where(comment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CommentDeleteOne{builder}
}

// Query returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{
		config: c.config,
	}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id int) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id int) *Comment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthor queries the author edge of a Comment.
func (c *CommentClient) QueryAuthor(co *Comment) *KhumuUserQuery {
	query := &KhumuUserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(khumuuser.Table, khumuuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.AuthorTable, comment.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryArticle queries the article edge of a Comment.
func (c *CommentClient) QueryArticle(co *Comment) *ArticleQuery {
	query := &ArticleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.ArticleTable, comment.ArticleColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStudyArticle queries the studyArticle edge of a Comment.
func (c *CommentClient) QueryStudyArticle(co *Comment) *StudyArticleQuery {
	query := &StudyArticleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(studyarticle.Table, studyarticle.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.StudyArticleTable, comment.StudyArticleColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Comment.
func (c *CommentClient) QueryParent(co *Comment) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, comment.ParentTable, comment.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Comment.
func (c *CommentClient) QueryChildren(co *Comment) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, comment.ChildrenTable, comment.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLike queries the like edge of a Comment.
func (c *CommentClient) QueryLike(co *Comment) *LikeCommentQuery {
	query := &LikeCommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(comment.Table, comment.FieldID, id),
			sqlgraph.To(likecomment.Table, likecomment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, comment.LikeTable, comment.LikeColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CommentClient) Hooks() []Hook {
	return c.hooks.Comment
}

// KhumuUserClient is a client for the KhumuUser schema.
type KhumuUserClient struct {
	config
}

// NewKhumuUserClient returns a client for the KhumuUser from the given config.
func NewKhumuUserClient(c config) *KhumuUserClient {
	return &KhumuUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `khumuuser.Hooks(f(g(h())))`.
func (c *KhumuUserClient) Use(hooks ...Hook) {
	c.hooks.KhumuUser = append(c.hooks.KhumuUser, hooks...)
}

// Create returns a create builder for KhumuUser.
func (c *KhumuUserClient) Create() *KhumuUserCreate {
	mutation := newKhumuUserMutation(c.config, OpCreate)
	return &KhumuUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of KhumuUser entities.
func (c *KhumuUserClient) CreateBulk(builders ...*KhumuUserCreate) *KhumuUserCreateBulk {
	return &KhumuUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for KhumuUser.
func (c *KhumuUserClient) Update() *KhumuUserUpdate {
	mutation := newKhumuUserMutation(c.config, OpUpdate)
	return &KhumuUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *KhumuUserClient) UpdateOne(ku *KhumuUser) *KhumuUserUpdateOne {
	mutation := newKhumuUserMutation(c.config, OpUpdateOne, withKhumuUser(ku))
	return &KhumuUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *KhumuUserClient) UpdateOneID(id string) *KhumuUserUpdateOne {
	mutation := newKhumuUserMutation(c.config, OpUpdateOne, withKhumuUserID(id))
	return &KhumuUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for KhumuUser.
func (c *KhumuUserClient) Delete() *KhumuUserDelete {
	mutation := newKhumuUserMutation(c.config, OpDelete)
	return &KhumuUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *KhumuUserClient) DeleteOne(ku *KhumuUser) *KhumuUserDeleteOne {
	return c.DeleteOneID(ku.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *KhumuUserClient) DeleteOneID(id string) *KhumuUserDeleteOne {
	builder := c.Delete().Where(khumuuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &KhumuUserDeleteOne{builder}
}

// Query returns a query builder for KhumuUser.
func (c *KhumuUserClient) Query() *KhumuUserQuery {
	return &KhumuUserQuery{
		config: c.config,
	}
}

// Get returns a KhumuUser entity by its id.
func (c *KhumuUserClient) Get(ctx context.Context, id string) (*KhumuUser, error) {
	return c.Query().Where(khumuuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *KhumuUserClient) GetX(ctx context.Context, id string) *KhumuUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComments queries the comments edge of a KhumuUser.
func (c *KhumuUserClient) QueryComments(ku *KhumuUser) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ku.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(khumuuser.Table, khumuuser.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, khumuuser.CommentsTable, khumuuser.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(ku.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryArticles queries the articles edge of a KhumuUser.
func (c *KhumuUserClient) QueryArticles(ku *KhumuUser) *ArticleQuery {
	query := &ArticleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ku.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(khumuuser.Table, khumuuser.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, khumuuser.ArticlesTable, khumuuser.ArticlesColumn),
		)
		fromV = sqlgraph.Neighbors(ku.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStudyArticles queries the studyArticles edge of a KhumuUser.
func (c *KhumuUserClient) QueryStudyArticles(ku *KhumuUser) *StudyArticleQuery {
	query := &StudyArticleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ku.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(khumuuser.Table, khumuuser.FieldID, id),
			sqlgraph.To(studyarticle.Table, studyarticle.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, khumuuser.StudyArticlesTable, khumuuser.StudyArticlesColumn),
		)
		fromV = sqlgraph.Neighbors(ku.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLike queries the like edge of a KhumuUser.
func (c *KhumuUserClient) QueryLike(ku *KhumuUser) *LikeCommentQuery {
	query := &LikeCommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ku.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(khumuuser.Table, khumuuser.FieldID, id),
			sqlgraph.To(likecomment.Table, likecomment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, khumuuser.LikeTable, khumuuser.LikeColumn),
		)
		fromV = sqlgraph.Neighbors(ku.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *KhumuUserClient) Hooks() []Hook {
	return c.hooks.KhumuUser
}

// LikeCommentClient is a client for the LikeComment schema.
type LikeCommentClient struct {
	config
}

// NewLikeCommentClient returns a client for the LikeComment from the given config.
func NewLikeCommentClient(c config) *LikeCommentClient {
	return &LikeCommentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `likecomment.Hooks(f(g(h())))`.
func (c *LikeCommentClient) Use(hooks ...Hook) {
	c.hooks.LikeComment = append(c.hooks.LikeComment, hooks...)
}

// Create returns a create builder for LikeComment.
func (c *LikeCommentClient) Create() *LikeCommentCreate {
	mutation := newLikeCommentMutation(c.config, OpCreate)
	return &LikeCommentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LikeComment entities.
func (c *LikeCommentClient) CreateBulk(builders ...*LikeCommentCreate) *LikeCommentCreateBulk {
	return &LikeCommentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LikeComment.
func (c *LikeCommentClient) Update() *LikeCommentUpdate {
	mutation := newLikeCommentMutation(c.config, OpUpdate)
	return &LikeCommentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LikeCommentClient) UpdateOne(lc *LikeComment) *LikeCommentUpdateOne {
	mutation := newLikeCommentMutation(c.config, OpUpdateOne, withLikeComment(lc))
	return &LikeCommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LikeCommentClient) UpdateOneID(id int) *LikeCommentUpdateOne {
	mutation := newLikeCommentMutation(c.config, OpUpdateOne, withLikeCommentID(id))
	return &LikeCommentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LikeComment.
func (c *LikeCommentClient) Delete() *LikeCommentDelete {
	mutation := newLikeCommentMutation(c.config, OpDelete)
	return &LikeCommentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LikeCommentClient) DeleteOne(lc *LikeComment) *LikeCommentDeleteOne {
	return c.DeleteOneID(lc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LikeCommentClient) DeleteOneID(id int) *LikeCommentDeleteOne {
	builder := c.Delete().Where(likecomment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LikeCommentDeleteOne{builder}
}

// Query returns a query builder for LikeComment.
func (c *LikeCommentClient) Query() *LikeCommentQuery {
	return &LikeCommentQuery{
		config: c.config,
	}
}

// Get returns a LikeComment entity by its id.
func (c *LikeCommentClient) Get(ctx context.Context, id int) (*LikeComment, error) {
	return c.Query().Where(likecomment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LikeCommentClient) GetX(ctx context.Context, id int) *LikeComment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLikedBy queries the likedBy edge of a LikeComment.
func (c *LikeCommentClient) QueryLikedBy(lc *LikeComment) *KhumuUserQuery {
	query := &KhumuUserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := lc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(likecomment.Table, likecomment.FieldID, id),
			sqlgraph.To(khumuuser.Table, khumuuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, likecomment.LikedByTable, likecomment.LikedByColumn),
		)
		fromV = sqlgraph.Neighbors(lc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAbout queries the about edge of a LikeComment.
func (c *LikeCommentClient) QueryAbout(lc *LikeComment) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := lc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(likecomment.Table, likecomment.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, likecomment.AboutTable, likecomment.AboutColumn),
		)
		fromV = sqlgraph.Neighbors(lc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LikeCommentClient) Hooks() []Hook {
	return c.hooks.LikeComment
}

// StudyArticleClient is a client for the StudyArticle schema.
type StudyArticleClient struct {
	config
}

// NewStudyArticleClient returns a client for the StudyArticle from the given config.
func NewStudyArticleClient(c config) *StudyArticleClient {
	return &StudyArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `studyarticle.Hooks(f(g(h())))`.
func (c *StudyArticleClient) Use(hooks ...Hook) {
	c.hooks.StudyArticle = append(c.hooks.StudyArticle, hooks...)
}

// Create returns a create builder for StudyArticle.
func (c *StudyArticleClient) Create() *StudyArticleCreate {
	mutation := newStudyArticleMutation(c.config, OpCreate)
	return &StudyArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StudyArticle entities.
func (c *StudyArticleClient) CreateBulk(builders ...*StudyArticleCreate) *StudyArticleCreateBulk {
	return &StudyArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StudyArticle.
func (c *StudyArticleClient) Update() *StudyArticleUpdate {
	mutation := newStudyArticleMutation(c.config, OpUpdate)
	return &StudyArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StudyArticleClient) UpdateOne(sa *StudyArticle) *StudyArticleUpdateOne {
	mutation := newStudyArticleMutation(c.config, OpUpdateOne, withStudyArticle(sa))
	return &StudyArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StudyArticleClient) UpdateOneID(id int) *StudyArticleUpdateOne {
	mutation := newStudyArticleMutation(c.config, OpUpdateOne, withStudyArticleID(id))
	return &StudyArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StudyArticle.
func (c *StudyArticleClient) Delete() *StudyArticleDelete {
	mutation := newStudyArticleMutation(c.config, OpDelete)
	return &StudyArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StudyArticleClient) DeleteOne(sa *StudyArticle) *StudyArticleDeleteOne {
	return c.DeleteOneID(sa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StudyArticleClient) DeleteOneID(id int) *StudyArticleDeleteOne {
	builder := c.Delete().Where(studyarticle.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StudyArticleDeleteOne{builder}
}

// Query returns a query builder for StudyArticle.
func (c *StudyArticleClient) Query() *StudyArticleQuery {
	return &StudyArticleQuery{
		config: c.config,
	}
}

// Get returns a StudyArticle entity by its id.
func (c *StudyArticleClient) Get(ctx context.Context, id int) (*StudyArticle, error) {
	return c.Query().Where(studyarticle.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StudyArticleClient) GetX(ctx context.Context, id int) *StudyArticle {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryComments queries the comments edge of a StudyArticle.
func (c *StudyArticleClient) QueryComments(sa *StudyArticle) *CommentQuery {
	query := &CommentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := sa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(studyarticle.Table, studyarticle.FieldID, id),
			sqlgraph.To(comment.Table, comment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, studyarticle.CommentsTable, studyarticle.CommentsColumn),
		)
		fromV = sqlgraph.Neighbors(sa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthor queries the author edge of a StudyArticle.
func (c *StudyArticleClient) QueryAuthor(sa *StudyArticle) *KhumuUserQuery {
	query := &KhumuUserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := sa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(studyarticle.Table, studyarticle.FieldID, id),
			sqlgraph.To(khumuuser.Table, khumuuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, studyarticle.AuthorTable, studyarticle.AuthorColumn),
		)
		fromV = sqlgraph.Neighbors(sa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StudyArticleClient) Hooks() []Hook {
	return c.hooks.StudyArticle
}
