package majun

import(
    "fmt"
    "sort"
    "majunFunc"
    "math/rand"
    "time"
    "net/http"
    //"sync"
    "encoding/json"
    "strconv"
    "strings"
)



var cards=[]int{}
var onTable=[]int{}
var writer http.ResponseWriter
//var wg sync.WaitGroup
var jsonChan chan interface{}
var playerJson map[string]interface{}

//確認是否有胡牌
func checkListen(player *[]int)int{
    playerList:=*player
	check:=0
    bb:=0
    bbIndex:=0
    sort.Ints(playerList)
    checkList:=make([]int,len(playerList))
    var bbJump int
    copy(checkList,playerList)
    for k:=0;k<len(checkList);k++{
        if (!majunFunc.Contain(checkList[k+1:],checkList[k]) && check!=1) && (!majunFunc.Contain(checkList,checkList[k]+1) && !majunFunc.Contain(checkList,checkList[k]+2) && checkList[k]<40){
            k+=1
            continue
        }
    	j:=k
        for (j<len(checkList) && check!=1){
        	if majunFunc.Count(checkList,checkList[j])>=2{
                if majunFunc.Count(checkList,checkList[j])==3{
                    bbJump=3
                }else if majunFunc.Count(checkList,checkList[j])==4{
                    bbJump=4
                }else{
                    bbJump=2
                }
                bb=checkList[j]
                bbIndex=j
                check=1
                k=0
                for i:=0;i<2;i++{
                    majunFunc.Remove(&checkList,bb)
                }
                break
            } 
            j+=1
        }
        if len(checkList)==0 {
            return 1
        }
        if (majunFunc.Count(checkList,checkList[k])>=3){
        	aaa:=checkList[k]
            for i:=0;i<3;i++{
                majunFunc.Remove(&checkList,aaa)
            }
            if k>=(len(checkList)-1){
                if check==1 && len(checkList)!=0{
                	check=0
                    copy(checkList,playerList)
                    k=bbIndex+bbJump
                }
            }
            continue
        }
        if checkList[k]<40{
            if majunFunc.Contain(checkList,checkList[k]+1) && majunFunc.Contain(checkList,checkList[k]+2){
                majunFunc.Remove(&checkList,checkList[k]+2)
                majunFunc.Remove(&checkList,checkList[k]+1)
                majunFunc.Remove(&checkList,checkList[k])
                if k>=(len(checkList)-1){
                    if check==1 && len(checkList)!=0{
                        check=0
                        copy(checkList,playerList)
                        k=bbIndex+bbJump
                    }
                }
                continue
            }
        }
        if k>=(len(checkList)-1){
        	if check==1 && len(checkList)!=0{
        		check=0
                copy(checkList,playerList)
                k=bbIndex+bbJump
                continue
        	}else {
            	break
            }
        }
        //k+=1
    }
    if len(checkList)==0{
        return 1
    }else{
        return 0
    }
}
    

//將每張牌代入 若能胡牌則表示有聽這張牌    
func listenWhat(player *[]int) []int{
    playerList:=*player
    Listen:=[]int{}
    pai:=[]int{11,12,13,14,15,16,17,18,19,
         21,22,23,24,25,26,27,28,29,
         31,32,33,34,35,36,37,38,39,
         41,42,43,44,45,46,47}
    for i:=0;i<len(pai);i++{
        playerList=append(playerList,pai[i])
        if checkListen(&playerList)==1{
            Listen=append(Listen,pai[i])
            majunFunc.Remove(&playerList,pai[i])
        } else{
            majunFunc.Remove(&playerList,pai[i])
        }
    }
    return Listen
}
    

//計算每張手牌的分數 分數低者優先打出去
func score(player *[]int)([]float64){//score可以改善
    playerList:=*player
    scoreList:=make([]float64,0,len(playerList))
    dangerList:=danger(&playerList)
    for i:=0;i<len(playerList);i++{
        score:=10.0
        if playerList[i]>40{
            if majunFunc.Count(playerList,playerList[i])==4{
                score=score+400
            }else if majunFunc.Count(playerList,playerList[i])==3{
                score=score+300
            }else if majunFunc.Count(playerList,playerList[i])==2{
                score=score+150
            }else{
                score-=5
            }
        }else{
            if majunFunc.Count(playerList,playerList[i])==4{
                score+=400
            }else if majunFunc.Count(playerList,playerList[i])==3{
                score+=300
            }else if majunFunc.Count(playerList,playerList[i])==2{
                score+=150
            }
            if majunFunc.Contain(playerList,playerList[i]-1) && majunFunc.Contain(playerList,playerList[i]+1){
                if majunFunc.Count(playerList,playerList[i])==2{
                    score+=110
                }else{
                    score+=200
                }
            }else if majunFunc.Contain(playerList,playerList[i]-1) && majunFunc.Contain(playerList,playerList[i]-2) && (playerList[i]%10!=1){
                if majunFunc.Contain(playerList,playerList[i]-3){
                    score+=120
                }else{
                    if majunFunc.Count(playerList,playerList[i])==2{
                        score+=120
                    }else{
                        score+=200
                    }
                }
                    
            }else if majunFunc.Contain(playerList,playerList[i]+1) && majunFunc.Contain(playerList,playerList[i]+2) && (playerList[i]%10!=9){
                if majunFunc.Contain(playerList,playerList[i]+3){
                    score+=120
                }else{
                    if majunFunc.Count(playerList,playerList[i])==2{
                        score+=120
                    }else{
                        score+=200
                    }
                }
                    
            }else if majunFunc.Contain(playerList,playerList[i]-1) || majunFunc.Contain(playerList,playerList[i]+1){
                if majunFunc.Contain(playerList,playerList[i]-1){
                    score+=120
                    score=score*float64(1-(majunFunc.Count(onTable,playerList[i]-2)+majunFunc.Count(onTable,playerList[i]+1)+majunFunc.Count(playerList,playerList[i]-2)+majunFunc.Count(playerList,playerList[i]+1))/8)*1.3
                }else{
                    score=score+120
                    score=score*float64(1-(majunFunc.Count(onTable,playerList[i]+2)+majunFunc.Count(onTable,playerList[i]-1)+majunFunc.Count(playerList,playerList[i]+2)+majunFunc.Count(playerList,playerList[i]-1))/8)*1.3
                }
                    
            }else{
                if (majunFunc.Contain(playerList,playerList[i]+2) && (playerList[i]%10!=9)) || (majunFunc.Contain(playerList,playerList[i]-2) && (playerList[i]%10!=1)){
                    if majunFunc.Contain(playerList,playerList[i]+2){
                        score=score+100
                        score=score*float64(1-((majunFunc.Count(onTable,playerList[i]+1)+majunFunc.Count(playerList,playerList[i]+1))/4))*1.3
                    }else{
                        score=score+100
                        score=score*float64(1-((majunFunc.Count(onTable,playerList[i]-1)+majunFunc.Count(playerList,playerList[i]-1))/4))*1.3
                    }
                        
                }
                    
            }
        }
        if len(onTable)>40{
            scoreList=append(scoreList,score*dangerList[i])    
        }else{
            scoreList=append(scoreList,score)
        }
        
    }
        
    return scoreList
}

//計算每張手牌的危險係數 以此調整手牌的分數來判斷是否打出去
func danger(player *[]int)([]float64){//危險系數可以改進   依剩餘數量改變系數  
    playerList:=*player
    dangerList:=[]float64{}
    for i:=0;i<len(playerList);i++{
        var dangerScore float64
        if len(onTable)>=20{
            if playerList[i]>40{
                if majunFunc.Contain(onTable[len(onTable)-13:len(onTable)-1],playerList[i])  || majunFunc.Count(onTable,playerList[i])==3{
                    dangerList=append(dangerList,1)
                }else if len(onTable)>=30{
                    dangerScore=1+(0.2*float64(1-majunFunc.Count(onTable,playerList[i])/4))
                    dangerList=append(dangerList,dangerScore)
                }else{
                    dangerList=append(dangerList,1)
                }
            }else if majunFunc.Contain(onTable[len(onTable)-13:len(onTable)-1],playerList[i]){
                dangerList=append(dangerList,1.0+(0.1*float64(1-majunFunc.Count(onTable,playerList[i])/4)))
            }else{
                dangerList=append(dangerList,1+(0.3*float64(1-majunFunc.Count(onTable,playerList[i])/4)))
            }
        }else{
            dangerList=append(dangerList,1)
        }
    }
        
    return dangerList
}        
//碰牌或吃牌
func eatCard(player *[]int ,eat int,pai int){
    fmt.Fprintf(writer,"eatCard:%d\n",pai)
    playerList:=*player
    if eat==1{
        for i:=0;i<2;i++{
            majunFunc.Remove(&playerList,pai)
            onTable=append(onTable,pai)
        }
            
    }else if eat==2{
        for i:=0;i<3;i++{
            majunFunc.Remove(&playerList,pai)
            onTable=append(onTable,pai)
        }
            
    }else if eat==3{
        majunFunc.Remove(&playerList,pai+1)
        majunFunc.Remove(&playerList,pai+2)
        onTable=append(onTable,pai+1)
        onTable=append(onTable,pai+2)
    }else if eat==4{
        majunFunc.Remove(&playerList,pai-1)
        majunFunc.Remove(&playerList,pai-2)
        onTable=append(onTable,pai-1)
        onTable=append(onTable,pai-2)
    }else if eat==5{
        majunFunc.Remove(&playerList,pai-1)
        majunFunc.Remove(&playerList,pai+1)
        onTable=append(onTable,pai-1)
        onTable=append(onTable,pai+1)
    }else if eat==6{
        for i:=0;i<4;i++{
            majunFunc.Remove(&playerList,pai)
            onTable=append(onTable,pai)
        }
    }
    *player=playerList
        
}
    
    


//判斷是否要槓牌或碰牌
func checkAAA(player *[]int) bool{
    playerList:=*player
    minScore:=[]int{}
    if ((!majunFunc.Contain(playerList,onTable[len(onTable)-1]-1) && !majunFunc.Contain(playerList,onTable[len(onTable)-1]+1)) || onTable[len(onTable)-1]>40) && majunFunc.Count(playerList,onTable[len(onTable)-1])==3{
        fmt.Fprintf(writer,"槓牌:%d\n",onTable[len(onTable)-1])
        statusString+=" 槓: "+strconv.Itoa(onTable[len(onTable)-1])
        eatCard(&playerList,2,onTable[len(onTable)-1])
        takeCard:=majunFunc.Pop(&cards,rand.Intn(len(cards)-1))
        statusString+=" 摸: "+strconv.Itoa(takeCard)
        playerList=append(playerList,takeCard)
        
        scoreList:=score(&playerList)
            
        fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(playerList))
        //fmt.Println("%v",scoreList)
        min:=majunFunc.Min(scoreList)
        for i,v :=range scoreList{
            if v==min{
                minScore=append(minScore,i)
            }
        }   
        a:=len(minScore)
        if a>1{
            a=len(minScore)-1
        }   
        onTable=append(onTable,majunFunc.Pop(&playerList,minScore[rand.Intn(a)]))
        statusString+=" 打: "+strconv.Itoa(onTable[len(onTable)-1])+"<br>"
        fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(playerList))
        *player=playerList
        return true
    }else if (majunFunc.Count(playerList,onTable[len(onTable)-1])==2 && ((!majunFunc.Contain(playerList,onTable[len(onTable)-1]-1) && !majunFunc.Contain(playerList,onTable[len(onTable)-1]+1)) || onTable[len(onTable)-1]>40)) || (majunFunc.Contain(playerList,onTable[len(onTable)-1]-1) && (majunFunc.Contain(playerList,onTable[len(onTable)-1]+1)) && majunFunc.Count(playerList,onTable[len(onTable)-1])==3){
        fmt.Fprintf(writer,"碰牌:%d\n",onTable[len(onTable)-1])
        statusString+=" 碰: "+strconv.Itoa(onTable[len(onTable)-1])
        eatCard(&playerList,1,onTable[len(onTable)-1])
        scoreList:=score(&playerList)
            
        fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(playerList))
        //fmt.Println("%v",scoreList)
        min:=majunFunc.Min(scoreList)
        for i,v :=range scoreList{
            if v==min{
                minScore=append(minScore,i)
            }
        }   
        a:=len(minScore)
        if a>1{
            a=len(minScore)-1
        }    
        onTable=append(onTable,majunFunc.Pop(&playerList,minScore[rand.Intn(a)]))
        statusString+=" 打: "+strconv.Itoa(onTable[len(onTable)-1])+"<br>"
        fmt.Fprintf(writer,"%s",majunFunc.PrintOut(playerList))
        *player=playerList
        return true
    }else{
        return false
    }
}
    
    

//判斷是否有人胡牌    
func checkWin(player *[]int)int{
    playerList:=*player
    listen:=listenWhat(&playerList)
    listenString:=[]string{}
    if majunFunc.Contain(listen,onTable[len(onTable)-1]){
        fmt.Fprintf(writer,"聽\n")
        fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(listen))
        fmt.Fprintf(writer,"win\n")
        fmt.Fprintf(writer,"%d\n",onTable[len(onTable)-1])
        for i := range listen {
            number := listen[i]
            text := strconv.Itoa(number)
            listenString = append(listenString, text)
        }
        result := strings.Join(listenString, ",")
        statusString+=" 聽: "+result+"<br>"+" 贏: "+strconv.Itoa(onTable[len(onTable)-1])+"<br>"
        return 0
    }else{
        return 1
    }
}
    


//進行打牌動作
func playGame(player *[]int)int{
    playerList:=*player
    listen:=listenWhat(&playerList)
    listenString:=[]string{}
    minScore:=[]int{}
    var takeCard int
    if len(onTable)==0{    
        takeCard:=majunFunc.Pop(&cards,rand.Intn(len(cards)-1))
        statusString+=" 摸: "+strconv.Itoa(takeCard)
        playerList=append(playerList,takeCard)
        if len(listen)!=0{
            if majunFunc.Contain(listen,playerList[len(playerList)-1]){
                fmt.Fprintf(writer,"聽\n")
                fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(listen))
                fmt.Fprintf(writer,"自摸\n")
                for i := range listen {
                    number := listen[i]
                    text := strconv.Itoa(number)
                    listenString = append(listenString, text)
                }
                result := strings.Join(listenString, ",")
                statusString+=" 聽: "+result+"<br>"+" 自摸: "+strconv.Itoa(takeCard)+"<br>"
                return 1
            }      
        }
        scoreList:=score(&playerList)
        
        fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(playerList))
        //fmt.Println("%v",scoreList)
        min:=majunFunc.Min(scoreList)
        for i,v :=range scoreList{
            if v==min{
                minScore=append(minScore,i)
            }
        }
        a:=len(minScore)
        if a>1{
            a=len(minScore)-1
        }      
        onTable=append(onTable,majunFunc.Pop(&playerList,minScore[rand.Intn(a)]))
        statusString+=" 打: "+strconv.Itoa(onTable[len(onTable)-1])+"<br>"
        *player=playerList
        return 0
    }
    if len(listen)!=0{
        takeCard:=majunFunc.Pop(&cards,rand.Intn(len(cards)-1))
        statusString+=" 摸: "+strconv.Itoa(takeCard)
        playerList=append(playerList,takeCard)
        if majunFunc.Contain(listen,playerList[len(playerList)-1]){
            fmt.Fprintf(writer,"聽\n")
            fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(listen))
            fmt.Fprintf(writer,"自摸\n")
            fmt.Fprintf(writer,"%d\n",takeCard)
            for i := range listen {
                    number := listen[i]
                    text := strconv.Itoa(number)
                    listenString = append(listenString, text)
                }
                result := strings.Join(listenString, ",")
                statusString+=" 聽: "+result+"<br>"+" 自摸: "+strconv.Itoa(takeCard)+"<br>"
            return 1
        }else{
            scoreList:=score(&playerList)
            fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(playerList))
            //fmt.Println("%v",scoreList)
            min:=majunFunc.Min(scoreList)
            for i,v :=range scoreList{
                if v==min{
                    minScore=append(minScore,i)
                }
            }
            a:=len(minScore)
            if a>1{
                a=len(minScore)-1
            }      
            onTable=append(onTable,majunFunc.Pop(&playerList,minScore[rand.Intn(a)]))  
            *player=playerList     
            return 0
        }
            
    }else{//吃牌可以改善
        if majunFunc.Contain(playerList,onTable[len(onTable)-1]+1) && majunFunc.Contain(playerList,onTable[len(onTable)-1]+2) && !majunFunc.Contain(playerList,onTable[len(onTable)-1]+3) && onTable[len(onTable)-1]<40{
            if !majunFunc.Contain(playerList,onTable[len(onTable)-1]){
                statusString+=" 吃: "+strconv.Itoa(onTable[len(onTable)-1])
                eatCard(&playerList,3,onTable[len(onTable)-1])
            }else{
                takeCard=majunFunc.Pop(&cards,rand.Intn(len(cards)-1))
                playerList=append(playerList,takeCard)
                statusString+=" 摸: "+strconv.Itoa(takeCard)

            }
                
        }else if majunFunc.Contain(playerList,onTable[len(onTable)-1]-1) && majunFunc.Contain(playerList,onTable[len(onTable)-1]-2) && !majunFunc.Contain(playerList,onTable[len(onTable)-1]-3) && onTable[len(onTable)-1]<40{
            if !majunFunc.Contain(playerList,onTable[len(onTable)-1]){
                statusString+=" 吃: "+strconv.Itoa(onTable[len(onTable)-1])
                eatCard(&playerList,4,onTable[len(onTable)-1])
            }else{
                takeCard=majunFunc.Pop(&cards,rand.Intn(len(cards)-1))
                playerList=append(playerList,takeCard)
                statusString+=" 摸: "+strconv.Itoa(takeCard)
            }
 
        }else if majunFunc.Contain(playerList,onTable[len(onTable)-1]-1) && majunFunc.Contain(playerList,onTable[len(onTable)-1]+1) && onTable[len(onTable)-1]<40{
            if !majunFunc.Contain(playerList,onTable[len(onTable)-1]){
                statusString+=" 吃: "+strconv.Itoa(onTable[len(onTable)-1])
                eatCard(&playerList,5,onTable[len(onTable)-1])
            }else{
                takeCard=majunFunc.Pop(&cards,rand.Intn(len(cards)-1))
                playerList=append(playerList,takeCard)
                statusString+=" 摸: "+strconv.Itoa(takeCard)
            }
                
        }else{
            takeCard=majunFunc.Pop(&cards,rand.Intn(len(cards)-1))
            statusString+=" 摸: "+strconv.Itoa(takeCard)
            playerList=append(playerList,takeCard)
            for i:=0;i<len(playerList);i++{
                if ((!majunFunc.Contain(playerList,playerList[i]-1) && !majunFunc.Contain(playerList,playerList[i]+1)) || playerList[i]>40) && majunFunc.Count(playerList,playerList[i])==4{
                    fmt.Fprintf(writer,"暗槓\n")
                    statusString+=" 暗槓: "+strconv.Itoa(onTable[len(onTable)-1])
                    eatCard(&playerList,6,takeCard)
                    takeCard=majunFunc.Pop(&cards,rand.Intn(len(cards)-1))
                    playerList=append(playerList,takeCard)
                    statusString+=" 摸: "+strconv.Itoa(takeCard)
                    
                } 
            }
        }
        
           
        
        scoreList:=score(&playerList)
        min:=majunFunc.Min(scoreList)
        fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(playerList))
        //fmt.Println("%v",scoreList)
        
        for i,v :=range scoreList{
            if v==min{
                minScore=append(minScore,i)
            }
        }
        a:=len(minScore)
        if a>1{
            a-=1
        }   
        onTable=append(onTable,majunFunc.Pop(&playerList,minScore[rand.Intn(a)]))
        statusString+=" 打: "+strconv.Itoa(onTable[len(onTable)-1])+"<br>"
        *player=playerList
        return 0
    }
}
           
var statusString string

func Play(w http.ResponseWriter){
    player:=make([][]int,17,17)
    cards=cards[0:0]
    cards=append(cards,11,12,13,14,15,16,17,18,19,
      11,12,13,14,15,16,17,18,19,
      11,12,13,14,15,16,17,18,19,
      11,12,13,14,15,16,17,18,19,
      21,22,23,24,25,26,27,28,29,
      21,22,23,24,25,26,27,28,29,
      21,22,23,24,25,26,27,28,29,
      21,22,23,24,25,26,27,28,29,
      31,32,33,34,35,36,37,38,39,
      31,32,33,34,35,36,37,38,39,
      31,32,33,34,35,36,37,38,39,
      31,32,33,34,35,36,37,38,39,
      41,42,43,44,45,46,47,
      41,42,43,44,45,46,47,
      41,42,43,44,45,46,47,
      41,42,43,44,45,46,47)
    onTable=onTable[0:0]
    writer=w
	rand.Seed(time.Now().Unix())            
	for i :=0 ;i<4;i++{
	    for j:=0;j<4;j++{
	        for k:=0;k<4;k++{
	            player[j]=append(player[j],majunFunc.Pop(&cards,rand.Intn(len(cards)-1)))
	        }
	    }
	}

	playFinish:=0
	i:=0
	var winner int
    playerJson=make(map[string]interface{})
    jsonChan=make(chan interface{})
    playerJson["player1"]=player[0]
    playerJson["player2"]=player[1]
    playerJson["player3"]=player[2]
    playerJson["player4"]=player[3]
    playerJson["status"]="抓牌"
    jsonChan<-playerJson
	for (len(cards)>8 && playFinish==0){
        //wg.Add(1)
        statusString=""
		if len(onTable)!=0{
	        for j:=0;j<3;j++{
	        	if checkWin(&player[(i+j)%4])==0{
	        		playFinish=1
	                winner=(i+j)%4+1
                    statusString+="player "+strconv.Itoa(winner)
	                break
	        	}
            }    
            if checkAAA(&player[(i+1)%4]){
            	fmt.Fprintf(writer,"player%d 碰\n",(i+1)%4+1)
                i=(i+1+1)%4
                playerJson["player1"]=player[0]
                playerJson["player2"]=player[1]
                playerJson["player3"]=player[2]
                playerJson["player4"]=player[3]
                statusString+=" player "+strconv.Itoa(i)
                playerJson["status"]=statusString
                jsonChan<-playerJson
                //wg.Wait()
                continue
            } 
            if checkAAA(&player[(i+2)%4]){
            	fmt.Fprintf(writer,"player%d 碰\n",(i+2)%4+1)
                i=(i+2+1)%4
                playerJson["player1"]=player[0]
                playerJson["player2"]=player[1]
                playerJson["player3"]=player[2]
                playerJson["player4"]=player[3]
                statusString+=" player "+strconv.Itoa(i)
                playerJson["status"]=statusString
                jsonChan<-playerJson
                //wg.Wait()
                continue
            }
            if checkAAA(&player[(i+3)%4]){
            	fmt.Fprintf(writer,"player%d 碰\n",(i+3)%4+1)
                i=(i+3+1)%4
                playerJson["player1"]=player[0]
                playerJson["player2"]=player[1]
                playerJson["player3"]=player[2]
                playerJson["player4"]=player[3]
                statusString+=" player "+strconv.Itoa(i)
                playerJson["status"]=statusString
                jsonChan<-playerJson
                //wg.Wait()
                continue
            }  
	        
		}
	    if playFinish==1{
	    	break
	    }
        statusString+="player "+strconv.Itoa(i+1)     
	    fmt.Fprintf(writer,"player%d\n",(i+1))
	    playFinish=playGame(&player[i])
	    fmt.Fprintf(writer,"%s\n",majunFunc.PrintOut(player[i]))
	    if playFinish==1{
	    	winner=i+1
	        break
	    }   
	    if len(cards)<=8{
	        break
	    }
        //json放這
        
        playerJson["player1"]=player[0]
        playerJson["player2"]=player[1]
        playerJson["player3"]=player[2]
        playerJson["player4"]=player[3]
        playerJson["status"]=statusString
        jsonChan<-playerJson
        
        //wg.Wait()
	    i+=1
	    if i>=4{
	        i=0
	    }
	}
	    
	    
	if playFinish==1{
        playerJson["status"]=statusString
        jsonChan<-playerJson
	    fmt.Fprintf(writer,"winner: player%d",winner)
        return
	}
	if len(cards)<=8 && playFinish==0{
        playerJson["status"]="流局"
        jsonChan<-playerJson
	    fmt.Fprintf(writer,"流局")
        return
	}
}

func OnClick(w http.ResponseWriter){
    w.Header().Set("Content-Type","application/json")
    jsonData,_:=json.Marshal(<-jsonChan)
    w.Write(jsonData)
    //wg.Done()
}
