import {
  Controller,
  Get,
  Post,
  Put,
  Delete,
  Param,
  Body,
  UseGuards,
  UseFilters,
} from '@nestjs/common';
import { ReviewService } from './review.service';
import {
  ApiTags,
  ApiOperation,
  ApiParam,
  ApiBody,
  ApiResponse,
  ApiBearerAuth,
} from '@nestjs/swagger';
import { AdminGuard } from '../auth/admin.guard';
import {
  CreateReviewDto,
  UpdateReviewDto,
  AssignReviewersDto,
  ReviewDto,
} from './review.dto';
import { Review } from '@prisma/client';
import { JwtExceptionFilter } from 'src/auth/jwt-exception';

@ApiTags('Reviews')
@ApiBearerAuth()
@Controller('admin/reviews')
@UseFilters(JwtExceptionFilter)
@UseGuards(AdminGuard)
export class ReviewController {
  constructor(private readonly reviewService: ReviewService) {}

  @Post()
  @ApiOperation({ summary: 'Create a new review' })
  @ApiBody({ type: CreateReviewDto })
  @ApiResponse({
    status: 201,
    description: 'The review has been successfully created.',
  })
  async createReview(@Body() reviewData: CreateReviewDto): Promise<Review> {
    return this.reviewService.createReview(reviewData);
  }

  @Put(':id')
  @ApiOperation({ summary: 'Update an existing review' })
  @ApiParam({ name: 'id', type: 'string' })
  @ApiBody({ type: UpdateReviewDto })
  @ApiResponse({
    status: 200,
    description: 'The review has been successfully updated.',
  })
  async updateReview(
    @Param('id') id: string,
    @Body() reviewData: UpdateReviewDto,
  ): Promise<Review> {
    return this.reviewService.updateReview(id, reviewData);
  }

  @Delete(':id')
  @ApiOperation({ summary: 'Delete a review' })
  @ApiParam({ name: 'id', type: 'string' })
  @ApiResponse({
    status: 200,
    description: 'The review has been successfully deleted.',
  })
  async deleteReview(@Param('id') id: string): Promise<Review> {
    return this.reviewService.deleteReview(id);
  }

  @Get()
  @ApiOperation({ summary: 'Get all reviews' })
  @ApiResponse({
    status: 200,
    description: 'Returns all reviews.',
    type: [ReviewDto],
  })
  async getReviews(): Promise<ReviewDto[]> {
    return this.reviewService.getReviews();
  }

  @Post(':id/assign')
  @ApiOperation({ summary: 'Assign reviewers to a review' })
  @ApiParam({ name: 'id', type: 'string' })
  @ApiBody({ type: AssignReviewersDto })
  @ApiResponse({
    status: 201,
    description: 'The reviewers have been successfully assigned.',
  })
  async assignReviewers(
    @Param('id') id: string,
    @Body() assignReviewersDto: AssignReviewersDto,
  ): Promise<void> {
    await this.reviewService.assignReviewers(
      id,
      assignReviewersDto.reviewerIds,
    );
  }
}
