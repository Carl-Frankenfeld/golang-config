# golang-config

Config reads key:value pairs from a text file into a Go map. The format of the config text file is:
- One line per key:value pair
- keys and values seperated by ':' delimeter (can be changed)
- blank lines ignored
- lines prefaced with comment delimeter (defauly '#') ignored
- any text after comment delimeter ignored

Exposed Functions:
- ReadConfig (filename string): opens 'filename.ext', parses text and returns ConfigStruct (map[string]string)
- (c *ConfigStruct) Get (key string) or (c *ConfigStruct) GetStr (key string): returns the string value for key
- (c *ConfigStruct) GetInt (key string): converts and returns value for key to type int (or 0 and err on fail)
- (c *ConfigStruct) GetInt64 (key string): converts and returns value for key to type int64 (or 0 and err on fail)
- (c *ConfigStruct) GetFloat32 (key string): converts and returns value for key to type float32 (or 0 and err on fail)
- (c *ConfigStruct) GetFloat64 (key string): converts and returns value for key to type float64 (or 0 and err on fail)
- (c *ConfigStruct) GetBool (key string): converts and returns value for key to type boolean (or false and err on fail)
    Default Boolean (textfile) values are true, false, on, off, 1, 0, yes, no (case insensitive)

Example:

[config.txt](https://github.com/carlf2107/golang-config/files/9352298/config.txt)

[config_eg.txt](https://github.com/carlf2107/golang-config/files/9352314/config_eg.txt)

