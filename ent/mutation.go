// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echo-clean/ent/article"
	"echo-clean/ent/author"
	"echo-clean/ent/predicate"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeArticle = "Article"
	TypeAuthor  = "Author"
)

// ArticleMutation represents an operation that mutates the Article nodes in the graph.
type ArticleMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	title         *string
	content       *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	author        *int64
	clearedauthor bool
	done          bool
	oldValue      func(context.Context) (*Article, error)
	predicates    []predicate.Article
}

var _ ent.Mutation = (*ArticleMutation)(nil)

// articleOption allows management of the mutation configuration using functional options.
type articleOption func(*ArticleMutation)

// newArticleMutation creates new mutation for the Article entity.
func newArticleMutation(c config, op Op, opts ...articleOption) *ArticleMutation {
	m := &ArticleMutation{
		config:        c,
		op:            op,
		typ:           TypeArticle,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withArticleID sets the ID field of the mutation.
func withArticleID(id int64) articleOption {
	return func(m *ArticleMutation) {
		var (
			err   error
			once  sync.Once
			value *Article
		)
		m.oldValue = func(ctx context.Context) (*Article, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Article.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withArticle sets the old Article of the mutation.
func withArticle(node *Article) articleOption {
	return func(m *ArticleMutation) {
		m.oldValue = func(context.Context) (*Article, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ArticleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ArticleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Article entities.
func (m *ArticleMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ArticleMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ArticleMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Article.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *ArticleMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *ArticleMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *ArticleMutation) ResetTitle() {
	m.title = nil
}

// SetContent sets the "content" field.
func (m *ArticleMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *ArticleMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *ArticleMutation) ResetContent() {
	m.content = nil
}

// SetAuthorID sets the "author_id" field.
func (m *ArticleMutation) SetAuthorID(i int64) {
	m.author = &i
}

// AuthorID returns the value of the "author_id" field in the mutation.
func (m *ArticleMutation) AuthorID() (r int64, exists bool) {
	v := m.author
	if v == nil {
		return
	}
	return *v, true
}

// OldAuthorID returns the old "author_id" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldAuthorID(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAuthorID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAuthorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthorID: %w", err)
	}
	return oldValue.AuthorID, nil
}

// ResetAuthorID resets all changes to the "author_id" field.
func (m *ArticleMutation) ResetAuthorID() {
	m.author = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ArticleMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ArticleMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ArticleMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ArticleMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ArticleMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Article entity.
// If the Article object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ArticleMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ArticleMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// ClearAuthor clears the "author" edge to the Author entity.
func (m *ArticleMutation) ClearAuthor() {
	m.clearedauthor = true
	m.clearedFields[article.FieldAuthorID] = struct{}{}
}

// AuthorCleared reports if the "author" edge to the Author entity was cleared.
func (m *ArticleMutation) AuthorCleared() bool {
	return m.clearedauthor
}

// AuthorIDs returns the "author" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AuthorID instead. It exists only for internal usage by the builders.
func (m *ArticleMutation) AuthorIDs() (ids []int64) {
	if id := m.author; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAuthor resets all changes to the "author" edge.
func (m *ArticleMutation) ResetAuthor() {
	m.author = nil
	m.clearedauthor = false
}

// Where appends a list predicates to the ArticleMutation builder.
func (m *ArticleMutation) Where(ps ...predicate.Article) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ArticleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ArticleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Article, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ArticleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ArticleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Article).
func (m *ArticleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ArticleMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.title != nil {
		fields = append(fields, article.FieldTitle)
	}
	if m.content != nil {
		fields = append(fields, article.FieldContent)
	}
	if m.author != nil {
		fields = append(fields, article.FieldAuthorID)
	}
	if m.created_at != nil {
		fields = append(fields, article.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, article.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ArticleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case article.FieldTitle:
		return m.Title()
	case article.FieldContent:
		return m.Content()
	case article.FieldAuthorID:
		return m.AuthorID()
	case article.FieldCreatedAt:
		return m.CreatedAt()
	case article.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ArticleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case article.FieldTitle:
		return m.OldTitle(ctx)
	case article.FieldContent:
		return m.OldContent(ctx)
	case article.FieldAuthorID:
		return m.OldAuthorID(ctx)
	case article.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case article.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Article field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case article.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case article.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case article.FieldAuthorID:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthorID(v)
		return nil
	case article.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case article.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ArticleMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ArticleMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ArticleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Article numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ArticleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ArticleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ArticleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Article nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ArticleMutation) ResetField(name string) error {
	switch name {
	case article.FieldTitle:
		m.ResetTitle()
		return nil
	case article.FieldContent:
		m.ResetContent()
		return nil
	case article.FieldAuthorID:
		m.ResetAuthorID()
		return nil
	case article.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case article.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Article field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ArticleMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.author != nil {
		edges = append(edges, article.EdgeAuthor)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ArticleMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case article.EdgeAuthor:
		if id := m.author; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ArticleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ArticleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ArticleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedauthor {
		edges = append(edges, article.EdgeAuthor)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ArticleMutation) EdgeCleared(name string) bool {
	switch name {
	case article.EdgeAuthor:
		return m.clearedauthor
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ArticleMutation) ClearEdge(name string) error {
	switch name {
	case article.EdgeAuthor:
		m.ClearAuthor()
		return nil
	}
	return fmt.Errorf("unknown Article unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ArticleMutation) ResetEdge(name string) error {
	switch name {
	case article.EdgeAuthor:
		m.ResetAuthor()
		return nil
	}
	return fmt.Errorf("unknown Article edge %s", name)
}

// AuthorMutation represents an operation that mutates the Author nodes in the graph.
type AuthorMutation struct {
	config
	op             Op
	typ            string
	id             *int64
	name           *string
	created_at     *time.Time
	updated_at     *time.Time
	clearedFields  map[string]struct{}
	article        map[int64]struct{}
	removedarticle map[int64]struct{}
	clearedarticle bool
	done           bool
	oldValue       func(context.Context) (*Author, error)
	predicates     []predicate.Author
}

var _ ent.Mutation = (*AuthorMutation)(nil)

// authorOption allows management of the mutation configuration using functional options.
type authorOption func(*AuthorMutation)

// newAuthorMutation creates new mutation for the Author entity.
func newAuthorMutation(c config, op Op, opts ...authorOption) *AuthorMutation {
	m := &AuthorMutation{
		config:        c,
		op:            op,
		typ:           TypeAuthor,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAuthorID sets the ID field of the mutation.
func withAuthorID(id int64) authorOption {
	return func(m *AuthorMutation) {
		var (
			err   error
			once  sync.Once
			value *Author
		)
		m.oldValue = func(ctx context.Context) (*Author, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Author.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAuthor sets the old Author of the mutation.
func withAuthor(node *Author) authorOption {
	return func(m *AuthorMutation) {
		m.oldValue = func(context.Context) (*Author, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AuthorMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AuthorMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Author entities.
func (m *AuthorMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AuthorMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AuthorMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Author.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *AuthorMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *AuthorMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Author entity.
// If the Author object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AuthorMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *AuthorMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *AuthorMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *AuthorMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Author entity.
// If the Author object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AuthorMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *AuthorMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *AuthorMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *AuthorMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Author entity.
// If the Author object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AuthorMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *AuthorMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddArticleIDs adds the "article" edge to the Article entity by ids.
func (m *AuthorMutation) AddArticleIDs(ids ...int64) {
	if m.article == nil {
		m.article = make(map[int64]struct{})
	}
	for i := range ids {
		m.article[ids[i]] = struct{}{}
	}
}

// ClearArticle clears the "article" edge to the Article entity.
func (m *AuthorMutation) ClearArticle() {
	m.clearedarticle = true
}

// ArticleCleared reports if the "article" edge to the Article entity was cleared.
func (m *AuthorMutation) ArticleCleared() bool {
	return m.clearedarticle
}

// RemoveArticleIDs removes the "article" edge to the Article entity by IDs.
func (m *AuthorMutation) RemoveArticleIDs(ids ...int64) {
	if m.removedarticle == nil {
		m.removedarticle = make(map[int64]struct{})
	}
	for i := range ids {
		delete(m.article, ids[i])
		m.removedarticle[ids[i]] = struct{}{}
	}
}

// RemovedArticle returns the removed IDs of the "article" edge to the Article entity.
func (m *AuthorMutation) RemovedArticleIDs() (ids []int64) {
	for id := range m.removedarticle {
		ids = append(ids, id)
	}
	return
}

// ArticleIDs returns the "article" edge IDs in the mutation.
func (m *AuthorMutation) ArticleIDs() (ids []int64) {
	for id := range m.article {
		ids = append(ids, id)
	}
	return
}

// ResetArticle resets all changes to the "article" edge.
func (m *AuthorMutation) ResetArticle() {
	m.article = nil
	m.clearedarticle = false
	m.removedarticle = nil
}

// Where appends a list predicates to the AuthorMutation builder.
func (m *AuthorMutation) Where(ps ...predicate.Author) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AuthorMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AuthorMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Author, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AuthorMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AuthorMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Author).
func (m *AuthorMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AuthorMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, author.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, author.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, author.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AuthorMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case author.FieldName:
		return m.Name()
	case author.FieldCreatedAt:
		return m.CreatedAt()
	case author.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AuthorMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case author.FieldName:
		return m.OldName(ctx)
	case author.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case author.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Author field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AuthorMutation) SetField(name string, value ent.Value) error {
	switch name {
	case author.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case author.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case author.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Author field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AuthorMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AuthorMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AuthorMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Author numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AuthorMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AuthorMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AuthorMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Author nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AuthorMutation) ResetField(name string) error {
	switch name {
	case author.FieldName:
		m.ResetName()
		return nil
	case author.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case author.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Author field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AuthorMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.article != nil {
		edges = append(edges, author.EdgeArticle)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AuthorMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case author.EdgeArticle:
		ids := make([]ent.Value, 0, len(m.article))
		for id := range m.article {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AuthorMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedarticle != nil {
		edges = append(edges, author.EdgeArticle)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AuthorMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case author.EdgeArticle:
		ids := make([]ent.Value, 0, len(m.removedarticle))
		for id := range m.removedarticle {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AuthorMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedarticle {
		edges = append(edges, author.EdgeArticle)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AuthorMutation) EdgeCleared(name string) bool {
	switch name {
	case author.EdgeArticle:
		return m.clearedarticle
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AuthorMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Author unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AuthorMutation) ResetEdge(name string) error {
	switch name {
	case author.EdgeArticle:
		m.ResetArticle()
		return nil
	}
	return fmt.Errorf("unknown Author edge %s", name)
}