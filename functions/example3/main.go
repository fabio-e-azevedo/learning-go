package main

import "fmt"

// Neste exemplo, a função incrementador retorna uma função que incrementa um contador interno e retorna o valor atualizado.
// A função retornada é uma closure porque ela captura a variável contador do escopo da função incrementador.
// Isso permite que a função retornada mantenha o estado do contador entre chamadas.
func incrementador() func() int {
    contador := 0
    return func() int {
        contador++
        return contador
    }
}

func main() {
    // Criar uma função de incremento
    inc := incrementador()

    // Chamar a função de incremento várias vezes
    fmt.Println(inc()) // Saída: 1
    fmt.Println(inc()) // Saída: 2
    fmt.Println(inc()) // Saída: 3
}
