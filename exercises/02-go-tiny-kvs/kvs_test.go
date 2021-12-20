package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type entry struct {
	Key   string
	Value int
}

func initKVS(entries []entry) KVS {
	kvs := NewKVS()

	for _, e := range entries {
		kvs.Insert(e.Key, e.Value)
	}

	return kvs
}

func sorted(values []int) []int {
	sort.Ints(values)
	return values
}

func TestCount(t *testing.T) {
	cases := []struct {
		name    string
		entries []entry
		expect  int
	}{
		{
			name:    "empty",
			entries: []entry{},
			expect:  0,
		},
		{
			name: "one",
			entries: []entry{
				{Key: "bee", Value: 1},
			},
			expect: 1,
		},
		{
			name: "keys",
			entries: []entry{
				{Key: "bee", Value: 1},
				{Key: "honey", Value: 2},
				{Key: "wasp", Value: 3}},
			expect: 3,
		},
		{
			name: "duplicate keys",
			entries: []entry{
				{Key: "honey", Value: 1},
				{Key: "honey", Value: 2},
			},
			expect: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			kvs := initKVS(c.entries)
			assert.Equal(t, c.expect, kvs.Count())
		})
	}
}

func TestSearch(t *testing.T) {
	cases := []struct {
		name    string
		entries []entry
		key     string
		expect  []int
	}{
		{
			name:    "empty",
			entries: []entry{},
			key:     "honey",
			expect:  []int{},
		},
		{
			name: "found",
			entries: []entry{
				{Key: "bee", Value: 1},
			},
			key:    "bee",
			expect: []int{1},
		},
		{
			name: "not found",
			entries: []entry{
				{Key: "bee", Value: 1},
			},
			key:    "honey",
			expect: []int{},
		},
		{
			name: "multi values",
			entries: []entry{
				{Key: "honey", Value: 40},
				{Key: "honey", Value: 10},
				{Key: "honey", Value: 20},
				{Key: "bee", Value: 30},
			},
			key:    "honey",
			expect: []int{10, 20, 40},
		},
		{
			name: "common prefix",
			entries: []entry{
				{Key: "honey", Value: 60},
				{Key: "honeybee", Value: 70},
				{Key: "honeybee", Value: 80},
			},
			key:    "honeybee",
			expect: []int{70, 80},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			kvs := initKVS(c.entries)
			values := sorted(kvs.Search(c.key))
			assert.EqualValues(t, c.expect, values, "values")
		})
	}
}

func TestPrefixSearch(t *testing.T) {
	cases := []struct {
		name    string
		entries []entry
		prefix  string
		expect  []int
	}{
		{
			name:    "empty",
			entries: []entry{},
			prefix:  "honey",
			expect:  []int{},
		},
		{
			name: "matched",
			entries: []entry{
				{Key: "bee", Value: 1},
			},
			prefix: "bee",
			expect: []int{1},
		},
		{
			name: "prefix matched",
			entries: []entry{
				{Key: "bee", Value: 1},
			},
			prefix: "be",
			expect: []int{1},
		},
		{
			name: "prefix not matched",
			entries: []entry{
				{Key: "bee", Value: 1},
			},
			prefix: "bo",
			expect: []int{},
		},
		{
			name: "suffix matched",
			entries: []entry{
				{Key: "bee", Value: 1},
			},
			prefix: "ee",
			expect: []int{},
		},
		{
			name: "multiple matched",
			entries: []entry{
				{Key: "honey", Value: 30},
				{Key: "honeycomb", Value: 50},
				{Key: "honeybee", Value: 10},
				{Key: "bee", Value: 90},
			},
			prefix: "honey",
			expect: []int{10, 30, 50},
		},
		{
			name: "empty prefix",
			entries: []entry{
				{Key: "honey", Value: 30},
				{Key: "honeycomb", Value: 50},
				{Key: "honeybee", Value: 10},
				{Key: "bee", Value: 90},
			},
			prefix: "",
			expect: []int{10, 30, 50, 90},
		},
		{
			name: "🐝",
			entries: []entry{
				{Key: "🐝bee", Value: 88},
				{Key: "🐝honeybee", Value: 88},
			},
			prefix: "🐝",
			expect: []int{88, 88},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			kvs := initKVS(c.entries)
			values := sorted(kvs.PrefixSearch(c.prefix))
			assert.EqualValues(t, c.expect, values, "values")
		})
	}
}

func TestMassivelyInsert(t *testing.T) {
	// The number of entries in the data store.
	n := 5000000

	// Test entries
	entries := []entry{
		{Key: "honey", Value: 20},
		{Key: "honey", Value: 10},
		{Key: "honeycomb", Value: 30},
		{Key: "honeybee", Value: 40},
		{Key: "honeymoon", Value: 50},
		{Key: "honesty", Value: 60},
		{Key: "hot fuzz", Value: 70},
	}

	kvs := NewKVS()

	// Insert
	{
		// dummy data
		for i := 0; i < (n - len(entries)); i++ {
			kvs.Insert(fmt.Sprintf("word %d", i), i)
		}

		// test entries
		for _, entry := range entries {
			kvs.Insert(entry.Key, entry.Value)
		}
	}

	// Search
	{
		cases := []struct {
			key    string
			expect []int
		}{
			{
				key:    "honey",
				expect: []int{10, 20},
			},
			{
				key:    "honeycomb",
				expect: []int{30},
			},
			{
				key:    "hone",
				expect: []int{},
			},
		}
		t.Run("Search", func(t *testing.T) {
			for _, c := range cases {
				assert.EqualValues(t, c.expect, sorted(kvs.Search(c.key)), c.key)
			}
		})
	}
	{
		cases := []struct {
			prefix string
			expect []int
		}{
			{
				prefix: "honey",
				expect: []int{10, 20, 30, 40, 50},
			},
			{
				prefix: "hone",
				expect: []int{10, 20, 30, 40, 50, 60},
			},
			{
				prefix: "ho",
				expect: []int{10, 20, 30, 40, 50, 60, 70},
			},
			{
				prefix: "h",
				expect: []int{10, 20, 30, 40, 50, 60, 70},
			},
			{
				prefix: "hot",
				expect: []int{70},
			},
			{
				prefix: "bee",
				expect: []int{},
			},
		}
		t.Run("PrefixSearch", func(t *testing.T) {
			for _, c := range cases {
				assert.EqualValues(t, c.expect, sorted(kvs.PrefixSearch(c.prefix)), c.prefix)
			}
		})
	}
}

// --- Banchmarks

func BenchmarkKVS(b *testing.B) {
	benchmarkKVS(b, 1000)
	benchmarkKVS(b, 5000)
	benchmarkKVS(b, 10000)
	benchmarkKVS(b, 50000)
	benchmarkKVS(b, 100000)
	benchmarkKVS(b, 500000)
	benchmarkKVS(b, 1000000)
	benchmarkKVS(b, 5000000)
}

func benchmarkKVS(b *testing.B, n int) {
	var kvs KVS

	// Test entries
	entries := []entry{
		{Key: "honey", Value: 20},
		{Key: "honey", Value: 10},
		{Key: "honeycomb", Value: 30},
		{Key: "honeybee", Value: 40},
		{Key: "honeymoon", Value: 50},
		{Key: "honesty", Value: 60},
	}

	b.Run(fmt.Sprintf("N = %d", n), func(b *testing.B) {
		b.Run("Insert", func(b *testing.B) {
			kvs = NewKVS()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// dummy data
				for i := 0; i < (n - len(entries)); i++ {
					kvs.Insert(fmt.Sprintf("word %d", i), i)
				}

				// test entries
				for _, entry := range entries {
					kvs.Insert(entry.Key, entry.Value)
				}
			}
		})

		b.Run("Count", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				kvs.Count()
			}
		})

		b.Run("Search", func(b *testing.B) {
			cases := []struct {
				key string
			}{
				{
					key: "honey",
				},
				{
					key: "honeycomb",
				},
				{
					key: "hone",
				},
			}
			for i := 0; i < b.N; i++ {
				for _, c := range cases {
					kvs.Search(c.key)
				}
			}
		})

		b.Run("PrefixSearch", func(b *testing.B) {
			cases := []struct {
				prefix string
			}{
				{
					prefix: "honey",
				},
				{
					prefix: "hone",
				},
				{
					prefix: "ho",
				},
				{
					prefix: "h",
				},
				{
					prefix: "bee",
				},
			}
			for i := 0; i < b.N; i++ {
				for _, c := range cases {
					kvs.PrefixSearch(c.prefix)
				}
			}
		})
	})

}
