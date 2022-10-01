package msql

import (
	"database/sql"
	"errors"
	"reflect"

	"taego/lib/mtrace"

	"go.uber.org/zap"
)

type mrows struct {
	rows *sql.Rows
	err  error

	num uint32

	trace mtrace.Trace
}

func (mrows *mrows) Close() error {
	if mrows.rows == nil {
		return nil
	}
	mrows.trace.Done(zap.Uint32("queryNum", mrows.num))
	return mrows.rows.Close()
}

func (mrows *mrows) Scan(t any) error {
	if mrows.err != nil {
		return mrows.err
	}
	defer mrows.Close()

	reflectValue := reflect.ValueOf(t)
	if reflectValue.Kind() != reflect.Pointer {
		return errors.New("The sql.Scan function parameter Table in the msql package should be a pointer")
	}
	reflectValue = reflectValue.Elem()

	var (
		batch, golangType, pointer bool
		reflectType                = reflectValue.Type()
	)

	if reflectValue.Kind() == reflect.Slice {
		batch = true
		reflectType = reflectType.Elem()
	}

	switch reflectType.Kind() {
	case reflect.Struct:
	case reflect.Pointer:
		pointer = true
		reflectType = reflectType.Elem()
		if reflectType.Kind() != reflect.Struct {
			return errors.New("parameter type error")
		}
	default:
		golangType = true
	}

	var fieldAddr = func(reflectValue reflect.Value, colname string) (reflect.Value, reflect.Type) {
		num := reflectType.NumField()
		for i := 0; i < num; i++ {
			tagValue := reflectType.Field(i).Tag.Get("db")
			if tagValue == colname {
				field := reflectValue.Field(i)
				return field.Addr(), field.Type()
			}
		}
		return reflect.Zero(reflectType), reflect.TypeOf(reflect.Invalid)
	}

	if !batch {
		if golangType {
			v := reflect.New(reflect.PointerTo(reflectType))
			v.Elem().Set(reflect.ValueOf(t).Elem().Addr())
			result := v.Elem().Interface()
			for mrows.rows.Next() {
				if err := mrows.rows.Scan(result); err != nil {
					return err
				}
				mrows.num++
				break
			}
			return nil
		}

		cols, _ := mrows.rows.ColumnTypes()
		for mrows.rows.Next() {
			result := make([]any, len(cols))
			for i, col := range cols {
				fieldAddr, fieldType := fieldAddr(reflectValue, col.Name())
				if fieldAddr == reflect.Zero(reflectType) {
					return errors.New("unknow column to store")
				}
				v := reflect.New(reflect.PointerTo(fieldType))
				v.Elem().Set(fieldAddr)
				result[i] = v.Elem().Interface()
			}
			if err := mrows.rows.Scan(result...); err != nil {
				return err
			}
			mrows.num++
			break
		}

		return nil
	}

	cols, _ := mrows.rows.ColumnTypes()
	reflectValue.Set(reflect.MakeSlice(reflectValue.Type(), 0, 0))
	for mrows.rows.Next() {
		mrows.num++
		if golangType {

			v := reflect.New(reflectType)
			result := v.Elem().Addr().Interface()
			if err := mrows.rows.Scan(result); err != nil {
				return err
			}

			reflectValue.Set(reflect.Append(reflectValue, v.Elem()))

		} else {
			result := make([]any, len(cols))

			structValue := reflect.New(reflectType)
			for i, col := range cols {
				fieldAddr, fieldType := fieldAddr(structValue.Elem(), col.Name())
				if fieldAddr == reflect.Zero(reflectType) {
					return errors.New("unknow column to store")
				}
				v := reflect.New(reflect.PointerTo(fieldType))
				v.Elem().Set(fieldAddr)
				result[i] = v.Elem().Interface()
			}
			if err := mrows.rows.Scan(result...); err != nil {
				return err
			}

			if pointer {
				reflectValue.Set(reflect.Append(reflectValue, structValue.Elem().Addr()))
			} else {
				reflectValue.Set(reflect.Append(reflectValue, structValue.Elem()))
			}
		}
	}

	return nil
}
