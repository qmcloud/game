// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/qmcloud/game/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/qmcloud/game/ent/member"
	"github.com/qmcloud/game/ent/memberrank"
	"github.com/qmcloud/game/ent/oauthprovider"
	"github.com/qmcloud/game/ent/token"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Member is the client for interacting with the Member builders.
	Member *MemberClient
	// MemberRank is the client for interacting with the MemberRank builders.
	MemberRank *MemberRankClient
	// OauthProvider is the client for interacting with the OauthProvider builders.
	OauthProvider *OauthProviderClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Member = NewMemberClient(c.config)
	c.MemberRank = NewMemberRankClient(c.config)
	c.OauthProvider = NewOauthProviderClient(c.config)
	c.Token = NewTokenClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Member:        NewMemberClient(cfg),
		MemberRank:    NewMemberRankClient(cfg),
		OauthProvider: NewOauthProviderClient(cfg),
		Token:         NewTokenClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:           ctx,
		config:        cfg,
		Member:        NewMemberClient(cfg),
		MemberRank:    NewMemberRankClient(cfg),
		OauthProvider: NewOauthProviderClient(cfg),
		Token:         NewTokenClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Member.
//		Query().
//		Count(ctx)
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
	c.Member.Use(hooks...)
	c.MemberRank.Use(hooks...)
	c.OauthProvider.Use(hooks...)
	c.Token.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Member.Intercept(interceptors...)
	c.MemberRank.Intercept(interceptors...)
	c.OauthProvider.Intercept(interceptors...)
	c.Token.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *MemberMutation:
		return c.Member.mutate(ctx, m)
	case *MemberRankMutation:
		return c.MemberRank.mutate(ctx, m)
	case *OauthProviderMutation:
		return c.OauthProvider.mutate(ctx, m)
	case *TokenMutation:
		return c.Token.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// MemberClient is a client for the Member schema.
type MemberClient struct {
	config
}

// NewMemberClient returns a client for the Member from the given config.
func NewMemberClient(c config) *MemberClient {
	return &MemberClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `member.Hooks(f(g(h())))`.
func (c *MemberClient) Use(hooks ...Hook) {
	c.hooks.Member = append(c.hooks.Member, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `member.Intercept(f(g(h())))`.
func (c *MemberClient) Intercept(interceptors ...Interceptor) {
	c.inters.Member = append(c.inters.Member, interceptors...)
}

// Create returns a builder for creating a Member entity.
func (c *MemberClient) Create() *MemberCreate {
	mutation := newMemberMutation(c.config, OpCreate)
	return &MemberCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Member entities.
func (c *MemberClient) CreateBulk(builders ...*MemberCreate) *MemberCreateBulk {
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MemberClient) MapCreateBulk(slice any, setFunc func(*MemberCreate, int)) *MemberCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MemberCreateBulk{err: fmt.Errorf("calling to MemberClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MemberCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MemberCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Member.
func (c *MemberClient) Update() *MemberUpdate {
	mutation := newMemberMutation(c.config, OpUpdate)
	return &MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MemberClient) UpdateOne(m *Member) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMember(m))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MemberClient) UpdateOneID(id uuid.UUID) *MemberUpdateOne {
	mutation := newMemberMutation(c.config, OpUpdateOne, withMemberID(id))
	return &MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Member.
func (c *MemberClient) Delete() *MemberDelete {
	mutation := newMemberMutation(c.config, OpDelete)
	return &MemberDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MemberClient) DeleteOne(m *Member) *MemberDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MemberClient) DeleteOneID(id uuid.UUID) *MemberDeleteOne {
	builder := c.Delete().Where(member.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MemberDeleteOne{builder}
}

// Query returns a query builder for Member.
func (c *MemberClient) Query() *MemberQuery {
	return &MemberQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMember},
		inters: c.Interceptors(),
	}
}

// Get returns a Member entity by its id.
func (c *MemberClient) Get(ctx context.Context, id uuid.UUID) (*Member, error) {
	return c.Query().Where(member.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MemberClient) GetX(ctx context.Context, id uuid.UUID) *Member {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRanks queries the ranks edge of a Member.
func (c *MemberClient) QueryRanks(m *Member) *MemberRankQuery {
	query := (&MemberRankClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, id),
			sqlgraph.To(memberrank.Table, memberrank.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, member.RanksTable, member.RanksColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MemberClient) Hooks() []Hook {
	return c.hooks.Member
}

// Interceptors returns the client interceptors.
func (c *MemberClient) Interceptors() []Interceptor {
	return c.inters.Member
}

func (c *MemberClient) mutate(ctx context.Context, m *MemberMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MemberCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MemberUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MemberUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MemberDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Member mutation op: %q", m.Op())
	}
}

// MemberRankClient is a client for the MemberRank schema.
type MemberRankClient struct {
	config
}

// NewMemberRankClient returns a client for the MemberRank from the given config.
func NewMemberRankClient(c config) *MemberRankClient {
	return &MemberRankClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `memberrank.Hooks(f(g(h())))`.
func (c *MemberRankClient) Use(hooks ...Hook) {
	c.hooks.MemberRank = append(c.hooks.MemberRank, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `memberrank.Intercept(f(g(h())))`.
func (c *MemberRankClient) Intercept(interceptors ...Interceptor) {
	c.inters.MemberRank = append(c.inters.MemberRank, interceptors...)
}

// Create returns a builder for creating a MemberRank entity.
func (c *MemberRankClient) Create() *MemberRankCreate {
	mutation := newMemberRankMutation(c.config, OpCreate)
	return &MemberRankCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MemberRank entities.
func (c *MemberRankClient) CreateBulk(builders ...*MemberRankCreate) *MemberRankCreateBulk {
	return &MemberRankCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MemberRankClient) MapCreateBulk(slice any, setFunc func(*MemberRankCreate, int)) *MemberRankCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MemberRankCreateBulk{err: fmt.Errorf("calling to MemberRankClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MemberRankCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MemberRankCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MemberRank.
func (c *MemberRankClient) Update() *MemberRankUpdate {
	mutation := newMemberRankMutation(c.config, OpUpdate)
	return &MemberRankUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MemberRankClient) UpdateOne(mr *MemberRank) *MemberRankUpdateOne {
	mutation := newMemberRankMutation(c.config, OpUpdateOne, withMemberRank(mr))
	return &MemberRankUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MemberRankClient) UpdateOneID(id uint64) *MemberRankUpdateOne {
	mutation := newMemberRankMutation(c.config, OpUpdateOne, withMemberRankID(id))
	return &MemberRankUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MemberRank.
func (c *MemberRankClient) Delete() *MemberRankDelete {
	mutation := newMemberRankMutation(c.config, OpDelete)
	return &MemberRankDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MemberRankClient) DeleteOne(mr *MemberRank) *MemberRankDeleteOne {
	return c.DeleteOneID(mr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MemberRankClient) DeleteOneID(id uint64) *MemberRankDeleteOne {
	builder := c.Delete().Where(memberrank.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MemberRankDeleteOne{builder}
}

// Query returns a query builder for MemberRank.
func (c *MemberRankClient) Query() *MemberRankQuery {
	return &MemberRankQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMemberRank},
		inters: c.Interceptors(),
	}
}

// Get returns a MemberRank entity by its id.
func (c *MemberRankClient) Get(ctx context.Context, id uint64) (*MemberRank, error) {
	return c.Query().Where(memberrank.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MemberRankClient) GetX(ctx context.Context, id uint64) *MemberRank {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryMembers queries the members edge of a MemberRank.
func (c *MemberRankClient) QueryMembers(mr *MemberRank) *MemberQuery {
	query := (&MemberClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := mr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(memberrank.Table, memberrank.FieldID, id),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, memberrank.MembersTable, memberrank.MembersColumn),
		)
		fromV = sqlgraph.Neighbors(mr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MemberRankClient) Hooks() []Hook {
	return c.hooks.MemberRank
}

// Interceptors returns the client interceptors.
func (c *MemberRankClient) Interceptors() []Interceptor {
	return c.inters.MemberRank
}

func (c *MemberRankClient) mutate(ctx context.Context, m *MemberRankMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MemberRankCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MemberRankUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MemberRankUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MemberRankDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown MemberRank mutation op: %q", m.Op())
	}
}

// OauthProviderClient is a client for the OauthProvider schema.
type OauthProviderClient struct {
	config
}

// NewOauthProviderClient returns a client for the OauthProvider from the given config.
func NewOauthProviderClient(c config) *OauthProviderClient {
	return &OauthProviderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `oauthprovider.Hooks(f(g(h())))`.
func (c *OauthProviderClient) Use(hooks ...Hook) {
	c.hooks.OauthProvider = append(c.hooks.OauthProvider, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `oauthprovider.Intercept(f(g(h())))`.
func (c *OauthProviderClient) Intercept(interceptors ...Interceptor) {
	c.inters.OauthProvider = append(c.inters.OauthProvider, interceptors...)
}

// Create returns a builder for creating a OauthProvider entity.
func (c *OauthProviderClient) Create() *OauthProviderCreate {
	mutation := newOauthProviderMutation(c.config, OpCreate)
	return &OauthProviderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OauthProvider entities.
func (c *OauthProviderClient) CreateBulk(builders ...*OauthProviderCreate) *OauthProviderCreateBulk {
	return &OauthProviderCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OauthProviderClient) MapCreateBulk(slice any, setFunc func(*OauthProviderCreate, int)) *OauthProviderCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OauthProviderCreateBulk{err: fmt.Errorf("calling to OauthProviderClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OauthProviderCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OauthProviderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OauthProvider.
func (c *OauthProviderClient) Update() *OauthProviderUpdate {
	mutation := newOauthProviderMutation(c.config, OpUpdate)
	return &OauthProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OauthProviderClient) UpdateOne(op *OauthProvider) *OauthProviderUpdateOne {
	mutation := newOauthProviderMutation(c.config, OpUpdateOne, withOauthProvider(op))
	return &OauthProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OauthProviderClient) UpdateOneID(id uint64) *OauthProviderUpdateOne {
	mutation := newOauthProviderMutation(c.config, OpUpdateOne, withOauthProviderID(id))
	return &OauthProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OauthProvider.
func (c *OauthProviderClient) Delete() *OauthProviderDelete {
	mutation := newOauthProviderMutation(c.config, OpDelete)
	return &OauthProviderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OauthProviderClient) DeleteOne(op *OauthProvider) *OauthProviderDeleteOne {
	return c.DeleteOneID(op.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OauthProviderClient) DeleteOneID(id uint64) *OauthProviderDeleteOne {
	builder := c.Delete().Where(oauthprovider.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OauthProviderDeleteOne{builder}
}

// Query returns a query builder for OauthProvider.
func (c *OauthProviderClient) Query() *OauthProviderQuery {
	return &OauthProviderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOauthProvider},
		inters: c.Interceptors(),
	}
}

// Get returns a OauthProvider entity by its id.
func (c *OauthProviderClient) Get(ctx context.Context, id uint64) (*OauthProvider, error) {
	return c.Query().Where(oauthprovider.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OauthProviderClient) GetX(ctx context.Context, id uint64) *OauthProvider {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OauthProviderClient) Hooks() []Hook {
	return c.hooks.OauthProvider
}

// Interceptors returns the client interceptors.
func (c *OauthProviderClient) Interceptors() []Interceptor {
	return c.inters.OauthProvider
}

func (c *OauthProviderClient) mutate(ctx context.Context, m *OauthProviderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OauthProviderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OauthProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OauthProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OauthProviderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown OauthProvider mutation op: %q", m.Op())
	}
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `token.Intercept(f(g(h())))`.
func (c *TokenClient) Intercept(interceptors ...Interceptor) {
	c.inters.Token = append(c.inters.Token, interceptors...)
}

// Create returns a builder for creating a Token entity.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TokenClient) MapCreateBulk(slice any, setFunc func(*TokenCreate, int)) *TokenCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TokenCreateBulk{err: fmt.Errorf("calling to TokenClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TokenCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id uuid.UUID) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TokenClient) DeleteOneID(id uuid.UUID) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeToken},
		inters: c.Interceptors(),
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id uuid.UUID) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id uuid.UUID) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	return c.hooks.Token
}

// Interceptors returns the client interceptors.
func (c *TokenClient) Interceptors() []Interceptor {
	return c.inters.Token
}

func (c *TokenClient) mutate(ctx context.Context, m *TokenMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TokenCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TokenDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Token mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Member, MemberRank, OauthProvider, Token []ent.Hook
	}
	inters struct {
		Member, MemberRank, OauthProvider, Token []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
