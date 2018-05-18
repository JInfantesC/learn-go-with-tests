package arrays_slices

func Sum(numbers []int) (summa int){
    for _, num := range numbers {
        summa += num
    }
    return
}
func SumAll(numbersToSum ...[]int) (sums []int) {
    //Make crea un Slice tipo []int, con la logitud del segundo parametro.
    //sums=make([]int,len(numbersToSum))
    for _,numbers :=range(numbersToSum){
        sums=append(sums, Sum(numbers))
    }
    return
}
func SumAllTails(numbersToSum ...[]int) (sums []int) {
    for _,numbers :=range(numbersToSum){
        if len(numbers)==0 {
            sums=append(sums, 0)
        }else{
            //Slices can be sliced. sliceName[low:high] will return a Slices from low to high.
            sums=append(sums, Sum(numbers[1:]))
        }
    }
    return
}
