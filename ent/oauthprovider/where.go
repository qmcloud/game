// Code generated by ent, DO NOT EDIT.

package oauthprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/qmcloud/game/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldName, v))
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldClientID, v))
}

// ClientSecret applies equality check predicate on the "client_secret" field. It's identical to ClientSecretEQ.
func ClientSecret(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldClientSecret, v))
}

// RedirectURL applies equality check predicate on the "redirect_url" field. It's identical to RedirectURLEQ.
func RedirectURL(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldRedirectURL, v))
}

// Scopes applies equality check predicate on the "scopes" field. It's identical to ScopesEQ.
func Scopes(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldScopes, v))
}

// AuthURL applies equality check predicate on the "auth_url" field. It's identical to AuthURLEQ.
func AuthURL(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldAuthURL, v))
}

// TokenURL applies equality check predicate on the "token_url" field. It's identical to TokenURLEQ.
func TokenURL(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldTokenURL, v))
}

// AuthStyle applies equality check predicate on the "auth_style" field. It's identical to AuthStyleEQ.
func AuthStyle(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldAuthStyle, v))
}

// InfoURL applies equality check predicate on the "info_url" field. It's identical to InfoURLEQ.
func InfoURL(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldInfoURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContainsFold(FieldName, v))
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldClientID, v))
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldClientID, v))
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldClientID, vs...))
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldClientID, vs...))
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldClientID, v))
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldClientID, v))
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldClientID, v))
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldClientID, v))
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContains(FieldClientID, v))
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasPrefix(FieldClientID, v))
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasSuffix(FieldClientID, v))
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEqualFold(FieldClientID, v))
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContainsFold(FieldClientID, v))
}

// ClientSecretEQ applies the EQ predicate on the "client_secret" field.
func ClientSecretEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldClientSecret, v))
}

// ClientSecretNEQ applies the NEQ predicate on the "client_secret" field.
func ClientSecretNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldClientSecret, v))
}

// ClientSecretIn applies the In predicate on the "client_secret" field.
func ClientSecretIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldClientSecret, vs...))
}

// ClientSecretNotIn applies the NotIn predicate on the "client_secret" field.
func ClientSecretNotIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldClientSecret, vs...))
}

// ClientSecretGT applies the GT predicate on the "client_secret" field.
func ClientSecretGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldClientSecret, v))
}

// ClientSecretGTE applies the GTE predicate on the "client_secret" field.
func ClientSecretGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldClientSecret, v))
}

// ClientSecretLT applies the LT predicate on the "client_secret" field.
func ClientSecretLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldClientSecret, v))
}

// ClientSecretLTE applies the LTE predicate on the "client_secret" field.
func ClientSecretLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldClientSecret, v))
}

// ClientSecretContains applies the Contains predicate on the "client_secret" field.
func ClientSecretContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContains(FieldClientSecret, v))
}

// ClientSecretHasPrefix applies the HasPrefix predicate on the "client_secret" field.
func ClientSecretHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasPrefix(FieldClientSecret, v))
}

// ClientSecretHasSuffix applies the HasSuffix predicate on the "client_secret" field.
func ClientSecretHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasSuffix(FieldClientSecret, v))
}

// ClientSecretEqualFold applies the EqualFold predicate on the "client_secret" field.
func ClientSecretEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEqualFold(FieldClientSecret, v))
}

// ClientSecretContainsFold applies the ContainsFold predicate on the "client_secret" field.
func ClientSecretContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContainsFold(FieldClientSecret, v))
}

// RedirectURLEQ applies the EQ predicate on the "redirect_url" field.
func RedirectURLEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldRedirectURL, v))
}

// RedirectURLNEQ applies the NEQ predicate on the "redirect_url" field.
func RedirectURLNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldRedirectURL, v))
}

// RedirectURLIn applies the In predicate on the "redirect_url" field.
func RedirectURLIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldRedirectURL, vs...))
}

// RedirectURLNotIn applies the NotIn predicate on the "redirect_url" field.
func RedirectURLNotIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldRedirectURL, vs...))
}

// RedirectURLGT applies the GT predicate on the "redirect_url" field.
func RedirectURLGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldRedirectURL, v))
}

// RedirectURLGTE applies the GTE predicate on the "redirect_url" field.
func RedirectURLGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldRedirectURL, v))
}

// RedirectURLLT applies the LT predicate on the "redirect_url" field.
func RedirectURLLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldRedirectURL, v))
}

// RedirectURLLTE applies the LTE predicate on the "redirect_url" field.
func RedirectURLLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldRedirectURL, v))
}

// RedirectURLContains applies the Contains predicate on the "redirect_url" field.
func RedirectURLContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContains(FieldRedirectURL, v))
}

// RedirectURLHasPrefix applies the HasPrefix predicate on the "redirect_url" field.
func RedirectURLHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasPrefix(FieldRedirectURL, v))
}

// RedirectURLHasSuffix applies the HasSuffix predicate on the "redirect_url" field.
func RedirectURLHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasSuffix(FieldRedirectURL, v))
}

// RedirectURLEqualFold applies the EqualFold predicate on the "redirect_url" field.
func RedirectURLEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEqualFold(FieldRedirectURL, v))
}

// RedirectURLContainsFold applies the ContainsFold predicate on the "redirect_url" field.
func RedirectURLContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContainsFold(FieldRedirectURL, v))
}

// ScopesEQ applies the EQ predicate on the "scopes" field.
func ScopesEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldScopes, v))
}

// ScopesNEQ applies the NEQ predicate on the "scopes" field.
func ScopesNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldScopes, v))
}

// ScopesIn applies the In predicate on the "scopes" field.
func ScopesIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldScopes, vs...))
}

// ScopesNotIn applies the NotIn predicate on the "scopes" field.
func ScopesNotIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldScopes, vs...))
}

// ScopesGT applies the GT predicate on the "scopes" field.
func ScopesGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldScopes, v))
}

// ScopesGTE applies the GTE predicate on the "scopes" field.
func ScopesGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldScopes, v))
}

// ScopesLT applies the LT predicate on the "scopes" field.
func ScopesLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldScopes, v))
}

// ScopesLTE applies the LTE predicate on the "scopes" field.
func ScopesLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldScopes, v))
}

// ScopesContains applies the Contains predicate on the "scopes" field.
func ScopesContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContains(FieldScopes, v))
}

// ScopesHasPrefix applies the HasPrefix predicate on the "scopes" field.
func ScopesHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasPrefix(FieldScopes, v))
}

// ScopesHasSuffix applies the HasSuffix predicate on the "scopes" field.
func ScopesHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasSuffix(FieldScopes, v))
}

// ScopesEqualFold applies the EqualFold predicate on the "scopes" field.
func ScopesEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEqualFold(FieldScopes, v))
}

// ScopesContainsFold applies the ContainsFold predicate on the "scopes" field.
func ScopesContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContainsFold(FieldScopes, v))
}

// AuthURLEQ applies the EQ predicate on the "auth_url" field.
func AuthURLEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldAuthURL, v))
}

// AuthURLNEQ applies the NEQ predicate on the "auth_url" field.
func AuthURLNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldAuthURL, v))
}

// AuthURLIn applies the In predicate on the "auth_url" field.
func AuthURLIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldAuthURL, vs...))
}

// AuthURLNotIn applies the NotIn predicate on the "auth_url" field.
func AuthURLNotIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldAuthURL, vs...))
}

// AuthURLGT applies the GT predicate on the "auth_url" field.
func AuthURLGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldAuthURL, v))
}

// AuthURLGTE applies the GTE predicate on the "auth_url" field.
func AuthURLGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldAuthURL, v))
}

// AuthURLLT applies the LT predicate on the "auth_url" field.
func AuthURLLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldAuthURL, v))
}

// AuthURLLTE applies the LTE predicate on the "auth_url" field.
func AuthURLLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldAuthURL, v))
}

// AuthURLContains applies the Contains predicate on the "auth_url" field.
func AuthURLContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContains(FieldAuthURL, v))
}

// AuthURLHasPrefix applies the HasPrefix predicate on the "auth_url" field.
func AuthURLHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasPrefix(FieldAuthURL, v))
}

// AuthURLHasSuffix applies the HasSuffix predicate on the "auth_url" field.
func AuthURLHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasSuffix(FieldAuthURL, v))
}

// AuthURLEqualFold applies the EqualFold predicate on the "auth_url" field.
func AuthURLEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEqualFold(FieldAuthURL, v))
}

// AuthURLContainsFold applies the ContainsFold predicate on the "auth_url" field.
func AuthURLContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContainsFold(FieldAuthURL, v))
}

// TokenURLEQ applies the EQ predicate on the "token_url" field.
func TokenURLEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldTokenURL, v))
}

// TokenURLNEQ applies the NEQ predicate on the "token_url" field.
func TokenURLNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldTokenURL, v))
}

// TokenURLIn applies the In predicate on the "token_url" field.
func TokenURLIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldTokenURL, vs...))
}

// TokenURLNotIn applies the NotIn predicate on the "token_url" field.
func TokenURLNotIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldTokenURL, vs...))
}

// TokenURLGT applies the GT predicate on the "token_url" field.
func TokenURLGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldTokenURL, v))
}

// TokenURLGTE applies the GTE predicate on the "token_url" field.
func TokenURLGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldTokenURL, v))
}

// TokenURLLT applies the LT predicate on the "token_url" field.
func TokenURLLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldTokenURL, v))
}

// TokenURLLTE applies the LTE predicate on the "token_url" field.
func TokenURLLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldTokenURL, v))
}

// TokenURLContains applies the Contains predicate on the "token_url" field.
func TokenURLContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContains(FieldTokenURL, v))
}

// TokenURLHasPrefix applies the HasPrefix predicate on the "token_url" field.
func TokenURLHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasPrefix(FieldTokenURL, v))
}

// TokenURLHasSuffix applies the HasSuffix predicate on the "token_url" field.
func TokenURLHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasSuffix(FieldTokenURL, v))
}

// TokenURLEqualFold applies the EqualFold predicate on the "token_url" field.
func TokenURLEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEqualFold(FieldTokenURL, v))
}

// TokenURLContainsFold applies the ContainsFold predicate on the "token_url" field.
func TokenURLContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContainsFold(FieldTokenURL, v))
}

// AuthStyleEQ applies the EQ predicate on the "auth_style" field.
func AuthStyleEQ(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldAuthStyle, v))
}

// AuthStyleNEQ applies the NEQ predicate on the "auth_style" field.
func AuthStyleNEQ(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldAuthStyle, v))
}

// AuthStyleIn applies the In predicate on the "auth_style" field.
func AuthStyleIn(vs ...uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldAuthStyle, vs...))
}

// AuthStyleNotIn applies the NotIn predicate on the "auth_style" field.
func AuthStyleNotIn(vs ...uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldAuthStyle, vs...))
}

// AuthStyleGT applies the GT predicate on the "auth_style" field.
func AuthStyleGT(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldAuthStyle, v))
}

// AuthStyleGTE applies the GTE predicate on the "auth_style" field.
func AuthStyleGTE(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldAuthStyle, v))
}

// AuthStyleLT applies the LT predicate on the "auth_style" field.
func AuthStyleLT(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldAuthStyle, v))
}

// AuthStyleLTE applies the LTE predicate on the "auth_style" field.
func AuthStyleLTE(v uint64) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldAuthStyle, v))
}

// InfoURLEQ applies the EQ predicate on the "info_url" field.
func InfoURLEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEQ(FieldInfoURL, v))
}

// InfoURLNEQ applies the NEQ predicate on the "info_url" field.
func InfoURLNEQ(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNEQ(FieldInfoURL, v))
}

// InfoURLIn applies the In predicate on the "info_url" field.
func InfoURLIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldIn(FieldInfoURL, vs...))
}

// InfoURLNotIn applies the NotIn predicate on the "info_url" field.
func InfoURLNotIn(vs ...string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldNotIn(FieldInfoURL, vs...))
}

// InfoURLGT applies the GT predicate on the "info_url" field.
func InfoURLGT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGT(FieldInfoURL, v))
}

// InfoURLGTE applies the GTE predicate on the "info_url" field.
func InfoURLGTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldGTE(FieldInfoURL, v))
}

// InfoURLLT applies the LT predicate on the "info_url" field.
func InfoURLLT(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLT(FieldInfoURL, v))
}

// InfoURLLTE applies the LTE predicate on the "info_url" field.
func InfoURLLTE(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldLTE(FieldInfoURL, v))
}

// InfoURLContains applies the Contains predicate on the "info_url" field.
func InfoURLContains(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContains(FieldInfoURL, v))
}

// InfoURLHasPrefix applies the HasPrefix predicate on the "info_url" field.
func InfoURLHasPrefix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasPrefix(FieldInfoURL, v))
}

// InfoURLHasSuffix applies the HasSuffix predicate on the "info_url" field.
func InfoURLHasSuffix(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldHasSuffix(FieldInfoURL, v))
}

// InfoURLEqualFold applies the EqualFold predicate on the "info_url" field.
func InfoURLEqualFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldEqualFold(FieldInfoURL, v))
}

// InfoURLContainsFold applies the ContainsFold predicate on the "info_url" field.
func InfoURLContainsFold(v string) predicate.OauthProvider {
	return predicate.OauthProvider(sql.FieldContainsFold(FieldInfoURL, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OauthProvider) predicate.OauthProvider {
	return predicate.OauthProvider(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OauthProvider) predicate.OauthProvider {
	return predicate.OauthProvider(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OauthProvider) predicate.OauthProvider {
	return predicate.OauthProvider(sql.NotPredicates(p))
}
