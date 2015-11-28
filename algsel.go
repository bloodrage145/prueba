package main

import(
    "fmt"
    "math/rand"
    "time"
    "os"
    "strconv"
)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max - min) + min
}

func division (a[] int, start, end, index int) int {
	pivotval, storeindex := a[index], start
	a[index], a [end] = a[end], a[index]
	for i := start; i < end; i++ {
		if a[i] > pivotval {
			a[storeindex], a[i] = a[i], a[storeindex]
			storeindex++
		}
	}
	a[end], a[storeindex] = a[storeindex], a [end]
	return storeindex
}

func QuickSelect (a[] int, start, end, n int) int {
if start == end {
	return a[start]
}else{
	index := random(start, end)
	pivotIndex := division (a, start, end, index)
	if n  == pivotIndex{
		return a[n]
	}else{
		if n < pivotIndex {
				return QuickSelect(a, start, pivotIndex - 1, n)
			}else{
				return QuickSelect(a, pivotIndex + 1, end, n)
			}
	}
}
}

func MedianofMedians (a[] int, start, end, n int) int {
if start == end {
	return a[start]
}else{
	indexValue := Median(a, len(a))
	var index int
	for i:= start; i <= end;i++ {
		if indexValue == a[i]{
			index = i
		}
	}
	pivotIndex := division (a, start, end, index)
	if n  == pivotIndex{
		return a[n]
	}else{
		if n < pivotIndex {
				return MedianofMedians(a, start, pivotIndex - 1, n)
			}else{
				return MedianofMedians(a, pivotIndex + 1, end, n)
			}
	}
}
}

func insertionSort(a []int, n int) []int {
    for i := 1; i < n; i++ {
        value := a[i]
        j := i - 1
        for j >= 0 && a[j] > value {
            a[j+1] = a[j]
            j = j - 1
        }
        a[j+1] = value
    }
    return a
}

func Partition(arr []int, from, to int) []int {
  var a []int
  for i := from; i <= to;i++{
  	a = append(a, arr[i])
  }
  return a
}

func Median(arr []int, n int) int {
  var medians []int

  if(n == 0) {
  	return arr[0]
  }

  for i := 0; i < n; i += 5 {
    if i + 5 > n && n % 5 != 0 {
      tmp := []int(insertionSort(Partition(arr,i,n-1),n-i))
      ind := n-i
      medians = append(medians,tmp[(ind-1) / 2])
    } else {
      medians = append(medians,insertionSort(Partition(arr,i,i+4),5)[2])
    }
  }

  mofm := QuickSelect(medians,0,len(medians)-1,(len(medians)/2))
  
  return mofm
}

func main() {
	option := 1
    // open output file
    fo, err := os.Create("times.txt")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

	for {
		fmt.Println("1) Para Quick Select, 2) Para Mediana de Medianas, 3) Para salir")
		fmt.Scanln(&option)

		if option == 1 {
			var large, i int = 10000, 10000
			for large <= 1000000{
				if large == 100000 {
					i = 20000
				}
				if large == 200000 {
					i = 50000
				}
				if large == 500000 {
					i = 100000
				}
				var tim_prom float64 = 0
				for iter:= 1; iter <= 30; iter++ {
					var a = make([]int, large)
					for i:= 0; i < large; i++ {
						a[i] = random(0, large)
					}
					k_element := random(0, large - 1)
					start := time.Now()
					q := QuickSelect (a, 0, large - 1, k_element)
					_ = q
					end := time.Duration(time.Now().Sub(start)).Seconds()
					elapsed := float64(end)
					tim_prom = tim_prom + elapsed
					_ = tim_prom
				}
				text := "Para un arreglo de largo" + " " + strconv.Itoa(large) + " QS demoro: " + strconv.FormatFloat(tim_prom/30, 'G', 16, 64) + "\n"
				fmt.Println(text)
				write := []byte(text)
                if _, err := fo.Write(write); err != nil {
                    panic(err)
                }
				large += i
			}
		}
		
		if option == 2 {
			var large, i int = 10, 1
		    for large <= 1000000{
				if large == 100000 {
					i = 20000
				}
				if large == 200000 {
					i = 50000
				}
				if large == 500000 {
					i = 100000
				}
				var tim_prom float64 = 0
				for iter:= 1; iter <= 30; iter++ {
					var a, rep = make([]int, large), make([]bool, large + 1)
					for i:= 0; i < large; i++ {
   						rand := random(0, large)
						if rep[rand] == false {
						    a[i] = rand
						    rep[rand] = true
					    }else{
					    	i--
					    }
					}
					k_element := random(0, large - 1)
					start := time.Now()
					m := MedianofMedians(a, 0, large - 1, k_element)
					_ = m
					end := time.Duration(time.Now().Sub(start)).Seconds()
					elapsed := float64(end)
					tim_prom = tim_prom + elapsed
					_ = tim_prom
				}
				text := "Para un arreglo de largo" + " " + strconv.Itoa(large) + " MdM demoro: " + strconv.FormatFloat(tim_prom/30, 'G', 16, 64) + "\n"
				fmt.Println(text)
				write := []byte(text)
                if _, err := fo.Write(write); err != nil {
                    panic(err)
                }
				large += i
			}
		}
		
		if option == 3 {
			break;
		}
	}
    fo.Close()
}
