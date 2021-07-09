# Anotações sobre a aula 3

## Dockercompose

Podemos executar alguns scripts quando nosso container sobe, para isso basta definirmos o `entrypoint` na docker-compose.yaml e passar o path onde o arquivo a ser executado se encontra.

## NextJS

### Arquivos de chamada global

No next temos alguns arquivos que sempre são chamados, independente das páginas que estivermos acessando, sendo alguns deles:
- `_app.tsx`
- `_document.tsx`

### Tipos de páginas

Existem alguns tipos de páginas no NextJS, como:
- Server Side Renders Runtime: Utilizam os métodos `getInitialProps` ou `getServerSideProps` e são dinâmicas pois seus valores podem variar mesmo sem ter que buildar novamente;
- Static rendered: são páginas estáticas;
- Static rendered + JSON: também são páginas estáticas, mas puxam informações de algun(s) JSON(s);
- Incremental Static Regeneration: Mesmo comportamento do anterior, mas ele gera novamente o html de tempos em tempos

### Bugs

Achei um bug extremamente esquisito quando usei arquivos com nomes de parâmetros, onde eu não conseguia buildar de jeito nenhum o projeto. O único jeito de fazer ele buildar foi depois de alterar alguma coisa dentro das funções `getStaticProps` e `getStaticPaths`. As vezes, apenas deletar um caractere de algo não funcionou, mas adicionar uma nova linha, sim... Bizarro, muito bizarro

## NestJS

## TypeORM

Comandos para manipular as migrations
- `migration:run`: Cria as migrations
- `migration:revert`: Faz rollback das últimas migrations