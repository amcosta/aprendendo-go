# Challenge 22: Conversor de Temperaturas

## Introdução

### A Escala Celsius
A água é o elemento mais importante para a vida na terra. A escala Celsius possui o ponto zero na temperatura que a água 
congela e 100 na temperatura que a água ferve. As medidas então são feitas em graus Celsius (°C).

### A Escala Fahrenheit
Daniel Gabriel Fahrenheit escolheu como ponto zero, a temperatura de congelamento de uma mistura de água e sal e o ponto 
máximo (96) a temperatura de um homem sadio. Desta forma o congelamento da água pura ocorre em 32° Fahrenheit (F) e a 
ebulição em 212°F.

### A Escala Kelvin
William Tomson (conhecido como Lord Kelvin) estudando o comportamento do gases, descobriu a menor temperatura que um corpo 
poderia atingir, que seria equivalente a -273°C. A partir daí determinou o ponto zero de sua escala. Criou assim o que 
chamamos de escala absoluta, pois utiliza um fenômeno universal como referência. Nela a água congela em 273 Kelvin (K) e 
ferve a 373 K - repare que não utilizamos graus, pois esta é a escala absoluta e não uma comparação entre fenômenos como 
as outras escalas.

> Fonte: [https://www.infoescola.com/fisica/conversao-de-escalas-termometricas/](https://www.infoescola.com/fisica/conversao-de-escalas-termometricas/)  

## Proposta

A proposta do desafio é criar um programa que faça a conversão de temperaturas, o usuário deve informar dois parâmetros:
- O primeiro é a medida de origem da conversão;
- O segundo é a medida de destino da conversão;

No exemplo abaixo o usuário deseja converter 12 graus Celsius para a escala Fahrenheit.
````sh
# Entrada
programa.php 12C F

# Saída
53.6F
````

### Desafio 1: Validar a entrada

- O programa deve obrigatoriamente receber dois parâmetros, caso contrário deve lançar uma exceção;
- As temperaturas são representadas por letras **(C: Celsius, F: Fahrenheit e K: Kelvin)**, o programa só aceita essas letras 
e em maiúsculo, caso ocorra alguma violação dessas regras o programa deve lançar uma exceção;
- O primeiro parâmetro informa a medida destino em formato numérico (pode aceitar valores negativos e decimais), seguido pela 
letra que representa a temperatura;
- O segundo parâmetro deve ser somente a letra em maiúsculo da medida de destino;

### Desafio 2: Escala Celsius
Converter de Celsius para Fahrenheit usando a formúla: **F = 1.8 * C + 32**

Convertendo 37ºC para Fahrenheit  
> F = 1.8 * 37 + 32  
> F = 66.6 + 32  
> F = 98.6  

Converter de Celsius para Kelvin usando a formúla: **K = C + 273.15**

Convertendo 37°C para a escala Kelvin
> K = 37 + 273.15  
> K = 310.15 

### Desafio 3: Escala Fahrenheit
Converter de Fahrenheit para Celsius usando a formúla: **C = (F - 32) / 1.8**

Exemplo: Convertendo 95°F para Celsius
> C = (95 - 32) / 1.8  
> C = 63 / 1.8  
> C = 35  

Converter de Fahrenheit para Kelvin usando a formúla: **K = (F - 32) / 1.8 + 273.15**

Exemplo: Convertendo 59°F para Kelvin
> K = (59 - 32) / 1.8 + 273.15  
> K = 27 / 1.8 + 273.15  
> K = 15 + 273.15  
> K = 288.15  

### Desafio 4: Escala Kelvin
Converter de Kelvin para Celsius usando a formúla: **K = C + 273.15**

Exemplo: Convertendo 250K para Celsius
> 250 = C + 273.15  
> C = -250 +273.15  
> C = 23.15ºC  

Converter de Kelvin para Fahrenheit usando a formúla: **F = (K − 273,15) * 1.8 + 32**

Exemplo: Convertendo 255.37K para Fahrenheit
> F = (255.37 − 273,15) * 1.8 + 32  
> F = -17.78 * 1.8 + 32  
> F = -32 +32  
> F =  0  
