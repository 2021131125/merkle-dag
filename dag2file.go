package merkledag

// Hash to file
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	
	data, err := store.Get(hash)
	if err != nil {
		return nil, err
	}

	fileContent, err := ParseTreeData(data, path, hp)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}


