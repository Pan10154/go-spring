/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package conf

import (
	"github.com/go-spring/spring-core/conf/prop"
	"github.com/go-spring/spring-core/conf/toml"
	"github.com/go-spring/spring-core/conf/yaml"
)

func init() {
	NewReader(prop.Read, ".properties")
	NewReader(yaml.Read, ".yaml")
	NewReader(toml.Read, ".toml")
}

var readers = make(map[string]Reader)

// Reader 属性列表解析器，将字节数组解析成 map 结构。
type Reader func(b []byte) (map[string]interface{}, error)

// NewReader 注册属性列表解析器，ext 是解析器支持的(标准)文件扩展名。
func NewReader(r Reader, ext string) {
	readers[ext] = r
}
