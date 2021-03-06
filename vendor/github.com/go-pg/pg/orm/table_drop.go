package orm

type DropTableOptions struct {
	IfExists bool
	Cascade  bool
}

func DropTable(db DB, model interface{}, opt *DropTableOptions) error {
	return NewQuery(db, model).DropTable(opt)
}

type dropTableQuery struct {
	q   *Query
	opt *DropTableOptions
}

func (q *dropTableQuery) Copy() *dropTableQuery {
	return &dropTableQuery{
		q:   q.q.Copy(),
		opt: q.opt,
	}
}

func (q *dropTableQuery) Query() *Query {
	return q.q
}

func (q *dropTableQuery) AppendTemplate(b []byte) ([]byte, error) {
	cp := q.Copy()
	cp.q = cp.q.Formatter(dummyFormatter{})
	return cp.AppendQuery(b)
}

func (q *dropTableQuery) AppendQuery(b []byte) ([]byte, error) {
	if q.q.stickyErr != nil {
		return nil, q.q.stickyErr
	}
	if q.q.model == nil {
		return nil, errModelNil
	}

	b = append(b, "DROP TABLE "...)
	if q.opt != nil && q.opt.IfExists {
		b = append(b, "IF EXISTS "...)
	}
	b = q.q.appendFirstTable(b)
	if q.opt != nil && q.opt.Cascade {
		b = append(b, " CASCADE"...)
	}

	return b, q.q.stickyErr
}
