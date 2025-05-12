<!-- ---
marp: true
theme: default
paginate: true
class: lead
style: |
  section {
    font-size: 1.1em;
  }
--- -->

# Curso de Extensão: Linguagem C  
### Ricardo Lüders  
_Revisão: 2025_

---

## Apresentação

**Ricardo Lüders**  
🐙 [github.com/rluders](https://github.com/rluders)  
💼 [linkedin.com/in/ricardoluders](https://www.linkedin.com/in/ricardoluders/)

---

## Sobre o curso

- 5 encontros (20 horas/aula)
- Introdução prática à linguagem C

**80% Prática**  
**20% Teoria**

---

## Objetivos

- Apresentar a sintaxe da linguagem C
- Estrutura básica de um programa
- Boas práticas de programação
- Capacitar o aluno a criar seus próprios programas

---

## Material de apoio

- Curso de Linguagem C – UFMG: http://www.ead.eee.ufmg.br/cursos/C/
- "C How to Program" – Deitel & Deitel
- "C Completo e Total" – Herbert Schildt

---

## Sobre a linguagem C

- Criada nos anos 70 por Dennis Ritchie
- Combina características de alto e baixo nível
- Muito utilizada ainda hoje
- Linguagem **estruturada** e **case sensitive**

```c
int Soma, SOMA, soma; // todas diferentes
```

---

## Prática: Case sensitive

1. Declare quatro variáveis chamadas `soma`, `Soma`, `SOMA` e `SoMa`, com valores diferentes.
2. Imprima cada uma com `printf`.

---

## Estrutura do código

1. Declaração das bibliotecas
2. Declaração de variáveis globais
3. Declaração de funções
4. Função principal

```c
#include <stdio.h>

int main() {
    return 0;
}
```

---

## Palavras reservadas

```text
auto, break, case, char, const, continue, default, do,
double, else, enum, extern, float, for, goto, if,
int, long, register, return, short, signed, sizeof,
static, struct, switch, typedef, union, unsigned, void, volatile, while
```

---

## Prática: Palavra reservada

1. Tente usar uma palavra reservada como nome de variável
2. Observe o erro de compilação

---

## Hello World

```c
#include <stdio.h>

int main() {
    printf("Hello, World!\n");
    return 0;
}
```

---

## Análise do exemplo

- `#include <stdio.h>`: biblioteca padrão de entrada e saída
- `main()`: função principal
- `printf`: imprime na tela
- `return 0`: encerra o programa

---

## Prática: Hello personalizado

1. Modifique o programa para mostrar seu nome

```c
printf("Olá, eu sou <SEU NOME>!\n");
```

---

## Tipos de dados (detalhado)

| Tipo   | Bits | scanf | Intervalo             |
|--------|------|--------|------------------------|
| char   | 8    | %c     | -128 a 127             |
| int    | 16   | %i     | -32.768 a 32.767       |
| float  | 32   | %f     | ~3.4E±38               |
| double | 64   | %lf    | ~1.7E±308              |

---

## Usando printf com tipos

```c
int idade = 25;
float altura = 1.75;
char letra = 'A';

printf("Idade: %d\n", idade);
printf("Altura: %.2f\n", altura);
printf("Letra: %c\n", letra);
```

- `%d` ou `%i` para inteiros
- `%f` para float
- `%lf` para double
- `%c` para caracteres

---

## Prática: Declaração

1. Crie variáveis dos tipos `char`, `int`, `float` e `double`
2. Inicialize e imprima cada uma

---

## Constantes

- Valores fixos no programa
- Tipos:
  - Básicas, hexadecimais, octais, strings, especiais

```c
#define PI 3.14
const int idade = 18;
```

---

## Prática: Constantes

1. Crie uma constante para armazenar o número de alunos
2. Imprima usando `printf`

---

## Comentários e identação

```c
// Comentário de uma linha
/* Comentário de várias linhas */
```

> Identação é a prática de organizar o código-fonte com espaços/tabulações para torná-lo mais legível.

---

## Prática: Comentários e indentação

1. Escreva um código desalinhado e comente
2. Reescreva corretamente indentado

(... o restante do conteúdo segue inalterado ...)

---

## Operadores

- Aritméticos: `+`, `-`, `*`, `/`, `%`
- Atribuição: `=`, `+=`, `-=`
- Relacionais: `==`, `!=`, `>`, `<`
- Lógicos: `&&`, `||`, `!`

---

## Exemplo: operadores

```c
int a = 10, b = 3;
int soma = a + b;
```

---

## Prática: operadores

1. Realize soma, subtração, divisão, multiplicação e módulo
2. Imprima cada resultado

---

## Expressões e precedência

```c
c = a * (b + d) / e;
```

- Tipos diferentes são convertidos conforme regras
- Parênteses influenciam a ordem de execução

---

## Prática: Precedência

1. Teste expressões com e sem parênteses
2. Observe a mudança no resultado

---

## Estruturas de controle: if-else

```c
if (nota >= 7) {
    printf("Aprovado");
} else {
    printf("Reprovado");
}
```

---

## Prática: if-else

1. Leia uma nota com `scanf`
2. Implemente a lógica acima

---

## Estrutura de controle: switch

```c
switch (opcao) {
  case 1:
    printf("Iniciar"); break;
  case 2:
    printf("Sair"); break;
  default:
    printf("Opção inválida");
}
```

---

## Prática: switch

1. Implemente um menu simples com opções

---

## Repetição: while e for

```c
int i = 0;
while (i < 5) {
    printf("%d", i);
    i++;
}

for (int j = 0; j < 5; j++) {
    printf("%d", j);
}
```

---

## Prática: laços

1. Imprima os números de 1 a 10 com `for`
2. Faça o mesmo com `while`

---

## Strings

```c
char nome[50];
scanf("%s", nome);
printf("Olá, %s!", nome);
```

Funções úteis:

- `strcpy`, `strcat`, `strlen`, `strcmp`

---

## Prática: Strings

1. Leia dois nomes e compare com `strcmp`
2. Informe se são iguais

---

## Matrizes

```c
int numeros[5] = {1, 2, 3, 4, 5};
```

Matrizes 2D:

```c
int matriz[3][3] = {
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9}
};
```

---

## Prática: Matrizes

1. Preencha um vetor com 5 valores
2. Calcule a média

---

## Funções

```c
int soma(int a, int b) {
    return a + b;
}
```

- Separação de lógica
- Possuem parâmetros e retorno

---

## Prática: Funções

1. Crie funções para soma, subtração, multiplicação e divisão
2. Chame cada uma a partir do `main`

---

## Visibilidade de variáveis

```c
int global = 100;

int main() {
    int local = 5;
    printf("%d %d", global, local);
    return 0;
}
```

---

## Prática: Visibilidade

1. Crie uma variável global `contador`
2. Incremente-a em uma função chamada várias vezes

---

## Criando bibliotecas

`matematica.h`

```c
int soma(int a, int b);
```

`matematica.c`

```c
int soma(int a, int b) {
    return a + b;
}
```

No `main.c`:

```c
#include "matematica.h"
```

---

## Prática: Biblioteca

1. Crie uma biblioteca com 4 operações básicas
2. Inclua e use em um arquivo `main.c`

---

## Licença

Este material está licenciado sob a:

**Creative Commons Attribution-ShareAlike 4.0 International (CC BY-SA 4.0)**

[https://creativecommons.org/licenses/by-sa/4.0/](https://creativecommons.org/licenses/by-sa/4.0/)

Você pode:
- Usar, copiar, adaptar e redistribuir
- Mesmo para fins comerciais
- Desde que atribua crédito a **Ricardo Lüders**
- E licencie obras derivadas com a mesma licença

---

# Obrigado!

**Ricardo Lüders**  
📧 [github.com/rluders](https://github.com/rluders)  
💼 [linkedin.com/in/ricardoluders](https://www.linkedin.com/in/ricardoluders/)