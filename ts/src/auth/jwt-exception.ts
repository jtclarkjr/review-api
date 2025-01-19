import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpException,
} from '@nestjs/common';
import { Response } from 'express';

interface HttpExceptionResponse {
  statusCode: number;
  message: string | string[];
  error?: string;
}

@Catch(HttpException)
export class JwtExceptionFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const status = exception.getStatus();
    const exceptionResponse = exception.getResponse();

    const message = this.getMessage(exceptionResponse);

    response.status(status).json({
      statusCode: status,
      message,
    });
  }

  private getMessage(exceptionResponse: string | object): string {
    if (typeof exceptionResponse === 'string') {
      return exceptionResponse;
    }

    const response = exceptionResponse as HttpExceptionResponse;
    return response.message instanceof Array
      ? response.message.join(', ')
      : response.message;
  }
}
