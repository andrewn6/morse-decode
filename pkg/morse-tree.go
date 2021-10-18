package pkg

import (
	"errors"
	"log"
	"strings"
)
type BST struct {
	Root *Node
}

func NewMorseTree() (*BST, error) {
	encodings := map[string]string{
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
		"P": ".--.",
		"Q": "--.-",
		"R": ".-.",
		"S": "...",
		"T": "-",
		"U": "..-",
		"V": "...-",
		"W": ".--",
		"X": "-..-",
		"Y": "-.--",
		"Z": "--..",
	}
	bst := NewTree()
	for char, morse := range encodings {
			if err := bst.Insert(morse, char); err!= nil {
					return nil, err
			}
	}

	return bst, nil
}

func (bst *BST) checkMorse(morse string) error {
		if len(strings.ReplaceAll(strings.ReplaceAll(morse, ".", ""), "-", "")) > 0 {
				return errors.New("Morse strings can only contain di and dah characters!")
	}

	return ""
}

func (bst *BST) Decode(morse string) string {
		if err := bst.checkMorse(morse); err != nil {
				log.Println(err)
				return ""
		}

		var result string

		switch {
		case morse[0:1] == ".":
			r, err := bst.decode(morse[1:], bst.Root.Dash)
			if err != nil {
					log.Printf("%s (%s), using '_' as placeholder\n", err, morse)
					return "_"
			}
			result = r
		}

		return result
}

func (bst *BST) decode(morse string, currentNode *Node) (string, error) {
		if len(morse) == 0 {
				result := currentNode.Char
				return result, nil
		}

		if morse[0:1] == "." {
				if currentNode.Dot != nil {
					return bst.decode(morse[1:], currentNode.Dot), nil
				}
		}

		if morse[0:1] == "-" {
			dash := currentNode.Dash
			if dash != nil {
				return bst.decode(morse[1:], dash), nil
				}
		}

		return "", errors.New("morse pattern not found")
}

func (bst* BST) Insert(morse string, char string) error {
		if err := bst.checkMorse(morse); err != nil {
				return err
		}

		root := bst.Root

		if morse[0:1] == "." {
				if root.Dot == nil {
						root.Dot = &Node{}
				}

				return bst.insert(morse[1:], char, root.Dot)
		}

		if morse[0:1] == "-" {
				if root.Dash == nil {
						root.Dash = &Node{}
				}

				return bst.insert(morse[1:], char, root.Dash)
		}

		return nil
}

func (bst *BST) insert(morse, char string, currentNode *Node) error {
		if len(morse) == 0 {
			currentNode.Char = char
			return nil
		}

		if morse[0:1] == "." {
				if currentNode.Dot == nil {
						currentNode.Dot = &Node{}
				}

				return bst.insert(morse[1:], char, currentNode.Dot)
		}

		if morse[0:1] == "-" {
				if currentNode.Dash == nil {
					currentNode.Dash = &Node{}
				}

				return bst.insert(morse[1:], char, currentNode.Dash)
		}

		return nil
}
func NewTree() *BST {
		bst := new(BST)
		bst.Root = NewNode("", "")
		return bst
}