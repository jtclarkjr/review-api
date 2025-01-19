import { Module } from '@nestjs/common';
import { EmployeeService } from './employee.service';
import { PrismaModule } from 'prisma/prisma.module';
import { AuthModule } from '../auth/auth.module'; // Import AuthModule
import { EmployeeController } from './employee.controller';

@Module({
  imports: [PrismaModule, AuthModule], // Add AuthModule here
  controllers: [EmployeeController],
  providers: [EmployeeService],
})
export class EmployeeModule {}
