# Curso de Go

Esse material é relativo ao curso de Go. Os códigos aqui disponibilizados foram baseados inicialmente nos distribuídos por [Wilson Junior](https://github.com/wpjunior).

- https://github.com/wpjunior/go-course
- https://github.com/wpjunior/go-course-2
- https://github.com/wpjunior/go-course-3

Os materiais nas pastas são exercícios que reforçam o aprendizado para cada dia de curso.

## Desafios

Dentro da pasta de cada dia, existe um conjunto de desafios específico de cada tópico. Por exemplo, no dia 1, existe o desafio `04-funcoes` que fica na pasta `00-challenges/04-funcoes`. Para resolver o desafio, basta editar o arquivo main.go que existe na pasta e executar (de dentro da pasta) o seguinte comando: `make -f ../Makefile`. Veja a seguir um passo a passo completo.

```
% cd dia1
% cd 04-funcoes # aqui você viu como é uma função em Go
% cd ../00-challenges/04-funcoes
% vi main.go
% make -f ../Makefile
go fmt main.go
go run main.go < main.in | diff -w main.out -
1c1
< 5
---
> 2
make: *** [main.in] Error 1 # esse erro indica que o seu desafio ainda não foi concluído. Tente de novo!
```

## Referências

A seguir, uma lista de referências que podem complementar os estudos de Go.
- HAKIMIAN, Parsia. Hacking-with-Go. [s.l.]: Github, [s.d.]. Disponível em: <https://github.com/parsiya/Hacking-with-Go>.
- awesome-golang-security. [s.l.]: Github, [s.d.]. Disponível em: <https://github.com/guardrailsio/awesome-golang-security>.
- BlackHat Go Code Samples. [s.l.]: Github, [s.d.]. Disponível em: <https://github.com/blackhat-go/bhg>.
- JAMES, Chris. Aprenda Golang com testes. Disponível em: <https://larien.gitbook.io/aprenda-go-com-testes/>.
- HAKIMIAN, Parsia. Hacking with Go. [s.l.]: Github, [s.d.]. Disponível em: <https://github.com/parsiya/Hacking-with-Go>.