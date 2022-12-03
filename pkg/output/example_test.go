package output

func ExampleKeyValues() {
	kvs := []KVPair{
		{Key: "key", Val: "val"},
		{Key: "key2", Val: "val"},
		{Key: "key3", Val: "val"},
	}
	KeyValues("*>> ", kvs)
	// Output:
	//*>> key : val
	//*>> key2: val
	//*>> key3: val
}
