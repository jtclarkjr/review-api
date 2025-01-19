import { Module } from '@nestjs/common';
import { AuthModule } from './auth/auth.module';
import { PrismaModule } from 'prisma/prisma.module';
import { EmployeeModule } from './employee/employee.module';
import { ReviewModule } from './review/review.module';
import { FeedbackModule } from './feedback/feedback.module';

@Module({
  imports: [
    AuthModule,
    PrismaModule,
    EmployeeModule,
    ReviewModule,
    FeedbackModule,
  ],
})
export class AppModule {}
