import {
  BeforeInsert,
  Column,
  CreateDateColumn,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { v4 as uuidv4 } from 'uuid';
import { CreditCard } from './credit-card.embbeded';
import { OrderItem } from './order-item.entity';

export enum OrderStatus {
  Approved = 'approved',
  Pending = 'pending',
}

@Entity({ name: 'orders' })
export class Order {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  total: number;

  @CreateDateColumn({ type: 'timestamp' })
  created_at: Date;

  /// Estratégia para deixar o código mais limpo e mais elegante. Aqui, o que acontece, é que
  /// ao invés de eu criar uma entidade, eu posso separar determinados dados de uma tabela ja existente
  /// em um outro objeto para manipulação aqui no código.
  /// No caso atual, a tabela 'orders' possue vários campos, e dentre eles, alguns são em relação ao
  /// cartão de crédito. Portanto, para facilitar a manipulação dos dados do cartão, podemos criar
  /// um objeto que equivale à seus campos na tabela, a parte. Esse objeto pode ser linkado
  /// utilizando o código abaixo, passando como primeiro parâmetro o tipo do objeto e como segundo
  /// o prefixo para poder diferenciar os campos que são pertencentes à esse nosso objeto
  @Column(() => CreditCard, { prefix: '' })
  credit_card: CreditCard;

  /// O cascade faz com que, na hora de de passar os itens, ele já crie em ambas as tabelas automaticamente
  @OneToMany(() => OrderItem, (item) => item.order, { cascade: true })
  items: OrderItem[];

  @Column()
  status: OrderStatus = OrderStatus.Pending;

  @BeforeInsert()
  beforeInsertActions() {
    this.generateId();
    this.calculateTotal();
  }

  generateId() {
    if (this.id) {
      return;
    }

    this.id = uuidv4();
  }

  calculateTotal() {
    return (this.total = this.items.reduce(
      (sum, item) => sum + item.price * item.quantity,
      0,
    ));
  }
}
