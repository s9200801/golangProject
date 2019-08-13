package majunFunc

func Contain(slice []int,card int)bool{
    for _,v:=range slice{
        if v==card{
            return true
        }
    }
    return false
}

func Remove(slice []int,card int) []int{
    for i,v := range slice{
        if v==card{
            slice=append(slice[:i],slice[i+1:]...)
            return slice
        }
    }
    return slice
}

func Count(slice []int,card int) int{
    count:=0
    for _,v:= range slice{
        if v==card{
            count+=1
        }
        
    }
    return count
}

func Pop(slice []int,index int) []int{
    slice=append(slice[:index],slice[index+1:]...)
    return slice
}