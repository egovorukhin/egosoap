package ecxecute_sql_query

const SoapAction = "CUCM:DB ver=11.5 executeSQLQuery"

type InputBody struct {
	SqlQuery SqlQuery `xml:"ns:executeSQLQuery"`
}

type SqlQuery struct {
	Sql      string `xml:"sql"`
	Sequence string `xml:"sequence,attr"`
}

func SetInputBody(sql, sequence string) InputBody {
	return InputBody{
		SqlQuery: SqlQuery{
			Sql:      sql,
			Sequence: sequence,
		},
	}
}
