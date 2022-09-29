package morm

import (
	"database/sql"
	"reflect"

	"taego/lib/merrors"
	"taego/lib/mlog"
	"taego/lib/mtrace"
)

type Rows interface {
	Close() error

	Scan(any) error
}

type mrows struct {
	rows *sql.Rows
	err  error

	trace mtrace.Trace
}

func (mrows *mrows) Close() error {
	defer mrows.trace.Done()
	if mrows.rows == nil {
		return nil
	}
	return mrows.rows.Close()
}

/*
type user struct{}

Scan(&us)

1.us=[]*user{}
2.us=[]user{}
3.us=user{}
4.us=[]string{}
5.us=""
*/
func (mrows *mrows) Scan(t any) error {
	if mrows.err != nil {
		return mrows.err
	}
	defer mrows.Close()

	reflectValue := reflect.ValueOf(t)
	if reflectValue.Kind() != reflect.Pointer {
		return merrors.Get(merrors.ERROR_MSQL_SCAN_PAREM)
	}
	reflectValue = reflectValue.Elem()

	mlog.Infof("reflectValue: %+v", reflectValue.Interface())
	mlog.Infof("reflectValueKind: %+v", reflectValue.Kind().String())

	var (
		batch, golangType bool
		reflectType       = reflectValue.Type()
	)

	if reflectValue.Kind() == reflect.Slice {
		batch = true
		reflectType = reflectType.Elem()
	}

	switch reflectType.Kind() {
	case reflect.Struct:
		reflectValue = reflect.New(reflectType).Elem()
	case reflect.Pointer:
		reflectValue = reflect.New(reflectType.Elem()).Elem()
	default:
		golangType = true
	}

	mlog.Infof("batch: %v, golangType: %v", batch, golangType)

	if !batch {
		if golangType {
			v := reflect.New(reflect.PointerTo(reflectType))
			v.Elem().Set(reflect.ValueOf(t).Elem().Addr())
			result := v.Elem().Interface()

			for mrows.rows.Next() {
				if err := mrows.rows.Scan(result); err != nil {
					return err
				}
				break
			}
			mlog.Infof("%+v", *(result.(*string)))
			return nil
		}
		return nil
	}

	return nil
}
