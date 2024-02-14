package meta

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

type QueryTestSuite struct {
	suite.Suite
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}

func (s *QueryTestSuite) TestReadPage() {
	db, err := sql.Open("sqlite", "file:ent?mode=memory&_fk=1&_pragma=foreign_keys(1)")
	s.Assert().Nil(err)
	s.Assert().NotNil(db)
	defer db.Close()

	client := enttest.NewClient(s.T(), enttest.WithOptions(ent.Driver(dialect_sql.OpenDB("sqlite3", db))))
	defer client.Close()

	client.Meta.Create().SetName("[some artist]manga 1 here.zip").Save(context.Background())
	client.Meta.Create().SetName("[some artist]manga 2 here.zip").Save(context.Background())
	client.Meta.Create().SetName("[some artist]manga 3 here.zip").Save(context.Background())
	client.Meta.Create().SetName("[some artist]manga 4 here.zip").Save(context.Background())
	client.Meta.Create().SetName("[some artist]manga 5 here.zip").SetActive(false).Save(context.Background())

	tags, err := ReadPage(context.Background(), client, QueryParams{
		SortBy:      SortFieldName,
		SortOrder:   SortOrderAscending,
		Page:        0,
		ItemPerPage: 30,
	})
	s.Assert().Nil(err)

	s.Assert().Equal(4, len(tags))

	s.Assert().Equal("[some artist]manga 1 here.zip", tags[0].Name)
	s.Assert().Equal("[some artist]manga 2 here.zip", tags[1].Name)
	s.Assert().Equal("[some artist]manga 3 here.zip", tags[2].Name)
	s.Assert().Equal("[some artist]manga 4 here.zip", tags[3].Name)
}

func (s *QueryTestSuite) TestReadPageFavoriteOnly() {
	db, err := sql.Open("sqlite", "file:ent?mode=memory&_fk=1&_pragma=foreign_keys(1)")
	s.Assert().Nil(err)
	s.Assert().NotNil(db)
	defer db.Close()

	client := enttest.NewClient(s.T(), enttest.WithOptions(ent.Driver(dialect_sql.OpenDB("sqlite3", db))))
	defer client.Close()

	client.Meta.Create().SetName("[some artist]manga 1 here.zip").SetFavorite(true).Save(context.Background())
	client.Meta.Create().SetName("[some artist]manga 2 here.zip").SetFavorite(true).Save(context.Background())
	client.Meta.Create().SetName("[some artist]manga 3 here.zip").Save(context.Background())
	client.Meta.Create().SetName("[some artist]manga 4 here.zip").Save(context.Background())

	tags, err := ReadPage(context.Background(), client, QueryParams{
		FavoriteOnly: true,
		SortBy:       SortFieldName,
		SortOrder:    SortOrderAscending,
		Page:         0,
		ItemPerPage:  30,
	})
	s.Assert().Nil(err)

	s.Assert().Equal(2, len(tags))

	s.Assert().Equal("[some artist]manga 1 here.zip", tags[0].Name)
	s.Assert().Equal("[some artist]manga 2 here.zip", tags[1].Name)
}
