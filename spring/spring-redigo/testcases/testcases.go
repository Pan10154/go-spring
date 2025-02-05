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

package testcases

import (
	"context"
	"testing"

	"github.com/go-spring/spring-core/redis"
	"github.com/go-spring/spring-redigo"
	g "github.com/gomodule/redigo/redis"
)

func RunCase(t *testing.T, fn func(t *testing.T, ctx context.Context, c redis.Client)) {

	conn, err := g.Dial("tcp", ":6379")
	if err != nil {
		t.Fatal(err)
	}

	c := SpringRedigo.NewClient(conn)
	ctx := context.Background()
	defer c.FlushAll(ctx)
	fn(t, ctx, c)
}
