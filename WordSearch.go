package main

import "fmt"

// Binary Search Tree
type BSTree struct {
	root *BSTNode
}

// BST Node with children
type BSTNode struct {
	word  string
	left  *BSTNode
	right *BSTNode
}

// Algorithm to insert node into search tree
func insertNode(tree *BSTree, node *BSTNode) {

	curNode := tree.root

	for curNode != nil {

		for i := 0; i < get_min(node, curNode); i++ {
			// if the node is higher in priority, else if node is lower in priority
			if string(node.word[i]) < string(curNode.word[i]) {
				// if the right child isn't empty
				if curNode.right != nil {
					curNode = curNode.right
					break
				} else {
					curNode.right = node
					break
				}
				// fmt.Println(string(node.word) + " is higher alphabetically than " + string(curNode.word))
				// break
			} else if string(node.word[i]) > string(curNode.word[i]) {
				// if the left child isn't empty
				if curNode.left != nil {
					curNode = curNode.left
					break
				} else {
					curNode.left = node
					break
				}
				// fmt.Println(string(node.word) + " is lower alphabetically than " + string(curNode.word))
				// break
			}
		}
		fmt.Println(string(curNode.word))
		// curNode = curNode.right
	}
}

// finds shorter word
func get_min(node *BSTNode, curNode *BSTNode) int {
	if len(node.word) > len(curNode.word) {
		return len(curNode.word)
	} else {
		return len(node.word)
	}
}

// Algorithm to create the search tree alphabetically
func buildTree(arr []string) BSTree {

	newTree := BSTree{nil}

	for i := 0; i < len(arr); i++ {

		// var newNode BSTNode
		newNode := BSTNode{arr[i], nil, nil}

		if i == 0 {
			newTree.root = &newNode
		} else {
			insertNode(&newTree, &newNode)
		}
	}

	return newTree
}

// search the created tree for the word
// returns the word if found, otherwise "Could not be found"
func searchTree(tree *BSTree, word string) string {

}

func main() {

	var arrayList = []string{"fade", "wood", "full", "cluttered", "hard", "save", "measure", "ski", "shame", "cloistered", "advice", "early",
		"windy", "annoying", "work", "acoustic", "prickly", "happen", "frightened", "spot", "tasteful", "roll", "admire", "stain", "magic", "horn",
		"degree", "soap", "reason", "tawdry", "equal", "expensive", "tray", "rush", "brush", "earthy", "unwieldy", "whine", "luxuriant", "offend", "quaint",
		"fluttering", "actor", "hulking", "mindless", "debt", "rigid", "harmony", "slope", "cross", "lunch", "divergent", "evasive", "outgoing", "look", "irate",
		"icicle", "manage", "nonchalant", "eye", "jolly", "snow", "list", "nostalgic", "elfin", "skinny", "show", "basin", "claim", "fair", "soup", "imagine", "injure",
		"ceaseless", "verse", "obtainable", "wasteful", "cry", "testy", "determined", "venomous", "inconclusive", "productive", "irritate", "loving", "wretched", "uninterested",
		"bomb", "obtain", "pickle", "abounding", "agreement", "corn", "matter", "gleaming", "tow", "hover", "voiceless", "mailbox", "special", "sofa", "board", "crush", "stocking",
		"disapprove", "slip", "authority", "lettuce", "quickest", "infamous", "cast", "roomy", "suggestion", "powerful", "soggy", "stormy", "fat", "unbiased", "scarce",
		"average", "recognise", "bulb", "normal", "applaud", "stage", "imperfect", "scared", "allow", "tin", "knowledgeable", "false", "loud", "absorbed", "whistle", "truck",
		"fretful", "waiting", "small", "awesome", "domineering", "tour", "true", "spark", "slimy", "pedal", "cynical", "disagree", "bird", "pray", "notice", "toad", "zealous",
		"practice", "rural", "secretary", "health", "yell", "eyes", "digestion", "compare", "road", "closed", "mundane", "choke", "gun", "invent", "rescue", "interrupt",
		"crook", "trite", "lively", "famous", "root", "fallacious", "alert", "treatment", "ice", "friend", "turkey", "seashore", "concern", "pink", "dinner", "stream",
		"violet", "thing", "angle", "gaudy", "next", "team", "itch", "workable", "thinkable", "pen", "light", "pan", "rainy", "mature", "squeal", "yielding", "basketball",
		"aboriginal", "needle", "mass", "dress", "wink", "flower", "rhythm", "noise", "representative", "astonishing", "handy", "empty", "tasty", "rose", "bouncy", "bewildered",
		"disillusioned", "sad", "caring", "auspicious", "excellent", "nine", "inexpensive", "letters", "chop", "humdrum", "towering", "woozy", "bat", "harm", "standing",
		"dazzling", "mean", "deserve", "lopsided", "somber", "wave", "gainful", "bumpy", "wiry", "overflow", "gray", "tap", "word", "knowledge", "functional", "stove", "touch",
		"press", "eggnog", "damaged", "peck", "trip", "rebel", "beam", "money", "zebra", "appear", "orange", "cheap", "sick", "uneven", "unkempt", "earth", "exuberant",
		"muddled", "ablaze", "start", "powder", "tightfisted", "extra-small", "relation", "sidewalk", "squash", "crowded", "leg", "tip", "grain", "squeamish", "abhorrent", "name",
		"space", "willing", "button", "elite", "income", "refuse", "lighten", "malicious", "flame", "skip", "muddle", "sturdy", "melt", "maniacal", "island", "filthy", "bustling", "lame"}

	tree := buildTree(arrayList)
	newNode := BSTNode{"button", nil, nil}
	// insertNode(&tree, &newNode)
}
