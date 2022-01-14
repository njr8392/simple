package main

/* PROBLEM
	Given a sentence text (A sentence is a string of space-separated words) in the following format:

First letter is in upper case.
Each word in text are separated by a single space.
Your task is to rearrange the words in text such that all words are rearranged in an increasing order of their lengths. If two words have the same length, arrange them in their original order.

Return the new text following the format shown above.

Leetcode problem 1451
*/

func arrangeWords(text string) string {
    text = strings.ToLower(string(text[0])) + text[1:]
    words := strings.Split(text, " ")
    sort.SliceStable(words, func(i,j int)bool{
        return len(words[i]) < len(words[j])
    })
    words[0] = strings.ToUpper(string(words[0][0])) + words[0][1:]
    return strings.Join(words, " ")

}
