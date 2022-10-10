## Tempo polinomiale n$^2$

|  | 1 | 2 | 3 | 4 | 5 |
| --- | --- | --- | --- | --- | --- |
| ⁍ | 1 | 2 | 3 | 434232 | non ho voglia |
| ⁍ |  |  |  |  |  |
- Algoritmi polinomiali
Gli algoritmi che lavorano in tempo limitato da un polinomio (rispetto alla lunghezza dell’input)
- Algoritmi Esponenziali
Gli algoritmi utilizzati che usano tempo esponenziale rispetto alla lunghezza degli input.

---

## Investire in hardware o in algoritmi?

algoritmi.

---

## Strutture indicizzate (array)

Un array é una collezione di elementi tutti dello stesso tipo, ciascuno dei quali é accessibile in base alla posizione `i`

Caratteristiche tipiche:

- Memorizzato in una porzione contigua di memoria
- Accesso mediante indice
- Tempo di accesso indipendente dalla posizione del dato

Limitazione:

- Struttura statica: non é possibile aggiungere nuove posizioni

> Nota:
Nei singoli linguaggi di programmazionie alcune caratteristiche degli array e delle relative variabili possono essere differenti (come in go con le slices)
> 

Variabili Array:

- Sono riferimenti (puntatori) all’array

```go
B <- A
B[1] <- 0
stampa(A)
```

### Procedura azzera

```go
PROCEDURA Azzera(Array X[0, ..., n-1])
for i <- 0 to n-1 do
	x[i] <- 0
```

---

## Ricerca in un Array

### Definizione del problema

Input : Array A, elemento x

Output : Indice i t.c. A[i] = x, -1 se A non contiene x

### Ricerca Sequenziale

```go
ALGORITMO ricercaSequenziale(Array A[0, ..., n-1], elemento) indice
	i <- 0
	WHILE i<n AND A[i] != x DO
		i <- i + 1
	IF i = n THEN RETURN -1
					 ELSE RETURN i
```

### Complessitá di Tempo

O(n)

### Complessitá di Spazio

se i costi di confronto sono = 1 allora la complessitá di tempo é O(n)

---

### Ricerca di array ordinato (binary search)

se l’array é ordinato in maniera non decrescente.

avendo un array ordinato, quello che posso fare é controllare se dividendolo a metá controllando che  x < o x > di A[meta]

---

### Algoritmo ricorsivo di binary search

```go
FUNZIONE RicercaRic(Array A, indice sx, indice dx, elemento x) -> indice
	// porzione vuota
	if dx <= sx then return -1
	else 
		m <- (sx + dx) / 2
		if x = A[m] then return m
		else if x < A[m] then 
											return RicercaRic(A, sx, m, x)
		else
				return ricercaRic(A, m+1, dx, x)
// DA IMPLEMENTARE!! 
```

```go
ALGORITMO RicercaBinaria(array A, elemento x) -> indice

	return RicercaRic(A, D, n, x)

```

Esercizio:

trovare x tale che il numero di chiamate sia 2logn

spazio logaritmico, tempo logaritmico

---