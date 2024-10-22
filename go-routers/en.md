---
theme: ./theme.slides.json
author: Ricardo Lüders
date: March 14, 2024
---

# Go Routers Comparison

A detailed analysis of popular Go routers

**Presented by:**  
Ricardo Lüders

- [linkedin.com/in/ricardoluders](https://www.linkedin.com/in/ricardoluders/)
- [github.com/rluders](https://github.com/rluders)


---

## Who Am I?

- **Name**: Ricardo Lüders
- **Software Engineer** with a focus on Go, microservices, and architecture
- Passionate about clean code, scalability, and software design
- Experienced in:
  - Go, PHP (Laravel), and web technologies
  - Microservices architecture
  - Container and Kubernetes clusers
- **Connect with me:**
  - LinkedIn: [linkedin.com/in/ricardoluders](https://www.linkedin.com/in/ricardoluders/)
  - GitHub: [github.com/rluders](https://github.com/rluders)

---

## Go Purist Philosophy

- Go promotes simplicity and minimalism
- "No dependencies unless necessary"
- Focus on reducing maintenance and complexity
- Introducing libraries only when the benefit justifies the cost

---

## Introduction

Go offers several popular HTTP routers

---

## The Options

We'll compare it against popular options:

- HttpRouter
- Chi
- Gin
- Mux
- Echo
- Go Router (1.22+)

---

## HttpRouter

- Excellent performance and simplicity
- Low memory footprint
- Minimal dependencies
- But... no longer maintained

[Example Code](./examples/httprouter.go)

---

## Chi

- Lightweight and minimalist
- Compatible with default http.Handler
- Very low memory footprint
- Excellent middleware support
- No lock-in, low dependencies

[Example Code](./examples/chi.go)

---

## Gin

- High performance
- Moderate dependencies
- Strong middleware and error handling
- Possible lock-in due to context wrapping
- Moderate package size

[Example Code](./examples/gin.go)

---

## Gorilla Mux

- Moderate performance
- Flexible regex routing
- High memory and package size
- Manual error handling
- Not ideal for Go’s purist philosophy

[Example Code](./examples/gorilla-mux.go)

---

## Echo

- Very high performance
- More built-in features (e.g., WebSockets)
- Excellent middleware and error handling
- Moderate lock-in due to utilities
- Larger package size

[Example Code](./examples/echo.go)

---

## Go Router

- High performance
- No additional dependencies (part of std library)
- No package size overhead
- Very low memory footprint
- Manual middleware setup (no built-in support)
- Path-based routing with basic param handling
- Compatible with `http.Handler`
- Very purist, aligns with Go's philosophy of minimalism
- No lock-in risk

[Example Code](./examples/go-std.go)

---

## Criteria for Choosing a Router

1. Performance
2. Dependency weight
3. Package size
4. Memory footprint
5. Middleware support
6. Param handling
7. Lock-in risk

---

## Criteria Explanations

- **Performance**: How fast the router processes requests and routes them.
- **Dependency**: How many external libraries the router relies on.
- **Package Size**: The overall size of the package, affecting deployment and resource use.
- **Memory Footprint**: How much memory the router uses during execution.
- **Middleware Support**: How easily the router allows adding middleware (like logging or authentication).
- **Param Handling**: How well the router handles route parameters and URL variables.
- **Error Handling**: The router’s built-in capabilities for managing errors.
- **Lock-in Risk**: The difficulty of switching to a different router in the future (lower risk is better).

---

## Table of Criteria

| Criteria           | Chi | Echo | Gin | Gorilla Mux | Go Router | HttpRouter |
|--------------------|-----|------|-----|-------------|-----------|------------|
| Performance        | 4   | 5    | 5   | 3           | 4         | 5          |
| Dependency         | 4   | 3    | 3   | 2           | 5         | 4          |
| Package Size       | 4   | 3    | 3   | 2           | 5         | 4          |
| Memory Footprint   | 5   | 4    | 3   | 2           | 5         | 5          |
| Middleware Support | 4   | 5    | 4   | 4           | 2         | 2          |
| Param Handling     | 4   | 5    | 5   | 4           | 2         | 5          |
| Error Handling     | 3   | 5    | 4   | 3           | 2         | 2          |
| Lock-in Risk       | 5   | 3    | 3   | 5           | 5         | 4          |

> Note: 1 is poor, 5 is excelent.

---

## Criteria Weights

| Criteria           | Weight |
|--------------------|--------|
| Performance        | 25%    |
| Dependency         | 20%    |
| Lock-in Risk       | 15%    |
| Middleware Support | 15%    |
| Package Size       | 10%    |
| Memory Footprint   | 10%    |
| Param Handling     | 10%    |
| Error Handling     | 5%     |

---

## Top 3 Routers

1. **Chi**
   - Weighted Score: 460
   - Strengths: Balanced performance, low dependency, strong middleware support, low lock-in risk

2. **Echo**
   - Weighted Score: 450
   - Strengths: Excellent performance, middleware support, advanced parameter handling, good error handling

3. **Go Router (1.22)**
   - Weighted Score: 435
   - Strengths: Zero dependencies, low memory footprint, aligns with Go’s purist philosophy, low lock-in risk

**Note**: HttpRouter (445) ranked third but is deprecated, so it's excluded from the viable options.

---

## Results

- **Chi** offers the best balance of performance, flexibility, and minimal dependencies, making it an ideal choice for most Go projects.
- **Echo** is an excellent alternative for projects needing rich features, middleware, and superior parameter handling.
- **Go Router (1.22)** aligns perfectly with Go's purist philosophy, offering a simple, dependency-free option for lightweight applications.

> Each router has strengths suited to different project needs, and the final choice should align with project priorities such as performance, dependencies, and future flexibility.

---

## Thank You!

- Questions? Feel free to ask!
- Connect with me on:
  - LinkedIn: [linkedin.com/in/ricardoluders](https://www.linkedin.com/in/ricardoluders/)
  - GitHub: [github.com/rluders](https://github.com/rluders)
- Looking forward to further discussions and insights.