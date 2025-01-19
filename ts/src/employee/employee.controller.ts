import {
  Controller,
  Get,
  Post,
  Put,
  Delete,
  Param,
  Body,
  UseGuards,
} from '@nestjs/common';
import { AdminGuard } from '../auth/admin.guard';
import { EmployeeService } from './employee.service';
import {
  CreateEmployeeDto,
  UpdateEmployeeDto,
  EmployeeResponseDto,
} from './employee.dto';
import {
  ApiTags,
  ApiOperation,
  ApiBearerAuth,
  ApiBody,
  ApiResponse,
  ApiParam,
} from '@nestjs/swagger';

@ApiTags('Admin Employees')
@ApiBearerAuth()
@UseGuards(AdminGuard)
@Controller('admin/employees')
export class EmployeeController {
  constructor(private readonly employeeService: EmployeeService) {}

  @Post()
  @ApiOperation({ summary: 'Add a new employee' })
  @ApiBody({ type: CreateEmployeeDto })
  @ApiResponse({
    status: 201,
    description: 'Employee created',
    type: EmployeeResponseDto,
  })
  @ApiResponse({
    status: 400,
    description: 'An employee with the same name or email already exists.',
  })
  async createEmployee(
    @Body() employeeData: CreateEmployeeDto,
  ): Promise<EmployeeResponseDto> {
    return this.employeeService.createEmployee(employeeData);
  }

  @Get()
  @ApiOperation({ summary: 'View all employees' })
  @ApiResponse({
    status: 200,
    description: 'A list of employees',
    type: [EmployeeResponseDto],
  })
  async getEmployees(): Promise<EmployeeResponseDto[]> {
    return this.employeeService.getEmployees();
  }

  @Put(':id')
  @ApiOperation({ summary: 'Update an existing employee' })
  @ApiParam({ name: 'id', type: 'string' })
  @ApiBody({ type: UpdateEmployeeDto })
  @ApiResponse({
    status: 200,
    description: 'Employee updated',
    type: EmployeeResponseDto,
  })
  @ApiResponse({
    status: 400,
    description: 'An employee with the same name or email already exists.',
  })
  async updateEmployee(
    @Param('id') id: string,
    @Body() employeeData: UpdateEmployeeDto,
  ): Promise<EmployeeResponseDto> {
    return this.employeeService.updateEmployee(id, employeeData);
  }

  @Delete(':id')
  @ApiOperation({ summary: 'Delete an employee' })
  @ApiParam({ name: 'id', type: 'string' })
  @ApiResponse({
    status: 200,
    description: 'Employee deleted',
    type: EmployeeResponseDto,
  })
  async deleteEmployee(@Param('id') id: string): Promise<EmployeeResponseDto> {
    return this.employeeService.deleteEmployee(id);
  }
}
