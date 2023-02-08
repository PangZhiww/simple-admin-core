// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gofrs/uuid"
	"github.com/suyuan32/simple-admin-core/pkg/ent/member"
	"github.com/suyuan32/simple-admin-core/pkg/ent/memberrank"
)

// Member is the model entity for the Member schema.
type Member struct {
	config `json:"-"`
	// ID of the ent.
	// UUID
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// status 1 normal 2 ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Member's login name | 登录名
	Username string `json:"username,omitempty"`
	// Password | 密码
	Password string `json:"password,omitempty"`
	// Nickname | 昵称
	Nickname string `json:"nickname,omitempty"`
	// Member Rank ID | 会员等级ID
	RankID uint64 `json:"rank_id,omitempty"`
	// Mobile number | 手机号
	Mobile string `json:"mobile,omitempty"`
	// Email | 邮箱号
	Email string `json:"email,omitempty"`
	// Avatar | 头像路径
	Avatar string `json:"avatar,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberQuery when eager-loading is set.
	Edges MemberEdges `json:"edges"`
}

// MemberEdges holds the relations/edges for other nodes in the graph.
type MemberEdges struct {
	// Rank holds the value of the rank edge.
	Rank *MemberRank `json:"rank,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RankOrErr returns the Rank value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberEdges) RankOrErr() (*MemberRank, error) {
	if e.loadedTypes[0] {
		if e.Rank == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: memberrank.Label}
		}
		return e.Rank, nil
	}
	return nil, &NotLoadedError{edge: "rank"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Member) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case member.FieldStatus, member.FieldRankID:
			values[i] = new(sql.NullInt64)
		case member.FieldUsername, member.FieldPassword, member.FieldNickname, member.FieldMobile, member.FieldEmail, member.FieldAvatar:
			values[i] = new(sql.NullString)
		case member.FieldCreatedAt, member.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case member.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Member", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Member fields.
func (m *Member) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case member.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case member.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				m.CreatedAt = value.Time
			}
		case member.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				m.UpdatedAt = value.Time
			}
		case member.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = uint8(value.Int64)
			}
		case member.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				m.Username = value.String
			}
		case member.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				m.Password = value.String
			}
		case member.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				m.Nickname = value.String
			}
		case member.FieldRankID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rank_id", values[i])
			} else if value.Valid {
				m.RankID = uint64(value.Int64)
			}
		case member.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				m.Mobile = value.String
			}
		case member.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				m.Email = value.String
			}
		case member.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				m.Avatar = value.String
			}
		}
	}
	return nil
}

// QueryRank queries the "rank" edge of the Member entity.
func (m *Member) QueryRank() *MemberRankQuery {
	return NewMemberClient(m.config).QueryRank(m)
}

// Update returns a builder for updating this Member.
// Note that you need to call Member.Unwrap() before calling this method if this Member
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Member) Update() *MemberUpdateOne {
	return NewMemberClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Member entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Member) Unwrap() *Member {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Member is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Member) String() string {
	var builder strings.Builder
	builder.WriteString("Member(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("created_at=")
	builder.WriteString(m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(m.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", m.Status))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(m.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(m.Password)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(m.Nickname)
	builder.WriteString(", ")
	builder.WriteString("rank_id=")
	builder.WriteString(fmt.Sprintf("%v", m.RankID))
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(m.Mobile)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(m.Email)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(m.Avatar)
	builder.WriteByte(')')
	return builder.String()
}

// Members is a parsable slice of Member.
type Members []*Member

func (m Members) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
