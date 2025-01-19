import { ApiProperty } from '@nestjs/swagger';
import { EmployeeResponseDto } from 'src/employee/employee.dto';

export class CreateReviewDto {
  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w' })
  employeeId: string;

  @ApiProperty({ example: 'Excellent performance in Q1' })
  review: string;
}

export class UpdateReviewDto {
  @ApiProperty({ example: 'Excellent performance in Q1', required: false })
  review?: string;

  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w', required: false })
  employeeId?: string;
}

export class AssignReviewersDto {
  @ApiProperty({
    example: ['cl0m3s93z0000xv8xkhc79p4w', 'cl0m3s93z0000xv8xkhc79p4x'],
  })
  reviewerIds: string[];
}

export class ReviewDto {
  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w' })
  id: string;

  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w' })
  employeeId: string;

  @ApiProperty({ example: 'Excellent performance in Q1' })
  review: string;

  @ApiProperty({ type: [EmployeeResponseDto] })
  reviewers: EmployeeResponseDto[];
}
