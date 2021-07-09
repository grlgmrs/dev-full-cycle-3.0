import { ClassSerializerInterceptor, ValidationPipe } from '@nestjs/common';
import { NestFactory, Reflector } from '@nestjs/core';
import { AppModule } from './app.module';
import { EntityNotFoundExceptionFilter } from './exception-filter/entity-not-found';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.useGlobalFilters(new EntityNotFoundExceptionFilter());

  /// Pipe para validação global nos DTOs
  app.useGlobalPipes(new ValidationPipe({ errorHttpStatusCode: 422 })); // Por padrão, retorna o erro 400, mas vamos mudar pro 422

  /// Em conjunto com o campo 'Exclude' ou outros decorators nas entidades (e nas falsas entidades), evita que esses campos marcados sejam retornados como
  /// resposta
  app.useGlobalInterceptors(new ClassSerializerInterceptor(app.get(Reflector)));
  await app.listen(3000);
}
bootstrap();
