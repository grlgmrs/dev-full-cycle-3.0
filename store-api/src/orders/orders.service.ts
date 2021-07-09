import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Product } from 'src/products/entities/product.entity';
import { In, Repository } from 'typeorm';
import { CreateOrderDto } from './dto/create-order.dto';
import { UpdateOrderDto } from './dto/update-order.dto';
import { Order } from './entities/order.entity';

@Injectable()
export class OrdersService {
  constructor(
    @InjectRepository(Order) private orderRepo: Repository<Order>,
    @InjectRepository(Product) private productRepo: Repository<Product>,
  ) {}

  async create(createOrderDto: CreateOrderDto) {
    const order = this.orderRepo.create(createOrderDto);
    const productPrices = (
      await this.productRepo.find({
        select: ['id', 'price'],
        where: {
          /// "In" verifica se o valor é igual à um dos que foram passados
          id: In(order.items.map((item) => item.product_id)),
        },
      })
    ).reduce(
      (prices, product) => ({ ...prices, [product.id]: product.price }),
      {},
    );

    order.items.forEach((item) => {
      item.price = productPrices[item.product_id];
    });

    return this.orderRepo.save(order);
  }

  findAll() {
    return `This action returns all orders`;
  }

  findOne(id: string) {
    return `This action returns a #${id} order`;
  }

  update(id: string, updateOrderDto: UpdateOrderDto) {
    return `This action updates a #${id} order`;
  }

  remove(id: string) {
    return `This action removes a #${id} order`;
  }
}
