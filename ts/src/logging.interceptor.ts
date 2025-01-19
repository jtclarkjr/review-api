import {
  CallHandler,
  ExecutionContext,
  Injectable,
  NestInterceptor,
} from '@nestjs/common';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable()
export class LoggingInterceptor implements NestInterceptor {
  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    return next.handle().pipe(
      catchError((err) => {
        // Log the error here
        console.error(
          '------------------------------------------\nError response:',
          {
            statusCode: err.status,
            message: err.message,
            timestamp: new Date().toISOString(),
            path: context.switchToHttp().getRequest().url,
          },
        );
        return throwError(() => err);
      }),
    );
  }
}
