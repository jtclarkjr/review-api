import { Request } from 'express';
import { Employee } from '@prisma/client';

export interface AuthenticatedRequest extends Request {
  user: Employee;
}