## Algoritmo moltiplicazione alla russa

moltiplicando(a) = 19

moltiplicatore(b) = 114

moltiplico il moltiplicando * 2 e il moltiplicatore / 2
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

$a*b = 2a * b/2$ se divisione tra reali 

$a *b = \{2a * [b/2] 
, 2a*[b/2] + a$ , la prima se b é pari, la seconda se b é dispari e $\neq$ 1

### Calcolo time complexity moltiplicazione alla russa

```go
ALGORITMO MoltiplicazioneRussa(intero a, intero b) intero
	prod <- 0                              
	WHILE b > 0 DO
		if b dispari THEN
			prod <- prod + a
		b <- b/2
		a <- a*2
return prod

/* time complexity
Linea 1, 7 -> 1 volta.   ( 2 )
Linea 2 -> u + 1 volte.  ( u + 1 )
Linea 3, 5, 6 -> u volte ( 3u )
Linea 4 -> <= u volte    ( <= u )
```

| b | u |
| --- | --- |
| 0 | 0 |
| 1 | 1 |
| 2 | 2 |
| 3 | 2 |
| 4 | 3 |
| 5 | 3 |

$u = [log_2b] + 1$

$T(a, b) = 5([log_2b] + 1) +3$

$T(a, b) = 5[log_2b] + 8$

---

### Paragone time complexity tra somma iterata e metodo russo

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e33959d0-9ad6-4be3-b5a4-f745379314ef/Untitled.png)

---

## Algoritmo potenza

```go
ALGORITMO Potenza(intero x, intero y) intero
	power <- 1
	WHILE y>0 DO
		power <- power*x
		y <- y-1
return power 
// time complexity:
// 3y + 3
// space complexity:
// 3 (x, y, power)
```

se y é pari e diverso da 0 → $(x^{y/2})^2$
se y é dispari → $(x^{y-1/2})^2 *x$
se y é 0 → 1

---

### Calcolo Complessitá tempo

```go
AlGORITMO Potenza(intero x, intero y) intero
	if y = 0 then
		return 1
	else
		power <- Potenza(x, y/2) //  divisione intera
		power <- power * power
		if y é dispari then
			power <- power * x
		return power

// time complexity:
// T(x, y)
// se y = 0 1, 2                              costo: 2
// se y > 0 1, 2, 3, 4, 5, 7                  costo: 5
// linea 6 <= 1 volta                         costo: <= 1
// linee eseguite per calcolare (x, y/2).     costo: T(x, [y/2])
//                                            ------------------
//                                             <= 6 + T(x, [y/2])
```

T(x, y) ≤ { 2 se y = 0, T( x, [y/2] + 6 )} , altrimenti equazione di ricorrenza

T(x, 1) = T(x, 0) + 6 = 2 + 6 = 8

T(x, y) = {2 se y = 0, 8 se y = 1, T(x, y/2) + 6 altrimenti}

T(x, y) = T(x, y/2) + 6 = T(x, y/2$^2$) + 6 + 6

T(x, y/2$^k$) + 6k → y/2$^k$ = 1 dove 2$^k$ = y , k = $log_2n$ ( logaritmica )

T( x, y ) ≤ $6log_2y + 8$

---

### Calcolo Complessitá Spazio

```go
//chiamo la funzione potenza
potenza(2, 6)
//chiamo poi
potenza(2, 3)
//chiamo poi
potenza(2, 1)
//chiamo poi
potenza(2, 0)
```