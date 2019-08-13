package majunFunc
/*
import(
    "fmt"
)*/

func Contain(slice []int,card int)bool{
    for _,v:=range slice{
        if v==card{
            return true
        }
    }
    return false
}

func Remove(slice *[]int,card int){
    sliceCopy:=*slice
    for i,v := range sliceCopy{
        if v==card{
            sliceCopy=append(sliceCopy[:i],sliceCopy[i+1:]...)
            *slice=sliceCopy
            return
        }
    }
}

func Count(slice []int,card int) float64{
    count:=0.0
    for _,v:= range slice{
        if v==card{
            count+=1
        }
        
    }
    return count
}

func Pop(slice *[]int,index int) (int){
    sliceCopy:=*slice
    x:=sliceCopy[index]
    sliceCopy=append(sliceCopy[:index],sliceCopy[index+1:]...)
    *slice=sliceCopy
    return x
}

func Min(slice []float64) float64{
    min:=slice[0]
    for i:=1;i<len(slice);i++{
        if slice[i]<min{
            min=slice[i]
        }
    }
    return min
}

func PrintOut(slice []int) string{
    output:=""
    for i:=0;i<len(slice);i++{
        if slice[i]/10<4{
            switch slice[i]%10{
            case 1:
                output+="一"
            case 2:
                output+="二"
            case 3:
                output+="三"
            case 4:
                output+="四"
            case 5:
                output+="五"
            case 6:
                output+="六"
            case 7:
                output+="七"
            case 8:
                output+="八"
            case 9:
                output+="九"
            }
        }else{
            output+="  "
        }
    }
    output+="\n"
    for i:=0;i<len(slice);i++{
        if slice[i]/10<4{
            switch slice[i]/10{
            case 1:
                output+="萬"
            case 2:
                output+="筒"
            case 3:
                output+="條"
            }
        }else{
            switch slice[i]%10{
            case 1:
                output+="東"
            case 2:
                output+="南"
            case 3:
                output+="西"
            case 4:
                output+="北"
            case 5:
                output+="中"
            case 6:
                output+="發"
            case 7:
                output+="白"
            }
        }
    }
    output+="\n"
    return output

}