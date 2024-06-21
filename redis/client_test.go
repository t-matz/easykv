/*
 * This file is part of easyKV.
 * © 2016 The easyKV Authors
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package redis

import (
	"testing"

	"github.com/t-matz/easykv/testutils"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type FilterSuite struct{}

var _ = Suite(&FilterSuite{})

func (s *FilterSuite) TestGetValues(t *C) {
	c, err := New([]string{"localhost:6379"})
	if err != nil {
		t.Error(err)
	}

	c.client.Do("SET", "/premtest/database/url", "www.google.de")
	c.client.Do("SET", "/premtest/database/user", "Boris")
	c.client.Do("SET", "/remtest/database/hosts/0/name", "test1")
	c.client.Do("SET", "/remtest/database/hosts/0/ip", "192.168.0.1")
	c.client.Do("SET", "/remtest/database/hosts/0/size", "60")
	c.client.Do("SET", "/remtest/database/hosts/1/name", "test2")
	c.client.Do("SET", "/remtest/database/hosts/1/ip", "192.168.0.2")
	c.client.Do("SET", "/remtest/database/hosts/1/size", "80")

	testutils.GetValues(t, c)
}

func (s *FilterSuite) TestWatchPrefix(t *C) {
	c, err := New([]string{"localhost:6379"})
	if err != nil {
		t.Error(err)
	}
	testutils.WatchPrefixError(t, c)
}
