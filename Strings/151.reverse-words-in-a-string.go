func reverseWords(s string) string {
    
    var words [][]rune
    var word []rune
    var str []rune

    var result []rune


    for i, letter := range s{


        if i == len(s) -1 && letter == 32{
            break
        }
        
        if letter == 32 && len(str) == 0{
            continue
        }

        if !(letter == 32 && s[i+1] == 32){
            str = append(str, letter)
        }

    }

    for i, letter := range str{
        
        if  i == len(str) - 1{
            word = append(word, letter)
            words = append(words, word)
            break
        }

        if letter != 32 {
            word = append(word, letter)
        }else{
            words = append(words, word)
            word = []rune{}
        }

    }

    for i :=  len(words)-1; i >= 0; i--{
        result = append(result, words[i]...)
        if i != 0{
            result = append(result, 32)
        }
    }

    return string(result)
}