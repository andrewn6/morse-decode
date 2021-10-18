package pkg

import (
	"errors"
	"log"
	"strings"
)
type TREE struct {
	Root *Node
}

func NewMorseTree() (*TREE, error) {
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
	tree := NewTree()
	for char, morse := range encodings {
			if err := tree.Insert(morse, char); err!= nil {
					return nil, err
			}
	}

	return tree, nil
}

func (tree *TREE) checkMorse(morse string) errpr {
		if len(strings.ReplaceAll(strings.ReplaceAll(morse, ".", ""), "-", "")) > 0 {
				return errors.New("Morse strings can only contain di and dah characters!")
	}

	return nil
}

func (tree *TREE) Decode(morse string) string {
		if err := tree.checkMorse(morse); err != nil {
				log.Println(err)
				return ""
		}

		var result string

		switch {
		case morse[0:1] == ".":
			r, err := tree.decode(morse[1:], tree.Root.Dash)
			if err != nil {
					log.Printf("%s (%s), using '_' as placeholder\n", err, morse)
					return "_"
			}
			result = r
		}

		return result
}

func (tree *TREE) decode(morse string, currentNode *Node) (string, error) {
		if len(morse) == 0 {
				result := currentNode.Char
				return result, nil
		}

		if morse[0:1] == "." {
				if currentNode.Dot != nil {
						return tree.decode(morse[1:], currentNode.Dot)
				}
		}

		if morse[0:1] == "-" {
			dash := currentNode.Dash
			if dash != nil {
					return tree.decode(morse[1:], dash)
				}
		}

		return "", errors.New("Morse pattern not found!")
}

func (tree* TREE) Insert(morse string, char string) error {
		if err := tree.checkMorse(morse); err != nil {
				return err
		}

		root := tree.Root

		if morse[0:1] == "." {
				if root.Dot == nil {
						root.Dot = &Node{}
				}

				return tree.insert(morse[1:], char, root.Dot)
		}

		if morse[0:1] == "-" {
				if root.Dash == nil {
						root.Dash = &Node{}
				}

				return tree.insert(morse[1:], char, root.Dash)
		}

		return nil
}

func (tree *TREE) insert(morse, char string, currentNode *Node) error {
		if len(morse) == 0 {
			currentNode.Char = char
			return nil
		}

		if morse[0:1] == "." {
				if currentNode.Dot == nil {
						currentNode.Dot = &Node{}
				}

				return tree.insert(morse[1:], char, currentNode.Dot)
		}

		if morse[0:1] == "-" {
				if currentNode.Dash == nil {
					currentNode.Dash = &Node{}
				}

				return tree.insert(morse[1:], char, currentNode.Dash)
		}

		return nil
}
func NewTree() *TREE {
		tree := new(TREE)
		tree.Root = NewNode("", "")
		return tree
}