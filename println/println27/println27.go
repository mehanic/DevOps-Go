package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("You enter a dark room with two doors. Do you go through door #1 or door #2?")
    door, _ := reader.ReadString('\n')
    door = strings.TrimSpace(door)

    if door == "1" {
        fmt.Println("There's a giant bear here eating a cheese cake. What do you do?")
        fmt.Println("1. Take the cake.")
        fmt.Println("2. Scream at the bear.")
        fmt.Println("3. Punch the bear in the face.")

        bear, _ := reader.ReadString('\n')
        bear = strings.TrimSpace(bear)

        if bear == "1" {
            fmt.Println("The bear eats your face off. Good job!")
        } else if bear == "2" {
            fmt.Println("The bear eats your legs off. Good job!")
        } else if bear == "3" {
            fmt.Println("You knock the bear out!")
            fmt.Println("When the bear falls over, you see a passage behind him. You go through it.")
            fmt.Println("It's dark in here, and smells kind of funny. What do you do?")
            fmt.Println("1. Light a match.")
            fmt.Println("2. Feel for a lightswitch.")
            fmt.Println("3. Turn on your mining hat.")

            light, _ := reader.ReadString('\n')
            light = strings.TrimSpace(light)

            if light == "1" {
                fmt.Println("Whoops, looks like that smell was a gas leak. KABOOM!")
            } else if light == "2" {
                fmt.Println("As you feel along the wall for a switch, you run into an electric fence. ZAP!")
            } else if light == "3" {
                fmt.Println("Smart! You find the exit and head home.")
            } else {
                fmt.Println("You huddle up on the floor and die.")
            }
        } else {
            fmt.Printf("Well, doing %s is probably better. Bear runs away.\n", bear)
        }
    } else if door == "2" {
        fmt.Println("You stare into the endless abyss at Cthulhu's retina.")
        fmt.Println("1. Blueberries.")
        fmt.Println("2. Yellow jacket clothespins.")
        fmt.Println("3. Understanding revolvers yelling melodies.")

        insanity, _ := reader.ReadString('\n')
        insanity = strings.TrimSpace(insanity)

        if insanity == "1" || insanity == "2" {
            fmt.Println("Your body survives powered by a mind of jello. Good job!")
        } else {
            fmt.Println("The insanity rots your eyes into a pool of muck. Good job!")
        }
    } else {
        fmt.Println("You stumble around and fall on a knife and die. Good job!")
    }
}

// You enter a dark room with two doors. Do you go through door #1 or door #2?
// 1
// There's a giant bear here eating a cheese cake. What do you do?
// 1. Take the cake.
// 2. Scream at the bear.
// 3. Punch the bear in the face.
// 2
// The bear eats your legs off. Good job!


// You enter a dark room with two doors. Do you go through door #1 or door #2?
// 2
// You stare into the endless abyss at Cthulhu's retina.
// 1. Blueberries.
// 2. Yellow jacket clothespins.
// 3. Understanding revolvers yelling melodies.
// 1
// Your body survives powered by a mind of jello. Good job!