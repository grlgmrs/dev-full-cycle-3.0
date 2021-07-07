# Anotações sobre a aula 1

## Apache Kafka

### Dinâmica de funcionamento

![Dinâmica de funcionamento Kafka](./assets/dinamica_funcionamento_apache_kafka.png)

- Producer: é o sistema que vai enviar o evento/mensagem. 
- Apache Kafka: Funcionamento estilo cluster, onde cada máquina é conhecida como Broker.
- Broker: Cada broker tem seu próprio banco de dados.
- Costumer: Lê os valores do Kafka, procurando eventos/mensagens específicas.

### Tópicos

Podemos criar um "tópico" no Kafka, que seria semelhante à um canal, onde serão armazenadas um tipo de mensagens, como por exemplo, um tópico de vendas. As mensagens são persistidas em disco, não em RAM.
![Tópicos Kafka](./assets/kafka_topicos.png)

### Registro

São a forma das mensagens enviadas nos tópicos.
![Registros Kafka](./assets/registros_kafka.png)

- Header: Podemos passar métadados, para utilizar essa informação de alguma forma;
- Key: Podemos utilizar para as mensagens sempre terem uma ordem desejada;
- Value: Payload da mensagem;
- Timestamp: timestamp da mensagem.

### Partições

As partições são fragmentos de um tópico, sendo assim, elas seriam como uma "lista dividida". Nossos dados originais são: `[a, b, c, d, e, f]`. No caso de termos dois Brokers, poderia acontecer de dividirmos nossa partição em duas, sendo assim `[a, b, c]` e `[d, e, f]`. Sendo assim, no Broker A, teriamos `[a, b, c]` e no Broker B teriamos `[d, e, f]`.

Cada tópico pode ter uma ou mais partições, e estas partições podem estar espalhadas em um ou mais Brokers, sendo assim, caso um Broker caia, você não perderá acesso à todos os dados de um tópico.

Porém, podemos setar um `Replication Factor = x`, que irá replicar nossa partição em `x`, tendo um comportamento semelhante aos pods do Kubernetes, mantendo uma resiliencia dos dados, já que caso um dos Brokers que armazena a partição caia, você ainda terá acesso à uma cópia por outro Broker.

![Exemplo replication factor](./assets/particao_replication_factor.png)

### Consumer

Independentemente de quantos Brokers as partições de um tópico estejam, o Consumer lê as partições sequencialmente (ou seja, lê o tópico) em busca da informação procurada.

Também podemos ter `Consumers Groups`, onde criamos mais de uma máquina rodando exatamente o mesmo software, como vários pods de um mesmo objeto no Kubernetes, com o objetivo de fracionar a responsabilidade de processar as partições. Então, podemos fazer com que meu Consumer A leia a partição 1, ou as partições de 1 a 4, enquanto meu consumer B leia a partição 2, ou as partições de 5 a 10. O objetivo dos `Consumers Groups` é evitar um gargalo no sistema, onde chegam mais mensagens do que efetivamente são processadas.