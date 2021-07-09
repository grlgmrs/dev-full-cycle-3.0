# Esse arquivo faz uma modificação da imagem postgres

FROM postgres:11

# muda o usuário do postgres para o id 1000, que equivale ao usuário da máquina e
# evita erros quando trabalharmos com volumes
RUN usermod -u 1000 postgres