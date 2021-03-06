package data

// Code generated by gen.py. DO NOT EDIT.

func (m *nativeSeriesMeta) read(cache *seriesIterCache) iterator {
	switch m.rValue.Type().Elem() {
	case typeFloat64:
		return m.readFloat64(cache)
	case typeInt64:
		return m.readInt64(cache)
	case typeInt:
		return m.readInt(cache)
	case typeString:
		return m.readString(cache)
	case typeBool:
		return m.readBool(cache)
	case typeTimestampMillis:
		return m.readTimestampMillis(cache)
	case typeTimestampMicros:
		return m.readTimestampMicros(cache)
	default:
		return m.fallbackRead(cache)
	}
}

// float64

// typed series builder

type nativeSeriesBuilderFloat64 struct {
	column  Column
	data    []float64
	notNull []bool
}

func newNativeSeriesBuilderFloat64(columnName ColumnName) *nativeSeriesBuilderFloat64 {
	return &nativeSeriesBuilderFloat64{
		column: Column{
			Name: columnName,
			Type: typeFloat64,
		},
		data:    []float64{},
		notNull: []bool{},
	}
}

func (b *nativeSeriesBuilderFloat64) Column() Column {
	return b.column
}

func (b *nativeSeriesBuilderFloat64) Reserve(capacity int) {
	if capacity > cap(b.data) {
		size := len(b.data)

		newData := make([]float64, size, capacity)
		copy(newData, b.data)
		b.data = newData

		newNotNull := make([]bool, size, capacity)
		copy(newNotNull, b.notNull)
		b.notNull = newNotNull
	}
}

func (b *nativeSeriesBuilderFloat64) Size() int {
	return len(b.data)
}

func (b *nativeSeriesBuilderFloat64) Append(value interface{}) {
	if value != nil {
		b.AppendFloat64(value.(float64), true)
	} else {
		b.AppendFloat64(float64(0), false)
	}
}

func (b *nativeSeriesBuilderFloat64) AppendFloat64(value float64, notNull bool) {
	b.data = append(b.data, value)
	b.notNull = append(b.notNull, notNull)
}

func (b *nativeSeriesBuilderFloat64) Build() *Series {
	return mustNewNativeSeriesFromSlice(b.column.Name, b.data, b.notNull)
}

var _ seriesBuilderFloat64 = (*nativeSeriesBuilderFloat64)(nil)

// typed iterator

func (m *nativeSeriesMeta) readFloat64(_ *seriesIterCache) iterator {
	return &nativeSeriesIterFloat64{
		data:    m.rValue.Interface().([]float64),
		notNull: m.notNull,
		pos:     -1,
	}
}

type nativeSeriesIterFloat64 struct {
	data    []float64
	notNull []bool
	pos     int
}

func (i *nativeSeriesIterFloat64) Next() bool {
	i.pos++
	return i.pos < len(i.data)
}

func (i *nativeSeriesIterFloat64) Float64() (float64, bool) {
	return i.data[i.pos], i.notNull[i.pos]
}

func (i *nativeSeriesIterFloat64) Value() interface{} {
	if val, ok := i.Float64(); ok {
		return val
	} else {
		return nil
	}
}

var _ iterFloat64 = (*nativeSeriesIterFloat64)(nil)

// int64

// typed series builder

type nativeSeriesBuilderInt64 struct {
	column  Column
	data    []int64
	notNull []bool
}

func newNativeSeriesBuilderInt64(columnName ColumnName) *nativeSeriesBuilderInt64 {
	return &nativeSeriesBuilderInt64{
		column: Column{
			Name: columnName,
			Type: typeInt64,
		},
		data:    []int64{},
		notNull: []bool{},
	}
}

func (b *nativeSeriesBuilderInt64) Column() Column {
	return b.column
}

func (b *nativeSeriesBuilderInt64) Reserve(capacity int) {
	if capacity > cap(b.data) {
		size := len(b.data)

		newData := make([]int64, size, capacity)
		copy(newData, b.data)
		b.data = newData

		newNotNull := make([]bool, size, capacity)
		copy(newNotNull, b.notNull)
		b.notNull = newNotNull
	}
}

func (b *nativeSeriesBuilderInt64) Size() int {
	return len(b.data)
}

func (b *nativeSeriesBuilderInt64) Append(value interface{}) {
	if value != nil {
		b.AppendInt64(value.(int64), true)
	} else {
		b.AppendInt64(int64(0), false)
	}
}

func (b *nativeSeriesBuilderInt64) AppendInt64(value int64, notNull bool) {
	b.data = append(b.data, value)
	b.notNull = append(b.notNull, notNull)
}

func (b *nativeSeriesBuilderInt64) Build() *Series {
	return mustNewNativeSeriesFromSlice(b.column.Name, b.data, b.notNull)
}

var _ seriesBuilderInt64 = (*nativeSeriesBuilderInt64)(nil)

// typed iterator

func (m *nativeSeriesMeta) readInt64(_ *seriesIterCache) iterator {
	return &nativeSeriesIterInt64{
		data:    m.rValue.Interface().([]int64),
		notNull: m.notNull,
		pos:     -1,
	}
}

type nativeSeriesIterInt64 struct {
	data    []int64
	notNull []bool
	pos     int
}

func (i *nativeSeriesIterInt64) Next() bool {
	i.pos++
	return i.pos < len(i.data)
}

func (i *nativeSeriesIterInt64) Int64() (int64, bool) {
	return i.data[i.pos], i.notNull[i.pos]
}

func (i *nativeSeriesIterInt64) Value() interface{} {
	if val, ok := i.Int64(); ok {
		return val
	} else {
		return nil
	}
}

var _ iterInt64 = (*nativeSeriesIterInt64)(nil)

// int

// typed series builder

type nativeSeriesBuilderInt struct {
	column  Column
	data    []int
	notNull []bool
}

func newNativeSeriesBuilderInt(columnName ColumnName) *nativeSeriesBuilderInt {
	return &nativeSeriesBuilderInt{
		column: Column{
			Name: columnName,
			Type: typeInt,
		},
		data:    []int{},
		notNull: []bool{},
	}
}

func (b *nativeSeriesBuilderInt) Column() Column {
	return b.column
}

func (b *nativeSeriesBuilderInt) Reserve(capacity int) {
	if capacity > cap(b.data) {
		size := len(b.data)

		newData := make([]int, size, capacity)
		copy(newData, b.data)
		b.data = newData

		newNotNull := make([]bool, size, capacity)
		copy(newNotNull, b.notNull)
		b.notNull = newNotNull
	}
}

func (b *nativeSeriesBuilderInt) Size() int {
	return len(b.data)
}

func (b *nativeSeriesBuilderInt) Append(value interface{}) {
	if value != nil {
		b.AppendInt(value.(int), true)
	} else {
		b.AppendInt(0, false)
	}
}

func (b *nativeSeriesBuilderInt) AppendInt(value int, notNull bool) {
	b.data = append(b.data, value)
	b.notNull = append(b.notNull, notNull)
}

func (b *nativeSeriesBuilderInt) Build() *Series {
	return mustNewNativeSeriesFromSlice(b.column.Name, b.data, b.notNull)
}

var _ seriesBuilderInt = (*nativeSeriesBuilderInt)(nil)

// typed iterator

func (m *nativeSeriesMeta) readInt(_ *seriesIterCache) iterator {
	return &nativeSeriesIterInt{
		data:    m.rValue.Interface().([]int),
		notNull: m.notNull,
		pos:     -1,
	}
}

type nativeSeriesIterInt struct {
	data    []int
	notNull []bool
	pos     int
}

func (i *nativeSeriesIterInt) Next() bool {
	i.pos++
	return i.pos < len(i.data)
}

func (i *nativeSeriesIterInt) Int() (int, bool) {
	return i.data[i.pos], i.notNull[i.pos]
}

func (i *nativeSeriesIterInt) Value() interface{} {
	if val, ok := i.Int(); ok {
		return val
	} else {
		return nil
	}
}

var _ iterInt = (*nativeSeriesIterInt)(nil)

// string

// typed series builder

type nativeSeriesBuilderString struct {
	column  Column
	data    []string
	notNull []bool
}

func newNativeSeriesBuilderString(columnName ColumnName) *nativeSeriesBuilderString {
	return &nativeSeriesBuilderString{
		column: Column{
			Name: columnName,
			Type: typeString,
		},
		data:    []string{},
		notNull: []bool{},
	}
}

func (b *nativeSeriesBuilderString) Column() Column {
	return b.column
}

func (b *nativeSeriesBuilderString) Reserve(capacity int) {
	if capacity > cap(b.data) {
		size := len(b.data)

		newData := make([]string, size, capacity)
		copy(newData, b.data)
		b.data = newData

		newNotNull := make([]bool, size, capacity)
		copy(newNotNull, b.notNull)
		b.notNull = newNotNull
	}
}

func (b *nativeSeriesBuilderString) Size() int {
	return len(b.data)
}

func (b *nativeSeriesBuilderString) Append(value interface{}) {
	if value != nil {
		b.AppendString(value.(string), true)
	} else {
		b.AppendString("", false)
	}
}

func (b *nativeSeriesBuilderString) AppendString(value string, notNull bool) {
	b.data = append(b.data, value)
	b.notNull = append(b.notNull, notNull)
}

func (b *nativeSeriesBuilderString) Build() *Series {
	return mustNewNativeSeriesFromSlice(b.column.Name, b.data, b.notNull)
}

var _ seriesBuilderString = (*nativeSeriesBuilderString)(nil)

// typed iterator

func (m *nativeSeriesMeta) readString(_ *seriesIterCache) iterator {
	return &nativeSeriesIterString{
		data:    m.rValue.Interface().([]string),
		notNull: m.notNull,
		pos:     -1,
	}
}

type nativeSeriesIterString struct {
	data    []string
	notNull []bool
	pos     int
}

func (i *nativeSeriesIterString) Next() bool {
	i.pos++
	return i.pos < len(i.data)
}

func (i *nativeSeriesIterString) String() (string, bool) {
	return i.data[i.pos], i.notNull[i.pos]
}

func (i *nativeSeriesIterString) Value() interface{} {
	if val, ok := i.String(); ok {
		return val
	} else {
		return nil
	}
}

var _ iterString = (*nativeSeriesIterString)(nil)

// bool

// typed series builder

type nativeSeriesBuilderBool struct {
	column  Column
	data    []bool
	notNull []bool
}

func newNativeSeriesBuilderBool(columnName ColumnName) *nativeSeriesBuilderBool {
	return &nativeSeriesBuilderBool{
		column: Column{
			Name: columnName,
			Type: typeBool,
		},
		data:    []bool{},
		notNull: []bool{},
	}
}

func (b *nativeSeriesBuilderBool) Column() Column {
	return b.column
}

func (b *nativeSeriesBuilderBool) Reserve(capacity int) {
	if capacity > cap(b.data) {
		size := len(b.data)

		newData := make([]bool, size, capacity)
		copy(newData, b.data)
		b.data = newData

		newNotNull := make([]bool, size, capacity)
		copy(newNotNull, b.notNull)
		b.notNull = newNotNull
	}
}

func (b *nativeSeriesBuilderBool) Size() int {
	return len(b.data)
}

func (b *nativeSeriesBuilderBool) Append(value interface{}) {
	if value != nil {
		b.AppendBool(value.(bool), true)
	} else {
		b.AppendBool(false, false)
	}
}

func (b *nativeSeriesBuilderBool) AppendBool(value bool, notNull bool) {
	b.data = append(b.data, value)
	b.notNull = append(b.notNull, notNull)
}

func (b *nativeSeriesBuilderBool) Build() *Series {
	return mustNewNativeSeriesFromSlice(b.column.Name, b.data, b.notNull)
}

var _ seriesBuilderBool = (*nativeSeriesBuilderBool)(nil)

// typed iterator

func (m *nativeSeriesMeta) readBool(_ *seriesIterCache) iterator {
	return &nativeSeriesIterBool{
		data:    m.rValue.Interface().([]bool),
		notNull: m.notNull,
		pos:     -1,
	}
}

type nativeSeriesIterBool struct {
	data    []bool
	notNull []bool
	pos     int
}

func (i *nativeSeriesIterBool) Next() bool {
	i.pos++
	return i.pos < len(i.data)
}

func (i *nativeSeriesIterBool) Bool() (bool, bool) {
	return i.data[i.pos], i.notNull[i.pos]
}

func (i *nativeSeriesIterBool) Value() interface{} {
	if val, ok := i.Bool(); ok {
		return val
	} else {
		return nil
	}
}

var _ iterBool = (*nativeSeriesIterBool)(nil)

// TimestampMillis

// typed series builder

type nativeSeriesBuilderTimestampMillis struct {
	column  Column
	data    []TimestampMillis
	notNull []bool
}

func newNativeSeriesBuilderTimestampMillis(columnName ColumnName) *nativeSeriesBuilderTimestampMillis {
	return &nativeSeriesBuilderTimestampMillis{
		column: Column{
			Name: columnName,
			Type: typeTimestampMillis,
		},
		data:    []TimestampMillis{},
		notNull: []bool{},
	}
}

func (b *nativeSeriesBuilderTimestampMillis) Column() Column {
	return b.column
}

func (b *nativeSeriesBuilderTimestampMillis) Reserve(capacity int) {
	if capacity > cap(b.data) {
		size := len(b.data)

		newData := make([]TimestampMillis, size, capacity)
		copy(newData, b.data)
		b.data = newData

		newNotNull := make([]bool, size, capacity)
		copy(newNotNull, b.notNull)
		b.notNull = newNotNull
	}
}

func (b *nativeSeriesBuilderTimestampMillis) Size() int {
	return len(b.data)
}

func (b *nativeSeriesBuilderTimestampMillis) Append(value interface{}) {
	if value != nil {
		b.AppendTimestampMillis(value.(TimestampMillis), true)
	} else {
		b.AppendTimestampMillis(TimestampMillis(0), false)
	}
}

func (b *nativeSeriesBuilderTimestampMillis) AppendTimestampMillis(value TimestampMillis, notNull bool) {
	b.data = append(b.data, value)
	b.notNull = append(b.notNull, notNull)
}

func (b *nativeSeriesBuilderTimestampMillis) Build() *Series {
	return mustNewNativeSeriesFromSlice(b.column.Name, b.data, b.notNull)
}

var _ seriesBuilderTimestampMillis = (*nativeSeriesBuilderTimestampMillis)(nil)

// typed iterator

func (m *nativeSeriesMeta) readTimestampMillis(_ *seriesIterCache) iterator {
	return &nativeSeriesIterTimestampMillis{
		data:    m.rValue.Interface().([]TimestampMillis),
		notNull: m.notNull,
		pos:     -1,
	}
}

type nativeSeriesIterTimestampMillis struct {
	data    []TimestampMillis
	notNull []bool
	pos     int
}

func (i *nativeSeriesIterTimestampMillis) Next() bool {
	i.pos++
	return i.pos < len(i.data)
}

func (i *nativeSeriesIterTimestampMillis) TimestampMillis() (TimestampMillis, bool) {
	return i.data[i.pos], i.notNull[i.pos]
}

func (i *nativeSeriesIterTimestampMillis) Value() interface{} {
	if val, ok := i.TimestampMillis(); ok {
		return val
	} else {
		return nil
	}
}

var _ iterTimestampMillis = (*nativeSeriesIterTimestampMillis)(nil)

// TimestampMicros

// typed series builder

type nativeSeriesBuilderTimestampMicros struct {
	column  Column
	data    []TimestampMicros
	notNull []bool
}

func newNativeSeriesBuilderTimestampMicros(columnName ColumnName) *nativeSeriesBuilderTimestampMicros {
	return &nativeSeriesBuilderTimestampMicros{
		column: Column{
			Name: columnName,
			Type: typeTimestampMicros,
		},
		data:    []TimestampMicros{},
		notNull: []bool{},
	}
}

func (b *nativeSeriesBuilderTimestampMicros) Column() Column {
	return b.column
}

func (b *nativeSeriesBuilderTimestampMicros) Reserve(capacity int) {
	if capacity > cap(b.data) {
		size := len(b.data)

		newData := make([]TimestampMicros, size, capacity)
		copy(newData, b.data)
		b.data = newData

		newNotNull := make([]bool, size, capacity)
		copy(newNotNull, b.notNull)
		b.notNull = newNotNull
	}
}

func (b *nativeSeriesBuilderTimestampMicros) Size() int {
	return len(b.data)
}

func (b *nativeSeriesBuilderTimestampMicros) Append(value interface{}) {
	if value != nil {
		b.AppendTimestampMicros(value.(TimestampMicros), true)
	} else {
		b.AppendTimestampMicros(TimestampMicros(0), false)
	}
}

func (b *nativeSeriesBuilderTimestampMicros) AppendTimestampMicros(value TimestampMicros, notNull bool) {
	b.data = append(b.data, value)
	b.notNull = append(b.notNull, notNull)
}

func (b *nativeSeriesBuilderTimestampMicros) Build() *Series {
	return mustNewNativeSeriesFromSlice(b.column.Name, b.data, b.notNull)
}

var _ seriesBuilderTimestampMicros = (*nativeSeriesBuilderTimestampMicros)(nil)

// typed iterator

func (m *nativeSeriesMeta) readTimestampMicros(_ *seriesIterCache) iterator {
	return &nativeSeriesIterTimestampMicros{
		data:    m.rValue.Interface().([]TimestampMicros),
		notNull: m.notNull,
		pos:     -1,
	}
}

type nativeSeriesIterTimestampMicros struct {
	data    []TimestampMicros
	notNull []bool
	pos     int
}

func (i *nativeSeriesIterTimestampMicros) Next() bool {
	i.pos++
	return i.pos < len(i.data)
}

func (i *nativeSeriesIterTimestampMicros) TimestampMicros() (TimestampMicros, bool) {
	return i.data[i.pos], i.notNull[i.pos]
}

func (i *nativeSeriesIterTimestampMicros) Value() interface{} {
	if val, ok := i.TimestampMicros(); ok {
		return val
	} else {
		return nil
	}
}

var _ iterTimestampMicros = (*nativeSeriesIterTimestampMicros)(nil)
