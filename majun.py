# -*- coding: utf-8 -*-
"""
Created on Tue Jul 16 20:59:12 2019

@author: user
"""
import random
player=[]
cards=[11,12,13,14,15,16,17,18,19,
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
      41,42,43,44,45,46,47,]
onTable=[]

#確認是否有胡牌
def checkListen(playerList):
    check=0
    k=0
    bb=0
    bbIndex=0
    playerList.sort()
    checkList=playerList.copy()
    while(k<len(checkList)):
        j=k
        if (checkList[k] not in (checkList[k+1:]) and check!=1) and ((checkList[k]+1 not in checkList) and (checkList[k]+2 not in checkList) and checkList[k]<40):
            k=k+1
            continue
        while(j<len(checkList) and check!=1):
            if checkList.count(checkList[j])>=2:
                if checkList.count(checkList[j])==3:
                    bbJump=3
                elif checkList.count(checkList[j])==4:
                    bbJump=4
                else:
                    bbJump=2
                bb=checkList[j]
                bbIndex=j
                check=1
                k=0
                for i in range(2):
                    checkList.remove(bb)
                break
            j=j+1
        if len(checkList)==0:
            return 1
        if (checkList.count(checkList[k])>=3):
            aaa=checkList[k]
            for i in range(3):
                checkList.remove(aaa)
            if k>=(len(checkList)-1):
                if check==1 and len(checkList)!=0:
                    
                    check=0
                    checkList=playerList.copy()
                    k=bbIndex+bbJump
            continue
        if checkList[k]<40:
            if (checkList[k]+1 in checkList) and (checkList[k]+2 in checkList):
                checkList.remove(checkList[k]+1)
                checkList.remove(checkList[k]+2)
                checkList.remove(checkList[k])
                if k>=(len(checkList)-1):
                    if check==1 and len(checkList)!=0:
                        
                        check=0
                        checkList=playerList.copy()
                        k=bbIndex+bbJump
                continue       
        if k>=(len(checkList)-1):
            if check==1 and len(checkList)!=0:
                check=0
                checkList=playerList.copy()
                k=bbIndex+bbJump
                continue 
            else :break
        k=k+1
    if len(checkList)==0:
        return 1
    else:
        return 0

#將每張牌代入 若能胡牌則表示有聽這張牌    
def listenWhat(playerList):
    Listen=[]
    pai=[11,12,13,14,15,16,17,18,19,
         21,22,23,24,25,26,27,28,29,
         31,32,33,34,35,36,37,38,39,
         41,42,43,44,45,46,47]
    for i in range(len(pai)):
        playerList.append(pai[i])
        if checkListen(playerList)==1:
            Listen.append(pai[i])
            playerList.remove(pai[i])
        else:
            playerList.remove(pai[i])
    
    return Listen

#計算每張手牌的分數 分數低者優先打出去
def score(playerList):#score可以改善
    scoreList=[]
    dangerList=danger(playerList)
    for i in range(len(playerList)):
        score=10
        if playerList[i]>40:
            if playerList.count(playerList[i])==4:
                score=score+400
            elif playerList.count(playerList[i])==3:
                score=score+300
            elif playerList.count(playerList[i])==2:
                score=score+150
            else:
                score+=5
        else:
            if playerList.count(playerList[i])==4:
                score=score+400
            elif playerList.count(playerList[i])==3:
                score=score+300
            elif playerList.count(playerList[i])==2:
                score=score+150
            if (playerList[i]-1 in playerList) and (playerList[i]+1 in playerList):
                if playerList.count(playerList[i])==2:
                    score+=110
                else:
                    score=score+200
            elif ((playerList[i]-1 in playerList) and (playerList[i]-2 in playerList)) and (playerList[i]%10!=1):
                if playerList[i]-3 in playerList:
                    score+=110
                else:
                    if playerList.count(playerList[i])==2:
                        score+=110
                    else:
                        score=score+200
            elif ((playerList[i]+1 in playerList) and (playerList[i]+2 in playerList)) and (playerList[i]%10!=9):
                if playerList[i]+3 in playerList:
                    score+=110
                else:
                    if playerList.count(playerList[i])==2:
                        score+=110
                    else:
                        score=score+200
            elif (playerList[i]-1 in playerList) or (playerList[i]+1 in playerList):
                if playerList[i]-1 in playerList:
                    score=score+120
                    score=score*(1-(onTable.count(playerList[i]-2)+onTable.count(playerList[i]+1)+playerList.count(playerList[i]-2)+playerList.count(playerList[i]+1))/8)*1.3
                else:
                    score=score+120
                    score=score*(1-(onTable.count(playerList[i]+2)+onTable.count(playerList[i]-1)+playerList.count(playerList[i]+2)+playerList.count(playerList[i]-1))/8)*1.3
            else:
                if (playerList[i]+2 in playerList) and (playerList[i]%10!=9) or ((playerList[i]-2 in playerList) and (playerList[i]%10!=1)):
                    if playerList[i]+2 in playerList:
                        score=score+100
                        score=score*(1-((onTable.count(playerList[i]+1)+playerList.count(playerList[i]+1))/4))*1.3
                    else:
                        score=score+100
                        score=score*(1-((onTable.count(playerList[i]-1)+playerList.count(playerList[i]-1))/4))*1.3
                        

        scoreList.append(score*dangerList[i])
    return scoreList

#計算每張手牌的危險係數 以此調整手牌的分數來判斷是否打出去
def danger(playerList):#危險系數可以改進   依剩餘數量改變系數  
    dangerList=[]
    for i in range(len(playerList)):
        if len(onTable)>=20:
            if playerList[i]>40:
                if playerList[i] in onTable[-1:-13] or onTable.count(playerList[i])==3:
                    dangerList.append(1)
                elif len(onTable)>=30:
                    dangerList.append(1+(0.2*(1-onTable.count(playerList[i])/4)))
                else:
                    dangerList.append(1)
            elif playerList[i] in onTable[-1:-13]:
                dangerList.append(1+(0.1*(1-onTable.count(playerList[i])/4)))
            else:
                dangerList.append(1+(0.3*(1-onTable.count(playerList[i])/4)))
        else:
            dangerList.append(1)
    return dangerList
        
#碰牌或吃牌
def eatCard(playerList,eat,pai):
    
    print("eatCard"+str(pai))
    
    if eat==1:
        for i in range(2):
            playerList.remove(pai)
            onTable.append(pai)
    elif eat==2:
        for i in range(3):
            playerList.remove(pai)
            onTable.append(pai)
    elif eat==3:
        playerList.remove(pai+1)
        playerList.remove(pai+2)
        onTable.append(pai+1)
        onTable.append(pai+2)
    elif eat==4:
        playerList.remove(pai-1)
        playerList.remove(pai-2)
        onTable.append(pai-1)
        onTable.append(pai-2)
    elif eat==5:
        playerList.remove(pai-1)
        playerList.remove(pai+1)
        onTable.append(pai-1)
        onTable.append(pai+1)


#判斷是否要槓牌或碰牌
def checkAAA(playerList):
    
    if (((onTable[-1]-1 not in playerList) and (onTable[-1]+1 not in playerList)) or onTable[-1]>40) and playerList.count(onTable[-1])==3:
        print("槓牌:"+str(onTable[-1]))
        eatCard(playerList,2,onTable[-1])
        takeCard=cards.pop(random.randint(0,len(cards)-1))
        playerList.append(takeCard)
        
        scoreList=score(playerList)
            
        print(playerList)
        print(scoreList)
            
        minScore=[i for i,v in enumerate(scoreList) if v==min(scoreList)]
        onTable.append(playerList.pop(random.choice(minScore)))
        print(playerList)
        return 1
    elif playerList.count(onTable[-1])==2 and (((onTable[-1]-1 not in playerList) and (onTable[-1]+1 not in playerList)) or onTable[-1]>40) or ((onTable[-1]-1 in playerList) and (onTable[-1]+1 in playerList) and playerList.count(onTable[-1])==3):
        print("碰牌:"+str(onTable[-1]))
        eatCard(playerList,1,onTable[-1])
        scoreList=score(playerList)
            
        print(playerList)
        print(scoreList)
            
        minScore=[i for i,v in enumerate(scoreList) if v==min(scoreList)]
        onTable.append(playerList.pop(random.choice(minScore)))
        print(playerList)
        return 1
    else:return 0

#判斷是否有人胡牌    
def checkWin(playerList):
    listen=listenWhat(playerList)
    if onTable[-1] in listen:
        print("聽")
        print(listen)
        print("win")
        print(onTable[-1])
        return 0
    else:
        return 1


#進行打牌動作
def playGame(playerList):
    listen=listenWhat(playerList)
    if len(onTable)==0:
        takeCard=cards.pop(random.randint(0,len(cards)-1))
        playerList.append(takeCard)
        scoreList=score(playerList)
        
        print(playerList)
        print(scoreList)
        
        minScore=[i for i,v in enumerate(scoreList) if v==min(scoreList)]
        onTable.append(playerList.pop(random.choice(minScore)))
        
        return 1

    if len(listen)!=0:
        takeCard=cards.pop(random.randint(0,len(cards)-1))
        playerList.append(takeCard)
        if playerList[-1] in listen:
            print("聽")
            print(listen)
            print("win itself")
            print(takeCard)
            return 0
        else:
            scoreList=score(playerList)
            print(playerList)
            print(scoreList)
            minScore=[i for i,v in enumerate(scoreList) if v==min(scoreList)]
            onTable.append(playerList.pop(random.choice(minScore)))       
            return 1
    else:#吃牌可以改善
        if (((onTable[-1]-1 not in playerList) and (onTable[-1]+1 not in playerList)) or onTable[-1]>40) and playerList.count(onTable[-1])==3:
            eatCard(playerList,2,onTable[-1])
            takeCard=cards.pop(random.randint(0,len(cards)-1))
            playerList.append(takeCard)
        elif playerList.count(onTable[-1])==2 or ((onTable[-1]-1 in playerList) and (onTable[-1]+1 in playerList) and playerList.count(onTable[-1])==3):
            eatCard(playerList,1,onTable[-1])
        elif (onTable[-1]+1 in playerList) and (onTable[-1]+2 in playerList) and (onTable[-1]+3 not in playerList) and onTable[-1]<40:
            if onTable[-1] not in playerList:
                eatCard(playerList,3,onTable[-1])
            else:
                takeCard=cards.pop(random.randint(0,len(cards)-1))
                playerList.append(takeCard)
        elif (onTable[-1]-1 in playerList) and (onTable[-1]-2 in playerList) and (onTable[-1]-3 not in playerList) and onTable[-1]<40:
            if onTable[-1] not in playerList:
                eatCard(playerList,4,onTable[-1])
            else:
                takeCard=cards.pop(random.randint(0,len(cards)-1))
                playerList.append(takeCard)
        elif (onTable[-1]-1 in playerList) and (onTable[-1]+1 in playerList) and onTable[-1]<40:
            if onTable[-1] not in playerList:
                eatCard(playerList,5,onTable[-1])
            else:
                takeCard=cards.pop(random.randint(0,len(cards)-1))
                playerList.append(takeCard)
        else:
            takeCard=cards.pop(random.randint(0,len(cards)-1))
            playerList.append(takeCard)
        
        scoreList=score(playerList)
        
        print(playerList)
        print(scoreList)
        
        minScore=[i for i,v in enumerate(scoreList) if v==min(scoreList)]
        onTable.append(playerList.pop(random.choice(minScore)))
        
        return 1
        

for i in range(4):
    player.append([])
                
for i in range(4):
    for j in range(4):
        for k in range(4):
            player[j].append(cards.pop(random.randint(0,len(cards)-1)))
for i in range(4):
    player[i].sort()
    print(player[i])


playFinish=1
i=0
while(len(cards)>8 and playFinish==1):
    if len(onTable)!=0:
        
        j=i
        for j in range(3):
            if checkWin(player[(i+j)%4])==0:
                playFinish=0
                winner=(i+j)%4+1
                break
            
        if checkAAA(player[(i+1)%4]):
        
            print("player"+str((i+1)%4+1)+" AAA")
            k=(i+1+1)%4
            i=k
            continue
        if checkAAA(player[(i+2)%4]):
            print("player"+str((i+2)%4+1)+" AAA")
            k=(i+2+1)%4
            i=k
            continue
        if checkAAA(player[(i+3)%4]):
            print("player"+str((i+3)%4+1)+" AAA")
            k=(i+3+1)%4
            i=k
            continue
    if playFinish==0:      
                break
    print("player"+str(i+1))
    playFinish=playGame(player[i])
    print(player[i])
    if playFinish==0:
        winner=i+1
        break
    if len(cards)<=8:
        break
    
    i+=1
    if i>=4:
        i=0
    
if playFinish==0:
    print("winner: player"+str(winner))
if len(cards)<=8 and playFinish==1:
    print("not finish")
    
