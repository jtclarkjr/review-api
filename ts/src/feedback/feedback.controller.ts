import {
  Controller,
  Get,
  Post,
  Param,
  Body,
  UseGuards,
  Req,
} from '@nestjs/common';
import { FeedbackService } from './feedback.service';
import {
  CreateFeedbackDto,
  FeebackReviewsDto,
  FeedbackResponseDto,
} from './feedback.dto';

import {
  ApiTags,
  ApiOperation,
  ApiBearerAuth,
  ApiBody,
  ApiResponse,
  ApiParam,
} from '@nestjs/swagger';
import { UserGuard } from 'src/auth/user.guard';
import { AuthenticatedRequest } from 'src/auth/authenticated-request.interface';

@ApiTags('Feedback')
@ApiBearerAuth()
@UseGuards(UserGuard)
@Controller('feedback')
export class FeedbackController {
  constructor(private readonly feedbackService: FeedbackService) {}

  @Get()
  @ApiOperation({ summary: 'Get reviews assigned to current user' })
  @ApiResponse({
    status: 200,
    description: 'Returns the reviews for the current user.',
  })
  async getUserReviews(@Req() req: AuthenticatedRequest): Promise<FeebackReviewsDto[]> {
    const reviewerId = req.user.id;
    return this.feedbackService.getUserReviews(reviewerId);
  }

  @Post()
  @ApiOperation({ summary: 'Submit feedback for a performance review' })
  @ApiBody({ type: CreateFeedbackDto })
  @ApiResponse({
    status: 201,
    description: 'Feedback submitted',
    type: FeedbackResponseDto,
  })
  async createFeedback(
    @Req() req: AuthenticatedRequest,
    @Body() feedbackData: CreateFeedbackDto,
  ): Promise<FeedbackResponseDto> {
    const reviewerId = req.user.id;
    return this.feedbackService.createFeedback(reviewerId, feedbackData);
  }

  @Get('me')
  @ApiOperation({ summary: 'Get feedbacks submitted by the current user' })
  @ApiResponse({
    status: 200,
    description: 'A list of feedbacks',
    type: [FeedbackResponseDto],
  })
  async getFeedbacks(@Req() req: AuthenticatedRequest): Promise<FeedbackResponseDto[]> {
    const reviewerId = req.user.id;
    return this.feedbackService.getFeedbacks(reviewerId);
  }

  @Get('review/:reviewId')
  @ApiOperation({ summary: 'Get feedbacks for a specific review' })
  @ApiParam({ name: 'reviewId', type: 'string' })
  @ApiResponse({
    status: 200,
    description: 'A list of feedbacks for the review',
    type: [FeedbackResponseDto],
  })
  async getReviewFeedbacks(
    @Param('reviewId') reviewId: string,
  ): Promise<FeedbackResponseDto[]> {
    return this.feedbackService.getReviewFeedbacks(reviewId);
  }
}
