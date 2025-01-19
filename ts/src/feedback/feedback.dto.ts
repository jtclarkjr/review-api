import { ApiProperty } from '@nestjs/swagger';

export class CreateFeedbackDto {
  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w' })
  reviewId: string;

  @ApiProperty({
    example: 'Very good performance, but needs to improve time management',
  })
  feedback: string;
}

export class FeebackReviewsDto {
  id: string;
  employeeName: string;
  review: string;
}

export class FeedbackResponseDto {
  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w' })
  id: string;

  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w' })
  reviewId: string;

  @ApiProperty({ example: 'cl0m3s93z0000xv8xkhc79p4w' })
  reviewerId: string;

  @ApiProperty({
    example: 'Very good performance, but needs to improve time management',
  })
  feedback: string;
}
