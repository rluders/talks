
# Comparação de Routers em Go

Uma análise detalhada dos routers mais populares em Go

**Apresentado por:**  
Ricardo Lüders

- [linkedin.com/in/ricardoluders](https://www.linkedin.com/in/ricardoluders/)
- [github.com/rluders](https://github.com/rluders)


---

## Quem sou eu?

- **Nome**: Ricardo Lüders
- **Engenheiro de Software** com foco em Go, microsserviços e arquitetura
- Apaixonado por código limpo, escalabilidade e design de software
- Experiência em:
  - Go, PHP (Laravel) e tecnologias web
  - Arquitetura de microsserviços
  - Containers e clusters Kubernetes
- **Conecte-se comigo:**
  - LinkedIn: [linkedin.com/in/ricardoluders](https://www.linkedin.com/in/ricardoluders/)
  - GitHub: [github.com/rluders](https://github.com/rluders)

---

## Filosofia Purista do Go

- Go promove simplicidade e minimalismo
- "Sem dependências, a menos que sejam necessárias"
- Foco em reduzir manutenção e complexidade
- Introduzir bibliotecas apenas quando o benefício justificar o custo

---

## Introdução

Go oferece vários routers HTTP populares

---

## As opções

Vamos comparar as opções mais populares:

- HttpRouter
- Chi
- Gin
- Mux
- Echo
- Go Router (1.22+)

---

## HttpRouter

- Excelente performance e simplicidade
- Baixo uso de memória
- Mínimas dependências
- Mas... não é mais mantido

[Exemplo de Código](./examples/httprouter.go)

---

## Chi

- Leve e minimalista
- Compatível com o `http.Handler` padrão
- Uso de memória muito baixo
- Excelente suporte a middleware
- Sem lock-in, baixas dependências

[Exemplo de Código](./examples/chi.go)

---

## Gin

- Alta performance
- Dependências moderadas
- Forte suporte a middleware e tratamento de erros
- Possível lock-in devido ao wrapping de contextos
- Tamanho médio do pacote

[Exemplo de Código](./examples/gin.go)

---

## Gorilla Mux

- Performance moderada
- Roteamento flexível com regex
- Alto uso de memória e tamanho de pacote
- Tratamento de erros manual
- Não é ideal para a filosofia purista do Go

[Exemplo de Código](./examples/gorilla-mux.go)

---

## Echo

- Performance muito alta
- Mais recursos embutidos (ex: WebSockets)
- Excelente suporte a middleware e tratamento de erros
- Moderado risco de lock-in devido a utilitários
- Tamanho de pacote maior

[Exemplo de Código](./examples/echo.go)

---

## Go Router

- Alta performance
- Sem dependências adicionais (parte da biblioteca padrão)
- Sem sobrecarga no tamanho do pacote
- Uso de memória muito baixo
- Configuração manual de middleware (sem suporte embutido)
- Roteamento baseado em paths com tratamento básico de parâmetros
- Compatível com `http.Handler`
- Muito purista, alinhado com a filosofia de minimalismo do Go
- Sem risco de lock-in

[Exemplo de Código](./examples/go-std.go)

---

## Critérios para Escolher um Router

1. Performance
2. Peso das dependências
3. Tamanho do pacote
4. Uso de memória
5. Suporte a middleware
6. Tratamento de parâmetros
7. Risco de lock-in

---

## Explicações dos Critérios

- **Performance**: Quão rápido o router processa e roteia as requisições.
- **Dependências**: Quantas bibliotecas externas o router depende.
- **Tamanho do Pacote**: O tamanho total do pacote, afetando o deployment e uso de recursos.
- **Uso de Memória**: Quanto de memória o router utiliza durante a execução.
- **Suporte a Middleware**: Quão facilmente o router permite adicionar middlewares (como logging ou autenticação).
- **Tratamento de Parâmetros**: Quão bem o router lida com parâmetros de rota e variáveis de URL.
- **Tratamento de Erros**: As capacidades do router para gerenciar erros.
- **Risco de Lock-in**: A dificuldade de trocar de router no futuro (quanto menor o risco, melhor).

---

## Tabela de Critérios

| Critério                 | Chi | Echo | Gin | Gorilla Mux | Go Router | HttpRouter |
|--------------------------|-----|------|-----|-------------|-----------|------------|
| Performance              | 4   | 5    | 5   | 3           | 4         | 5          |
| Dependências             | 4   | 3    | 3   | 2           | 5         | 4          |
| Tamanho do Pacote        | 4   | 3    | 3   | 2           | 5         | 4          |
| Uso de Memória           | 5   | 4    | 3   | 2           | 5         | 5          |
| Suporte a Middleware     | 4   | 5    | 4   | 4           | 2         | 2          |
| Tratamento de Parâmetros | 4   | 5    | 5   | 4           | 2         | 5          |
| Tratamento de Erros      | 3   | 5    | 4   | 3           | 2         | 2          |
| Risco de Lock-in         | 5   | 3    | 3   | 5           | 5         | 4          |

> Nota: 1 é ruim, 5 é excelente.

---

## Pesos dos Critérios

| Critério                 | Peso |
|--------------------------|------|
| Performance              | 25%  |
| Dependências             | 20%  |
| Risco de Lock-in         | 15%  |
| Suporte a Middleware     | 15%  |
| Tamanho do Pacote        | 10%  |
| Uso de Memória           | 10%  |
| Tratamento de Parâmetros | 10%  |
| Tratamento de Erros      | 5%   |

---

## Top 3 Routers

1. **Chi**
   - Pontuação Ponderada: 460
   - Pontos Fortes: Balanceamento entre performance, baixa dependência, forte suporte a middleware e baixo risco de lock-in

2. **Echo**
   - Pontuação Ponderada: 450
   - Pontos Fortes: Excelente performance, suporte a middleware, tratamento avançado de parâmetros, bom tratamento de erros

3. **Go Router (1.22)**
   - Pontuação Ponderada: 435
   - Pontos Fortes: Zero dependências, baixo uso de memória, alinhado com a filosofia purista do Go, baixo risco de lock-in

**Nota**: HttpRouter (445) ficou em terceiro lugar, mas está depreciado, então foi excluído das opções viáveis.

---

## Resultados

- **Chi** oferece o melhor equilíbrio entre performance, flexibilidade e dependências mínimas, tornando-o uma escolha ideal para a maioria dos projetos em Go.
- **Echo** é uma excelente alternativa para projetos que necessitam de recursos ricos, middleware e tratamento superior de parâmetros.
- **Go Router (1.22)** alinha-se perfeitamente à filosofia purista do Go, oferecendo uma opção simples, sem dependências, para aplicações leves.

> Cada router tem pontos fortes adequados para diferentes necessidades de projeto, e a escolha final deve estar alinhada com as prioridades do projeto, como performance, dependências e flexibilidade futura.

---

## Obrigado!

- Perguntas? Sinta-se à vontade para perguntar!
- Conecte-se comigo em:
  - LinkedIn: [linkedin.com/in/ricardoluders](https://www.linkedin.com/in/ricardoluders/)
  - GitHub: [github.com/rluders](https://github.com/rluders)
- Ansioso para mais discussões e insights.
