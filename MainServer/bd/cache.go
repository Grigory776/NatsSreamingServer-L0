package bd

import (
	"github.com/jmoiron/sqlx"
)

func RestoreCache(con *sqlx.DB)(map[string]Order, error){
	var res = map[string]Order{}
	uids,err := GetAllUid(con)
	if err != nil {
		return res, err
	}
	for _, val := range uids{
		tmp, fl := GetEntry(con,val)
		if fl {
			res[val] = *tmp
		}
	}
	return res, nil
}

