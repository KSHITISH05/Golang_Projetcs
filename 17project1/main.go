package main
import(
	"fmt"
)
func getName() string  {
	name:=""
	fmt.Println("Welcome to Tims casino...")
	fmt.Printf("Enter your name:")

	//you want to collect in put
	_,err:=fmt.Scanln(&name)
	if err!=nil {
		return ""
	}
    fmt.Printf("Welcome %s,lets play \n",name)
	return name
}

func getBet(balance uint)uint  {
	var bet uint
	//in go we do not have while loop syntex insted we have different kind of for loop
	for true{
		fmt.Printf("Enter your Bet,or 0 to quit (balance =$%d):",balance)
		fmt.Scan(&bet)
		if bet>balance {
			fmt.Println("Bet can not be larger than balance")
		}else{
			break
		}
	}
	
	return bet
}

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArr := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArr = append(symbolArr, symbol)
		}
	}
	return symbolArr
}

func getRandomNumber(min int,max int)int  {
	randomNumber:=rand.Intn(n:max-min+1)+min
	return randomNumber;
}

// func getRandomNumber(min int, max int) int {
// 	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
// 	return rand.Intn(max-min+1) + min
// }

func getSpin(reel string,rows int,cols int)[][]string  {
	result:=[][]string{}
	for i := 0; i < rows; i++ {
		result=append(result, []string{})
	}
	for col:=0;col<cols;col++{
		selected:=map[int]bool{}
		for row := 0; row < rows; row++ {
		    for true{
				randomIndex:=getRandomNumber(0,len(reel)-1)
				_,exist:=selected[randomIndex]
				if !exist{
					selected[randomIndex]=true
					result[row]=append(result[row], reel[randomIndex])
					break;
				}
			}
		}
	}
	return result
}


func main()  {
   
    symbols:=map[string]uint{
		"A":4,//most rare
		"B":7,
		"C":12,
		"D":20,//most frequent
	}
	multipliers:=map[string]uint{
		"A":20,
		"B":10,
		"C":5,
		"D":2,
	}
	symbolArr:=generateSymbolArray(symbols)
    spin:= getSpin(symbolArr,3,3)
	symbolArr:=generateSymbolArray(symbols)
	fmt.Println(symbolArr)


	var balance uint =200;
	
	getName()
	fmt.Println("hello")
	for balance>0{
		bet:=getBet(uint(balance))
		if bet==0{
			break
		}
		balance-=bet
	}
	fmt.Println("You are left with balance of ,$%d\n",balance)
}