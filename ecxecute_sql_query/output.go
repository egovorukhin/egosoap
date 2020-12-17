package ecxecute_sql_query

type OutputBody struct {
	SqlQueryResponse SqlQueryResponse `xml:"executeSQLQueryResponse"`
}

type SqlQueryResponse struct {
	Return Return `xml:"return"`
	Ns     string `xml:"ns,attr"`
}

type Return struct {
	Rows interface{} `xml:"row"`
}

func SetOutputBody(rows interface{}) *OutputBody {
	return &OutputBody{
		SqlQueryResponse: SqlQueryResponse{
			Return: Return{
				Rows: rows,
			},
		},
	}
}
