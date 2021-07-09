import slugify from 'slugify';
import {
  BeforeInsert,
  Column,
  CreateDateColumn,
  Entity,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { v4 as uuidv4 } from 'uuid';

@Entity({ name: 'products' })
export class Product {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  name: string;

  @Column()
  description: string;

  @Column()
  image_url: string;

  @Column()
  slug: string;

  @Column()
  price: number;

  @CreateDateColumn({ type: 'timestamp' })
  created_at: Date;

  @BeforeInsert()
  beforeInsertActions() {
    this.generateId();
    this.generateSlug();
  }

  generateId() {
    if (this.id) {
      return;
    }

    this.id = uuidv4();
  }

  generateSlug() {
    if (this.slug) {
      return;
    }

    this.slug = slugify(this.name);
  }
}
