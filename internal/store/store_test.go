package store

import (
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestCreateStore(t *testing.T) {
	lengthOfStoreMap := 0
	t.Log("Given the need to test the creation of store")

	t.Logf("\tTest 0:\tWhen creating the store it must be initialised with empty state")
	store := CreateStore()
	if len(store) == lengthOfStoreMap {
		t.Logf("\t%s\tShould receive the lenght of map in store as %d", succeed, lengthOfStoreMap)
	} else {
		t.Errorf("\t%s\tShould receive the lenght of map in store as %d", failed, lengthOfStoreMap)
	}

}
