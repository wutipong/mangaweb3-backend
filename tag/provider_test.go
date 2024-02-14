package tag

import (
	"context"
	"database/sql"
	"testing"

	dialect_sql "entgo.io/ent/dialect/sql"
	"github.com/stretchr/testify/suite"
	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/enttest"
	_ "modernc.org/sqlite"
)

type ProviderTestSuite struct {
	suite.Suite
}

func TestProviderTestSuite(t *testing.T) {
	suite.Run(t, new(ProviderTestSuite))
}

func (s *ProviderTestSuite) TestReadPage() {
	db, err := sql.Open("sqlite", "file:ent?mode=memory&_fk=1&_pragma=foreign_keys(1)")
	s.Assert().Nil(err)
	s.Assert().NotNil(db)
	defer db.Close()

	client := enttest.NewClient(s.T(), enttest.WithOptions(ent.Driver(dialect_sql.OpenDB("sqlite3", db))))
	defer client.Close()

	Init(client)

	client.Tag.Create().SetName("Tag 1").SetFavorite(false).Save(context.Background())
	client.Tag.Create().SetName("Tag 2").SetFavorite(false).Save(context.Background())
	client.Tag.Create().SetName("Tag 3").SetFavorite(false).Save(context.Background())

	tags, err := ReadPage(context.Background(), false, "", 0, 30)

	s.Assert().Nil(err)
	s.Assert().Equal(3, len(tags))
}

func (s *ProviderTestSuite) TestReadPagePageCount() {
	db, err := sql.Open("sqlite", "file:ent?mode=memory&_fk=1&_pragma=foreign_keys(1)")
	s.Assert().Nil(err)
	s.Assert().NotNil(db)
	defer db.Close()

	client := enttest.NewClient(s.T(), enttest.WithOptions(ent.Driver(dialect_sql.OpenDB("sqlite3", db))))
	defer client.Close()

	Init(client)

	client.Tag.Create().SetName("Tag 1").SetFavorite(false).Save(context.Background())
	client.Tag.Create().SetName("Tag 2").SetFavorite(false).Save(context.Background())
	client.Tag.Create().SetName("Tag 3").SetFavorite(false).Save(context.Background())

	tags, err := ReadPage(context.Background(), false, "", 0, 2)

	s.Assert().Nil(err)
	s.Assert().Equal(2, len(tags))

	tags, err = ReadPage(context.Background(), false, "", 1, 2)
	s.Assert().Nil(err)
	s.Assert().Equal(1, len(tags))
}

func (s *ProviderTestSuite) TestReadPagePageWithSearch() {
	db, err := sql.Open("sqlite", "file:ent?mode=memory&_fk=1&_pragma=foreign_keys(1)")
	s.Assert().Nil(err)
	s.Assert().NotNil(db)
	defer db.Close()

	client := enttest.NewClient(s.T(), enttest.WithOptions(ent.Driver(dialect_sql.OpenDB("sqlite3", db))))
	defer client.Close()

	Init(client)

	client.Tag.Create().SetName("Name 1").SetFavorite(false).Save(context.Background())
	client.Tag.Create().SetName("Name 2").SetFavorite(false).Save(context.Background())
	client.Tag.Create().SetName("Tag 3").SetFavorite(false).Save(context.Background())

	tags, err := ReadPage(context.Background(), false, "name", 0, 30)

	s.Assert().Nil(err)
	s.Assert().Equal(2, len(tags))
	s.Assert().Equal("Name 1", tags[0].Name)
	s.Assert().Equal("Name 2", tags[1].Name)
}

func (s *ProviderTestSuite) TestReadPagePageWithSearchFavoriteOnly() {
	db, err := sql.Open("sqlite", "file:ent?mode=memory&_fk=1&_pragma=foreign_keys(1)")
	s.Assert().Nil(err)
	s.Assert().NotNil(db)
	defer db.Close()

	client := enttest.NewClient(s.T(), enttest.WithOptions(ent.Driver(dialect_sql.OpenDB("sqlite3", db))))
	defer client.Close()

	Init(client)

	client.Tag.Create().SetName("Name 1").SetFavorite(true).Save(context.Background())
	client.Tag.Create().SetName("Name 2").SetFavorite(false).Save(context.Background())
	client.Tag.Create().SetName("Tag 3").SetFavorite(false).Save(context.Background())

	tags, err := ReadPage(context.Background(), true, "name", 0, 30)

	s.Assert().Nil(err)
	s.Assert().Equal(1, len(tags))
	s.Assert().Equal("Name 1", tags[0].Name)
	s.Assert().Equal(true, tags[0].Favorite)
}
