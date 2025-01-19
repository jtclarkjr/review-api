import { NestFactory } from '@nestjs/core';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { AppModule } from './app.module';
import * as fs from 'fs';
import * as YAML from 'yaml';
import { JwtExceptionFilter } from './auth/jwt-exception';
import { LoggingInterceptor } from './logging.interceptor';
import { LogLevel } from '@nestjs/common';

async function bootstrap() {
  const logLevels: LogLevel[] =
    process.env.NODE_ENV === 'production'
      ? ['error', 'warn']
      : ['error', 'warn']; // ['log', 'error', 'warn', 'debug', 'verbose']

  const app = await NestFactory.create(AppModule, {
    logger: logLevels,
  });
  app.useGlobalFilters(new JwtExceptionFilter());
  app.useGlobalInterceptors(new LoggingInterceptor());

  const config = new DocumentBuilder()
    .setTitle('Employee Performance Review API')
    .setDescription('API for managing employees and performance reviews')
    .setVersion('1.0')
    .addBearerAuth()
    .build();

  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('api', app, document);

  // Convert the document to YAML and save it to a file
  const yamlString = YAML.stringify(document);
  fs.writeFileSync('./swagger.yaml', yamlString, 'utf8');

  await app.listen(3000);
  console.log(`Application is running on: ${await app.getUrl()}`);
}

bootstrap();
