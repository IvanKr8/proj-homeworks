package main

import (
	"fmt"
	"time"
)

type Character struct {
	Name     string
	Age      int
	BodyType string
	Cowardly bool
}

func chooseCharacter() {
	Characters := []Character{
		{
			Name:     "Ethan",
			Age:      22,
			BodyType: "Athletic",
			Cowardly: false,
		},
		{
			Name:     "Mason",
			Age:      17,
			BodyType: "Thin",
			Cowardly: false,
		},
		{
			Name:     "Caleb",
			Age:      29,
			BodyType: "Fat",
			Cowardly: true,
		},
	}
	fmt.Printf("Okey, now you can choose a character to play. You have 3 characters to choose from:\n")
	for i, Character := range Characters {
		var cowardly string
		if Character.Cowardly == true {
			cowardly = "yes"
		} else {
			cowardly = "no"
		}
		fmt.Printf("%d) %s\nAge: %d\nBody type: %s\nCowardly: %s\n", i+1, Character.Name, Character.Age, Character.BodyType, cowardly)
		fmt.Println()
	}
	var character string
	var selectedCharacter *Character
	fmt.Println("Please, tell me the name of the character")
	character = Input()
	for {
		switch character {
		case "Ethan":
			print()
			fmt.Println("Loading...")
			selectedCharacter = &Characters[0]
			time.Sleep(1 * time.Second)
			start(selectedCharacter)
			return
		case "Mason":
			fmt.Println("Loading...")
			selectedCharacter = &Characters[1]
			time.Sleep(1 * time.Second)
			start(selectedCharacter)
			return
		case "Caleb":
			fmt.Println("Loading...")
			selectedCharacter = &Characters[2]
			time.Sleep(1 * time.Second)
			start(selectedCharacter)
			return
		default:
			fmt.Printf("I dont know this character, try please again 111111")
			character = Input()
		}
	}
}

func endGame() {
	fmt.Println("A day later...")
	time.Sleep(1 * time.Second)
	fmt.Println("--- CITY NEWS ---")
	fmt.Println("A maniac fatally stabbed a woman on the street!")
	fmt.Println()
	fmt.Println("Last night a terrible event took place in a quiet suburb. An unknown maniac attacked a woman when she was returning home. Neighbors heard her cries for help, but arrived at the scene too late.")
	fmt.Println("The victim was fatally stabbed in the heart, police said. The motive for the crime is not yet clear, and investigators are working to clarify all the circumstances.")
	fmt.Println("Local residents are urged to be vigilant and not to walk alone on the streets at night. The police promise to make every effort to find and arrest this dangerous criminal.")
	fmt.Println()
	fmt.Println()
	fmt.Println("Game is over! You lose :^(")
}

func attackManiac(selectedCharacter *Character) {
	if selectedCharacter.BodyType != "Athletic" && selectedCharacter.Cowardly != false {
		fmt.Println()
		fmt.Println("You decide to emerge from hiding and attack the maniac. Your heart pounds harder as you step out of the closet and approach him.")
		fmt.Println("You lunge at him, trying to defend yourself and Sarah. However, the maniac proves to be stronger and adeptly evades your attacks.")
		fmt.Println("Suddenly, you feel a sharp pain as the maniac delivers a deadly blow to you. Your strength leaves you, and you collapse to the ground, consciousness slowly slipping away.")
		time.Sleep(1 * time.Second)
		fmt.Println("You died...")
		time.Sleep(1 * time.Second)
		fmt.Println("Game over. You lose :(")
	} else {
		fmt.Println()
		fmt.Println("You decide to emerge from hiding and attack the maniac. Your heart races as you step out of the closet and approach him.")
		fmt.Println("You act swiftly, using all your strength and skills to defend yourself and Sarah. The fight with the maniac becomes intense, but you refuse to back down.")
		fmt.Println("Ultimately, thanks to your prowess, you manage to intercept his attack and, ultimately, deliver a decisive blow. The maniac falls lifeless before you.")
		fmt.Println("You feel a mix of relief and adrenaline, realizing that you were able to defend yourself and your friend from danger.")
		fmt.Println("You are saved. Almost everybody is alive) You win :D")
	}

}

func prayFunc() {
	fmt.Println("You remain still, praying that the maniac won't discover you. The footsteps draw nearer to the closet, and you feel even more terrified.")
	time.Sleep(1 * time.Second)
	fmt.Println("Suddenly, you hear a police officer's voice from outside: \"Open the door, this is the police!\" Relief washes over you as you realize help has arrived.")
	fmt.Println()
	fmt.Println("You are saved. Everybody is alive! You win :D")
}

func hideInCloset(selectedCharacter *Character) {
	fmt.Println("You decide to hide in the closet, hoping the maniac won't find you. Your heart is pounding like crazy as you hear footsteps in the hallway.")
	fmt.Println("Come out and try to attack, write `attack`")
	fmt.Println("Stay still and pray, type `pray`")
	SwitchHandler("attack", "pray", func() {
		attackManiac(selectedCharacter)
	}, prayFunc)
}

func getHelp() {
	fmt.Println("You quickly attempt to call for help. Fortunately, your neighbor, who works in the police force, happened to be passing by your house and heard your cries for assistance. He promptly arrives at the scene and, upon hearing your story, calls for backup. Soon, the police arrive and arrest the maniac, ensuring your safety.")
	time.Sleep(1 * time.Second)
	fmt.Println()
	fmt.Println("Neighbor: \"Thank goodness I heard your cries! Are you okay?\"")
	fmt.Println("Game is over. You won :D")
}

func EnterHouseSituation(selectedCharacter *Character) {
	fmt.Println("You came to Sarah's house. The maniac ended up in the same house as you.")
	fmt.Println("Try to call for help, write `help`")
	fmt.Println("Attempt to hide, write `hide`")
	SwitchHandler("help", "hide", getHelp, func() {
		hideInCloset(selectedCharacter)
	})
}

func phoneTalking(selectedCharacter *Character) {
	fmt.Printf("Beep...\n")
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, it's me Sarah! Can you come to my place? I'm walking down Justin Street right now and it seems to me that this man behind me is following me...")
	fmt.Println("To go to Sarah, enter `go`.")
	fmt.Println("To stay at home enter `stay`.")
	SwitchHandler("go", "stay", func() {
		EnterHouseSituation(selectedCharacter)
	}, endGame)
}

func start(selectedCharacter *Character) {
	fmt.Println("Game is started! Good luck :D")
	fmt.Println()
	fmt.Println("You are sitting at home, and suddenly your friend calls you...")
	fmt.Println("To pick up the phone, write `pick`")
	fmt.Println("To hang up, write `hang`")
	SwitchHandler("pick", "hang", func() {
		phoneTalking(selectedCharacter)
	}, endGame)
}
