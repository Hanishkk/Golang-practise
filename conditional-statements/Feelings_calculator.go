

//write a feelings calculator
//happy, sad, neutral, depressed, lonely

//write a feelings calculator with extended version
//mood

//write a feelings calculator with taste
//ex-sweet,tasteless, normal, sour, fasting

//write a feelings calculator with the medicine
//ex-viagra, sleeping pill, paracetamol, coccaine, marijuana


package conditional

import "fmt"

func Mood() {
	Mood := "Happy"

	switch {
	case Mood == "Happy":
		fmt.Println("I am feeling happy")
	case Mood == "Sad":
		fmt.Println("I am feeling Sad")
	case Mood == "Neutral":
		fmt.Println("I am normal")
	case Mood == "Depressed":
		fmt.Println("I am feeling nothing")
	case Mood == "Lonely":
		fmt.Println("I am feeling like left alone")
	}
}


func Taste(x string) {
	switch x {
	case "Happy":
		{
			fmt.Println("Want to have some sweets")
		}

	case "Sad":
		{
			fmt.Println("Tasteless")
		}

	case "Neutral":
		{
			fmt.Println("Normal taste as usual")
		}

	case "Depressed":
		{
			fmt.Println("Sour")
		}

	case "Lonely":
		{
			fmt.Println("Not Hungry")
		}

	default:
		{
			fmt.Println("Dead person")
		}
	}
}


func Medicine(x string) {
	switch x {
	case "Happy":
		{
			fmt.Println("Viagra needed")
		}
	case "Sad":
		{
			fmt.Println("Sleeping Pills needed")
		}
	case "Neutral":
		{
			fmt.Println("Paracetamol needed")
		}

	case "Depressed":
		{
			fmt.Println("Coccaine needed")
		}

	case "Lonely":
		{
			fmt.Println("Marijuana needed")
		}

	default:
		{
			fmt.Println("Person died")
		}
	}
}


func Extend(yourFeeling string) string{
	const value= 100
	switch yourFeeling{
	case "Happy":
		{
			result:= value * 2
			Finalresult:= fmt.Sprintf("My actual feeling is>>%d", result)
			return Finalresult
		}

	case "sad":
		{
			result:= value / 2
			Finalresult:= fmt.Sprintf("My actual feeling is>>%d", result)
			return Finalresult
		}

	case "Neutral":
		{
			result:= value % 2
			Finalresult:= fmt.Sprintf("My actual feeling is>>%d", result)
			return Finalresult
		}

	case "Depressed":
		{
			result:= 0
			Finalresult:=fmt.Sprintf("My actual feeling is>>%d", result)
			return Finalresult
		}

	case "Lonely":
		{
			result:= 1
			Finalresult:= fmt.Sprintf("My actual feeling is >>%d", result)
			return Finalresult
		}

	default:
		{
			return "Check your requirement"
		}
	} 
}