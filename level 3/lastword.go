package piscine

func LastWord(s string) string {
    if s == " "{
        return "\n"
    }
    word := ""
    words := []string{}

    for _,char:=range s{
        if char == ' ' && word != ""{
            words=append(words,word)
            word=""
        } else if char != ' '{
            word=word + string(char)
        }
    }
    if word != ""{
        words=append(words,word)
    }
    
    return words[len(words)-1] + "\n"
}