import { ApiProperty } from '@nestjs/swagger';

export class CreateEmployeeDto {
  @ApiProperty({ example: 'John Doe' })
  name: string;

  @ApiProperty({ example: 'john.doe@example.com' })
  email: string;

  @ApiProperty({ example: 'password123' })
  password: string;

  @ApiProperty({ example: false })
  isAdmin: boolean;
}

export class UpdateEmployeeDto {
  @ApiProperty({ example: 'John Doe' })
  name?: string;

  @ApiProperty({ example: 'john.doe@example.com' })
  email?: string;

  @ApiProperty({ example: 'password123' })
  password?: string;

  @ApiProperty({ example: false })
  isAdmin?: boolean;
}

export class EmployeeResponseDto {
  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w' })
  id: string;

  @ApiProperty({ example: 'John Doe' })
  name: string;

  @ApiProperty({ example: 'john.doe@example.com' })
  email: string;

  @ApiProperty({ example: false })
  isAdmin: boolean;
}
