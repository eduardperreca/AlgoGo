## Definizione di Algoritmo

### Definizione

<aside>
💡 Insieme ordinato e finito di passi eseguibili e non ambigui che definiscono un procedimento che termina.

</aside>

---

### Algoritmi di Approssimizzazione (random)

Un esempio é quella del calcolo di un integrale, immaginando di calcolarlo disegnando in un quadrato che rappresenta l’area della proiezione della funzione quanti punti stanno al disotto e al disopra, i punti che stanno al disotto sono parte della funzione quindi fanno parte della soluzione esprimibile in percentuale.

---

### Algoritmi per problemi difficili con l’approssimizzazione

Posso accettare un valore anche errato purché il margine rimanga molto basso, per esempio garantire se un numero é primo o no.

---

### Visione Matematica

Un algoritmo $a$ é rappresentabile come:
funzione $f_a = D_I \rightarrow D_S$
Dove $D_I$ é il dominio delle istanze e $D_S$ é il dominio delle soluzioni

---

Problema $P$
Istanza $x \in D_I$
soluzione $f(x) \in D_S$

Questo é risolvibile quando la $f_a(x)$ = $f(x)$

$a$ risolve $P \iff$$f_a(x) = f(x)$

---

$D_I = N\times N$ (istanza $(x, y) \in \N \times \N$)

$D_S = \N$ $f(x) = x*y$

---

## Definizione di Programma

<aside>
💡 Insieme ordinato e finito di istruzioni scritte secondo le regole di uno specifico linguaggio di programmazione

</aside>

### Algoritmo moltiplicazione(intero a, intero b) $\rightarrow$ intero

```go
Algoritmo moltiplicazione(intero a, intero b) intero
	return a*b
```

---

### Progettazione di un algoritmo

<aside>
💡 Dato un problema progetto un algoritmo che lo risolve

</aside>

Bisogna controllare:

Correttezza:

Dato un algoritmo $A$ e un problema $P$ dimostrare che $A$ risolve $P$

Efficienza:

Valutare la quantitá di risorse utilizzate da un algoritmo

### Come valuto un di un algoritmo?

Eseguo l’algoritmo e controllo con dei testing la valutazione a posteriori.

Un altro metodo é la valutazione a priori, stimando in fase di progettazione la correttezza e le risorse utilizzate controllandone l’efficienza.

---

### Algoritmo moltiplicazione basilare

19 x 114 = 2166 ma come si fa?
moltiplico ogni cifra del moltiplicatore per il moltiplicando metodo elementare

per calcolare il numero di cifre da cui é formata un numero bisogna fare il log con la base il numero del numero.

---

## Algoritmo Somma Iterata

$a * b = a + a ... + a + a$ ($b$ volte)

```go
/* a, b >= 0 */
ALGORITMO moltiblicazione(intero a, intero b) -> intero
	prod <- 0
	WHILE(b > 0) DO
		prod <- prod+a
		b <- b-1
return prod
/* ----------------- */
func moltiplicazione(int a, int b) int{
	for b; b>0; b--{
		prod += a
	}
	return prod
}
```

time complexity = O(n)

space complexity = (idk)

$b_i = b - i$

$prod_i = a*i$

---

controlliamo per induzione

caso base:

$i = 0$

$b_0 = b$

$prod_0 = 0$

passo induttivo:
$i - 1 = i$

$b_i = b_{i-1} = b - (i-1) - 1 = b - i$

$prod_i = prod_{i-1} + a = a * (i-1)+ a = a*i$

$b_i = 0 \rightarrow$ se $i = b$ iterazione $b: prod_b = a*b$

---

### Calcolo time complexity algoritmo somma iterata

Il tempo dipende da quante iterazioni faccio, legata al valore di b

se b = 0:

- linea 1, 2, 3
  costo 3

se b > 0:

- linea 1, 5
  costo 2
- linea 3, 4
  costo 2b
- linea 2
  costo b + 1

somma 3b + 3 = b volte
$T(a, b) = 3b + 3$

---

## Algoritmo moltiplicazione alla russa

moltiplicando(a) = 19

moltiplicatore(b) = 114

moltiplico il moltiplicando \* 2 e il moltiplicatore / 2
prendo quelle coppie di valori dove il moltiplicatore é dispari e sommo il valore associato, il valore é lo stesso.

```go
ALGORITMO MoltiplicazioneRussa(intero a, intero b) intero
	prod <- 0
	WHILE b > 0 DO
		if b dispari THEN
			prod <- prod + a
		b <- b/2
		a <- a*2
return prod
```
