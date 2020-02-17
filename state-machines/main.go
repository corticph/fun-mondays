package main

import (
	"fmt"
)

type SqlQueryBuilder struct {
	query string
}

func (builder SqlQueryBuilder) Select(properties string) SelectQuery {
	return SelectQuery{
		query: "SELECT " + properties,
	}
}

type SelectQuery struct {
	query string
}

func (query SelectQuery) From(table string) FromQuery {
	return FromQuery{
		query: query.query + " FROM " + table,
	}
}

type FromQuery struct {
	query string
}

func (query FromQuery) Where(operator string) WhereQuery {
	return WhereQuery{
		query: query.query + " WHERE " + operator,
	}
}

func (query FromQuery) String() string {
	return query.query + ";"
}

type WhereQuery struct {
	query string
}

func (query WhereQuery) Or(operator string) WhereQuery {
	return WhereQuery{
		query: query.query + " OR " + operator,
	}
}

func (query WhereQuery) And(operator string) WhereQuery {
	return WhereQuery{
		query: query.query + " AND " + operator,
	}
}

func (query WhereQuery) String() string {
	return query.query + ";"
}

func main() {
	query := (&SqlQueryBuilder{}).
		Select("*").
		From("users").
		Where("id = 1").
		And("role = 'ADMIN'").
		Or("email = 'lmj@corti.ai'").String()


	fmt.Println(query)
}