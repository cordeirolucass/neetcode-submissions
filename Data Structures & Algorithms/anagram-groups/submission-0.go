

func groupAnagrams(strs []string) [][]string {
    tabelahash := make(map[string][]string)
    for i := 0; i < len(strs); i++ {
        chaves := []byte(strs[i])
        sort.Slice(chaves, func(a, b int) bool {return chaves[a] < chaves[b]})
        chaveStr := string(chaves)
        tabelahash[chaveStr] = append(tabelahash[chaveStr], strs[i])
    }
    resposta := [][]string{}
    for _, grupo := range tabelahash {
        resposta = append(resposta, grupo)
    }
    return resposta
}
