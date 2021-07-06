# Anotações sobre a aula 1

## Golang
Para iniciarmos o projeto, utilizaremos:
```sh
$ go mod init github.com/codeedu/codebank
```
E em seguia, precisaremos baixar nossos módulos conforme necessário. Para isso, utilizaremos:
```sh
$ go get -v ${PATH}
```
Não entendi bem como dele decide baixar os módulos, mas aparentemente, os arquivos que importam os módulos devem estar na mesma pasta especificada em ${PATH}