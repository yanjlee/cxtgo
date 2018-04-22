package cxtgo

// Params are additional parameters which can be send for specific options per exchange
type Params map[string]interface{}

// MergeParams merges the parameters
func MergeParams(params ...Params) Params {
	mergedParams := Params{}
	for _, value := range params {
		for k, v := range value {
			mergedParams[k] = v
		}
	}
	return mergedParams
}

// GetInt return key in params as an int. Indicating if the actual value was an int.
func (p Params) GetInt(key string) (int, bool) {
	v, ok := p[key].(int)
	return v, ok
}

// GetBool return key in params as an bool. Indicating if the actual value was an bool.
func (p Params) GetBool(key string) (bool, bool) {
	v, ok := p[key].(bool)
	return v, ok
}

// GetInt32 return key in params as an int32. Indicating if the actual value was an int32.
func (p Params) GetInt32(key string) (int32, bool) {
	v, ok := p[key].(int32)
	return v, ok
}

// GetInt64 return key in params as an int64. Indicating if the actual value was an int64.
func (p Params) GetInt64(key string) (int64, bool) {
	v, ok := p[key].(int64)
	return v, ok
}

// GetString return key in params as a string. Indicating if the actual value was a string.
func (p Params) GetString(key string) (string, bool) {
	v, ok := p[key].(string)
	return v, ok
}

// GetFloat32 return key in params as a float32. Indicating if the actual value was a float32.
func (p Params) GetFloat32(key string) (float32, bool) {
	v, ok := p[key].(float32)
	return v, ok
}

// GetFloat64 return key in params as a float64. Indicating if the actual value was a float64.
func (p Params) GetFloat64(key string) (float64, bool) {
	v, ok := p[key].(float64)
	return v, ok
}
