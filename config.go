package config

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// map of keys and values
// access values through methods
type ConfigStruct struct {
	Values map[string]string
}

// Reads configuration set out in file 'fname' into map [string]string
// returns pointer to new struct instance, error on fail
func ReadConfig(fname string) (*ConfigStruct, error) {
	//create new instance
	c := new(ConfigStruct)
	c.Values = make(map[string]string)

	cfile, err := os.Open(fname)
	if err != nil {
		nerr := errors.New("config.Readconfig - could not open config file " + fname)
		return nil, nerr
	}
	defer cfile.Close()

	if err := c.readFile(bufio.NewReader(cfile)); err != nil {
		nerr := errors.New("config.Readconfig - could not read config file " + fname)
		return nil, nerr
	}
	return c, nil
}

// Populates config map. Takes pointer to bufio created in Setconfig
// Reads file line by line until EOF, store returns in configStruct values map
// returns error on fail
func (c *ConfigStruct) readFile(buf *bufio.Reader) error {
	for {
		l, err := buf.ReadString('\n')
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		default:
			ok, key, val := parseLine(l)
			if ok {
				c.Values[key] = val
			}
		}
	}
}

// Parses line into keys and values
// returns ok if valid line with key, value
func parseLine(l string) (bool, string, string) {
	ok := true
	key := ""
	val := ""
	var del string = ":"     // delimeter between keys and values
	var comment string = "#" // ignore everything to right of comment delimeter
	l = strings.TrimSpace(l) // remove leading and trailing whitespace

	//strip out comments
	if strings.Contains(l, comment) {
		l = l[:strings.Index(l, comment)]
		if len(l) == 0 {
			ok = false
			return false, "", ""
		}
	}

	delimindex := strings.Index(l, del)
	if delimindex < 1 {
		ok = false
		return false, "", ""
	}

	key = strings.TrimSpace(l[:delimindex])
	val = strings.TrimSpace(l[delimindex+1:])

	if key == "" || val == "" {
		ok = false
		return false, "", ""
	}
	return ok, key, val
}

// Get returns the value mapped to key as a string
func (c *ConfigStruct) Get(key string) string {
	return c.Values[key]
}

// GetStr returns the value mapped to key as a string
func (c *ConfigStruct) GetStr(key string) string {
	return c.Values[key]
}

// GetInt converts and returns the value mapped to key as an int
// returns error on fail
func (c *ConfigStruct) GetInt(key string) (int, error) {
	err1 := errors.New("config.GetInt: could not convert value to Int")
	i, err := strconv.Atoi(c.Values[key])
	if err != nil {
		return 0, err1
	}
	return i, nil
}

// GetInt64 converts and returns the value mapped to key as an int64
// returns error on fail
func (c *ConfigStruct) GetInt64(key string) (int64, error) {
	err1 := errors.New("config.GetInt64: could not convert value to Int64")
	i, err := strconv.ParseInt(c.Values[key], 10, 64)
	if err != nil {
		return 0, err1
	}
	return i, nil
}

// GetIntFloat32 converts and returns the value mapped to key as a float32
// returns error on fail
func (c *ConfigStruct) GetFloat32(key string) (float32, error) {
	err1 := errors.New("config.GetFloat32: could not convert value to float32")
	f64, err := strconv.ParseFloat(c.Values[key], 32)
	i := float32(f64)
	if err != nil {
		return 0, err1
	}
	return i, nil
}

// GetIntFloat64 converts and returns the value mapped to key as a float64
// returns error on fail
func (c *ConfigStruct) GetFloat64(key string) (float64, error) {
	err1 := errors.New("config.GetFloat64: could not convert value to float64")
	i, err := strconv.ParseFloat(c.Values[key], 64)
	if err != nil {
		return 0, err1
	}
	return i, nil
}

// GetIntFloat64 converts and returns the value mapped to key as a boolean
// returns error on fail
func (c *ConfigStruct) GetBool(key string) (bool, error) {
	err1 := errors.New("config.GetBool: could not convert value to boolean")
	var r bool = false
	found := false
	boolmap := map[string]bool{
		"true":  true,
		"false": false,
		"on":    true,
		"off":   false,
		"1":     true,
		"0":     false,
		"yes":   true,
		"no":    false,
	}
	s := strings.ToLower(c.Values[key])
	for k, v := range boolmap {
		if s == k {
			r = v
			found = true
		}
	}
	if !found {
		return false, err1
	}
	return r, nil
}
