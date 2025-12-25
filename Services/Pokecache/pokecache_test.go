package Pokecache

import (
	"reflect"
	"testing"
	"time"
)

func TestCache_Add(t *testing.T) {
	type args struct {
		key string
		val []byte
	}
	tests := map[string]struct {
		got    args
		expect args
	}{
		"Normal": {
			got:    args{key: "https://example.com", val: []byte("testdata")},
			expect: args{key: "https://example.com", val: []byte("testdata")},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			pokecache := NewCache(10 * time.Second)
			pokecache.Add(tc.got.key, tc.got.val)

			got := pokecache.store[tc.got.key]
			if !reflect.DeepEqual(got.val, tc.expect.val) {
				t.Errorf("Add() = %v, expect %v", got.val, tc.expect.val)
			}
		})
	}
}

func TestCache_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := map[string]struct {
		got    args
		expect args
	}{
		"Normal": {
			got:    args{key: "https://example.com"},
			expect: args{key: "https://example.com"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			pokecache := NewCache(10 * time.Second)
			pokecache.store[tt.expect.key] = cacheEntry{createdAt: time.Now(), val: []byte("testdata")}

			gotByte, gotBool := pokecache.Get(tt.got.key)
			if !reflect.DeepEqual(gotByte, []byte("testdata")) {
				t.Errorf("Get() got = %v, want %v", gotByte, []byte("testdata"))
			}
			if !gotBool {
				t.Errorf("Get() got1 = %v, want %v", gotBool, true)
			}
		})
	}
}
