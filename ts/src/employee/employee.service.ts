import { Injectable, BadRequestException } from '@nestjs/common';
import { PrismaService } from 'prisma/prisma.service';
import {
  CreateEmployeeDto,
  UpdateEmployeeDto,
  EmployeeResponseDto,
} from './employee.dto';
import { Employee } from '@prisma/client';

@Injectable()
export class EmployeeService {
  constructor(private prisma: PrismaService) {}

  private async checkEmployeeExists(
    name: string,
    email: string,
    excludeId?: string,
  ): Promise<void> {
    const conditions: any = {
      OR: [{ name }, { email }],
    };

    if (excludeId) {
      conditions.NOT = {
        id: excludeId,
      };
    }

    const existingEmployee = await this.prisma.employee.findFirst({
      where: conditions,
    });

    if (existingEmployee) {
      throw new BadRequestException(
        'An employee with the same name or email already exists.',
      );
    }
  }

  async createEmployee(data: CreateEmployeeDto): Promise<EmployeeResponseDto> {
    await this.checkEmployeeExists(data.name, data.email);

    const employee = await this.prisma.employee.create({ data });
    return this.mapToEmployeeResponseDto(employee);
  }

  async getEmployees(): Promise<EmployeeResponseDto[]> {
    const employees = await this.prisma.employee.findMany();
    return employees.map(this.mapToEmployeeResponseDto);
  }

  async updateEmployee(
    id: string,
    data: UpdateEmployeeDto,
  ): Promise<EmployeeResponseDto> {
    await this.checkEmployeeExists(data.name, data.email, id);

    const existingEmployee = await this.prisma.employee.findUnique({
      where: { id },
    });

    if (!existingEmployee) {
      throw new BadRequestException('Employee not found.');
    }

    const employee = await this.prisma.employee.update({
      where: { id },
      data,
    });
    return this.mapToEmployeeResponseDto(employee);
  }

  async deleteEmployee(id: string): Promise<EmployeeResponseDto> {
    const existingEmployee = await this.prisma.employee.findUnique({
      where: { id },
    });

    if (!existingEmployee) {
      throw new BadRequestException('Employee not found.');
    }

    const employee = await this.prisma.employee.delete({
      where: { id },
    });
    return this.mapToEmployeeResponseDto(employee);
  }

  private mapToEmployeeResponseDto(employee: Employee): EmployeeResponseDto {
    return {
      id: employee.id,
      name: employee.name,
      email: employee.email,
      isAdmin: employee.isAdmin,
    };
  }
}
