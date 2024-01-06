
func binary_search(array []int64, to_search int64) bool {
	found := false
	low := 0
	high := len(array) - 1
	for low <= high {
	  mid := (low + high) / 2
	  if array[mid] == to_search {
		found = true
		break
	  }
	  if array[mid] > to_search {
		high = mid - 1
	  } else {
		low = mid + 1
	  }
	}
	return found
  }


func Test_binary_search(t *testing.T) {
	array := []int64{1, 3, 4, 5}
	to_search1 := 3
	if !binary_search(array, int64(to_search1)) {
	  t.Errorf("expected true, got false")
	}
  
	to_search2 := 6
	if binary_search(array, int64(to_search2)) {
	  t.Errorf("expected false, got true")
	}
  }




func lower_bound(array []int64, to_search int64) int {
	index := -1
	low := 0
	high := len(array) - 1
	for low <= high {
	  mid := (low + high) / 2
	  if array[mid] <= to_search {
		low = mid + 1
		index = mid
	  } else {
		high = mid - 1
	  }
	}
	return index
  }
  
  // Test
  func Test_lower_bound(t *testing.T) {
  
	array := []int64{1, 3, 4, 5}
	to_search1 := 0
	to_search2 := 2
	to_search3 := 5
	if lower_bound(array, int64(to_search1)) != -1 {
	  t.Errorf("test case failed")
	}
	if lower_bound(array, int64(to_search2)) != 0 {
	  t.Errorf("test case failed")
	}
	if lower_bound(array, int64(to_search3)) != 3 {
	  t.Errorf("test case failed")
	}
  }
  