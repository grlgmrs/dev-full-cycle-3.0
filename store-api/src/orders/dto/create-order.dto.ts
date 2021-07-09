import { Type } from 'class-transformer';
import {
  MaxLength,
  MinLength,
  IsString,
  IsNotEmpty,
  IsInt,
  IsObject,
  ValidateNested,
  Min,
  IsUUID,
  IsArray,
  ArrayMinSize,
} from 'class-validator';
import { Product } from 'src/products/entities/product.entity';
import { MinCallback, Exists } from 'src/validators';

class CreditCardDto {
  @MaxLength(16)
  @MinLength(16)
  @IsString()
  @IsNotEmpty()
  number: string;

  @MaxLength(255)
  @IsString()
  @IsNotEmpty()
  name: string;

  /// Foi feito dessa forma para ser executado na hora da validação, e não
  /// no momento que a aplicação foi criada
  @MinCallback(() => new Date().getMonth() + 1)
  @IsInt()
  @IsNotEmpty()
  expiration_month: number;

  @MinCallback(() => new Date().getFullYear())
  @IsInt()
  @IsNotEmpty()
  expiration_year: number;

  @MaxLength(4)
  @IsString()
  @IsNotEmpty()
  cvv: string;
}

export class OrderItemDto {
  @Min(1)
  @IsInt()
  @IsNotEmpty()
  quantity: number;

  @Exists(Product)
  /// versão do UUID
  @IsUUID('4')
  @IsString()
  @IsNotEmpty()
  product_id: string;
}

export class CreateOrderDto {
  @Type(() => CreditCardDto)
  /// Serve para ele fazer as validações dentro do CreditCardDto, também
  @ValidateNested()
  @IsObject()
  @IsNotEmpty()
  credit_card: CreditCardDto;

  @Type(() => OrderItemDto)
  /// Precisamos passar o { each: true } para ele validar todos os itens do array
  @ValidateNested({ each: true })
  @ArrayMinSize(1)
  @IsArray()
  @IsNotEmpty()
  items: OrderItemDto[];
}
