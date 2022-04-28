package esquery

// QueryString represents a query of type "query_string"
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-query-string-query.html
type QueryStringQuery struct {
	query string
	defaultField string
}

// Map returns a map representation of the query, thus implementing the
// Mappable interface.
func (q *QueryStringQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"query_string": map[string]interface{}{
			"query": q.query,
			"default_field": q.defaultField,
		},
	}
}

// QueryString sets that the value query_string
func QueryString(query string) *QueryStringQuery {
	return &QueryStringQuery{
		query: query,
	}
}